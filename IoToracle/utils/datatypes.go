package utils

import (
	abi "IoToracle/abitogo"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type DataType struct {
	Type    string
	Topic   string
	TAfter  uint64
	TBefore uint64
}

type FetchedBoolIoTResult struct {
	Result    bool
	Timestamp uint64
}

type FetchedBigIntIoTResult struct {
	Result    big.Int
	Timestamp uint64
}

type IoTBoolResult struct {
	Result bool
}

type IoTBigIntResult struct {
	Result big.Int
}

// package dataType message correctly.

// comes in as bytes need to transform into

// TODO function that builds a bool IoTRequest

// TODO function that builds an big.Int IoTRequest

// IoTBigIntResult function that returns the fetched

// TODO function that subscribes for a bool result

// TODO function that subscribes for an int64 result

func FetchIoTData(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) {
	unpack := ConvertToDataTypeStruct(eventReleaseRequestDetails, id)
	Requests[id].UnPackedDataType = *unpack
	if unpack.Type == "bool" {
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic
		result := UnpackIoTBoolResult(packedResult)
		// check timestamp is in window given in DataType
		if checkBoolTimeStamp(result, id) {
			packed := IoTBoolResult{Result: result.Result}           // IoTBoolResult struct created to be packed for marshal
			packedJson := packBoolToJson(packed)                     // packed into []byte ready to be sent back
			fmt.Println("final packed result: ", string(packedJson)) // print to terminal
			Requests[id].IoTResult = packedJson                      // add packed data to Requests[id] map
		} else {
			log.Error("Error... timestamp not in time window required.")
		}

	} else if unpack.Type == "int" {
		// call IoTFetch big.int
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic
		result := UnpackIoTBigIntResult(packedResult)
		// check timestamp is in window given in DataType
		if checkBigIntTimeStamp(result, id) {
			packed := IoTBigIntResult{Result: result.Result}         // IoTBoolResult struct created to be packed for marshal
			packedJson := packBigIntToJson(packed)                   // packed into []byte ready to be sent back
			fmt.Println("final packed result: ", string(packedJson)) // print to terminal
			Requests[id].IoTResult = packedJson                      // add packed data to Requests[id] map
		} else {
			log.Error("Error... timestamp not in time window required.")
		}
	} else {
		log.Error("error")
	}
}

func checkBoolTimeStamp(result FetchedBoolIoTResult, id uint64) bool {
	if result.Timestamp >= Requests[id].UnPackedDataType.TAfter &&
		result.Timestamp <= Requests[id].UnPackedDataType.TBefore {
		return true
	}
	return false
}

func checkBigIntTimeStamp(result FetchedBigIntIoTResult, id uint64) bool {
	if result.Timestamp >= Requests[id].UnPackedDataType.TAfter &&
		result.Timestamp <= Requests[id].UnPackedDataType.TBefore {
		return true
	}
	return false
}
