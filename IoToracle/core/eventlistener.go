package core

import (
	abi "IoToracle/abitogo"
	c "IoToracle/config"
	"IoToracle/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"sync"
	"time"
)

// ***TODOList*** (not yet needed)
// TODO catch up with old request that node may have placed bid on
// TODO subscribe OracleJoined
// TODO Subscribe OracleLeft

// ***OracleRequestContract***

// SubscribeToOracleRequestContractEvents Listens to all event methods in OracleRequestContract
func SubscribeToOracleRequestContractEvents(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	var w sync.WaitGroup
	w.Add(5)
	go EventOpenForBids(client, &w, nodeInfo)
	go EventBidPlaced(client, &w)
	go EventStatusChange(client, &w)
	go EventReleaseRequestDetails(client, &w, nodeInfo)
	go EventOraclePaid(client, &w)
	w.Wait()
}

// EventOpenForBids subscribe to OpenForBids
func EventOpenForBids(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelOpenForBids := make(chan *abi.OracleRequestContractOpenForBids)
	// Start a goroutine which watches new events

	sub, err := orcInstance.WatchOpenForBids(watchOpts, channelOpenForBids)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventOpenForBids := <-channelOpenForBids:
			utils.OPENFORBIDSMESSAGE(eventOpenForBids)

			request, id := utils.ConvertOpenForBidsData(eventOpenForBids.Arg0, eventOpenForBids.Arg1)
			utils.Requests[id] = &request

			time.Sleep(time.Second)
			tx, err := utils.TxPlaceBid(client, nodeInfo, big.NewInt(int64(id)))
			if err != nil {
				log.Fatal(err)
			}
			utils.NEWBID()
			utils.PRINTTXHASH(tx)
			utils.ENDBID()

		}
	}
	// Receive events from the channel

}

// EventBidPlaced subscribe to PlaceBid
func EventBidPlaced(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelBidPlaced := make(chan *abi.OracleRequestContractBidPlaced)
	// Start a goroutine which watches new events

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
			utils.BIDPLACEDMEASSGE(eventBidPlaced)
		}
	}
	// Receive events from the channel

}

// EventReleaseRequestDetails subscribe to ReleaseRequestDetails
func EventReleaseRequestDetails(client *ethclient.Client, wg *sync.WaitGroup, nodeInfo utils.OracleNodeInfo) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelReleaseRequestDetails := make(chan *abi.OracleRequestContractReleaseRequestDetails)
	// Start a goroutine which watches new events

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
			id := eventReleaseRequestDetails.Arg0.Uint64()
			utils.RELEASEREQUESTDETAILS(eventReleaseRequestDetails)
			// call fetch to IoT
			utils.FetchIoTData(eventReleaseRequestDetails, id)
			fmt.Println("Dummy response...")
			// call a tx to send response to Aggregator Oracle

			//utils.TxReceiveResponse(client, nodeInfo, eventReleaseRequestDetails.Arg0, fetchedResult)

		}
	}
	// Receive events from the channel
}

// EventStatusChange Subscribe to StatusChange
func EventStatusChange(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelStatusChange := make(chan *abi.OracleRequestContractStatusChange)
	// Start a goroutine which watches new events

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
			utils.STATUSCHANGEMESSAGE(eventStatusChange)
		}
	}
	// Receive events from the channel
}

func EventOraclePaid(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	orcInstance, err := abi.NewOracleRequestContract(c.ORACLEREQUESTCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelOraclePaid := make(chan *abi.OracleRequestContractOraclePaid)
	// Start a goroutine which watches new events

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
			utils.ORACLEPAID(eventOraclePaid)
		}
	}
	// Receive events from the channel
}

// ***AggregatorContract***

// SubscribeToAggregationContractEvents listens to all events in AggregationContract
func SubscribeToAggregationContractEvents(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	var w sync.WaitGroup
	w.Add(2)
	go EventResponseReceived(client, &w)
	go EventAggregationComplete(client, &w)
	//go EventLogHahses(client, &w)
	w.Wait()
}

// EventResponseReceived Subscribe to ResponseReceived
func EventResponseReceived(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelResponseReceived := make(chan *abi.AggregatorContractResponseReceived)
	// Start a goroutine which watches new events

	sub, err := aggInstance.WatchResponseReceived(watchOpts, channelResponseReceived)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventResponseReceived := <-channelResponseReceived:
			utils.RESPONSERECIEVED(eventResponseReceived)
		}
	}
	// Receive events from the channel
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
	// Setup a channel for results

	channelAggregationCompleted := make(chan *abi.AggregatorContractAggregationCompleted)
	// Start a goroutine which watches new events

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
			utils.AGGREGATIONCOMPLETE(eventAggregationCompleted)

		}
	}
	// Receive events from the channel
}

func EventLogHahses(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	aggInstance, err := abi.NewAggregatorContract(c.AGGREGATIONCONTRACTADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results

	channelLogHashes := make(chan *abi.AggregatorContractLogHashes)
	// Start a goroutine which watches new events
	sub, err := aggInstance.WatchLogHashes(watchOpts, channelLogHashes)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLogHahses := <-channelLogHashes:
			fmt.Println("OG hash: ", eventLogHahses.Arg0)
			fmt.Println("Result hash: ", eventLogHahses.Arg1)

		}
	}
}
