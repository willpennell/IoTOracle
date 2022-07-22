package config

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0x529D508A34cEb8b4dCfa0E5695118877d2e25C11")
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0x92Bec34538867819A5757b8Bf1A2AD1308a9c77f")

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
