package utils

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var IOTORACLECONTRACTADDRESS = common.HexToAddress("0xadbC910AC565881C029c668C74012252EB142378")

func IOTORACLECONTRACTABI() abi.ABI {
	scabi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))

	if err != nil {
		log.Fatalln(err)
	}
	return scabi
}
