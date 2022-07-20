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

type OracleNodeInfo struct {
	NodeAddress         common.Address
	JoinedOracleNetwork bool
	PrivateKey          string
}

var lock sync.Mutex

// InitOracle creates hardcoded oracle details inorder to join the network
func InitOracle(client *ethclient.Client) OracleNodeInfo {

	nodeInfo := OracleNodeInfo{
		NodeAddress:         common.HexToAddress(getNodeAddress()),
		JoinedOracleNetwork: getJoinedOracleNetwork(),
		PrivateKey:          getPrivateKey(),
	}
	LeaveOracleNetwork(client, &nodeInfo)

	JoinOracleNetwork(client, &nodeInfo)
	return nodeInfo
}

// JoinOracleNetwork function to call join oracle
func JoinOracleNetwork(client *ethclient.Client, info *OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == false {
		TxJoinAsOracle(client, *info)
		ORACLENODEHASJOINED()
		info.JoinedOracleNetwork = true
		setHasOracleJoinedNetwork(info.JoinedOracleNetwork)
		return true
	}
	return false
}

// LeaveOracleNetwork Leave the oracle network
func LeaveOracleNetwork(client *ethclient.Client, info *OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == true {
		TxLeaveOracleNetwork(client, *info)
		ORACLENODEHASLEFT()
		info.JoinedOracleNetwork = false
		setHasOracleJoinedNetwork(info.JoinedOracleNetwork)
		return true
	}
	return false
}

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}

func getNodeAddress() string {
	loadEnv()
	return os.Getenv("NODEADDRESS")
}

func getJoinedOracleNetwork() bool {
	var parseBool *bool
	err := Load(&parseBool)
	if err != nil {
		log.Fatal(err)
	}
	return *parseBool
}

// ***Getters & Setters***

func getPrivateKey() string {
	loadEnv()
	return os.Getenv("PRIVATEKEY")
}

func setNodeAddress(nodeAdress string) {
	loadEnv()
	err := os.Setenv("NODEADDRESS", nodeAdress)
	if err != nil {
		log.Fatal(err)
	}
}

func setHasOracleJoinedNetwork(in bool) {
	err := Save(in)
	if err != nil {
		log.Fatal(err)
	}
}

func setPrivateKey(pk string) {
	loadEnv()
	err := os.Setenv("PRIVATEKEY", pk)
	if err != nil {
		log.Fatal(err)
	}
}

// ***load and save bool value for oracle joined network ***

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

func Marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

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

func Unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
