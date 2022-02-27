package ChatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type chat struct {
	users []net.Conn
}

func Connect() {

	dataStream, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dataStream.Close()

	AllConn := chat{
		users: []net.Conn{},
	}
	for {
		conn, err := dataStream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		AllConn.users = append(AllConn.users, conn)
		go func() {
			// your very own scope
			HandleConnection(conn, &AllConn)

		}()
	}
}

func HandleConnection(connection net.Conn, chat *chat) {
	connection.Write([]byte("Hello, please enter your name: "))
	Name, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	Name = strings.TrimSpace(Name)
	for {
		connection.Write([]byte("Enter message: "))
		msg, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		msg = strings.TrimSpace(msg)
		log.Println(Name, msg)

		for _, client := range chat.users { //does not work as goroutine
			client.Write([]byte(fmt.Sprintf("%s said: %s \n", Name, msg)))
		}
	}
	connection.Close()
}
