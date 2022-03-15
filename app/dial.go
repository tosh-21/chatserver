package ChatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

	ReadFromServer(userconnection)      //server sends first message for name prompt
	reader := bufio.NewReader(os.Stdin) //initialize reader

	go func() { //run infinite for loop concurrently to retrieve any messages from other users
		for {
			ReadFromServer(userconnection)
		}
	}()

	input, err := reader.ReadString('\n') //user enters name

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	WriteToServer(userconnection, input) //send name to server

	for { //main messaging block

		input, err := reader.ReadString('\n') //user enters messsage
		if err != nil {
			log.Fatal("Error while reading input")
		}
		WriteToServer(userconnection, input) //message sent to server

		if strings.TrimSpace(input) == "end" {
			break
		}

	}

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
