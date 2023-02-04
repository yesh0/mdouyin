package services

import (
	"common/kitex_gen/douyin/rpc/feedservice"
	"fmt"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Feed feedservice.Client

func Init() error {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		return fmt.Errorf("unable to create etcd resolver: %s", err.Error())
	}

	Feed, err = feedservice.NewClient("feed", client.WithResolver(r))
	if err != nil {
		return fmt.Errorf("unable to create feed client: %s", err)
	}

	return nil
}
