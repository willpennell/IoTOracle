package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("We are connected")
	_ = client
	/*contractAddress := common.HexToAddress(" 0x7d79D4256C9Ed0B416250b3d5B6CD5365bc93188")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}*/
}
