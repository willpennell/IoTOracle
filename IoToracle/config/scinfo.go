package config

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0xa576ed0C02C563C04ef02b7CaCF4b4Ca8711dCcB")
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0xC1c7Ee473A1eb05b274099b865AcE0bB6E24FFf4")

func ORACLEREQUESTCONTRACTABI() abi.ABI {
	scabi, err := abi.JSON(strings.NewReader(ioto.OracleRequestContractMetaData.ABI))
	if err != nil {
		log.Fatalln(err)
	}
	return scabi
}

func AGGREGATORCONTRACTABI() abi.ABI {
	agabi, err := abi.JSON(strings.NewReader(ioto.AggregatorContractMetaData.ABI))
	if err != nil {
		log.Fatalln(err)
	}
	return agabi
}
