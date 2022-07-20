package core

import (
	c "IoToracle/config"
	u "IoToracle/utils"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sync"
)

// ***TODOList*** (not yet needed)
// TODO catch up with old request that node may have placed bid on
// TODO subscribe OracleJoined
// TODO Subscribe OracleLeft

// ***OracleRequestContract***

func SubscribeToORCEvents(client *ethclient.Client, orcABI abi.ABI, wg *sync.WaitGroup,
	requests u.RequestsSingleton, info u.OracleNodeInfo) {
	defer wg.Done()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.ORACLEREQUESTCONTRACTADDRESS},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	u.SUBSCRIBEDMESSAGE(c.ORACLEREQUESTCONTRACTADDRESS.Hex())

	for {

		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			openForBids, err := orcABI.Unpack("OpenForBids", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			request, requestId := u.ConvertOpenForBidsData(openForBids[0], openForBids[1])
			// print out new request
			requests[requestId] = request
			u.NEWREQUEST()
			u.REQUESTMESSAGE(request)
			u.ENDREQUEST()
			tx, err := u.TxPlaceBid(client, info, big.NewInt(int64(requestId)))
			if err != nil {
				log.Fatal(err)
			}
			u.NEWBID()
			u.PRINTTXHASH(tx)
			u.ENDBID()

		}
	}

}

// TODO subscribe to OpenForBids

// TODO subscribe to PlaceBid

// TODO subscribe to ReleaseRequestDetails

// TODO Subscribe to StatusChange

// ***AggregatorContract***

// TODO Subscribe to ResponseReceived

// TODO Subscribe to AggregationComplete
