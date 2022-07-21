package main

import (
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
	client := establishConnection()
	nodeInfo := p.InitOracle(client)
	// run catch up

	// goroutine for subscribing to events

	go core.SubscribeToOracleRequestContractEvents(client, &wg, nodeInfo)
	go core.SubscribeToAggregationContractEvents(client, &wg)
	wg.Wait()
}
