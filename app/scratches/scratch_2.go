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

	NewUser := GetUserData()

	if val, ok := chatServer1.Users[NewUser.UserName]; !ok {
		log.Println(val)
		chatServer1.Users[NewUser.UserName] = NewUser
	} else {
		fmt.Println("Username already exists")
	}
	
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

func GetUserData() User {
	UN := PromptQuestion("Please enter your Username: ")
	PW := PromptQuestion("Please enter your Password: ")
	SN := PromptQuestion("Please enter your Screen Name: ")

	UserData := User{
		UserName:   UN,
		ScreenName: SN,
		Password:   PW,
	}

	return UserData

}
