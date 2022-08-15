// Package utils helper files and functions
package utils

import (
	abi "IoToracle/abitogo"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

// ConvertOpenForBidsData converts the input from events into a Request struct
func ConvertOpenForBidsData(a interface{}, b interface{}, c interface{}) (Request, uint64) {
	requestId, _ := a.(*big.Int)        // covert RequestId
	requestIdConv := requestId.Uint64() // converts from big.Int to uint64 for ID index in Requests map
	dataType, _ := b.([]byte)           // gets dataType info from event as []byte
	aggType, _ := c.(*big.Int)
	aggTypeConv := aggType.Int64()
	// adds converted types to struct
	request := Request{
		RequestId:       requestIdConv,
		DataType:        dataType,
		IotId:           nil,
		AggregationType: aggTypeConv,
	}
	return request, requestIdConv //returns the request struct and ID
}

// AddIoTIDToRequests add incoming IoTID and assign to correct id in Request map
func AddIoTIDToRequests(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails) []byte {
	Requests[eventReleaseRequestDetails.Arg0.Uint64()].IotId = eventReleaseRequestDetails.Arg1
	SaveRequestJson()
	return Requests[eventReleaseRequestDetails.Arg0.Uint64()].IotId
}

// ConvertToDataTypeStruct convert data from emitted event from json into Datatype struct
func ConvertToDataTypeStruct(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) *DataType {
	err := json.Unmarshal(eventReleaseRequestDetails.Arg2,
		&Requests[id].UnPackedDataType)
	if err != nil {
		log.Fatal(err)
	}
	return &Requests[id].UnPackedDataType
}

// UnpackIoTBoolResult unmarshal json into FetchedBoolIoTResult struct
func UnpackIoTBoolResult(packedResult []byte) FetchedBoolIoTResult {
	var resultAndTimeStamp FetchedBoolIoTResult
	err := json.Unmarshal(packedResult, &resultAndTimeStamp)
	if err != nil {
		log.Fatal(err)
	}
	return resultAndTimeStamp
}

func UnpackBool(packedResult []byte) bool {
	var ioTBool bool
	err := json.Unmarshal(packedResult, &ioTBool)
	if err != nil {
		log.Fatal(err)
	}
	return ioTBool
}

func UnpackInt(packedResult []byte) int {
	var ioTInt int
	err := json.Unmarshal(packedResult, &ioTInt)
	if err != nil {
		log.Fatal(err)
	}
	return ioTInt
}

// UnpackIoTBigIntResult unmarshal json into FetchedBigIntIoTResult struct
func UnpackIoTBigIntResult(packedResult []byte) FetchedBigIntIoTResult {
	var jsonObject interface{}
	var resultAndTimestamp FetchedBigIntIoTResult
	err := json.Unmarshal(packedResult, &jsonObject)
	if err != nil {
		log.Fatal(err)
	}
	proxy := createProxy(interfaceToMap(jsonObject))       // creates a proxy to get the correct types from the json IoT int data
	resultAndTimestamp.Result = int(proxy.Result)          // assigns the proxy values to the correct struct
	resultAndTimestamp.Timestamp = uint64(proxy.Timestamp) // assigns the proxy values to the correct struct
	return resultAndTimestamp
}

func interfaceToMap(v interface{}) map[string]interface{} {
	return v.(map[string]interface{})
}
func createProxy(v map[string]interface{}) Proxy {
	result := v["result"]
	timestamp := v["timestamp"]
	proxy := Proxy{
		Result:    result.(float64),
		Timestamp: timestamp.(float64),
	}
	return proxy
}

// packBoolToJson marshal into json from IoTBoolResult to []byte
func packBoolToJson(IoTResult FetchedBoolIoTResult) []byte {
	packed, err := json.Marshal(IoTResult.Result)
	if err != nil {
		log.Fatal(err)
	}
	return packed
}

// packBigIntToJson marshal into json from IoTBigIntResult to []byte
func packBigIntToJson(IoTResult FetchedBigIntIoTResult) []byte {
	packed, err := json.Marshal(IoTResult.Result)
	if err != nil {
		log.Fatal(err)
	}
	return packed
}

// TopicBuilder builds topic string from IoTID and Topic = IoTID/Topic for use with MQTT client
func TopicBuilder(unpack *DataType, id uint64) string {
	IoTtopic := fmt.Sprintf("%v/%v", string(Requests[id].IotId), unpack.Topic)
	return IoTtopic
}
