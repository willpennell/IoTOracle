// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract OracleRequesterContract {
    //

    address payable public owner;

    Request[100] public requests;
    mapping (address => bool) public oracles;
    uint requestCounter = 1;
    uint pendingCounter = 1;
    enum Status{NOTSSTARTED, PENDING, COMPLETE}
    // a single Request structure
    struct Request {
        uint256 requestID;
        address requester;
        address callbackAddress;
        bytes callbackFID;
        bytes IoTID;
        bytes dataType;
        bytes32 pHash;
        //bytes requiredResult;
        uint32 numberOfOracles;
        mapping (address => bool) oraclesForRequest;
        uint32 oracleCounter;
        Status status;
    }

    event OpenForBids(uint256, bytes);
    event BidPlaced(address);
    event ReleaseRequestDetails(uint256, bytes, bytes);
    event StatusChange(Status, string);
    event OracleJoined(address, string);
    event OracleLeft(address, string);

    constructor() {
        owner = payable(msg.sender);

    }

    fallback() external{}

    function joinAsOracle() public oracleNotJoined() returns(bool){
        oracles[msg.sender] = true;
        emit OracleJoined(msg.sender, "Welcome new node!");
        return oracles[msg.sender];
    }

    function leaveOracleNetwork() public oracleHasJoined() returns(bool) {
        delete oracles[msg.sender];
        emit OracleLeft(msg.sender, "Node has left, Goodbye");
        return true;
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
        uint32 _numberOfOracles) public returns(uint256) {
        // assign id to request
        uint256 requestID = requestCounter;
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
        emit StatusChange(requests[requestID].status, "PENDING");


        // create a hash of the _dataType and _requiredResult:
        // we will use the actual result to test upon delivery.
        bytes32 phash = keccak256(abi.encodePacked(_dataType, _requiredResult)); // add required result to hash
        requests[requestID].pHash = phash;
        // increment requestId ready for next create request function.

        requestCounter++;

        emit OpenForBids(requestID, requests[requestID].dataType);

        return requestID;
    }

    // @notice off-chain oracle node places bid to fetch data, once threshold is reached, emits event details.
    // @dev
    // @param requestId oracle node wishes to bid on
    // @return true to indicate bid successful
    function placeBid(uint256 _requestID) public oracleHasJoined() returns(bool)  {
        require(requests[_requestID].oracleCounter < requests[_requestID].numberOfOracles);
        requests[_requestID].oraclesForRequest[msg.sender] = true;
        requests[_requestID].oracleCounter++;
        // also add a timeout
        emit BidPlaced(msg.sender);
        if (requests[_requestID].numberOfOracles == requests[_requestID].oracleCounter) {
            emit ReleaseRequestDetails(requests[_requestID].requestID,
                requests[_requestID].IoTID,
                requests[_requestID].dataType);
            return true;
        }
        return oracles[msg.sender];
    }

    //

    // function deliveryResponse() {} returns result to user smart contract
    function deliverResponse(uint256 _requestID, bytes memory _finalResult) public {
        // return single aggregation result
        _finalResult;
        requests[_requestID].status = Status.COMPLETE;
    emit StatusChange(requests[_requestID].status, "COMPLETE");
    }

    function cancelRequest(uint256 _requestID) external cancelCheck(_requestID){
        delete requests[_requestID];
        requestCounter = 1;

    }

    // getters
    function getOracleForRequest(uint256 _requestID, address _orcAddr) external view returns(bool){
        return requests[_requestID].oraclesForRequest[_orcAddr];
    }
    function getNumberOfOracles(uint256 _requestID) external view returns(uint32){
        return requests[_requestID].numberOfOracles;
    }
    function getPHash(uint256 _requestID) external view returns(bytes32){
        return requests[_requestID].pHash;
    }
    function getDataType(uint256 _requestID) external view returns(bytes memory) {
        return requests[_requestID].dataType;
    }

    // @notice modifier to confirm calling address is an oracle node address already acknowledged
    modifier oracleHasJoined() {
        require(oracles[msg.sender] == true);
        _;
    }
    // @notice modifier to confirm calling address is not yet an oracle node address acknowledged
    modifier oracleNotJoined() {
        require(oracles[msg.sender] != true);
        _;
    }
    //
    modifier cancelCheck(uint256 _requestID) {
        require(requests[_requestID].callbackAddress == msg.sender);
        _;
    }
}
