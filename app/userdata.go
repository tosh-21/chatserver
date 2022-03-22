package ChatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func PromptQuestion(question string, conn net.Conn) string {
	WriteToServer(conn, question)
	reader := bufio.NewReader(conn)       //initialize reader
	input, err := reader.ReadString('\n') //user enters messsage
	if err != nil {
		log.Fatalf("Error while reading input: %s", err.Error())
	}
	return strings.TrimSpace(input)
}

func GetUserName(conn net.Conn) User {
	UN := PromptQuestion("Please enter your Username: ", conn)

	UserData := User{
		UserName: UN,
	}
	return UserData

}

func GetPassword(conn net.Conn) User {
	PW := PromptQuestion("Please enter your Password: ", conn)

	UserData := User{
		Password: PW,
	}
	return UserData
}

func GetScreenName(conn net.Conn) User {
	SN := PromptQuestion("Please enter your Screen Name: ", conn)

	UserData := User{
		ScreenName: SN,
	}
	return UserData
}

func (chat *ChatServer) VerifyUserName(conn net.Conn) string {
	var NewName string
	NewUser := GetUserName(conn)
	if _, found := chat.Users[NewUser.UserName]; found {
		WriteToServer(conn, fmt.Sprintf("%s is taken. ", NewUser.UserName))
		NewName = chat.VerifyUserName(conn)
	} else {
		NewName = NewUser.UserName
		WriteToServer(conn, fmt.Sprintf("%s is available, Welcome! \n", NewUser.UserName))
	}
	return NewName

}

func (chat *ChatServer) VerifyScreenName(conn net.Conn) string {
	var NewScreenName string
	NewUser := GetScreenName(conn)
	if chat.Users == nil {
		NewScreenName = NewUser.ScreenName
		WriteToServer(conn, fmt.Sprintf("%s is available. Hello! \n", NewScreenName))
	} else {
		for _, users := range chat.Users {
			if NewUser.ScreenName == users.ScreenName {
				WriteToServer(conn, fmt.Sprintf("%s is taken. ", NewUser.ScreenName))
				NewScreenName = chat.VerifyScreenName(conn)
				return NewScreenName
			} else {
				NewScreenName = NewUser.ScreenName
				WriteToServer(conn, fmt.Sprintf("%s is available. Hello! \n", NewScreenName))
			}
		}

	}
	return NewScreenName
}

func (chat *ChatServer) CreateNewUser(conn net.Conn) User {
	var NewUser User

	NewUser.UserName = chat.VerifyUserName(conn)
	NewUser.Password = GetPassword(conn).Password
	NewUser.ScreenName = chat.VerifyScreenName(conn)

	return NewUser

}
