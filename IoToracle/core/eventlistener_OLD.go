package core

import (
	p "IoToracle/utils"
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

func CatchUpPreviousRequests(client *ethclient.Client, contractABI abi.ABI, reqMap p.RequestsSingleton) {
	//contractAddress := common.HexToAddress("0x2cfAfFC93599Ae55c4aa1Bd513314b8b0A0460EF")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(319),
		Addresses: []common.Address{p.IOTORACLECONTRACTADDRESS},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		requestIoTInfo, err := contractABI.Unpack("RequestIoTInfo", vLog.Data)

		if err != nil {
			log.Fatal(err)
		}

		p.DASHEDLINE()
		p.DATAMESSAGE()

		req := p.ConvertRequestTypes(
			requestIoTInfo[0],
			requestIoTInfo[1],
			requestIoTInfo[2],
			requestIoTInfo[3])

		p.REQUESTMESSAGE(req)
		reqMap[req.RequestId] = req
	}
}

func subscribeToEvents(client *ethclient.Client, contractABI abi.ABI, reqMap p.RequestsSingleton, wg *sync.WaitGroup) {
	defer wg.Done()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{p.IOTORACLECONTRACTADDRESS},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	p.SUBSCRIBEDMESSAGE(p.IOTORACLECONTRACTADDRESS.Hex())

	for {

		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			requestIoTInfo, err := contractABI.Unpack("RequestIoTInfo", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			p.NEWREQUEST()
			p.DATAMESSAGE()

			req := p.ConvertRequestTypes(
				requestIoTInfo[0],
				requestIoTInfo[1],
				requestIoTInfo[2],
				requestIoTInfo[3])

			p.REQUESTMESSAGE(req)
			p.ENDREQUEST()
			reqMap[req.RequestId] = req
		}
	}

}
