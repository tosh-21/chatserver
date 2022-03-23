package ChatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

func (chat *ChatServer) VerifyPassword(conn net.Conn, NewUser User) bool {
	if NewUser.Password == chat.Users[NewUser.UserName].Password {
		return true
	} else {
		return false
	}
}

func (chat *ChatServer) CheckUserName(conn net.Conn) (string, bool) {
	var NewName string
	NewUser := GetUserName(conn)
	if _, found := chat.Users[NewUser.UserName]; found {
		NewUser.Password = GetPassword(conn).Password
		if chat.VerifyPassword(conn, NewUser) {
			WriteToServer(conn, fmt.Sprintf("Password found, welcome."))
			WriteToServer(conn, fmt.Sprintf("Your screen name is %s \n", chat.Users[NewUser.UserName].ScreenName))
			return NewUser.UserName, true
		} else {
			WriteToServer(conn, fmt.Sprintf("Password incorrect, try again: \n"))
			NewUser.Password = GetPassword(conn).Password
			if chat.VerifyPassword(conn, NewUser) {
				WriteToServer(conn, fmt.Sprintf("Password found, welcome."))
				WriteToServer(conn, fmt.Sprintf("Your screen name is %s \n", chat.Users[NewUser.UserName].ScreenName))
				return NewUser.UserName, true
			} else {
				WriteToServer(conn, fmt.Sprintf("Password incorrect, goodbye."))
				os.Exit(3)
			}
		}
	} else {
		NewName = NewUser.UserName
		WriteToServer(conn, fmt.Sprintf("%s is available, Welcome! \n", NewUser.UserName))
	}
	return NewName, false

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

func (chat *ChatServer) VerifyUser(conn net.Conn) User {
	var NewUser User
	var Verification bool
	UserMap := make(map[string]User)
	NewUser.UserName, Verification = chat.CheckUserName(conn)
	if Verification == false {
		NewUser.Password = GetPassword(conn).Password
		NewUser.ScreenName = chat.VerifyScreenName(conn)
		UserMap[NewUser.UserName] = NewUser
		chat.Users = UserMap
	}

	return NewUser

}
