package core

import (
	"IoToracle/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
)

// TODO loop through the outstanding requests that are currently set to 1

// TODO flag for no request details

// TODO event listener for

func OutstandingRequests(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	wg.Done()
	w := sync.WaitGroup{}
	w.Add(2)
	fmt.Println(len(utils.Requests))
	for _, element := range utils.Requests {
		if element.Status == 1 {

			// TODO fetch IoT Data
			fmt.Println("waiting in the outstanding requests")
			utils.SimpleFetchIoT(&element.UnPackedDataType, element.RequestId)
			//fmt.Println("data fetched...")
			//// TODO send back to smart contract
			//fmt.Println("Sending to the commitResponse: ", element.RequestId, element.CommitHash)
			//utils.TxCommitResponse(client, nodeInfo, big.NewInt(int64(element.RequestId)), element.CommitHash)
			//// TODO send response
			//EventCommitsPlaced(client, &w, nodeInfo)
			//// TODO Aggregation complete
			//EventRevealsPlaced(client, &w, nodeInfo)
			//EventAggregationComplete(client, &w)
		}

	}
}

// TODO complete the fetch requests
