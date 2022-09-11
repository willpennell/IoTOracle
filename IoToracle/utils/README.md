**converttypes.go** \
Used to convert incoming and outcoming data correctly
```go
// ConvertOpenForBidsData converts the input from events into a Request struct
func ConvertOpenForBidsData(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) (Request, uint64) {
	requestId, _ := a.(*big.Int)        // covert RequestId
	requestIdConv := requestId.Uint64() // converts from big.Int to uint64 for ID index in Requests map
	dataType, _ := b.([]byte)           // gets dataType info from event as []byte
	aggType, _ := c.(*big.Int)
	aggTypeConv := aggType.Int64()
	timestamp, _ := d.(*big.Int)
	timestampConv := timestamp.Int64()
	elapsedTime, _ := e.(*big.Int)
	elapsedTimeConv := elapsedTime.Int64()

	// adds converted types to struct
	request := Request{
		RequestId:       requestIdConv,
		DataType:        dataType,
		IotId:           nil,
		AggregationType: aggTypeConv,
		Timestamp:       timestampConv,
		ElapsedTime:     elapsedTimeConv,
		AppealFlag:      0,
	}
	return request, requestIdConv //returns the request struct and ID
}
```

\
**fetchdata.go** \
Used to create the correct struct for the aggregation type from the request and call the MQTT broker \
```go
// DataType struct of datatype that comes in from orc contract request
type DataType struct {
	Topic   string `json:"topic"`   // Topic used for identifying the MQTT topic combined with IoTId
	TAfter  uint64 `json:"tAfter"`  // timestamp after, used to check that fetched result is after this
	TBefore uint64 `json:"tBefore"` // timestamp before, used to check that fetched result is before this
}
```
```go
// FetchIoTData unpacks DataType and then subscribes MQTT broker
func FetchIoTData(eventReleaseRequestDetails *abi.OracleRequestContractReleaseRequestDetails, id uint64) {
	unpack := ConvertToDataTypeStruct(eventReleaseRequestDetails, id) // convert DataType []byte into DataType struct
	Requests[id].UnPackedDataType = *unpack                           // add unpacked DataType struct to Requests map
	SaveRequestJson()
	SimpleFetchIoT(&Requests[id].UnPackedDataType, id, GetClientID())
	SaveRequestJson()
}
```
\
**hashGeneration.go** \
Used to generate a commit hash\
```go
func GenerateHash(id uint64, secret []byte, result []byte) [32]byte {
	commitHash := crypto.Keccak256Hash(secret, result)
	Requests[id].CommitHash = commitHash
	return commitHash
}
```
**init.go** \
Used in the oraclecore.go file to initalise parts\
```go
// JoinOracleNetwork function to call join oracle function in smart contract
func JoinOracleNetwork(client *ethclient.Client, info *OracleNodeInfo) bool {
	if info.JoinedOracleNetwork == false { // makes sure node is currently not connected to network
		TxJoinAsOracle(client, *info)                       // tx function to call orc contract joinAsOracle() function
		ORACLENODEHASJOINED()                               // prints message
		info.JoinedOracleNetwork = true                     // change flag so we know oracle is now in network
		setHasOracleJoinedNetwork(info.JoinedOracleNetwork) // saves value to .tmp file
		// as .env cannot change value after program terminated
		return true
	}
	return false // returns false if already in network
}
```
**logging.go** \
Used to log info
```go
func TXLog(tx *types.Transaction) {
	txLogFile, err := os.OpenFile(TX_LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer txLogFile.Close()
	txLogFile.WriteString(PRINTTXHASH(tx))
}
```
**mqttclient.go** \
Used to connect to MQTT broker as a client
```go
// StartMQTTClient connects to MQTT client and subscribes to particular topic
func StartMQTTClient(id uint64, topic string, clientID string) []byte {
	fmt.Println(topic)
	//defer w.Done()
	subWait.Add(1)                 // add a counter to global sync.WaitGroup function
	opts := SetOpts(clientID)      // create set options and return options object and assign to 'opts'
	client := mqtt.NewClient(opts) // start new MQTT client with opts
	// creates token when connected if token has error then throw panic
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// subscription to topic
	setClient(client)
	sub(client, topic) // pass in client and topic to subscribe
	//client.Unsubscribe(topic)
	var tenSecs uint = 10 * 1000
	client.Disconnect(uint((Requests[id].ElapsedTime*2)*1000) + tenSecs) // disconnects after 250 milliseconds but sync.Wait() in sub will pause until complete first
	return packedResult                                                  // returns the message from pub
}
```
**printer.go** \
Used to print messages to terminal \
**randomString.go** \
Used to generate random string for secret in commit hash\
```go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

```
**requestsglobal.go** \
Used as the global Request Structure, one exists and is updated by all areas \
```go
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
	Timestamp        int64
	ElapsedTime      int64
	AppealFlag       uint8
	CommitFlag       uint8
	RevealFlag       uint8
}
```
**timer.go** \
used to keep a local timer from when the request comes in, 
if node commits and timer runs out they can appeal against unresponsive nodes
```go
func StampTimer(id uint64, timestamp int64, elapsedTime int64, client *ethclient.Client, nodeInfo OracleNodeInfo) bool {
	t := time.Now()
	fmt.Println(Requests[id].Status)
	fmt.Println("Timestamp: ", timestamp)
	fmt.Println("Elapsed Time: ", elapsedTime)
	fmt.Println("Time now: ", t.Unix())

	for t.Unix() < (timestamp+elapsedTime) && Requests[id].Status == 1 {
		t = time.Now()
		time.Sleep(time.Second)
		continue
	}
	if Requests[id].Status == 2 {
		return true
	}
	// call Timeout Appeal
	if Requests[id].AppealFlag == 1 && Requests[id].CommitFlag == 1 {
		TXtimeoutCommitsAppeal(client, nodeInfo, big.NewInt(int64(id)))
		// call commitAppeal
	} else if Requests[id].AppealFlag == 2 && Requests[id].RevealFlag == 1 {
		TXtimeoutRevealsAppeal(client, nodeInfo, big.NewInt(int64(id)))
	} else {
		forceCloseConnection()
		log.Println(log.Lshortfile, log.LstdFlags, "No appeal can be made")
	}
	return true
}
```
**txcreation.go**
Used to create transactions and send on chain \
```go
// Auth creates auth binding for transactions
func Auth(client *ethclient.Client, info OracleNodeInfo) *bind.TransactOpts {
	// creates transaction binding with private key and chainID
	// auth is returned, and we can add values to its attributes
	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey(info), ChainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(Nonce(client, info))) // add nonce value to auth
	auth.Value = big.NewInt(0)                          // add value to send when calling certain functions that are payable **some are**
	auth.GasLimit = uint64(6721975)                     // gas limit
	auth.GasPrice = GasForFuncCall(client)              // gas price for transaction
	return auth                                         // returns bind.TransactOpts object
}
```
```go
// OracleRequestContractInstance creates instance of OracleRequestContract
func OracleRequestContractInstance(client *ethclient.Client) *abi.OracleRequestContract {
	ORCInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	return ORCInstance
}

// TxJoinAsOracle creates tx to join oracle network
func TxJoinAsOracle(client *ethclient.Client, info OracleNodeInfo) *types.Transaction {
	auth := Auth(client, info)                               // auth object for transaction based on client and OracleNodeInfo
	stake := math.BigPow(10, 18)                             // 1 ether stake, when leaving the network it is returned.
	auth.Value = stake                                       // add to Value opt
	orcJoinAsOracle := OracleRequestContractInstance(client) // orc contract instance
	tx, err := orcJoinAsOracle.JoinAsOracle(auth)            // tx function call to JoinAsOracle, returns a tx
	if err != nil {
		log.Fatal(err)
	}
	TXLog(tx)
	return tx // return tx types.Transaction
}

```