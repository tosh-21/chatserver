package main

import (
	"os"
	"socketProgram/app"
)

func main() {
	if os.Getenv("mode") == "accept" {
		ChatServer.Connect()
	} else {
		ChatServer.Dial()
	}
}
