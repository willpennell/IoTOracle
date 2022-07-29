package utils

import "sync"

var once sync.Once

type Request struct {
	RequestId        uint64
	DataType         []byte
	IotId            []byte
	UnPackedDataType DataType
	IoTResult        []byte
	Status           uint8 // 0 = not started, 1 = pending, 2 = complete
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
