package GoEndpointManager
import "errors"

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
func (this *InMemEndpointManager) GetEndpoint(serviceID string) (host, port string, err error){

	ep := this.defaultEndpoints[serviceID]
	if ep != nil {
		host = ep.Host
		port = ep.Port
		err = nil
		return
	}
	err = ErrNotSetDefautEndpoint
	return
}

///SetDefaultEntpoint set endpoint of service by id
func (this *InMemEndpointManager)SetDefaultEntpoint(serviceID, host, port string) (err error) {
	this.defaultEndpoints[serviceID] = &Endpoint{host, port}
	return
}
