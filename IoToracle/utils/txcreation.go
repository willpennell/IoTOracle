package utils

import (
	abi "IoToracle/abitogo"
	c "IoToracle/config"
	"crypto/ecdsa"
	"fmt"
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

// PrivateKey creates private key object to sign objects
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

// Nonce creates a nonce for tx
func Nonce(client *ethclient.Client, info OracleNodeInfo) uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), info.NodeAddress)
	if err != nil {
		log.Fatal(err)
	}
	return nonce
}

// GasForFuncCall creates gas price for tx
func GasForFuncCall(client *ethclient.Client) *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

// Auth creates auth binding for transactions
func Auth(client *ethclient.Client, info OracleNodeInfo) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey(info), ChainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(Nonce(client, info)))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = GasForFuncCall(client)
	return auth
}

// ****OracleRequestContract****

// OracleRequestContractInstance creates instance of OracleRequestContract
func OracleRequestContractInstance(client *ethclient.Client) *abi.OracleRequestContract {
	ORCInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return ORCInstance
}

// TxJoinAsOracle creates tx to join oracle network
func TxJoinAsOracle(client *ethclient.Client, info OracleNodeInfo) *types.Transaction {
	orcJoinAsOracle := OracleRequestContractInstance(client)
	tx, err := orcJoinAsOracle.JoinAsOracle(Auth(client, info))
	if err != nil {
		log.Fatal(err)
	}
	PRINTTXHASH(tx)
	return tx
}

// TxLeaveOracleNetwork run this at start of test to make sure the different parts of the contract are working
func TxLeaveOracleNetwork(client *ethclient.Client, info OracleNodeInfo) *types.Transaction {
	orcLeaveOracle := OracleRequestContractInstance(client)
	tx, err := orcLeaveOracle.LeaveOracleNetwork(Auth(client, info))
	if err != nil {
		log.Fatal(err)
	}
	PRINTTXHASH(tx)
	return tx
}

// TxPlaceBid creates tx to  place bid on request
func TxPlaceBid(client *ethclient.Client, info OracleNodeInfo, requestID *big.Int) (*types.Transaction, error) {
	orcPLaceBid := OracleRequestContractInstance(client)
	tx, err := orcPLaceBid.PlaceBid(Auth(client, info), requestID)
	if err != nil {
		log.Fatal(err)
	}
	//PRINTTXHASH(tx)
	return tx, nil
}

// ***ToDoList***
// TODO generate tx createRequest
// TODO generate tx deliverResponse
// TODO generate tx for getters

func TxPhash(client *ethclient.Client, info OracleNodeInfo, requestID *big.Int) {
	orcPLaceBid := OracleRequestContractInstance(client)
	call := bind.CallOpts{Context: context.Background()}
	pHash, err := orcPLaceBid.GetPHash(&call, requestID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Solidity HASH: ", pHash)

}

// Do not need at the moment

// ****AggregatorContract****

// AggregatorContractInstance creates instance of the aggregator contract
func AggregatorContractInstance(client *ethclient.Client) *abi.AggregatorContract {
	AggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return AggInstance
}

// TxReceiveResponse generate tx to send result fetched to receiveResponse
func TxReceiveResponse(client *ethclient.Client, info OracleNodeInfo, requestID *big.Int, fetchedResult []byte) *types.Transaction {
	aggReceiveResponse := AggregatorContractInstance(client)
	tx, err := aggReceiveResponse.ReceiveResponse(Auth(client, info), requestID, fetchedResult)
	if err != nil {
		log.Fatal(err)
	}
	//PRINTTXHASH(tx)
	return tx
}

// ***ToDoList***
