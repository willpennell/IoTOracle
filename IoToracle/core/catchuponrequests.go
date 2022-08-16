package core

import (
	"IoToracle/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

func OutstandingRequests(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	wg.Done()
	w := sync.WaitGroup{}
	w.Add(2)
	clientID := "request_oracle_catch_up"
	for _, element := range utils.Requests {

		if element.Status == 1 {

			// fetch IoT Data
			utils.SimpleFetchIoT(&element.UnPackedDataType, element.RequestId, clientID)
			// send back to smart contract
			utils.TxCommitResponse(client, nodeInfo, big.NewInt(int64(element.RequestId)), element.CommitHash)
		}

	}
}
