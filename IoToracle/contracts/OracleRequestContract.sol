// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./AggregatorContract.sol";

contract OracleRequestContract {
    //address that deploys contract
    address public owner;
    // address of aggregator contract
    address public aggregatorAddr;
    // Array to populate with requests that come in
    Request[100] public requests;
    // all mappings of oracles that have joined network
    mapping (address => bool) public oracles;
    // stakeBalance of each node, should be equal to 1 eth
    mapping (address => uint) public stakeBalance;
    // request and and pending counters, pending decreased when status changed to COMPLETE
    uint requestCounter = 1;
    uint pendingCounter = 1;
    // Status logs for each request
    enum Status{PENDING, COMPLETE, CANCELLED}
    // a single Request structure
    struct Request {
        uint256 requestID; //Id for each request
        address requester; // address which made the request
        address callbackAddress; // an address in which the response can be delivered to
        bytes callbackFID; // a function in the callback address that can be called
        bytes IoTID; // the ID of the IoT the user-SC wishes to gather data from
        bytes dataType; // data type, JSON that is decoded by the oracle node that will collect the data
        bytes32 pHash; // a hash of the data type and the actual result required.
        //bytes requiredResult; // left this out for security reason, we only apply it to the hash
        uint32 numberOfOracles; // number of oracles needed to fetch the data
        mapping (address => bool) oraclesForRequest; // must be true in order for oracle node to participate
        mapping (uint32 => address) oracleAddressAccess; // mapping so we can iterate through addresses
        uint32 oracleCounter; // index and counter for the number of oracles
        Status status; // status of the request
        uint8 cancelFlag; // cancel flag, if == 1 then do not continue with request
        uint fee; // fee sent to cover costs and extra for fetch the data
    }

    bool[100] public completedRequests;// true for each requestID once status changes to complete.

    event OpenForBids(uint256, bytes); // tells oracle nodes there is a request that needs bidding on
    event BidPlaced(address); // tells nodes that a bid has been placed by another node
    event ReleaseRequestDetails(uint256, bytes, bytes); // the details the node needs to gather
    event StatusChange(Status, string); // tells nodes in network of status change of requests
    event OracleJoined(address, string); // tells network that a new oracle has joined the network
    event OracleLeft(address, string); // tells the network that another oracle has left
    event OraclePaid(address, uint, string); // tells the network when an oracle has been paid for its services
    // ***payments***
    uint GAS_PRICE = 2 * 10**10; // cost of gas
    uint FEE = GAS_PRICE * 2; // fee to cover other function calls and pay each oracle
    uint MIN_FEE = FEE + (GAS_PRICE * 3); // minimum needed to cover all costs

    // @notice constructor when first deployed
    constructor() {
        owner = msg.sender; // set the owner of the contract to the address that deploys it
    }
    // @notice used if eth is sent to an invalid function etc
    fallback() external{}

    // @notice function called by addresses of Off-chain Oracle nodes that want to join the network
    // @return bool, true if successful otherwise reverts
    function joinAsOracle()
    public
    payable
    oracleNotJoined()
    valueEqual1Eth()
    returns(bool){
        oracles[address(msg.sender)] = true; // set to true as is now in network
        stakeBalance[msg.sender] = msg.value; // added value to stakeBalance as a record
        emit OracleJoined(msg.sender, "Welcome new node!"); // emit event so listeners can be notified
        return oracles[address(msg.sender)]; // return true
    }
    // @notice oracle node calls this function when they wish to leave network
    // @return bool
    function leaveOracleNetwork()
    public
    oracleHasJoined()
    stakeCheck()
    returns(bool) {
        delete oracles[msg.sender]; // deletes address from mapping
        payable(msg.sender).transfer(stakeBalance[msg.sender]); // transfers their stake back to them
        stakeBalance[msg.sender] = 0; // set stakeBalance back to 0 for the oracle node
        emit OracleLeft(msg.sender, "Node has left, Goodbye"); // emit event so listeners can be notified
        return true; // return true to confirm
    }

    // @notice user can call function create request
    // @dev
    // @param _callbackAddress could be different to address that called the function
    // @param _callbackFID a particular function in a smart contract that is needed to be called based on request
    // @param _IoTID the identifier of the IoT device oracle node will be subscribing to
    // @param _dataType, true false answer, an event, continuous data over time?
    // @param _requiredResult what is needed from IoT device
    // @param _numberOfOracles number of oracle nodes that fetch IoT data
    // @return the requestID so the user can keep track of request
    function createRequest(
        address _callbackAddress,
        bytes memory _callbackFID,
        bytes memory _IoTID,
        bytes memory _dataType,
        bytes memory _requiredResult,
        uint32 _numberOfOracles)
    external
    payable
    valueGTMinFeeCheck() // requires the value sent is >= MIN_FEE
    oddNumOracles(_numberOfOracles) // requires an odd number of oracles
    returns(uint256)
    {
        uint256 requestID = requestCounter; // assign id to request
        // log details of the request to array to cross reference later
        requests[requestID].requestID = requestID;
        requests[requestID].requester = msg.sender;
        requests[requestID].callbackAddress = _callbackAddress;
        requests[requestID].callbackFID = _callbackFID;
        requests[requestID].IoTID = _IoTID;
        requests[requestID].dataType = _dataType;
        //requests[requestId].requiredResult = _requiredResult; // do not need to store result put in a hash?
        requests[requestID].numberOfOracles = _numberOfOracles;
        requests[requestID].oracleCounter = 0;
        requests[requestID].status = Status.PENDING;
        requests[requestID].cancelFlag = 0;
        requests[requestID].fee = msg.value; // keeps track of the fee for the request
        emit StatusChange(requests[requestID].status, "PENDING"); // tells nodes in network of status change of requests
        // create a hash of the _dataType and _requiredResult:
        // we will use the actual result to test upon delivery.
        bytes32 phash = keccak256(abi.encodePacked(_dataType, _requiredResult));
        // add required result to hash, only time it is used, we assume that the user-SC knows what answer they are looking for
        requests[requestID].pHash = phash;
        // increment requestId ready for next create request function.
        requestCounter++;
        pendingCounter++;
        emit OpenForBids(requestID, requests[requestID].dataType);
        return requestID;
    }

    // @notice off-chain oracle node places bid to fetch data, once threshold is reached, emits event details.
    // @dev
    // @param requestId oracle node wishes to bid on
    // @return true to indicate bid successful
    function placeBid(
        uint256 _requestID)
    public
    oracleHasJoined() // requires that the oracle is in the network
    cancelFlagZero(_requestID) // requires requests cancel flag is zero
    orcCountLessThanNum(_requestID) // requires that the oracle counter is still les than the number of oracles needed
    returns(bool)
    {
        requests[_requestID].oraclesForRequest[msg.sender] = true; // add true for oracle in request so we know they can vote
        requests[_requestID].oracleAddressAccess[requests[_requestID].oracleCounter] = address(msg.sender); // map addr for iteration
        requests[_requestID].oracleCounter++; // increment the counter
        // TODO also add a timeout, if bids take to long
        emit BidPlaced(msg.sender); // send message to network that a bid has been placed
        if (requests[_requestID].numberOfOracles == requests[_requestID].oracleCounter) {
            // when the last oracle has submitted, we can emit the details
            emit ReleaseRequestDetails(
                requests[_requestID].requestID, // requestID
                requests[_requestID].IoTID, // IoT info, this is most of importance
                requests[_requestID].dataType); // DataType
            return true; // returns true to show completion of the bidding phase that is now over
        }
        return oracles[msg.sender]; // returns true for each oracle that bids
    }

    // @notice called by aggregator contract to deliver final response
    // @param _requestID, id of request
    // @param _finalResult, final aggregated result
    // @param _correctOracles, array of addresses of oracles to be paid for their work
    function deliverResponse(
        uint256 _requestID,
        bytes memory _finalResult,
        address[] memory _correctOracles)
    public
    cancelFlagZero(_requestID) // requires that the cancelFlag is = 0
    onlyAggregator() // only the aggregator contract can call this function
    returns(bool)
    {
        // return single aggregation result
        _finalResult;
        uint fee = requests[_requestID].fee / (requests[_requestID].numberOfOracles);
        // pay fees
        for (uint i=0; i<_correctOracles.length; i++) {
            // pay each share to each oracle
            emit OraclePaid(_correctOracles[i], fee, "Oracle has been Paid!"); // tells listeners oracle node has been paid
            payable(_correctOracles[i]).transfer(fee); // pays oracle their fee
            requests[_requestID].fee -= fee; // deduct the oracles portion from the logged amount until it reaches 0
            // and all oracles have been paid
        }
        requests[_requestID].status = Status.COMPLETE; // request is complete
        completedRequests[_requestID] = true; // add request id to completedRequests
        emit StatusChange(requests[_requestID].status, "COMPLETE"); // tell listeners that the status has changed
        pendingCounter--; // no longer pending so countered can be decremented
        return true; // return true to so it has finished
    }

    function cancelRequest(
        uint256 _requestID)
    external
    // need to make sure only the requester cancel the request
    onlyRequester(_requestID) // requires that only the requester can cancel
    cancelFlagZero(_requestID)
    returns(bool)
    {
        requests[_requestID].cancelFlag = 1;
        requests[_requestID].status = Status.CANCELLED;
        AggregatorContract(aggregatorAddr).cancelRequest(_requestID);
        return true;
    }
    // *** Setters ***
    function setAggregatorContract(
        address _aggAddr)
    public
    onlyOwner() // only owner can set aggregator contract
    returns(bool)
    {
        aggregatorAddr = _aggAddr;
        return true;
    }

    // *** Getters ***
    function getOracleForRequest(
        uint256 _requestID,
        address _orcAddr)
    external
    view
    onlyAggregator()
    returns(bool)
    {
        return requests[_requestID].oraclesForRequest[_orcAddr];
    }

    function getNumberOfOracles(
        uint256 _requestID)
    external
    view
    onlyAggregator()
    returns(uint32)
    {
        return requests[_requestID].numberOfOracles;
    }
    function getPHash(
        uint256 _requestID)
    external
    view
    onlyAggregator()
    returns(bytes32)
    {
        return requests[_requestID].pHash;
    }
    function getDataType(
        uint256 _requestID)
    external
    view
    onlyAggregator()
    returns(bytes memory)
    {
        return requests[_requestID].dataType;
    }
    // ***Modifiers***
    // @notice only owner of contract
    modifier onlyOwner()
    {
        require(msg.sender == owner);
        _;
    }
    // @notice only owner of aggregator contract
    modifier onlyAggregator()
    {
        require(msg.sender == aggregatorAddr);
        _;
    }
    // @notice modifier to confirm calling address is an oracle node address already acknowledged
    modifier oracleHasJoined()
    {
        require(oracles[msg.sender] == true);
        _;
    }
    // @notice modifier to confirm calling address is not yet an oracle node address acknowledged
    modifier oracleNotJoined()
    {
        require(oracles[msg.sender] != true);
        _;
    }
    // @notice requires request cancelFlag == 0
    modifier cancelFlagZero(
        uint256 _requestID)
    {
        require(requests[_requestID].cancelFlag == 0);
        _;
    }
    // @notice requires stakeBalance to equal exactly 1 ETH
    modifier stakeCheck()
    {
        require(stakeBalance[msg.sender] == 1 ether);
        _;
    }
    // @notice requires the value sent to be >= MIN_FEE set
    modifier valueGTMinFeeCheck()
    {
        // require msg.value has enough to pay oracles for each function call
        require(msg.value >= MIN_FEE);
        _;
    }
    // @notice requires an odd number of oracles so voting/aggregation doesnt become deadlocked
    modifier oddNumOracles(
        uint32 _numberOfOracles)
    {
        // makes sure the number that there is always an odd number of oracles so a majority can be reached.
        require(!(_numberOfOracles % 2 == 0));
        _;
    }
    // @notice requires a new oracle to send 1 eth to be held by contract until it leaves the network
    modifier valueEqual1Eth()
    {
        require(msg.value == 1 ether); // must send 1 eth that is to be held by contract
        _;
    }
    // @notice requires counter to be less than number required
    modifier orcCountLessThanNum(
        uint256 _requestID)
    {
        // makes sure oracle counter is less than number of oracles required so response can be submitted
        require(requests[_requestID].oracleCounter < requests[_requestID].numberOfOracles);
        _;
    }
    // @notice requires that only the request can call the function
    modifier onlyRequester(
        uint256 _requestID)
    {
        require(msg.sender == requests[_requestID].requester);
        _;
    }

}