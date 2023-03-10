package common

import (
	"common/utils"
	"net"

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

// From https://stackoverflow.com/questions/23558425/
func GetOutboundIP() net.IP {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		klog.Fatal(err)
	}
	for _, addr := range addresses {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			return ip.IP
		}
	}
	klog.Fatal("no outbound ip found")
	return nil
}

func WithRandomPort() server.Option {
	if addr, err := net.ResolveTCPAddr("tcp", GetOutboundIP().String()+":0"); err != nil {
		klog.Fatal(err)
		panic(err)
	} else {
		klog.Infof("listening to %v", addr)
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
