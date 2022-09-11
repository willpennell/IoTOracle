// Package core eventlistener.go is the core functionality of the system
package core

import (
	abi "IoToracle/abitogo"
	c "IoToracle/config"
	"IoToracle/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"sync"
	"time"
)

// ***TODOList*** (not yet needed)
// TODO catch up with old request that node may have placed bid on

// ***OracleRequestContract***

// SubscribeToOracleRequestContractEvents Listens to all event methods in OracleRequestContract
func SubscribeToOracleRequestContractEvents(client *ethclient.Client,
	wg *sync.WaitGroup,
	nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	var w sync.WaitGroup // start new sync.WaitGroup for the events in orc contract
	w.Add(6)
	go EventOpenForBids(client, &w, nodeInfo)           // go routine for open for bids event
	go EventBidPlaced(client, &w)                       // go routine for bids placed by oracle nodes
	go EventStatusChange(client, &w)                    // go routine for a change in status for requests
	go EventReleaseRequestDetails(client, &w, nodeInfo) // go routine for event that
	// emits request details for nodes to fetch
	go EventOraclePaid(client, &w) // go routine for paying oracles
	go EventUpdateTimestamp(client, &w, nodeInfo)
	w.Wait() // waits indefinitely as these events are persistent
}

// EventOpenForBids subscribe to OpenForBids
func EventOpenForBids(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	// creates a new orc contract instance that we can interact with
	if err != nil {
		log.Fatal(err)
	}
	// Watch options, just context used for this
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelOpenForBids := make(chan *abi.OracleRequestContractOpenForBids)
	// watch new event OpenForBids
	sub, err := orcInstance.WatchOpenForBids(watchOpts, channelOpenForBids)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe() // unsubscribes
	for {                   // continuous loop
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventOpenForBids := <-channelOpenForBids: // creates event object
			utils.OPENFORBIDSMESSAGE(eventOpenForBids) // prints message with data from event

			// creates a request struct
			request, id := utils.ConvertOpenForBidsData(eventOpenForBids.Arg0, eventOpenForBids.Arg1,
				eventOpenForBids.Arg2, eventOpenForBids.Arg3, eventOpenForBids.Arg4)

			utils.Requests[id] = &request // this is added to the Requests map
			utils.Requests[id].Status = 1 // change status to 1 to show this request is pending
			utils.SaveRequestJson()

			go utils.StampTimer(id, utils.Requests[id].Timestamp, utils.Requests[id].ElapsedTime, client, nodeInfo)
			//time.Sleep(time.Second)
			// call tx function to send to orc contract and place bid
			utils.TxPlaceBid(client, nodeInfo, big.NewInt(int64(id)))
			//go EventReleaseRequestDetails(client, wg, nodeInfo)
		}
	}
}

// EventBidPlaced subscribe to PlaceBid
func EventBidPlaced(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	// creates a new orc contract instance that we can interact with
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for an event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelBidPlaced := make(chan *abi.OracleRequestContractBidPlaced)
	// watch new event BidPlaced
	sub, err := orcInstance.WatchBidPlaced(watchOpts, channelBidPlaced)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventBidPlaced := <-channelBidPlaced:
			utils.BIDPLACEDMEASSGE(eventBidPlaced) // prints message
		}
	}
	// Receive events from the channel

}

// EventReleaseRequestDetails subscribe to ReleaseRequestDetails
func EventReleaseRequestDetails(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()

	// creates a new orc contract instance that we can interact with
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for an event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelReleaseRequestDetails := make(chan *abi.OracleRequestContractReleaseRequestDetails)
	// watch new event ReleaseRequestDetails
	sub, err := orcInstance.WatchReleaseRequestDetails(watchOpts, channelReleaseRequestDetails)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventReleaseRequestDetails := <-channelReleaseRequestDetails:
			time.Sleep(time.Second * 2)
			id := eventReleaseRequestDetails.Arg0.Uint64() // converts id to uint64
			if utils.Requests[id].Status == 1 {
				utils.AddIoTIDToRequests(eventReleaseRequestDetails) // adds IoTID to Requests map
				utils.SaveRequestJson()
				utils.RELEASEREQUESTDETAILS(eventReleaseRequestDetails) // prints message
				// call fetch to IoT
				//time.Sleep(time.Second * 10)

				utils.FetchIoTData(eventReleaseRequestDetails, id) // function to call MQTT broker
				// need to convert the result to hex
				utils.SaveRequestJson()
				// send tx function call ReceiveResponse in Aggregator contract with the result as a hex string
				txBool := utils.TxCommitResponse(client, nodeInfo, eventReleaseRequestDetails.Arg0, utils.Requests[id].CommitHash)
				if txBool == true {
					utils.Requests[id].CommitFlag = 1
					utils.Requests[id].AppealFlag = 1
					utils.SaveRequestJson()
				}

				// utils.TxReceiveResponse(client, nodeInfo, big.NewInt(int64(id)), common.
			} else {
				fmt.Println("Unlucky, didnt win bid for this request..")
			}

		}
	}
}

// EventUpdateTimestamp subscribe to UpdateTimestamp
func EventUpdateTimestamp(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	// creates a new orc contract instance that we can interact with
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for an event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelUpdateTimestamp := make(chan *abi.OracleRequestContractUpdateTimestamp)
	// watch new event BidPlaced
	sub, err := orcInstance.WatchUpdateTimestamp(watchOpts, channelUpdateTimestamp)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventUpdateTimestamp := <-channelUpdateTimestamp:
			id := eventUpdateTimestamp.Arg0.Uint64()
			newTimestamp := eventUpdateTimestamp.Arg1.Int64()
			newElapsedTime := eventUpdateTimestamp.Arg2.Int64()
			utils.Requests[id].Timestamp = newTimestamp
			utils.Requests[id].ElapsedTime = newElapsedTime
			utils.SaveRequestJson()
			go utils.StampTimer(id, utils.Requests[id].Timestamp, utils.Requests[id].ElapsedTime, client, nodeInfo)
		}
	}
	// Receive events from the channel

}

// EventStatusChange Subscribe to StatusChange
func EventStatusChange(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	// creates a new orc contract instance that we can interact with
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch options for event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelStatusChange := make(chan *abi.OracleRequestContractStatusChange)
	// watch new event Status Change
	sub, err := orcInstance.WatchStatusChange(watchOpts, channelStatusChange)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventStatusChange := <-channelStatusChange:
			utils.STATUSCHANGEMESSAGE(eventStatusChange) // prints status change message
			id := eventStatusChange.Arg0.Uint64()
			//if eventStatusChange.Arg1 == 0 {
			//	utils.Requests[id].Status = 1
			//} else
			if eventStatusChange.Arg1 == 1 {
				utils.Requests[id].Status = 2
			}
			utils.SaveRequestJson()
		}
	}
}

// EventOraclePaid listens for oracle paid messages
func EventOraclePaid(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	// creates a new orc contract instance that we can interact with
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch options
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event

	channelOraclePaid := make(chan *abi.OracleRequestContractOraclePaid)
	// watch new event OraclePaid
	sub, err := orcInstance.WatchOraclePaid(watchOpts, channelOraclePaid)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventOraclePaid := <-channelOraclePaid:
			utils.ORACLEPAID(eventOraclePaid) // prints message when OraclePaid
		}
	}
}

// ***AggregatorContract***

// SubscribeToAggregationContractEvents listens to all events in AggregationContract
func SubscribeToAggregationContractEvents(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	var w sync.WaitGroup
	w.Add(4)

	go EventAggregationComplete(client, &w) // go routine for AggregationComplete event
	go EventLogHashes(client, &w)           // go routine for LogHashes event
	go EventCommitsPlaced(client, &w, nodeInfo)
	go EventRevealsPlaced(client, &w, nodeInfo)
	w.Wait() // waits indefinitely
}

// EventCommitsPlaced TODO commits received
func EventCommitsPlaced(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelCommitsPlaced := make(chan *abi.AggregatorContractCommitsPlaced)
	// watch new event AggregationComplete
	sub, err := aggInstance.WatchCommitsPlaced(watchOpts, channelCommitsPlaced)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventCommitsPlaced := <-channelCommitsPlaced:
			time.Sleep(time.Second * 2)
			id := eventCommitsPlaced.Arg0.Uint64()
			if utils.Requests[id].Status == 1 && utils.Requests[id].CommitFlag == 1 {
				utils.COMMITSPLACEDMESSAGE(eventCommitsPlaced)

				if utils.Requests[id].AggregationType == 1 {
					shaBoolEncode := solsha3.Bool(utils.UnpackBool(utils.Requests[id].IoTResult))
					txBool := utils.TxRevealVoteResponse(client, nodeInfo, eventCommitsPlaced.Arg0, shaBoolEncode, utils.Requests[id].Secret)
					if txBool == true {
						utils.Requests[id].RevealFlag = 1
						utils.SaveRequestJson()
					}

				} else if utils.Requests[id].AggregationType == 2 {
					shaIntEncode := solsha3.Int256(utils.UnpackInt(utils.Requests[id].IoTResult))
					txBool := utils.TxRevealAverageResponse(client, nodeInfo, eventCommitsPlaced.Arg0, shaIntEncode, utils.Requests[id].Secret)
					if txBool == true {
						utils.Requests[id].RevealFlag = 1
						utils.SaveRequestJson()
					}

				}
				utils.Requests[id].AppealFlag = 2
				utils.SaveRequestJson()

				// prints aggregation complete message
			} else {
				fmt.Println("Failed to submit reveal.")
			}

		}
	}
}

// TODO reveals received
func EventRevealsPlaced(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelRevealsPlaced := make(chan *abi.AggregatorContractRevealsPlaced)
	// watch new event AggregationComplete
	sub, err := aggInstance.WatchRevealsPlaced(watchOpts, channelRevealsPlaced)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventRevealsPlaced := <-channelRevealsPlaced:
			utils.REVEALSPLACEDMESSAGE(eventRevealsPlaced)

		}
	}
}

// EventAggregationComplete Subscribe to AggregationComplete
func EventAggregationComplete(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelAggregationCompleted := make(chan *abi.AggregatorContractAggregationCompleted)
	// watch new event AggregationComplete
	sub, err := aggInstance.WatchAggregationCompleted(watchOpts, channelAggregationCompleted)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventAggregationCompleted := <-channelAggregationCompleted:
			utils.AGGREGATIONCOMPLETE(eventAggregationCompleted) // prints aggregation complete message
			utils.Requests[eventAggregationCompleted.Arg0.Uint64()].Status = 2
			fmt.Println("Voting status change: ", utils.Requests[eventAggregationCompleted.Arg0.Uint64()].Status)
			utils.SaveRequestJson()
		}
	}
}

// EventLogHashes func that listens for log hashes
func EventLogHashes(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch options for event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// chan for event
	channelLogHashes := make(chan *abi.AggregatorContractLogHashes)
	// watch new event LogHashes
	sub, err := aggInstance.WatchLogHashes(watchOpts, channelLogHashes)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLogHashes := <-channelLogHashes:
			utils.HASHLOGMESSAGE(eventLogHashes)
		}
	}
}
