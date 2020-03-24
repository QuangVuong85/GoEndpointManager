package GoEndpointManager

import (
	"fmt"
	"sync"

	etcdv3 "go.etcd.io/etcd/clientv3"
)

var etcdManagerSingleton *EtcdBackendEndpointManager
var once sync.Once

func GetEtcdBackendEndpointManagerSingleton(etcdEndpoints []string) *EtcdBackendEndpointManager {
	once.Do(func() {
		etcdManagerSingleton = NewEtcdBackendEndpointManager(etcdEndpoints)
		fmt.Println("Starting Backend Endpoint manager etcd  ", etcdEndpoints)

		if len(etcdEndpoints) == 0 {
			etcdManagerSingleton = nil
		}

		cfg := etcdv3.Config{
			Endpoints: etcdEndpoints,
		}
		aClient, err := etcdv3.New(cfg)
		if err != nil {
			etcdManagerSingleton = nil
		}

		etcdManagerSingleton.client = aClient
	})
	return etcdManagerSingleton
}
