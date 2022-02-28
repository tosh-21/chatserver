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

	ReadFromServer(userconnection) //server sends first message for name prompt
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			ReadFromServer(userconnection)
		}
	}()
	input, err := reader.ReadString('\n') //user enters name

	if err != nil {
		log.Fatal("Error while reading input")
	}

	WriteToServer(userconnection, input) //send name to server

	//ReadFromServer(userconnection) //checks for and prints any messages from any client during name prompt

	for {

		input, err := reader.ReadString('\n') //user enters messsage
		if err != nil {
			log.Fatal("Error while reading input")
		}
		WriteToServer(userconnection, input) //message sent to server

		//ReadFromServer(userconnection) //checks for and prints any messages from other users

	}

	//time.Sleep(1 * time.Second)

	//userconnection.Close()
}

func ReadFromServer(conn net.Conn) {
	//buffer initialization; read server value into buffer; print value as string
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println(string(buf))
}

func WriteToServer(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
