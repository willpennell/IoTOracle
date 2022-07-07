package utils

import "sync"

var once sync.Once

type RequestsSingleton map[uint64]Request

var (
	reqs RequestsSingleton
)

/*func MakeRequests() RequestsSingleton {
	once.Do(func() {
		reqs = make(RequestsSingleton)
	})
	return reqs
}*/
