package utils

import (
	"IoToracle/abitogo"
	"math/big"
)

/*func ConvertRequestTypes(a interface{}, b interface{}, c interface{}, d interface{}) Request {
	// convert requestID to uint64
	r1, _ := a.(*big.Int)
	r1conv := r1.Uint64()

	// convert to bool
	r2, _ := b.(bool)

	r3, _ := c.([]byte)

	r4, _ := d.(common.Address)

	req := Request{
		RequestId:     r1conv,
		RequireResult: r2,
		IotId:         r3,
		Requester:     r4,
	}
	return req
}*/

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

func ConvertIoTID(eventReleaseRequestDetails *abitogo.OracleRequestContractReleaseRequestDetails) []byte {
	IoTId := eventReleaseRequestDetails.Arg1
	Requests[eventReleaseRequestDetails.Arg0.Uint64()].IotId = IoTId
	return IoTId
}