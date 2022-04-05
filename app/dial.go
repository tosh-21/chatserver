package ChatServer

import (
	"bufio"
	"fmt"
	"io"
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

	reader := bufio.NewReader(os.Stdin) //initialize reader

	if os.Getenv("username") != "" {
		ReadFromServer(userconnection)
		name := os.Getenv("username") + "\n"
		WriteToConnection(userconnection, name)
	} else {
		PrintFromServer(userconnection)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		WriteToConnection(userconnection, input)
	}

	go func() { //run infinite for loop concurrently to retrieve any messages from other UserConnections
		for {
			PrintFromServer(userconnection)
		}
	}()

	for { //main messaging block

		input, err := reader.ReadString('\n') //user enters messsage
		if err != nil {
			log.Fatal("Error while reading input")
		}
		WriteToConnection(userconnection, input) //message sent to server

		if strings.TrimSpace(input) == "end" {
			break
		}

	}

	//userconnection.Close()
}

func PrintFromServer(conn net.Conn) {
	fmt.Println(ReadFromServer(conn))
}
func ReadFromServer(conn net.Conn) string {
	//buffer initialization; read server value into buffer; print value as string
	buf := make([]byte, 1024)
	conn.Read(buf)
	return string(buf)
}

func WriteToConnection(conn io.ReadWriter, message string) {
	conn.Write([]byte(message))
}
