package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func StampTimer(id uint64, timestamp int64, elapsedTime int64, client *ethclient.Client, nodeInfo OracleNodeInfo) bool {
	t := time.Now()
	// prints timer messages to terminal
	fmt.Println(Requests[id].Status)
	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Elapsed Time: ", elapsedTime)
	fmt.Println("Time now: ", t.Unix())
	// timer for requests, we need status to be 1 and current time less than timestamp and elapsed time.
	for t.Unix() < (timestamp+elapsedTime) && Requests[id].Status == 1 {
		t = time.Now()          // current time
		time.Sleep(time.Second) // sleep
		continue                // do nothing during timer
	}
	if Requests[id].Status == 2 {
		return true // if timer finishes and request is 2 we know it is complete of cancelled
	}
	// call Timeout Appeal
	if Requests[id].AppealFlag == 1 && Requests[id].CommitFlag == 1 {
		TXtimeoutCommitsAppeal(client, nodeInfo, big.NewInt(int64(id)))
		// call commitAppeal
	} else if Requests[id].AppealFlag == 2 && Requests[id].RevealFlag == 1 {
		// timer stopped due to reveal not commit
		TXtimeoutRevealsAppeal(client, nodeInfo, big.NewInt(int64(id)))
	} else {
		// close connection and no appeals can be made
		forceCloseConnection()
		log.Println(log.Lshortfile, log.LstdFlags, "No appeal can be made")
	}
	return true
}
