// Package utils helper files and functions
package utils

import (
	"sync"
)

// once var for singleton Requests map
var once sync.Once

// Request struct that holds all a requests information
type Request struct {
	RequestId        uint64   // RequestID
	DataType         []byte   // DataType []byte packed
	UnPackedDataType DataType // unpacked DataType so attributes can be used
	AggregationType  int64    // Aggregation type either 1 = voting, 2 = averaging
	IotId            []byte   // IoTId []byte packed
	IoTResult        []byte   // packed result
	Status           uint8    // 0 = not started, 1 = pending, 2 = complete
	Secret           []byte   // secret generated by RandString to use in commit hash
	CommitHash       [32]byte // commit hash
}

// RequestsSingleton map of ids and Request structs
type RequestsSingleton map[uint64]*Request

// var of RequestsSingleton
var (
	reqs RequestsSingleton
)

// MakeRequests function to make a single Request or return the only reqs
func MakeRequests() RequestsSingleton {
	once.Do(func() {
		reqs = make(RequestsSingleton) // only done on first run
	})
	return reqs // otherwise runs
}

var Requests = MakeRequests() // singleton Requests struct

// TODO log to output file
