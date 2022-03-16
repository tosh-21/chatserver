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

	reader := bufio.NewReader(os.Stdin) //initialize reader

	//debuguser := os.Getenv("username")
	//log.Println(debuguser)
	if os.Getenv("username") == "willy" {
		ReadFromServer(userconnection)
		WriteToServer(userconnection, "willy \n")
	} else if os.Getenv("username") == "kookoo" {
		WriteToServer(userconnection, "kookoo") //send name to server
	} else if os.Getenv("username") == "tosh" {
		WriteToServer(userconnection, "tosh")
	} else {
		PrintFromServer(userconnection)       //server sends first message for name prompt
		input, err := reader.ReadString('\n') //user enters name

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		WriteToServer(userconnection, input)
	}

	go func() { //run infinite for loop concurrently to retrieve any messages from other users
		for {
			PrintFromServer(userconnection)
		}
	}()

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

func PrintFromServer(conn net.Conn) {
	fmt.Println(ReadFromServer(conn))
}
func ReadFromServer(conn net.Conn) string {
	//buffer initialization; read server value into buffer; print value as string
	buf := make([]byte, 1024)
	conn.Read(buf)
	return string(buf)
}

func WriteToServer(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
