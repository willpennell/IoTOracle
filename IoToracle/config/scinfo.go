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
var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0x4A6283Dac19DB110172500485f751c29b85E5a10")

// AGGREGATIONCONTRACTADDRESS var that holds Aggregator contract address
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0xEAfa6a4eeD22BD592E0Ba588eA64bbD6bbDE7254")

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
