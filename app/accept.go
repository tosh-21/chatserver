package ChatServer

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type User struct {
	UserName   string
	ScreenName string
	Password   string
}
type ChatServer struct {
	UserConnections []net.Conn
	Users           map[string]User
}

func Connect() {

	dataStream, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dataStream.Close()

	chatServer := ChatServer{}
	chatServer.Users = make(map[string]User)

	for {
		conn, err := dataStream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		chatServer.UserConnections = append(chatServer.UserConnections, conn)
		go func() {
			// your very own scope
			chatServer.HandleConnection(conn, chatServer)

		}()
	}
}

func (chat *ChatServer) HandleConnection(connection net.Conn, chserv ChatServer) {
	//Ask for Username
	//Check if Username exists in chat.Users map
	//If User exists, ask for password
	//Check if Password = value of chat.Users map
	//If yes, grant access
	//If no 3 times, send emoji warning
	//If User does not exist, prompt for password and details
	//connection.Write([]byte("Hello, please enter your name: ")) //prompts user for name upon connection
	//Name, err := bufio.NewReader(connection).ReadString('\n')   //reads name
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//	return
	//}
	//
	//Name = strings.TrimSpace(Name)

	//NewUser := chat.VerifyUser(connection)
	//NewUserMap := make(map[string]User)
	//NewUserMap[NewUser.UserName] = NewUser
	//chat.Users = NewUserMap

	NewUser := chat.VerifyUser(connection)
	fmt.Println(chat.Users)
	for {
		//infinite loop for user's messages
		connection.Write([]byte("\n Enter message: "))           //prompts user for message
		msg, err := bufio.NewReader(connection).ReadString('\n') //reads message
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		msg = strings.TrimSpace(msg)

		for _, client := range chat.UserConnections {
			if msg == "end" {
				client.Write([]byte(fmt.Sprintf("\n %s has left the ChatServer", NewUser.ScreenName)))
			} else {
				client.Write([]byte(fmt.Sprintf("\n %s said: %s", NewUser.ScreenName, msg)))
			}
		}
	}
	connection.Close()
}
