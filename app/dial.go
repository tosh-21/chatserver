package SocketConnection

import (
	"fmt"
	"net"
)

func Dial() {

	address := "localhost:8080"
	connect, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error occurred: ", err)
	} else {
		fmt.Println("The connection was established to: ", connect)
	}
}
