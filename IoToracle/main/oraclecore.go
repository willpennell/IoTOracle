package main

import (
	p "IoToracle/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

func initOracle(client *ethclient.Client) {

	nodeInfo := p.OracleNodeInfo{
		NodeAddress:         common.HexToAddress("0x14FFcA7245f39F6517f901A33AD9DAc43Fc8d238"),
		JoinedOracleNetwork: false,
		PrivateKey:          "0x68db5a059e251dda4767627aa7528dba91c7425c826ff185a63d76687cb14056",
	}
	p.JoinOracleNetwork(client, nodeInfo)

}

// creates connection to Blockchain
func establishConnection() *ethclient.Client {

	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}
	p.DASHEDLINE()
	p.CONNECTIONMESSAGE()
	return client
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// initialise Global Vars
	//requests := make(p.RequestsSingleton)
	//OracleRequestContractABI := p.ORACLEREQUESTCONTRACTABI()
	//AggregationContractABI := p.AGGREGATORCONTRACTABI()
	client := establishConnection()

	initOracle(client)
	// run catch up
	//c.CatchUpPreviousRequests(client, OracleRequestContractABI, requests)
	// goroutine for subscribing to events

	//go c.SubscribeToEvents(client, OracleRequestContractABI, requests, &wg)

	//go c.ServerSetup(&wg)
	wg.Wait()
}
