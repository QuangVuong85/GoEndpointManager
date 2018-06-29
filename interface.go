package GoEndpointManager

//EnpointManagerIf interface of enpoint manager
type EnpointManagerIf interface{
	GetEndpoint(serviceID string) (host, port string, err error)
	SetDefaultEntpoint(serviceID, host, port string) (err error)
}