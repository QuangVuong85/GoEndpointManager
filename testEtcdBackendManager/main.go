package main

import (
	"log"

	"github.com/OpenStars/GoEndpointManager"
)

//Chạy 1 hoặc nhiều service với path dưới để test

func main() {
	Mngr := GoEndpointManager.NewEtcdBackendEndpointManager([]string{"10.60.1.20:2379"})
	Mngr.GetAllEndpoint("/trustkeys/socialnetwork/followuser/stringbs")
	// err := Mngr.SetDefaultEntpoint("/trustkeys/socialnetwork/followuser/stringbs", "10.110.69.96", "5656")
	// if err != nil {
	// 	log.Fatalln("err", err)
	// }
	host, port, err := Mngr.GetEndpoint("/trustkeys/socialnetwork/followuser/stringbs")
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("host", host, "port", port)
	// lstEndPoints, err := Mngr.GetAllEndpoint("/trustkeys/socialnetwork/followuser/stringbs")
	// if err != nil {
	// 	log.Fatalln("err", err)
	// }
	// for i, e := range lstEndPoints {
	// 	log.Println(i, "endpoints:", e.Host+":"+e.Port, "sid", e.ServiceID)

	// }

	// Mngr.SetDefaultEntpoint("/trustkeys/socialnetwork/followuser/stringbs", "10.110.69.96", "5656")
	// fmt.Println(Mngr.GetAllEndpoint("/trustkeys/socialnetwork"))
	// Mngr.Start()
	// for {
	// 	h, p, _ := Mngr.GetEndpoint("/trustkeys/socialnetwork/comment/stringbs")
	// 	fmt.Println("host:port: ", h, p)
	// 	fmt.Println("press enter to get endpoint again")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	text, _ := reader.ReadString('\n')
	// 	if text == "quit" {
	// 		break
	// 	}

	// }

}
