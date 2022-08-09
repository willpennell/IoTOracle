// Main this package houses the oraclecore.go file which starts the node
package main

import (
	"IoToracle/core"
	p "IoToracle/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"sync"
)

// establishConnection creates connection to Blockchain
func establishConnection() *ethclient.Client {
	// returns a client object
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}
	p.DASHEDLINE()        // prints message
	p.CONNECTIONMESSAGE() // prints message
	return client         // returns client object to main function
}

func main() {
	var wg sync.WaitGroup            // sync group for concurrent go routines to run
	wg.Add(2)                        // add a count to the waitGroup
	client := establishConnection()  // client blockchain object
	nodeInfo := p.InitOracle(client) // calls init functions for when first running the program

	// go routines
	go core.SubscribeToOracleRequestContractEvents(client, &wg, nodeInfo) // subscribes to events from Orc contract
	go core.SubscribeToAggregationContractEvents(client, &wg, nodeInfo)   // subscribes to events from Aggregator contract
	wg.Wait()                                                             // waits until the routines are complete, runs indefinitely
}
