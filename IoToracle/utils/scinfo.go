package utils

import (
	ioto "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
)

var IOTORACLECONTRACTADDRESS = common.HexToAddress("0xd2C08435e6107E60E057BbBFa12D53b9F98692fC")

func IOTORACLECONTRACTABI() abi.ABI {
	scabi, err := abi.JSON(strings.NewReader(ioto.IoTOracleContractMetaData.ABI))

	if err != nil {
		log.Fatalln(err)
	}
	return scabi
}
