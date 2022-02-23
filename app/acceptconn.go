package SocketConnection

import (
	"bufio"
	"fmt"
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
	buf := make([]byte, 1024)
	connection.Read(buf)
	Name, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	Name = strings.TrimSpace(Name)
	connection.Write([]byte("Please enter your message: "))
	for {
		msg, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
		}
		msg = strings.TrimSpace(msg)
		//connection.Write([]byte(fmt.Sprintf("%s said: %s \n", Name, msg)))
		for _, client := range chat.users {
			client.Write([]byte(fmt.Sprintf("%s said: %s ", Name, msg)))
		}
	}
	connection.Close()
}
