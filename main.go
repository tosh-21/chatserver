package main

import (
	"os"
	SocketConnection "socketProgram/app"
)

func main() {
	if os.Getenv("mode") == "accept" {
		SocketConnection.Connect()
	} else {
		SocketConnection.Dial()
	}
}
