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
		fmt.Println("The connection was established to: ", connect)
		dataStream, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dataStream.Close()
		for {
			UserConn, err := dataStream.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}

			go func() {
				HandleUserConnection(UserConn)
			}()
		}
	}
}

func HandleUserConnection(userconnection net.Conn) {
	userconnection.Write([]byte("Hello, welcome."))
	//buf := make([]byte, 1024)
	//userconnection.Read(buf)
	userconnection.Close()
}
