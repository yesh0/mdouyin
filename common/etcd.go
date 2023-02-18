package common

import (
	"common/utils"
	"net"
	"os"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// Sets etcd as the registry and selects a random to listen to.
//
// It panics when it cannot connect to an etcd registry.
func WithEtcdOptions(name RpcServiceName) []server.Option {
	r, err := etcd.NewEtcdRegistry([]string{utils.Env.Etcd})
	if err != nil {
		klog.Fatal(err)
	}

	return []server.Option{
		WithRandomPort(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: string(name),
		}),
		server.WithRegistry(r),
	}
}

func WithRandomPort() server.Option {
	var port string
	if isRandom, ok := os.LookupEnv("ENV_MDOUYIN_RANDOM_PORT"); ok &&
		(isRandom == "1" || isRandom == "yes" || isRandom == "true") {
		port = ":0"
	} else {
		port = ":3000"
	}

	if addr, err := net.ResolveTCPAddr("tcp", port); err != nil {
		klog.Fatal(err)
		panic(err)
	} else {
		return server.WithServiceAddr(addr)
	}
}

// Creates a new new resolver option with etcd.
//
// It panics when it cannot connect to an etcd server.
func WithEtcdResolver() client.Option {
	r, err := etcd.NewEtcdResolver([]string{utils.Env.Etcd})
	if err != nil {
		klog.Fatalf("unable to create etcd resolver: %s", err.Error())
	}
	return client.WithResolver(r)
}
