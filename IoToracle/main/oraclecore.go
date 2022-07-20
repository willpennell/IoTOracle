package main

import (
	c "IoToracle/config"
	"IoToracle/core"
	p "IoToracle/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

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
	requests := p.Reqs
	//OracleRequestContractABI := c.ORACLEREQUESTCONTRACTABI()
	//AggregationContractABI := c.AGGREGATORCONTRACTABI()
	client := establishConnection()

	nodeInfo := p.InitOracle(client)
	// run catch up
	//c.CatchUpPreviousRequests(client, OracleRequestContractABI, requests)
	// goroutine for subscribing to events

	go core.SubscribeToORCEvents(client, c.ORACLEREQUESTCONTRACTABI(), &wg,
		requests, nodeInfo)

	//go c.ServerSetup(&wg)
	//wg.Wait()
}
