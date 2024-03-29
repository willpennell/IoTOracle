// Package utils helper functions
package utils

import (
	abi "IoToracle/abitogo"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fatih/color"
	"log"
	"net"
)

// printer helper functions to make output to terminal look neat

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

func SERVERLISTENING(lis net.Listener) {
	color.Red("server listening at %v\n", lis.Addr())
}

func ORACLENODEHASJOINED() {
	color.Green("Node has joined network\n")
}

func ORACLENODEHASLEFT() {
	color.Red("NODE HAS LEFT NETWORK!\n")
}

func PRINTTXHASH(tx *types.Transaction) string {
	return fmt.Sprintf("Cost: %v\nHash: %v\n", tx.Cost(), tx.Hash())
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

func BIDPLACEDMEASSGE(eventBidPlaced *abi.OracleRequestContractBidPlaced) {
	color.Green("--------------------BID PLACED-----------------------")
	fmt.Println("Bid Placed:")
	fmt.Println("Oracle Node Address: ", eventBidPlaced.Arg0)
	REQUESTLINE()
}

func STATUSCHANGEMESSAGE(eventStatusChange *abi.OracleRequestContractStatusChange) {
	color.Green("--------------------NEW STATUS---------------------")
	color.Red("Request Status Change:\n")
	color.Red("Request ID: %v\n", eventStatusChange.Arg0)
	color.Red("Status Code: %v\n", eventStatusChange.Arg1)
	color.Red("TO: %v\n", eventStatusChange.Arg2)
	REQUESTLINE()
}

func RELEASEREQUESTDETAILS(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails) {
	color.Green("--------------------REQUEST DETAILS-----------------")
	fmt.Println("Release Request Details:")
	fmt.Println("RequestID: ", eventReleaseRequestDetails.Arg0)
	fmt.Println("IoTID: ", string(eventReleaseRequestDetails.Arg1))
	fmt.Println("Data Type: ", string(eventReleaseRequestDetails.Arg2))
	REQUESTLINE()
}

func COMMITSPLACEDMESSAGE(eventCommitsPlaced *abi.AggregatorContractCommitsPlaced) {
	color.Green("--------------------COMMITS PLACED-----------------")
	fmt.Println("RequestID: ", eventCommitsPlaced.Arg0)
	fmt.Println(eventCommitsPlaced.Arg1)
	REQUESTLINE()
}

func REVEALSPLACEDMESSAGE(eventRevealsPlaced *abi.AggregatorContractRevealsPlaced) {
	color.Green("--------------------REVEALS PLACED-----------------")
	fmt.Println("RequestID: ", eventRevealsPlaced.Arg0)
	fmt.Println(eventRevealsPlaced.Arg1)
	REQUESTLINE()
}

func AGGREGATIONCOMPLETE(eventAggregationCompleted *abi.AggregatorContractAggregationCompleted) {
	color.Green("--------------------AGGREGATION--------------------")
	fmt.Println("Aggregation Complete:")
	fmt.Println("RequestID: ", eventAggregationCompleted.Arg0)
	fmt.Println("Aggregation Message: ", eventAggregationCompleted.Arg1)
	REQUESTLINE()
}

func ORACLEPAID(eventOraclePaid *abi.OracleRequestContractOraclePaid) {
	color.Green("--------------------ORACLE PAID--------------------")
	color.HiYellow("Oracle Address: %v\n", eventOraclePaid.Arg0)
	color.HiYellow("Wei Earned: %v\n", eventOraclePaid.Arg1)
	color.HiYellow("Message: %v\n", eventOraclePaid.Arg2)
	REQUESTLINE()

}

func MQTTBROKERSUB(topic string) {
	color.Green("-----------------SUBSCRIBED TO MQTT TOPIC------------------")
	it := color.New(color.FgHiBlue).Add(color.Italic)
	_, err := it.Printf("Subscribed to topic: %s\n", topic)
	if err != nil {
		log.Println(err)
	}
	REQUESTLINE()

}

func MQTTMESSAGE(msg string) {
	MQTTHEADER()
	it := color.New(color.FgHiBlue).Add(color.Italic)
	_, err := it.Printf("%s\n", msg)
	if err != nil {
		log.Println(err)
	}
	REQUESTLINE()

}

func MQTTHEADER() {
	color.Green("-----------------MQTT MESSAGE------------------")
}

func HASHLOGMESSAGE(eventLogHashes *abi.AggregatorContractLogHashes) {
	color.Green("-----------------HASH MESSAGE------------------")
	it := color.New(color.FgHiRed).Add(color.Italic)

	_, err := it.Printf("commit hash: 0x%s\n", common.Bytes2Hex(eventLogHashes.Arg0[:]))
	if err != nil {
		log.Println(err)
	}
	_, err = it.Printf("smart contract hash: 0x%s\n", common.Bytes2Hex(eventLogHashes.Arg1[:]))
	if err != nil {
		log.Println(err)
	}
	REQUESTLINE()
}
