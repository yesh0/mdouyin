package snowy_test

import (
	"common/snowy"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdAtomic(t *testing.T) {
	assert.Nil(t, snowy.Init("127.0.0.1:2379"))
	rand.Seed(time.Now().Unix())

	randStr := strconv.FormatInt(rand.Int63(), 16)
	assert.Nil(t, snowy.EtcdSetKey("my-prefix"+randStr, 1, clientv3.NoLease))

	for i := 0; i < 3; i++ {
		randStr = strconv.FormatInt(rand.Int63(), 16)
		for j := 0; j < 128; j++ {
			count, err := snowy.EtcdIncrement("my-key-counter"+randStr, int16(0x3f), 3, 5*time.Second)
			assert.Nil(t, err)
			assert.Equal(t, int16((j+1)&0x3f), count)
		}
	}
}
