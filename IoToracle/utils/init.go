// @Author william Pennell

// Package utils - package for all the helper files and functions
package utils

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"sync"
)

// OracleNodeInfo struct that holds information on oracle info based on .env file
type OracleNodeInfo struct {
	NodeAddress         common.Address // address associated with node
	JoinedOracleNetwork bool           // flag to state nodes current status in the network
	PrivateKey          string         // private key to sign TX
}

var lock sync.Mutex // global lock for writing flag value to .tmp file
// InitOracle creates hardcoded oracle details inorder to join the network
func InitOracle(client *ethclient.Client) OracleNodeInfo {
	// nodeinfo object to hold info from .env
	nodeInfo := OracleNodeInfo{
		NodeAddress:         common.HexToAddress(getNodeAddress()),
		JoinedOracleNetwork: getJoinedOracleNetwork(),
		PrivateKey:          getPrivateKey(),
	}
	LeaveOracleNetwork(client, &nodeInfo) // restarts each time for clean session, leaves network
	JoinOracleNetwork(client, &nodeInfo)  //  then rejoins
	return nodeInfo                       // returns OracleNodeInfo object
}

// JoinOracleNetwork function to call join oracle function in smart contract
func JoinOracleNetwork(client *ethclient.Client, info *OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == false { // makes sure node is currently not connected to network
		TxJoinAsOracle(client, *info)                       // tx function to call orc contract joinAsOracle() function
		ORACLENODEHASJOINED()                               // prints message
		info.JoinedOracleNetwork = true                     // change flag so we know oracle is now in network
		setHasOracleJoinedNetwork(info.JoinedOracleNetwork) // saves value to .tmp file
		// as .env cannot change value after program terminated
		return true
	}
	return false // returns false if already in network
}

// LeaveOracleNetwork Leave the oracle network
func LeaveOracleNetwork(client *ethclient.Client, info *OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == true { // make sure oracle is already in network
		TxLeaveOracleNetwork(client, *info)                 // tx call to orc contract function leaveOracleNetwork()
		ORACLENODEHASLEFT()                                 // prints message
		info.JoinedOracleNetwork = false                    // change flag, so we know oracle node is no longer in network
		setHasOracleJoinedNetwork(info.JoinedOracleNetwork) // save flag state to .tmp file
		return true
	}
	return false // returns false if not in network
}

// LoadEnv helper func to load .env file
func LoadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}

// ***Getters & Setters***
// getNodeAddress helper func to get node address from .env
func getNodeAddress() string {
	LoadEnv()
	return os.Getenv("NODEADDRESS")
}

// getJoinedOracleNetwork helper func to get flag value from .tmp file
func getJoinedOracleNetwork() bool {
	var parseBool *bool
	err := Load(&parseBool) // parsed from .tmp file
	if err != nil {
		log.Fatal(err)
	}
	return *parseBool //return value
}

// getPrivateKey gets private key from .env file
func getPrivateKey() string {
	LoadEnv()
	return os.Getenv("PRIVATEKEY")
}

// setNodeAddress **Not needed as cannot write to .env after termination**
func setNodeAddress(nodeAdress string) {
	LoadEnv()
	err := os.Setenv("NODEADDRESS", nodeAdress)
	if err != nil {
		log.Fatal(err)
	}
}

// setHasOracleJoinedNetwork saves the flag value to .tmp file
func setHasOracleJoinedNetwork(in bool) {
	err := Save(in)
	if err != nil {
		log.Fatal(err)
	}
}

// setPrivateKey **Not needed as cannot write to .env after termination**
func setPrivateKey(pk string) {
	LoadEnv()
	err := os.Setenv("PRIVATEKEY", pk)
	if err != nil {
		log.Fatal(err)
	}
}

// ***load and save bool value for oracle joined network ***

// Save function to save flag to .tmp file
func Save(v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create("./JoinedAsOracle.tmp")
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(f, r)
	return err
}

// Marshal helper func to marshal flag for .tmp file
func Marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Load loads flag value from .tmp
func Load(v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open("./JoinedAsOracle.tmp")
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// Unmarshal new json decoder for .tmp flag
func Unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// GetBroker gets MQTT broker value from .env
func GetBroker() string {
	LoadEnv()
	return os.Getenv("MQTTBROKER")
}

// GetClientID gets MQTT client ID from .env
func GetClientID() string {
	LoadEnv()
	return os.Getenv("ORACLEMQTTCLIENTID")
}

// GetClientUsername gets MQTT client username from.env
func GetClientUsername() string {
	LoadEnv()
	return os.Getenv("ORACLEMQTTUSERNAME")
}

// GetClientPasswd gets MQTT client passwd from .env
func GetClientPasswd() string {
	LoadEnv()
	return os.Getenv("ORACLEMQTTPASSWD")
}
