// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IoTOracleContract {
    address owner;
    Request[100] public requests;
    uint256 requestIdCounter = 1;
    // uint64 unrespondedRequests = 0;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    constructor() {
        owner = msg.sender;
        requests[0].requester = owner;
    }

    struct Request {
        uint256 requestId;
        address requester;
        address callbackAddress;
        //bytes callbackFuncID;
        bytes IoTId;
        bool requiredResult;
        bytes32 pHash;
    }

    event RequestIoTInfo(
        uint id,
        bool requiredResult,
        bytes IotId,
        address requester,
        address callbackAddress);

    event Delivery();

    function makeRequest(
        address _callbackAddress,
        //bytes memory _callbackFuncId,
        bytes memory _IoTId,
        bool _requiredResult) public returns(uint256){
        // assign id to request
        uint256 reqId = requestIdCounter;

        // log details of the request to array to cross reference later
        requests[reqId].requestId = reqId;
        requests[reqId].requester = msg.sender;
        requests[reqId].callbackAddress = _callbackAddress;
        //requests[reqId].callbackFuncID = _callbackFuncId;
        requests[reqId].IoTId = _IoTId;
        requests[reqId].requiredResult = _requiredResult;

        bytes32 pHash = keccak256(abi.encodePacked(reqId, address(msg.sender), _IoTId, _requiredResult));
        requests[reqId].pHash = pHash;

        // push to eventListener to then pass to off-chain core to process
        emit RequestIoTInfo(reqId, _requiredResult, _IoTId, requests[reqId].requester, _callbackAddress);
        // increment Id counter
        requestIdCounter++;

        // return the id of request
        return reqId;
    }

    /*function deliverResponse() {
    }*/
}
