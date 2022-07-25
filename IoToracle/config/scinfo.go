package config

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var ORACLEREQUESTCONTRACTADDRESS = common.HexToAddress("0xB9d49Cea8d89577FFeFc966309a360566D08C4A4")
var AGGREGATIONCONTRACTADDRESS = common.HexToAddress("0x62927B0B5133238119A94644efF635C4f90d8Eca")

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
