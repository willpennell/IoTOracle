package utils

import (
	abi "IoToracle/abitogo"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

func ConvertOpenForBidsData(a interface{}, b interface{}) (Request, uint64) {
	// covert RequestId
	requestId, _ := a.(*big.Int)
	requestIdConv := requestId.Uint64()

	dataType, _ := b.([]byte)

	request := Request{
		RequestId: requestIdConv,
		DataType:  dataType,
		IotId:     nil,
	}
	return request, requestIdConv
}

// AddIoTIDToRequests add incoming IoTID and assign to correct id in Request map
func AddIoTIDToRequests(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails) []byte {
	Requests[eventReleaseRequestDetails.Arg0.Uint64()].IotId = eventReleaseRequestDetails.Arg1
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

func UnpackIoTBoolResult(packedResult []byte) FetchedBoolIoTResult {
	var resultAndTimeStamp FetchedBoolIoTResult
	err := json.Unmarshal(packedResult, &resultAndTimeStamp)
	if err != nil {
		log.Fatal(err)
	}
	return resultAndTimeStamp
}

func UnpackIoTBigIntResult(packedResult []byte) FetchedBigIntIoTResult {
	var resultAndTimeStamp FetchedBigIntIoTResult
	err := json.Unmarshal(packedResult, &resultAndTimeStamp)
	if err != nil {
		log.Fatal(err)
	}
	return resultAndTimeStamp
}

func packBoolToJson(IoTResult IoTBoolResult) []byte {
	packed, err := json.Marshal(IoTResult)
	if err != nil {
		log.Fatal(err)
	}
	return packed
}

func packBigIntToJson(IoTResult IoTBigIntResult) []byte {
	packed, err := json.Marshal(IoTResult)
	if err != nil {
		log.Fatal(err)
	}
	return packed
}

func TopicBuilder(unpack *DataType, id uint64) string {
	IoTtopic := fmt.Sprintf("%v/%v", string(Requests[id].IotId), unpack.Topic)
	return IoTtopic
}
