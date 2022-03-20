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

	NewUser := GetUserName()
	chatServer1.VerifyUserName(NewUser)
	NewUser.Password = GetPassword().Password
	NewUser.ScreenName = chatServer1.VerifyScreenName(GetScreenName().ScreenName)

	chatServer1.Users[NewUser.UserName] = NewUser

	fmt.Println(chatServer1)
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

func (chat *ChatServer) VerifyUserName(name User) User {
	var NewName User
	for _, users := range chat.Users {
		if name.UserName == users.UserName {
			fmt.Println("Username is taken, please enter another: ")
			NewName := GetUserName()
			chat.VerifyUserName(NewName)
			return NewName
		} else {
			NewName := name
			fmt.Printf("%s is available. Welcome! \n", NewName.UserName)
		}
	}
	return NewName
}

func (chat *ChatServer) VerifyScreenName(ScreenName string) string {
	var NewScreenName string
	for _, users := range chat.Users {
		if ScreenName == users.ScreenName {
			fmt.Printf("Screen Name is taken. ")
			NewScreenName := GetScreenName()
			chat.VerifyScreenName(NewScreenName.ScreenName)
			return NewScreenName.ScreenName
		} else {
			NewScreenName := ScreenName
			fmt.Printf("%s is available. Hello! \n", NewScreenName)
		}
	}
	return NewScreenName
}
