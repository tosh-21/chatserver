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

	if _, ok := chatServer1.Users[NewUser.UserName]; !ok {
		GetPassword()
		GetScreenName()
	} else {
		fmt.Println("Username already exists. Choose another one")
		main()
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
