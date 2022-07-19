package utils

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0xbc094BA231013BF6b1c46eC3C8BF0Ea784d38127")
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0xc232937E521e8dB741E05aBF19F9C1f537Ea1dB3")

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
