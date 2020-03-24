package GoEndpointManager

import "sync"

var etcdManagerSingleton *EtcdBackendEndpointManager
var once sync.Once

func GetEtcdBackendEndpointManagerSingleton(etcdEndpoints []string) *EtcdBackendEndpointManager {
	once.Do(func() {
		etcdManagerSingleton = NewEtcdBackendEndpointManager(etcdEndpoints)
fmt.Println("Starting Backend Endpoint manager etcd  ", o.EtcdEndpoints)


	if len(o.EtcdEndpoints) == 0 {
		etcdManagerSingleton = nil
	}

	cfg := etcdv3.Config{
		Endpoints: o.EtcdEndpoints,
	}
	aClient, err := etcdv3.New(cfg)
	if err != nil {
		etcdManagerSingleton = nil
	}
	o.client = aClient
	})
	return etcdManagerSingleton
}
