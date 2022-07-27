package utils

import (
	abi "IoToracle/abitogo"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type DataType struct {
	Type    string
	Topic   string
	TAfter  uint64
	TBefore uint64
}

type FetchedBoolResult struct {
	Result bool
}

type FetchedBigIntResult struct {
	Result big.Int
}

// package dataType message correctly.

// comes in as bytes need to transform into

// TODO function that builds a bool IoTRequest

func IoTBoolResult(result bool) FetchedBoolResult {
	return FetchedBoolResult{Result: result}
}

// TODO function that builds an big.Int IoTRequest

// IoTBigIntResult function that returns the fetched
func IoTBigIntResult(result big.Int) FetchedBigIntResult {
	return FetchedBigIntResult{Result: result}
}

// TODO function that subscribes for a bool result

// TODO function that subscribes for an int64 result

func FetchIoTData(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) {
	unpack := ConvertToJson(eventReleaseRequestDetails, id)
	if unpack.Type == "bool" {
		// call IoTfetch bool

	} else if unpack.Type == "int" {
		// call IoTFetch big.int
	} else {
		log.Error("error")
	}
}
