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

	CheckForMessage(userconnection) //server sends first message for name prompt

	go func() {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') //user enters name

		if err != nil {
			log.Fatal("Error while reading input")
		}

		userconnection.Write([]byte(input)) //send name to server
	}()

	go CheckForMessage(userconnection) //checks for and prints any messages from any client during name prompt

	for {

		go func() {
			reader := bufio.NewReader(os.Stdin)

			input, err := reader.ReadString('\n') //user enters messsage
			if err != nil {
				log.Fatal("Error while reading input")
			}
			userconnection.Write([]byte(input)) //message sent to server
		}()
		CheckForMessage(userconnection) //checks for and prints any messages from other users

	}

	//time.Sleep(1 * time.Second)

	//userconnection.Close()
}

func CheckForMessage(conn net.Conn) {
	//buffer initialization; read server value into buffer; print value as string
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println(string(buf))
}
