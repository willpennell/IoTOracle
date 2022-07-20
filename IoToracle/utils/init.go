package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OracleNodeInfo struct {
	NodeAddress         common.Address
	JoinedOracleNetwork bool
	PrivateKey          string
}

func InitOracle(client *ethclient.Client) {

	nodeInfo := OracleNodeInfo{
		NodeAddress:         common.HexToAddress("0x14FFcA7245f39F6517f901A33AD9DAc43Fc8d238"),
		JoinedOracleNetwork: false,
		PrivateKey:          "0x68db5a059e251dda4767627aa7528dba91c7425c826ff185a63d76687cb14056",
	}
	JoinOracleNetwork(client, nodeInfo)

}

func JoinOracleNetwork(client *ethclient.Client, info OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == false {
		TxJoinAsOracle(client, info)
	}
	return false
}
