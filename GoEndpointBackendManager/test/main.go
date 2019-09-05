package main

import (
	"log"
	"time"

	backendmanager "github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

type ServerSample struct {
	apiHost string
	apiPort string
}

func (s *ServerSample) Run() {
	log.Println("Server running")
	for {
		log.Println("Server call api to: ", s.apiHost, ":", s.apiPort)
		time.Sleep(time.Second * 2)
	}
}

func (s *ServerSample) OnChangeEndpoint(ep *backendmanager.EndPoint) {
	s.apiHost = ep.Host
	s.apiPort = ep.Port
	log.Println("Api thay  doi dia chi thanh ", ep.Host, ":", ep.Port)
}
func TestEndpoint() {
	ep := &backendmanager.EndPoint{
		Host: "127.0.0.1",
		Port: "2485",
	}
	log.Println(ep.IsGoodEndpoint())
}
func main() {
	TestEndpoint()
	// sv := &ServerSample{}
	// epManager := backendmanager.NewEndPointManager([]string{"127.0.0.1:2379"}, "/trustkeys/monitorbtc/counteruid")
	// err := epManager.LoadEndpoints()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	epManager.SetDefaultEnpoint("/trustkeys/monitorbtc/counteruid", "127.0.0.1", "21312", backendmanager.EThriftCompact)
	// 	sv.apiHost = "127.0.0.1"
	// 	sv.apiPort = "21312"
	// } else {
	// 	err, ep := epManager.GetEndPointType(backendmanager.EThriftCompact)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return
	// 	}
	// 	sv.apiHost = ep.Host
	// 	sv.apiPort = ep.Port
	// 	go epManager.EventChangeEndPoints(sv.OnChangeEndpoint)
	// 	sv.Run()
	// }
}
