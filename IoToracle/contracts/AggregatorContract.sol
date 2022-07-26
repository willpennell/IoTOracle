// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./OracleRequestContract.sol";

contract AggregatorContract {
    address public owner; // owner of contract
    OracleRequestContract orc; // instance of ORC contract
    Answer[100] public answers; // array of 100 Answer structs to keep track of answers
    struct Answer {
        uint256 requestID; // the request ID
        bytes dataType; // data type of request 'bool' 'temp' etc
        mapping (address => bool) oracleHasSubmitted; // authorised oracles, marked true when voted
        mapping (address => bytes) oracleResults; // add their result to a map with their address as index
        mapping (uint => address) oracleAddresses; // uint to address to match counter and iterate through
        uint256 oracleCounter; // keep an index of the oracles
        address[] correctOracles; // dynamic array to keep track of correct oracle address to be paid
        uint32 t; // a true counter;
        uint32 f; // a false counter: we need this under (n-1)/3
        uint cancelFlag; // cancelFlag == 0, == 1 if canceled
        bool lastOracle; // last oracle
    }

    event ResponseReceived(address, string); // emits logs of each response received
    event AggregationCompleted(uint256, string); // emits logs once the aggregation is complete
    event LogHashes(bytes32, bytes32); // emits logs of hashes
    // @notice constructor called when deployed
    constructor(address _addr)
    {
        owner = msg.sender; // set deploying address as owner
        orc = OracleRequestContract(_addr); // initialise the Oracle Request Contract contract
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
    onlyAuthorisedOraclesYetToVote(_requestID) // makes sure only oracles that are authorised call call this function,
    // also only if it has not yet submitted an answer
    cancelFlagCheck(_requestID) // requires cancelFlag == 0
    returns(bool)
    {
        // only logs answers for any number of oracles that are under the number required
        if (answers[_requestID].oracleCounter < orc.getNumberOfOracles(_requestID)) {
            answers[_requestID].requestID = _requestID; // assign correct id
            answers[_requestID].oracleResults[msg.sender] = _actualResult; // orcale nodes answer mapped to its address
            answers[_requestID].oracleHasSubmitted[msg.sender] = true; // it has now submitted.
            answers[_requestID].oracleAddresses[answers[_requestID].oracleCounter] = msg.sender; // map counter to oracle addr
            answers[_requestID].oracleCounter++; // once this is equal we know to enter next if statement
            answers[_requestID].cancelFlag = 0; // sets cancelFlag to 0
            emit ResponseReceived(msg.sender, "response logged"); // emits to listeners that a response has been logged
        }
        // if the counter has reached the number of oracles needed, on the last oracle,
        // we can call the deliverResponse with the aggregation() method with returns a single result
        if(answers[_requestID].oracleCounter == orc.getNumberOfOracles(_requestID)) {
            answers[_requestID].lastOracle = true; // flag to know all oracles have submitted a vote
            // call aggregator function, returns a final result and an array of correct oracles
            (bytes memory finalResult, address[] memory correctOracles) = aggregation(_requestID);
            orc.deliverResponse(_requestID, finalResult, correctOracles); // call the deliverResponse function
        }
        return answers[_requestID].oracleHasSubmitted[msg.sender]; // returns true
    }

    // @notice aggregates results on a voting system, also requires that false votes are not larger than a threshold
    // @dev
    // @param _requestID pass in requestID to index the answers
    // @return final result
    function aggregation(
        uint256 _requestID)
    internal
    cancelFlagCheck(_requestID) // requires cancelFlag == 0
    returns(
        bytes memory,
        address[] memory)
    {
        // local finalResult initialised
        bytes memory finalResult;
        //gets dataType from ORC request struct which is then used for a hash of the params to match the results
        answers[_requestID].dataType = orc.getDataType(_requestID);
        // loop through the oracle counter to get oracle address and determine their result
        for (uint i=0; i<answers[_requestID].oracleCounter; i++){
            address oracleAddress = answers[_requestID].oracleAddresses[i]; // assign oracleAddress
            // ensures actual result and required result are the same
            bytes32 responseHash = keccak256(abi.encodePacked(
                    answers[_requestID].dataType, answers[_requestID].oracleResults[oracleAddress]
                )); // response hash from answer
            emit LogHashes(orc.getPHash(_requestID), responseHash); // emits to listeners to 2 hashes, these should equal
            if (responseHash == orc.getPHash(_requestID)){
                answers[_requestID].t++; // if correct increment the true count
                answers[_requestID].correctOracles.push(oracleAddress); // adds correct oracles to an array
                if (finalResult.length <= 0) {
                    // first response will assign its answer to the actual answer
                    finalResult = answers[_requestID].oracleResults[oracleAddress]; // as these are all the same
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

    function cancelRequest(
        uint256 _requestID)
    public
    cancelFlagCheck(_requestID)
    onlyORC() // only ORC contract can call this function
    returns(bool)
    {
        answers[_requestID].cancelFlag = 1;
        return true;
    }
    // ***modifiers***
    modifier onlyAuthorisedOraclesYetToVote(
        uint256 _requestID)
    {
        // checks that the oracle is in request list
        require(orc.getOracleForRequest(_requestID, msg.sender));
        // checks that the oracle hasn't yet voted.
        require(answers[_requestID].oracleHasSubmitted[msg.sender] == false);
        _;
    }
    modifier onlyOwner()
    {
        require(msg.sender == owner);
        _;
    }
    modifier cancelFlagCheck(
        uint256 _requestID)
    {
        require(answers[_requestID].cancelFlag == 0);
        _;
    }
    modifier onlyORC()
    {
        require(msg.sender == address(orc));
        _;
    }

}






