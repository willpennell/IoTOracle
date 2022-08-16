package core

import (
	"IoToracle/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
	clientID := "request_oracle_catch_up"
	for _, element := range utils.Requests {
		if element.Status == 1 {

			// TODO fetch IoT Data
			utils.SimpleFetchIoT(&element.UnPackedDataType, element.RequestId, clientID)
			fmt.Println(element.CommitHash)
			// TODO send back to smart contract
			utils.TxCommitResponse(client, nodeInfo, big.NewInt(int64(element.RequestId)), element.CommitHash)
			// TODO send response
			//EventCommitsPlaced(client, &w, nodeInfo)
			// TODO Aggregation complete
		}

	}
}

// TODO complete the fetch requests
