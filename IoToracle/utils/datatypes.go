// Package utils helper files and functions
package utils

import (
	abi "IoToracle/abitogo"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

// DataType struct of datatype that comes in from orc contract request
type DataType struct {
	Type    string // type of request needed, bool or int
	Topic   string // Topic used for identifying the MQTT topic combined with IoTId
	TAfter  uint64 // timestamp after, used to check that fetched result is after this
	TBefore uint64 // timestamp before, used to check that fetched result is before this
}

// FetchedBoolIoTResult struct of fetched result that comes in json style from MQTT broker {"Request":true,"Timestamp":189822}
type FetchedBoolIoTResult struct {
	Result    bool   // result true or false
	Timestamp uint64 // timestamp of recorded result
}

// FetchedBigIntIoTResult struct of fetched result that comes in json from MQTT broker {"Result":26,"Timestamp":189822}
type FetchedBigIntIoTResult struct {
	Result    big.Int // result int could be temperature, elapsed time, counter etc
	Timestamp uint64  // timestamp of recorded result
}

// IoTBoolResult the result for the blockchain only needs the result and not the timestamp so package for json marshal
type IoTBoolResult struct {
	Result bool
}

// IoTBigIntResult the result for the blockchain only needs the result and not the timestamp so package for json marshal
type IoTBigIntResult struct {
	Result big.Int
}

// FetchIoTData unpacks DataType and then subscribes MQTT broker
func FetchIoTData(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) {
	unpack := ConvertToDataTypeStruct(eventReleaseRequestDetails, id) // convert DataType []byte into DataType struct
	Requests[id].UnPackedDataType = *unpack                           // add unpacked DataType struct to Requests map
	if unpack.Type == "bool" {
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic and start MQTT client
		result := UnpackIoTBoolResult(packedResult)               // unpacked the result into IoTBoolResult
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
		packedResult := StartMQTTClient(TopicBuilder(unpack, id)) // pass in full topic IoTID/Topic and start MQTT client
		result := UnpackIoTBigIntResult(packedResult)             // unpacked the result into IoTBigIntResult
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
