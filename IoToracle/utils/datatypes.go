package utils

import (
	abi "IoToracle/abitogo"
	"encoding/json"
	"log"
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

func ConvertToJson(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) *DataType {
	err := json.Unmarshal(eventReleaseRequestDetails.Arg2,
		&Requests[id].UnPackedDataType)
	if err != nil {
		log.Fatal(err)
	}
	return &Requests[id].UnPackedDataType
}
