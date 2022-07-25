// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./OracleRequesterContract.sol";

contract AggregatorContract {

    address payable public owner;
    // need an answers struct
    OracleRequesterContract orc;
    //orc = OracleRequesterContract() // address of deployed contract
    Answer[100] public answers; // array of 100 Answer structs
    struct Answer {
        uint256 requestID; // the request ID
        bytes dataType; // data type of request 'bool' 'temp' etc
        mapping (address => bool) oracleHasSubmitted; // authorised oracles, marked true when voted
        mapping (address => bytes) oracleResults; // add their result to a map of their address
        mapping (uint => address) oracleAddresses; // uint to address to match counter and iterate through
        address[] correctOracles;
        uint256 oracleCounter; // keep an index of the oracles
        uint32 t; // a true counter;
        uint32 f; // a false counter: we need this under (n-1)/3
        uint cancelFlag;
        bool lastOracle;
    }

    event ResponseReceived(address, string);
    event AggregationCompleted(uint256, string);
    event LogHashes(bytes32, bytes32);

    constructor(address _addr) {
        // initialise the Oracle Request Contract contract
        orc = OracleRequesterContract(_addr);
    }

    // @notice each oracle sends their response to through this function
    // filling up the Answer struct at index requestID.
    // @dev
    // @param _requestID ID of request
    // @param _actualResult the value the off-chain oracle has collected
    // @return bool, true to say answer has been submitted
    function receiveResponse(
        uint256 _requestID,
        bytes memory _actualResult
    )
    public
    onlyAuthorisedOraclesYetToVote(_requestID) cancelFlagCheck(_requestID)
    returns(bool){
        if (answers[_requestID].oracleCounter < orc.getNumberOfOracles(_requestID)) {
            answers[_requestID].requestID = _requestID;
            answers[_requestID].oracleResults[msg.sender] = _actualResult;
            answers[_requestID].oracleHasSubmitted[msg.sender] = true;
            answers[_requestID].oracleAddresses[answers[_requestID].oracleCounter] = msg.sender;
            answers[_requestID].oracleCounter++;
            answers[_requestID].cancelFlag = 0;
        }
        // if the counter has reached the number of oracles needed,
        // we can call the deliverResponse with the aggregation() method with returns a single result
        if(answers[_requestID].oracleCounter == orc.getNumberOfOracles(_requestID)) {
            // call aggregator function
            emit ResponseReceived(msg.sender, "response logged");
            answers[_requestID].lastOracle = true;
            (bytes memory finalResult, address[] memory correctOracles) = aggregation(_requestID);
            orc.deliverResponse(_requestID, finalResult, correctOracles);
        }
        return answers[_requestID].oracleHasSubmitted[msg.sender];
    }

    // @notice aggregates results on a voting system, also requires that false votes are not larger than a threshold
    // @dev
    // @param _requestID pass in requestID to index the answers
    // @return final result
    function aggregation(uint256 _requestID) internal cancelFlagCheck(_requestID) returns(bytes memory, address[] memory) {
        // local finalResult initialised
        bytes memory finalResult;
        //gets dataType from ORC request struct which is then used for a hash of the params to match the results
        answers[_requestID].dataType = orc.getDataType(_requestID);
        for (uint i=0; i<answers[_requestID].oracleCounter; i++){
            address oracleAddress = answers[_requestID].oracleAddresses[i];
            // ensures actual result and required result are the same
            bytes32 responseHash = keccak256(abi.encodePacked(
                    answers[_requestID].dataType, answers[_requestID].oracleResults[oracleAddress]
                ));
            emit LogHashes(orc.getPHash(_requestID), responseHash);
            if (responseHash == orc.getPHash(_requestID)){
                // if correct increment the true count
                answers[_requestID].t++;
                answers[_requestID].correctOracles.push(oracleAddress);
                if (finalResult.length <= 0) {
                    // first response will assign its answer to the actual answer
                    finalResult = answers[_requestID].oracleResults[oracleAddress];
                }
            } else {
                //otherwise increment the false count
                answers[_requestID].f++;
            }
        }
        // requires a threshold for false answers
        require(!(answers[_requestID].f > (answers[_requestID].oracleCounter - 1) / 3), 'too many incorrect nodes');

        emit AggregationCompleted(_requestID, "aggregation complete");
        return (finalResult, answers[_requestID].correctOracles);
    }
   /* function getHashes(uint256 _requestID, address _oracleAddress) public view returns(bytes32, bytes32) {
        bytes32 ogHash = orc.getPHash(_requestID);
        bytes32 orcHash = keccak256(abi.encodePacked(
                answers[_requestID].dataType, answers[_requestID].oracleResults[_oracleAddress]
        ));
        return (ogHash, orcHash);
    }
    function getResultInfo(uint256 _requestID, address _oracleAddress) public view returns(bytes memory, bytes memory){
        bytes memory dataType = orc.getDataType(_requestID);
        bytes memory answer = answers[_requestID].oracleResults[_oracleAddress];
        return (dataType, answer);
    }*/

    modifier onlyAuthorisedOraclesYetToVote(uint256 _requestID) {
        // checks that the oracle is in request list
        require(orc.getOracleForRequest(_requestID, msg.sender));
        // checks that the oracle hasn't yet voted.
        require(answers[_requestID].oracleHasSubmitted[msg.sender] == false);
        _;
    }

    function cancelRequest(uint256 _requestID) public cancelFlagCheck(_requestID) returns(bool){
        answers[_requestID].cancelFlag = 1;
        orc.cancelRequest(_requestID);
        return true;
    }

    modifier cancelFlagCheck(uint256 _requestID) {
        require(answers[_requestID].cancelFlag == 0);
        _;
    }

}






