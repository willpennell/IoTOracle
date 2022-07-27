package utils

import (
	abi "IoToracle/abitogo"
	"encoding/json"
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

func ConvertIoTID(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails) []byte {
	IoTId := eventReleaseRequestDetails.Arg1
	Requests[eventReleaseRequestDetails.Arg0.Uint64()].IotId = IoTId
	return IoTId
}

func ConvertToJson(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) *DataType {
	err := json.Unmarshal(eventReleaseRequestDetails.Arg2,
		&Requests[id].UnPackedDataType)
	if err != nil {
		log.Fatal(err)
	}
	return &Requests[id].UnPackedDataType
}
