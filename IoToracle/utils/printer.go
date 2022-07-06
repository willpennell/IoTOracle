package utils

import (
	"fmt"
)

func DASHEDLINE() {
	fmt.Println("------------------------------------------------------------------")
}

func CONNECTIONMESSAGE() {
	fmt.Println("Welcome, connection established")
}

func DATAMESSAGE() {
	fmt.Println("DATA:")
}

func REQUESTMESSAGE(req Request) {
	fmt.Println(req.RequestId)
	fmt.Println(req.RequireResult)
	fmt.Println(string(req.IotId))
	fmt.Println(req.Requester)
}
