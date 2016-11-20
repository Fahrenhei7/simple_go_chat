package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fahrenhei7/simple_chat/lib"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		fmt.Println("is host")
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		fmt.Println("is guest")
		lib.RunGuest(connIP)
	}
}
