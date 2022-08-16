// Package utils helper files and functions
package utils

import (
	abi "IoToracle/abitogo"
	c "IoToracle/config"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

// ChainID network id used to interact with blockchain
var ChainID = big.NewInt(5777)

// PrivateKey creates private key object to sign objects
func PrivateKey(info OracleNodeInfo) *ecdsa.PrivateKey {
	hexKey := info.PrivateKey                    // gets this from OracleNodeInfo struct
	privateKeyBin, err := hexutil.Decode(hexKey) // hex decoder function into []byte
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.ToECDSA(privateKeyBin) // convert to ECDSA object for tx signatures
	if err != nil {
		log.Fatal(nil)
	}
	return privateKey // returns private key object
}

// Nonce creates a nonce for tx
func Nonce(client *ethclient.Client, info OracleNodeInfo) uint64 {
	// nonce of address for tx
	nonce, err := client.PendingNonceAt(context.Background(), info.NodeAddress)
	if err != nil {
		log.Fatal(err)
	}
	return nonce // returns nonce uint64 value
}

// GasForFuncCall creates gas price for tx
func GasForFuncCall(client *ethclient.Client) *big.Int {
	// gets suggested gas price to send with tx
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice // returns big.Int value
}

// Auth creates auth binding for transactions
func Auth(client *ethclient.Client, info OracleNodeInfo) *bind.TransactOpts {
	// creates transaction binding with private key and chainID
	// auth is returned, and we can add values to its attributes
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey(info), ChainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(Nonce(client, info))) // add nonce value to auth
	auth.Value = big.NewInt(0)                          // add value to send when calling certain functions that are payable **some are**
	auth.GasLimit = uint64(6721975)                     // gas limit
	auth.GasPrice = GasForFuncCall(client)              // gas price for transaction
	return auth                                         // returns bind.TransactOpts object
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
	auth := Auth(client, info)                               // auth object for transaction based on client and OracleNodeInfo
	stake := math.BigPow(10, 18)                             // 1 ether stake, when leaving the network it is returned.
	auth.Value = stake                                       // add to Value opt
	orcJoinAsOracle := OracleRequestContractInstance(client) // orc contract instance
	tx, err := orcJoinAsOracle.JoinAsOracle(auth)            // tx function call to JoinAsOracle, returns a tx
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx // return tx types.Transaction
}

// TxLeaveOracleNetwork run this at start of test to make sure the different parts of the contract are working
func TxLeaveOracleNetwork(client *ethclient.Client, info OracleNodeInfo) *types.Transaction {
	orcLeaveOracle := OracleRequestContractInstance(client)          // orc contract instance
	tx, err := orcLeaveOracle.LeaveOracleNetwork(Auth(client, info)) // call LeaveOracleNetwork function
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx // return tx types.Transaction
}

// TxPlaceBid creates tx to  place bid on request
func TxPlaceBid(client *ethclient.Client, info OracleNodeInfo, requestID *big.Int) (*types.Transaction, error) {
	orcPLaceBid := OracleRequestContractInstance(client)           // orc contract instance
	tx, err := orcPLaceBid.PlaceBid(Auth(client, info), requestID) // call PlaceBid function
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx, nil // return tx types.Transaction and error value
}

// ****AggregatorContract****

// AggregatorContractInstance creates instance of the aggregator contract
func AggregatorContractInstance(client *ethclient.Client) *abi.AggregatorContract {
	// Aggregator contract instance
	AggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return AggInstance // returns Aggregator contract
}

// TxCommitResponse generate tx to send result fetched to commit
func TxCommitResponse(client *ethclient.Client, info OracleNodeInfo,
	requestID *big.Int,
	commitHash [32]byte) *types.Transaction {
	aggCommitResponse := AggregatorContractInstance(client)                                // Aggregator contract instance
	tx, err := aggCommitResponse.CommitResponse(Auth(client, info), requestID, commitHash) // call commitResponse function in Aggregator contract
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx // return tx types.Transaction

}

func TxRevealVoteResponse(client *ethclient.Client, info OracleNodeInfo,
	requestID *big.Int, ioTresult []byte, secret []byte) *types.Transaction {
	aggRevealVoteResponse := AggregatorContractInstance(client)
	tx, err := aggRevealVoteResponse.RevealVoteResponse(Auth(client, info), requestID, ioTresult, secret)
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx
}

func TxRevealAverageResponse(client *ethclient.Client, info OracleNodeInfo,
	requestID *big.Int, ioTresult []byte, secret []byte) *types.Transaction {
	aggRevealAverageResponse := AggregatorContractInstance(client)
	tx, err := aggRevealAverageResponse.RevealAverageResponse(Auth(client, info),
		requestID, ioTresult, secret)
	if err != nil {

		log.Fatal(err)
	}
	TXLog(tx)
	return tx
}

// ***ToDoList***
