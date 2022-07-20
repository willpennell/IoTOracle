package utils

import (
	abi "IoToracle/abitogo"
	c "IoToracle/config"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

var ChainID = big.NewInt(5777)

// TODO generate privatekey object

func PrivateKey(info OracleNodeInfo) *ecdsa.PrivateKey {
	hexKey := info.PrivateKey
	privateKeyBin, err := hexutil.Decode(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.ToECDSA(privateKeyBin)
	if err != nil {
		log.Fatal(nil)
	}
	return privateKey
}

// TODO generate nonce

func Nonce(client *ethclient.Client, info OracleNodeInfo) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), info.NodeAddress)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

// TODO generate gas
func GasForFuncCall(client *ethclient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

// TODO generate auth

func Auth(client *ethclient.Client, info OracleNodeInfo) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey(info), ChainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(Nonce(client, info)))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = GasForFuncCall(client)
	return auth
}

// TODO generate instance
// AggregatorContractInstance
func AggregatorContractInstance(client *ethclient.Client) *abi.AggregatorContract {
	AggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return AggInstance
}

// OracleRequestContractInstance OracleRequestContract
func OracleRequestContractInstance(client *ethclient.Client) *abi.OracleRequestContract {
	ORCInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return ORCInstance
}

// TODO generate tx JoinAsOracle
func TxJoinAsOracle(client *ethclient.Client, info OracleNodeInfo) *types.Transaction {
	orcJoinAsOracle := OracleRequestContractInstance(client)
	tx, err := orcJoinAsOracle.JoinAsOracle(Auth(client, info))
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

// TODO generate tx createRequest

// TODO generate tx placeBids

// TODO generate tx deliverResponse ???

// TODO generate tx for getters

// AggregatorContract

// TODO generate tx recieveResponse
