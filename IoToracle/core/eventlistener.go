package main

import (
	ioto "IoToracle/abitogo"
	p "IoToracle/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	p.DASHEDLINE()
	p.CONNECTIONMESSAGE()

	/*contractAddress := common.HexToAddress("0x2cfAfFC93599Ae55c4aa1Bd513314b8b0A0460EF")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(15),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		requestIoTInfo, err := contractAbi.Unpack("RequestIoTInfo", vLog.Data)
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
	}*/

	contractAddress := common.HexToAddress("0x2cfAfFC93599Ae55c4aa1Bd513314b8b0A0460EF")

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	fmt.Println("we here")
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			requestIoTInfo, err := contractAbi.Unpack("RequestIoTInfo", vLog.Data)
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

		}
	}

}
