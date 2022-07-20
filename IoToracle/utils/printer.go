package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
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
	fmt.Println(req.DataType)
	fmt.Println(string(req.IotId))
}

func SUBSCRIBEDMESSAGE(oracle string) {
	fmt.Printf("Subscribed to Events from %s - Smart Contract\n", oracle)
}

func SERVERLISTENING(lis net.Listener) {
	color.Red("server listening at %v\n", lis.Addr())
}

func ORACLENODEHASJOINED() {
	color.Green("Node has joined network\n")
}

func ORACLENODEHASLEFT() {
	color.Red("NODE HAS LEFT NETWORK!\n")
}

func PRINTTXHASH(tx *types.Transaction) {
	color.HiMagenta("Cost: %v\nHash: %v\n", tx.Cost(), tx.Hash())
}

func NEWBID() {
	color.Green("-----------------NEW BID PLACED----------------")
}

func ENDBID() {
	color.Green("-----------------------------------------------")
}
