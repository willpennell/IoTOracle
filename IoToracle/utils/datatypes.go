// Package utils helper files and functions
package utils

import (
	abi "IoToracle/abitogo"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// DataType struct of datatype that comes in from orc contract request
type DataType struct {
	Topic   string `json:"topic"`   // Topic used for identifying the MQTT topic combined with IoTId
	TAfter  uint64 `json:"tAfter"`  // timestamp after, used to check that fetched result is after this
	TBefore uint64 `json:"tBefore"` // timestamp before, used to check that fetched result is before this
}

// FetchedBoolIoTResult struct of fetched result that comes in json style from MQTT broker {"Request":true,"Timestamp":189822}
type FetchedBoolIoTResult struct {
	Result    bool   `json:"result"`    // result true or false
	Timestamp uint64 `json:"timestamp"` // timestamp of recorded result
}

type Proxy struct {
	Result    float64 `json:"result"`
	Timestamp float64 `json:"timestamp"`
}

// FetchedBigIntIoTResult struct of fetched result that comes in json from MQTT broker {"Result":26,"Timestamp":189822}
type FetchedBigIntIoTResult struct {
	Result    int    `json:"result"`    // result int could be temperature, elapsed time, counter etc
	Timestamp uint64 `json:"timestamp"` // timestamp of recorded result
}

// FetchIoTData unpacks DataType and then subscribes MQTT broker
func FetchIoTData(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) {
	unpack := ConvertToDataTypeStruct(eventReleaseRequestDetails, id) // convert DataType []byte into DataType struct
	Requests[id].UnPackedDataType = *unpack                           // add unpacked DataType struct to Requests map
	if Requests[id].AggregationType == 1 {
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic and start MQTT client
		result := UnpackIoTBoolResult(packedResult)               // unpacked the result into IoTBoolResult
		Requests[id].IoTResult = packBoolToJson(result)
		// check timestamp is in window given in DataType
		if checkBoolTimeStamp(result, id) {
			Requests[id].Secret = RandStringBytes(64)
			fmt.Println(string(Requests[id].Secret))
			tes := solsha3.Bool(UnpackBool(Requests[id].IoTResult))
			fmt.Println(tes)
			Requests[id].CommitHash = GenerateHash(id, Requests[id].Secret, tes)
			fmt.Println("Commit Hash: ", Requests[id].CommitHash)
		} else {
			log.Error("Error... timestamp not in time window required.")
		}

	} else if Requests[id].AggregationType == 2 {
		// call IoTFetch big.int
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic and start MQTT client
		jsonString := string(packedResult)
		result := UnpackIoTBigIntResult([]byte(jsonString)) // unpacked the result into IoTBigIntResult
		//fmt.Println(yo)
		Requests[id].IoTResult = packBigIntToJson(result)
		// check timestamp is in window given in DataType
		if checkBigIntTimeStamp(result, id) {
			Requests[id].Secret = RandStringBytes(64)
			fmt.Println(string(Requests[id].Secret))
			tes := solsha3.Int256(UnpackInt(Requests[id].IoTResult))
			fmt.Println(tes)
			Requests[id].CommitHash = GenerateHash(id, Requests[id].Secret, tes)
			fmt.Println("Commit Hash: ", Requests[id].CommitHash)
		} else {
			log.Error("Error... timestamp not in time window required.")
		}
	} else {
		log.Error("error")
	}
}

// checkBoolTimeStamp checks that timestamp is between given window
func checkBoolTimeStamp(result FetchedBoolIoTResult, id uint64) bool {
	if result.Timestamp >= Requests[id].UnPackedDataType.TAfter &&
		result.Timestamp <= Requests[id].UnPackedDataType.TBefore {
		return true
	}
	return false
}

// checkBigIntTimeStamp checks that timestamp is between given window
func checkBigIntTimeStamp(result FetchedBigIntIoTResult, id uint64) bool {
	if result.Timestamp >= Requests[id].UnPackedDataType.TAfter &&
		result.Timestamp <= Requests[id].UnPackedDataType.TBefore {
		return true
	}
	return false
}
