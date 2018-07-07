package GoEndpointManager
import (
	"errors"
	"sync"
)

var (
	//ErrNotSetDefautEndpoint not set default enpoints
	ErrNotSetDefautEndpoint = errors.New("Not set default endpoint")
)

type  Endpoint struct{
	Host string
	Port string
}
//EnpointManagerIf
type InMemEndpointManager struct {
	defaultEndpoints map[string]*Endpoint 
}

//GetEndpoint get endpoint by service id
func (o *InMemEndpointManager) GetEndpoint(serviceID string) (host, port string, err error){

	ep := o.defaultEndpoints[serviceID]
	if ep != nil {
		host = ep.Host
		port = ep.Port
		err = nil
		return
	}
	err = ErrNotSetDefautEndpoint
	return
}

//SetDefaultEntpoint set endpoint of service by id
func (o *InMemEndpointManager)SetDefaultEntpoint(serviceID, host, port string) (err error) {
	o.defaultEndpoints[serviceID] = &Endpoint{host, port}
	return
}


