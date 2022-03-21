package main

import "C"
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

func main() {
	chatServer1 := ChatServer{}
	chatServer1.Users = make(map[string]User)

	chatServer1.Users["kkalisek"] = User{
		UserName:   "kkalisek",
		ScreenName: "kookoo",
		Password:   "pw123",
	}

	chatServer1.CreateNewUser()

	fmt.Println(chatServer1.Users)
}

func PromptQuestion(question string) string {
	fmt.Printf("%s", question)
	reader := bufio.NewReader(os.Stdin)   //initialize reader
	input, err := reader.ReadString('\n') //user enters messsage
	if err != nil {
		log.Fatal("Error while reading input")
	}
	return strings.TrimSpace(input)
}

func GetUserName() User {
	UN := PromptQuestion("Please enter your Username: ")

	UserData := User{
		UserName: UN,
	}
	return UserData

}

func GetPassword() User {
	PW := PromptQuestion("Please enter your Password: ")

	UserData := User{
		Password: PW,
	}
	return UserData
}

func GetScreenName() User {
	SN := PromptQuestion("Please enter your Screen Name: ")

	UserData := User{
		ScreenName: SN,
	}
	return UserData
}

func (chat *ChatServer) VerifyUserName() string {
	var NewName string
	NewUser := GetUserName()
	if _, found := chat.Users[NewUser.UserName]; found {
		fmt.Printf("%s is taken. ", NewUser.UserName)
		NewName = chat.VerifyUserName()
	} else {
		NewName = NewUser.UserName
		fmt.Printf("%s is available, Welcome! \n", NewUser.UserName)
	}
	return NewName

}

func (chat *ChatServer) VerifyScreenName() string {
	var NewScreenName string
	NewUser := GetScreenName()
	for _, users := range chat.Users {
		if NewUser.ScreenName == users.ScreenName {
			fmt.Printf("%s is taken. ", NewUser.ScreenName)
			NewScreenName = chat.VerifyScreenName()
			return NewScreenName
		} else {
			NewScreenName = NewUser.ScreenName
			fmt.Printf("%s is available. Hello! \n", NewScreenName)
		}
	}
	return NewScreenName
}

func (chat *ChatServer) CreateNewUser() {
	var NewUser User
	NewUser.UserName = chat.VerifyUserName()
	NewUser.Password = GetPassword().Password
	NewUser.ScreenName = chat.VerifyScreenName()

	chat.Users[NewUser.UserName] = NewUser

}
