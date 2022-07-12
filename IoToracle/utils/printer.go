package utils

import (
	"fmt"
	"github.com/fatih/color"
	"net"
)

func DASHEDLINE() {
	fmt.Println("------------------------------------------------------------------")
}

func CONNECTIONMESSAGE() {
	fmt.Println("Welcome, connection established")
}

func DATAMESSAGE() {
	color.HiMagenta("DATA:")
}

func NEWREQUEST() {
	color.Green("--------------------NEW REQUEST--------------------")
}
func ENDREQUEST() {
	color.Green("---------------------------------------------------")
}

func REQUESTMESSAGE(req Request) {
	fmt.Println(req.RequestId)
	fmt.Println(req.RequireResult)
	fmt.Println(string(req.IotId))
	fmt.Println(req.Requester)
}

func SUBSCRIBEDMESSAGE(oracle string) {
	fmt.Printf("Subscribed to Events from %s - Smart Contract\n", oracle)
}

func SERVERLISTENING(lis net.Listener) {
	color.Red("server listening at %v", lis.Addr())
}
