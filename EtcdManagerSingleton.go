package GoEndpointManager

import "sync"

var etcdManagerSingleton *EtcdBackendEndpointManager
var once sync.Once

func GetEtcdBackendEndpointManagerSingleton(etcdEndpoints []string) *EtcdBackendEndpointManager {
	once.Do(func() {
		etcdManagerSingleton = NewEtcdBackendEndpointManager(etcdEndpoints)
	})
	return etcdManagerSingleton
}
