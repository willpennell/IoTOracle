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
	fmt.Println(Requests[id].Status)
	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Elapsed Time: ", elapsedTime)
	fmt.Println("Time now: ", t.Unix())

	for t.Unix() < (timestamp+elapsedTime) && Requests[id].Status == 1 {
		t = time.Now()
		time.Sleep(time.Second)
		continue
	}
	if Requests[id].Status == 2 {
		return true
	}
	// call Timeout Appeal
	if Requests[id].AppealFlag == 1 && Requests[id].CommitFlag == 1 {
		TXtimeoutCommitsAppeal(client, nodeInfo, big.NewInt(int64(id)))
		// call commitAppeal
	} else if Requests[id].AppealFlag == 2 && Requests[id].RevealFlag == 1 {
		TXtimeoutRevealsAppeal(client, nodeInfo, big.NewInt(int64(id)))
	} else {
		forceCloseConnection()
		log.Println(log.Lshortfile, log.LstdFlags, "No appeal can be made")
	}
	return true
}
