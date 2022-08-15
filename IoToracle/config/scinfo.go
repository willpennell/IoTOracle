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
var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0x3ce2Da376608573C842b041f3E76795211A4A987")

// AGGREGATIONCONTRACTADDRESS var that holds Aggregator contract address
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0x6E5968F43c06550c7880A115fd4e5a38c30FeF4F")

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
