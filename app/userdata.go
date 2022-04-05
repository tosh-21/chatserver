package ChatServer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func PromptQuestion(question string, conn net.Conn) string {
	WriteToConnection(conn, question)
	reader := bufio.NewReader(conn)       //initialize reader
	input, err := reader.ReadString('\n') //user enters messsage
	if err != nil {
		return fmt.Sprintf("Error while reading input: %s", err.Error())
	}
	return strings.TrimSpace(input)
}

func GetUserName(conn net.Conn) string {
	return PromptQuestion("Please enter your Username: ", conn)
}

func GetPassword(conn net.Conn) string {
	return PromptQuestion("Please enter your Password: ", conn)
}

func GetScreenName(conn net.Conn) string {
	return PromptQuestion("Please enter your Screen Name: ", conn)
}

func (chat *ChatServer) VerifyPassword(password string, username string) bool {
	if password == chat.Users[username].Password {
		return true
	} else {
		return false
	}
}

func (chat *ChatServer) CheckUserName(conn net.Conn) (string, bool) {
	var NewName string
	NewUser := GetUserName(conn)
	if _, found := chat.Users[NewUser]; found {
		NewUserPassword := GetPassword(conn)
		if chat.VerifyPassword(NewUserPassword, NewUser) {
			WriteToConnection(conn, fmt.Sprintf("Welcome."))
			WriteToConnection(conn, fmt.Sprintf("Your screen name is %s \n", chat.Users[NewUser].ScreenName))
			return NewUser, true
		} else {
			WriteToConnection(conn, fmt.Sprintf("Password incorrect, try again: \n"))
			NewUserPassword := GetPassword(conn)
			if chat.VerifyPassword(NewUserPassword, NewUser) {
				WriteToConnection(conn, fmt.Sprintf("Welcome."))
				WriteToConnection(conn, fmt.Sprintf("Your screen name is %s \n", chat.Users[NewUser].ScreenName))
				return NewUser, true
			} else {
				WriteToConnection(conn, fmt.Sprintf("Password incorrect, goodbye."))
				os.Exit(3)
			}
		}
	} else {
		NewName = NewUser
		WriteToConnection(conn, fmt.Sprintf("%s is available, Welcome! \n", NewUser))
	}
	return NewName, false

}

func (chat *ChatServer) VerifyScreenName(conn net.Conn) string {
	var NewScreenName string
	NewUserSN := GetScreenName(conn)
	if chat.Users == nil {
		NewScreenName = NewUserSN
		WriteToConnection(conn, fmt.Sprintf("%s is available. Hello! \n", NewScreenName))
	} else {
		for _, users := range chat.Users {
			if NewUserSN == users.ScreenName {
				WriteToConnection(conn, fmt.Sprintf("%s is taken. ", NewUserSN))
				NewScreenName = chat.VerifyScreenName(conn)
				return NewScreenName
			} else {
				NewScreenName = NewUserSN
				WriteToConnection(conn, fmt.Sprintf("%s is available. Hello! \n", NewScreenName))
			}
		}

	}
	return NewScreenName
}

func (chat *ChatServer) VerifyUser(conn net.Conn) User {
	var NewUser User
	var UserNameExists bool
	NewUser.UserName, UserNameExists = chat.CheckUserName(conn)
	if !UserNameExists {
		NewUser.Password = GetPassword(conn)
		NewUser.ScreenName = chat.VerifyScreenName(conn)
		chat.Users[NewUser.UserName] = NewUser
	}

	return NewUser

}
