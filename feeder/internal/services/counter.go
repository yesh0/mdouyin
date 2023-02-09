package services

import (
	"common"
	"common/kitex_gen/douyin/rpc/counterservice"
	"fmt"
)

var Counter counterservice.Client

func Init() (err error) {
	Counter, err = counterservice.NewClient(string(common.CounterServiceName), common.WithEtcdResolver())
	if err != nil {
		return fmt.Errorf("unable to create counter client: %s", err)
	}

	return nil
}
