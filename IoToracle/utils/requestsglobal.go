package utils

import "sync"

var once sync.Once

type Request struct {
	RequestId        uint64
	DataType         []byte
	IotId            []byte
	UnPackedDataType DataType
	Status           int8
}
type RequestsSingleton map[uint64]*Request

var (
	reqs RequestsSingleton
)

func MakeRequests() RequestsSingleton {
	once.Do(func() {
		reqs = make(RequestsSingleton)
	})
	return reqs
}

var Requests = MakeRequests()
