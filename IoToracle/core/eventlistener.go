package main

import (
	ioto "IoToracle/abitogo"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("______________________________")
	fmt.Println("We are connected")
	contractAddress := common.HexToAddress("0x5C4CE9bE2a3d5Eed375c72dDb90Bf1e2407C34b0")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(24),
		Addresses: []common.Address{contractAddress},
	}
	fmt.Println("______________________________")
	fmt.Println("QUERY:")
	fmt.Println(query)
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("______________________________")
	fmt.Println("LOGS:")
	fmt.Println(logs)
	contractAbi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("______________________________")
	fmt.Println("CONTRACT ABI:")
	fmt.Println(contractAbi)
	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		type event interface {
			id()             uint
			requiredResult()  bool
			IoTId()           []byte
			requester()       [20]byte
			callbackAddress() [20]byte
		}
		event()
		data, err := contractAbi.UnpackIntoInterface(,"RequestIoTInfo", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("------------------------")
		fmt.Println("DATA:")
		fmt.Println(data)
		/*fmt.Println(data.requester[:])
		fmt.Println(event.callbackAddress[:])
		fmt.Println(event.IoTId[:])
		fmt.Println(event.requiredResult)*/
	}

}
