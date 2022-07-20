package utils

//var once sync.Once

type Request struct {
	RequestId uint64
	DataType  []byte
	IotId     []byte
}
type RequestsSingleton map[uint64]Request

var (
	Reqs *RequestsSingleton
)

/*func MakeRequests() RequestsSingleton {
	once.Do(func() {
		reqs = make(RequestsSingleton)
	})
	return reqs
}

var Requests = MakeRequests()*/
