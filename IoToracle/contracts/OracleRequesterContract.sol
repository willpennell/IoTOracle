pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";


contract OracleRequesterContract is Ownable{
    //

    address owner;

    mapping (uint => Request) public requests;
    address[] public oracles;
    uint requestCounter = 1;
    uint pendingCounter = 1;
    struct Request {
        uint256 requestID;
        address callbackAddress;
        bytes callbackFID;
        bytes IoTID;
        bytes dataType;
        bytes32 pHash;
        bytes requiredResult;
        uint32 numberOfOracles;
        address[] oracles;
    }

    function createRequest(address _callbackAddress, bytes _IoTID, ){

    }








}
