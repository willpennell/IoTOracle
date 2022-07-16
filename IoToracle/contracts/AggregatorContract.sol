// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./OracleRequesterContract.sol";

contract AggregatorContract {

    // need an answers struct
    OracleRequesterContract orc;
    //orc = OracleRequesterContract() // address of deployed contract
    mapping (uint256 => Answer[]) answers;
    struct Answer {
        uint256 requestID;
        bytes result;
        mapping (address => bool) oracleHasSubmitted;

    }

    constructor() {
        orc = OracleRequesterContract();
    }



 /*   function receiveResponse(uint256 _requestID, bytes _actualResult) onlyAuthorisedOraclesYetToVote(_requestID) returns(bool){
        //answers[_requestID].;
        return _;
    }*/


    modifier onlyAuthorisedOraclesYetToVote(uint _requestID) {
        require(orc.requests[_requestID].oracles[msg.sender]);
        require(answers[_requestID].oracleHasSubmitted[msg.sender] == false);
        _;
    }

    //modifier oracleHasNotVoted() {


}
