package main

import (
	"fmt"

	"github.com/OpenStars/GoEndpointManager"
)

func main() {
	Mngr := GoEndpointManager.NewEtcdBackendEndpointManager([]string{"127.0.0.1:2379"})
	Mngr.SetDefaultEntpoint("/services/bigset/stringbigset", "127.0.0.1", "5656")
	Mngr.Start()
	h, p, _ := Mngr.GetEndpoint("/services/bigset/stringbigset")
	fmt.Println("host:port: ", h, p)
}
