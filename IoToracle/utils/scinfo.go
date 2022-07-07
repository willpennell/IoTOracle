package utils

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var IOTORACLECONTRACTADDRESS = common.HexToAddress("0x2cfAfFC93599Ae55c4aa1Bd513314b8b0A0460EF")

func IOTORACLECONTRACTABI() abi.ABI {
	scabi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))

	if err != nil {
		log.Fatalln(err)
	}
	return scabi
}
