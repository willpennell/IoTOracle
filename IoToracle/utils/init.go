package utils

import (
	abi "IoToracle/abitogo"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type OracleNodeInfo struct {
	NodeAddress         common.Address
	JoinedOracleNetwork bool
	PrivateKey          string
}

func JoinOracleNetwork(client *ethclient.Client, info OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == false {
		hexKey := info.PrivateKey
		privateKeyBin, err := hexutil.Decode(hexKey)
		if err != nil {
			log.Fatal(err)
		}
		privateKey, err := crypto.ToECDSA(privateKeyBin)
		if err != nil {
			log.Fatal(nil)
		}
		nonce, err := client.PendingNonceAt(context.Background(), info.NodeAddress)
		if err != nil {
			log.Fatal(err)
		}
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		chainID := big.NewInt(5777)
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice

		instance, err := abi.NewOracleRequestContract(ORACLEREQUESTCONTRACTADDRESS, client)
		if err != nil {
			log.Fatal(nil)
		}

		tx, err := instance.JoinAsOracle(auth)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("tx sent: %s", tx.Hash().Hex())
		info.JoinedOracleNetwork = true
		return true

	}
	return false
}
