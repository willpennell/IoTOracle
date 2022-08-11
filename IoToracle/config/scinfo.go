// Package config package for configuration
package config

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

// ORACLEREQUESTCONTRACTADDRESS var that holds orc contract address
var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0x1e15685dDE8bB6FfcE2Ec18D558013dfFf68320E")

// AGGREGATIONCONTRACTADDRESS var that holds Aggregator contract address
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0x649D0d036F5b794a605Dd417be131Db810C60aB8\n")

// ORACLEREQUESTCONTRACTABI to interact with orc ABI **not used**
func ORACLEREQUESTCONTRACTABI() abi.ABI {
	// creates an abi object
	scabi, err := abi.JSON(strings.NewReader(ioto.OracleRequestContractMetaData.ABI))
	if err != nil {
		log.Fatalln(err)
	}
	return scabi // returns abi object
}

// AGGREGATORCONTRACTABI  to interact with Aggregator contract abi **not used**
func AGGREGATORCONTRACTABI() abi.ABI {
	// creates an abi object
	agabi, err := abi.JSON(strings.NewReader(ioto.AggregatorContractMetaData.ABI))
	if err != nil {
		log.Fatalln(err)
	}
	return agabi // returns abi object
}
