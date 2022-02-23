package ChatServer

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
		fmt.Println("The connection was established to: ", address)

		defer connect.Close()

		HandleUserConnection(connect)

	}
}

func HandleUserConnection(userconnection net.Conn) {
	userconnection.Write([]byte("Hello, welcome."))
	buf := make([]byte, 1024)
	userconnection.Read(buf)
	userconnection.Close()
}
