package config

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0xA7e1C200E343213bF7A702A4d72032D56c2eA194")
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0x68caDf1B0Bf413cEc74c3cf1b9dc0a2Cf927E2Ad")

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
