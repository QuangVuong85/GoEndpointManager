package main

import (
	"fmt"

	"bufio"

	"os"

	"github.com/OpenStars/GoEndpointManager"
)

//Chạy 1 hoặc nhiều service với path dưới để test

func main() {
	Mngr := GoEndpointManager.NewEtcdBackendEndpointManager([]string{"127.0.0.1:2379"})
	Mngr.SetDefaultEntpoint("/services/bigset/stringbigset", "127.0.0.1", "5656")
	Mngr.Start()
	for {
		h, p, _ := Mngr.GetEndpoint("/services/bigset/stringbigset")
		fmt.Println("host:port: ", h, p)
		fmt.Println("press enter to get endpoint again")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "quit" {
			break
		}

	}

}
