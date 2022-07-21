package utils

import (
	abi "IoToracle/abitogo"
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
func REQUESTLINE() {
	color.Green("---------------------------------------------------")
}

func REQUESTMESSAGE(req Request) {
	fmt.Println("Request ID: ", req.RequestId)
	fmt.Println("Data Type: ", string(req.DataType))
	fmt.Println("IoT ID: ", string(req.IotId))
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
	color.Green("-----------------NEW BID PLACED----------------\n")
}

func ENDBID() {
	color.Green("-----------------------------------------------\n")
}

func OPENFORBIDSMESSAGE(eventOpenForBids *abi.OracleRequestContractOpenForBids) {
	NEWREQUEST()
	color.Cyan("RequestID: %v\n", eventOpenForBids.Arg0)
	color.Cyan("DataType: %v\n", string(eventOpenForBids.Arg1))
	color.Cyan("Open For Bidding...\n")
	REQUESTLINE()
}

func STATUSCHANGEMESSAGE(eventStatusChange *abi.OracleRequestContractStatusChange) {
	REQUESTLINE()
	color.Red("Request Status Change:\n")
	color.Red("Status Code: %v\n", eventStatusChange.Arg0)
	color.Red("TO: %v\n", eventStatusChange.Arg1)
	REQUESTLINE()
}
