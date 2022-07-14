pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";


contract OracleRequesterContract is Ownable{
    //

    address owner;
    Request[100] public requests;
    address[] public oracles;
    uint requestCounter = 1;
    uint pendingCounter = 1;
    struct Request {
        uint256 requestID;
        address requester;
        address callbackAddress;
        bytes callbackFID;
        bytes IoTID;
        bytes dataType;
        bytes32 pHash;
        bytes requiredResult;
    }

    function OracleRequesterContract() public payable onlyOwner() {
    }

    function placeBid() public returns(bool) {
        if (msg.sender )
        return true;
    }


}
