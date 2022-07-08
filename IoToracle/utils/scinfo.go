package utils

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var IOTORACLECONTRACTADDRESS = common.HexToAddress("0x50f4f1F6EEF4C74DE927823de88EB04716d84709")

func IOTORACLECONTRACTABI() abi.ABI {
	scabi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))

	if err != nil {
		log.Fatalln(err)
	}
	return scabi
}
