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
var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0xa576ed0C02C563C04ef02b7CaCF4b4Ca8711dCcB")

// AGGREGATIONCONTRACTADDRESS var that holds Aggregator contract address
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0xC1c7Ee473A1eb05b274099b865AcE0bB6E24FFf4")

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
