// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract OracleRequesterContract {
    //

    address payable public owner;

    Request[100] public requests;
    bool[100] public pendingRequests;
    Request[100] public completedRequests;
    mapping (address => bool) public oracles;
    uint requestCounter = 1;
    uint pendingCounter = 1;
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
    }

    event OpenForBids(uint256, bytes);
    event BidPlaced(address);
    event ReleaseRequestDetails(uint256, bytes, bytes);

    constructor() {
        owner = payable(msg.sender);

    }

    fallback() external{}

    function joinAsOracle() public oracleNotJoined() returns(bool){
        oracles[msg.sender] = true;
        return oracles[msg.sender];
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
        uint32 _numberOfOracles) public payable returns(uint256) {
        // assign id to request
        uint256 requestId = requestCounter;
        // log details of the request to array to cross reference later
        requests[requestId].requestID = requestId;
        requests[requestId].requester = msg.sender;
        requests[requestId].callbackAddress = _callbackAddress;
        requests[requestId].callbackFID = _callbackFID;
        requests[requestId].IoTID = _IoTID;
        requests[requestId].dataType = _dataType;
        //requests[requestId].requiredResult = _requiredResult;
        requests[requestId].numberOfOracles = _numberOfOracles;
        //requests[requestId].oracles;
        requests[requestId].oracleCounter = 0;


        // create a hash of the _dataType and _requiredResult:
        // we will use the actual result to test upon delivery.
        bytes32 phash = keccak256(abi.encodePacked(_dataType, _requiredResult));
        requests[requestId].pHash = phash;
        // increment requestId ready for next create request function.
        pendingRequests[requestId] = true;
        requestCounter++;

        emit OpenForBids(requestId, requests[requestId].dataType);

        return requestId;
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
    // TODO: MAKE SURE TO SET AGGREGATOR CONTRACT TO ONLY CALL THIS FUNCTION
    function deliverResponse(uint256 _requestID, bytes memory _finalResult) public {
        // return single aggregation result
    }

    // getters
    function getOracleForRequest(uint256 _requestID, address _orcAddr) public view returns(bool){
        return requests[_requestID].oraclesForRequest[_orcAddr];
    }
    function getNumberOfOracles(uint256 _requestID) public view returns(uint32){
        return requests[_requestID].numberOfOracles;
    }
    function getPHash(uint256 _requestID) public view returns(bytes32){
        return requests[_requestID].pHash;
    }
    function getDataType(uint256 _requestID) public view returns(bytes memory) {
        return requests[_requestID].dataType;
    }
    function moveCompletedRequest(uint256 _requestID) internal {
        completedRequests[_requestID] = requests[_requestID];

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
}
