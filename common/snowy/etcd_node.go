package snowy

import (
	"context"
	"fmt"
	"time"

	"github.com/godruoyi/go-snowflake"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	node_counter_key = "snowflake_node_counter"
	node_registry    = "/snowy/nodes/"
)

var etcd *clientv3.Client
var lease clientv3.LeaseID

func Init(endpoints ...string) error {
	if etcd != nil {
		return fmt.Errorf("etcd client already initialized")
	}

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		return err
	}
	etcd = etcdClient

	nodeId, err := etcdGetNodeId()

	snowflake.SetMachineID(uint16(nodeId))

	return err
}

func etcdGetNodeId() (int16, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := etcd.Grant(ctx, 60)
	if err != nil {
		return 0, err
	}

	lease = resp.ID
	for i := 0; i < 16 && ctx.Err() == nil; i++ {
		nodeId, err := EtcdIncrement(
			node_counter_key,
			(1<<int16(snowflake.MachineIDLength))-1,
			3,
			3*time.Second,
		)
		if err != nil {
			return 0, err
		}

		if err := EtcdSetKey(node_registry, int(nodeId), lease); err != nil {
			continue
		} else {
			responses, err := etcd.KeepAlive(context.Background(), lease)
			if err != nil {
				return 0, err
			}
			go func() {
				for _, ok := <-responses; ok; _, ok = <-responses {
				}
				panic("etcd is down")
			}()
			return nodeId, nil
		}
	}
	return 0, fmt.Errorf("unable to find a node id")
}

func EtcdSetKey(prefix string, id int, lease clientv3.LeaseID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	key := fmt.Sprintf("%s-%d", prefix, id)
	resp, err := etcd.Txn(ctx).
		If(clientv3.Compare(clientv3.LeaseValue(key), "=", clientv3.NoLease)).
		Then(clientv3.OpPut(key, "occupied", clientv3.WithLease(lease))).
		Commit()

	if err != nil {
		return err
	}
	if resp.Succeeded {
		return nil
	} else {
		return fmt.Errorf("key occupied")
	}
}

func initializeValue(ctx context.Context, key string, value []byte) error {
	_, err := etcd.Txn(ctx).
		If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(clientv3.OpPut(key, string(value))).Commit()
	if err != nil {
		return err
	}
	return nil
}

func EtcdIncrement(key string, mask int16, attempts int, timeout time.Duration) (int16, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var prev []byte
	next := []byte{0, 0}
	var current int16

	for i := 0; i < attempts; i++ {
		get, err := etcd.Get(ctx, key)
		if err != nil {
			return 0, err
		}
		if get.Count == 0 {
			prev = []byte{0, 0}
			initializeValue(ctx, key, prev)
		} else {
			prev = get.Kvs[0].Value
		}
		current = (int16(prev[0]) | int16(prev[1])<<8)
		current = (current + 1) & mask
		next[0] = byte(current & 0xff)
		next[1] = byte(current >> 8)

		txn, err := etcd.Txn(ctx).
			If(clientv3.Compare(clientv3.Value(key), "=", string(prev))).
			Then(clientv3.OpPut(key, string(next))).Commit()
		if err != nil {
			return 0, err
		}
		if txn.Succeeded {
			return current, nil
		}
	}
	return 0, fmt.Errorf("unable to increment value")
}
