package ChatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Dial() {

	address := "localhost:8080"
	time.Sleep(500 * time.Millisecond)
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

	buf := make([]byte, 1024) //Initial message from server prompting for name
	userconnection.Read(buf)
	fmt.Printf(string(buf)) //Server prints name
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error while reading input")
	}

	userconnection.Write([]byte(input)) //Send name to server

	buf = make([]byte, 1024)
	userconnection.Read(buf)
	fmt.Printf(string(buf)) //Server prints input request for first message

	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error while reading input")
		}
		userconnection.Write([]byte(input))
		buf = make([]byte, 1024)
		userconnection.Read(buf)
		fmt.Println(string(buf))

	}

	//time.Sleep(1 * time.Second)

	//userconnection.Close()
}
