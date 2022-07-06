package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Request struct {
	RequestId     uint64
	RequireResult bool
	IotId         []byte
	Requester     common.Address
}

func ConvertRequestTypes(a interface{}, b interface{}, c interface{}, d interface{}) Request {
	// convert requestID to uint64
	r1, _ := a.(big.Int)
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
}
