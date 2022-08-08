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
        mapping (address => bool) oracleHasSubmittedCommit; // authorised oracles, marked true when committed
        mapping(address => bool) oracleHasSubmittedReveal; // authorised oracles, marked true when revealed
        mapping (address => bytes32) oracleCommits; // add their result to a map with their address as index
        mapping (address => bool) oracleVoteReveals;
        mapping (address => int256) oracleAverageReveals;
        mapping (uint256 => address) oracleAddresses; // uint to address to match counter and iterate through
        uint256 oracleCounter; // keep an index of the oracles
        address[] correctRevealOracles; // dynamic array to keep track of correct oracle address to be paid
        address[] incorrectRevealOracles;

        uint32 t; // a true counter;
        uint32 f; // a false counter: we need this under (n-1)/3
        uint cancelFlag; // cancelFlag == 0, == 1 if canceled
        uint commitsFlag;
        uint revealsFlag;
        bool lastOracle; // last oracle

    }

    event ResponseReceived(address, string); // emits logs of each response received
    event AggregationCompleted(uint256, string); // emits logs once the aggregation is complete
    event LogHashes(bytes32, bytes32); // emits logs of hashes
    event CommitsPlaced(uint, string);
    // @notice constructor called when deployed
    constructor(address _addr)
    {
        owner = msg.sender; // set deploying address as owner
        orc = OracleRequestContract(_addr); // initialise the Oracle Request Contract contract
    }
    // @notice commitResponse
    function commitResponse(uint256 _requestID, bytes32 _commitHash )
    public
    onlyAuthorisedOraclesYetToCommit(_requestID)
    cancelFlagCheck(_requestID)
    commitFlagCheck(_requestID)
    returns(bool)
    {
        if (answers[_requestID].requestID == 0) {
            answers[_requestID].requestID = _requestID; // assign correct id
        }
        if (answers[_requestID].oracleCounter < orc.getNumberOfOracles(_requestID)) {
            answers[_requestID].oracleCommits[msg.sender] = _commitHash; // oracle nodes add commit to map
            answers[_requestID].oracleHasSubmittedCommit[msg.sender] = true; // it has now submitted.
            answers[_requestID].oracleAddresses[answers[_requestID].oracleCounter] = msg.sender; // map counter to oracle addr
            answers[_requestID].oracleCounter++; // once this is equal we know to enter next if statement
        }
        if (answers[_requestID].oracleCounter == orc.getNumberOfOracles(_requestID)) {
            answers[_requestID].commitsFlag = 1;
            answers[_requestID].oracleCounter = 0;
            emit CommitsPlaced(_requestID, "commits are placed");
        }
        return true;
    }

    function revealVoteResponse(uint256 _requestID, bool _result, bytes[] memory _secret)
    public
    onlyAuthorisedOraclesYetToReveal(_requestID)
    cancelFlagCheck(_requestID)
    commitCompleteFlagCheck(_requestID) // checks all commits are in and flag is equal to 1
    revealsFlagCheck(_requestID) // checks that reveal flag is 0 otherwise function will not run
    onlyVoteAggregationType(_requestID)
    returns(bool){
        // need to check the hash of the _result and _secret against answers[_requestID].oracleCommits[msg.sender]
        bytes32 revealHash = keccak256(abi.encode(_secret, _result));
        if (answers[_requestID].oracleCommits[msg.sender] == revealHash) {
            answers[_requestID].oracleVoteReveals[msg.sender] = _result;
            answers[_requestID].oracleHasSubmittedReveal[msg.sender] = true;
            answers[_requestID].correctRevealOracles.push(msg.sender);
            // return answers[_requestID].oracleHasSubmittedReveal[msg.sender];
        } else {
            answers[_requestID].incorrectRevealOracles.push(msg.sender);
            answers[_requestID].oracleHasSubmittedReveal[msg.sender] = true;
            //return false;
        }
        answers[_requestID].oracleCounter++;
        if (answers[_requestID].oracleCounter == (orc.getNumberOfOracles(_requestID) - 1)){
            voteAggregation(_requestID);
        }
        return true;
    }

    function revealAverageResponse(uint256 _requestID, int256 _result, bytes[] memory _secret)
    public
    onlyAuthorisedOraclesYetToReveal(_requestID)
    cancelFlagCheck(_requestID)
    commitCompleteFlagCheck(_requestID) // checks all commits are in and flag is equal to 1
    revealsFlagCheck(_requestID) // checks that reveal flag is 0 otherwise function will not run
    onlyAverageAggregationType(_requestID)
    returns(bool){
        // need to check the hash of the _result and _secret against answers[_requestID].oracleCommits[msg.sender]
        bytes32 revealHash = keccak256(abi.encode(_secret, _result));
        if (answers[_requestID].oracleCommits[msg.sender] == revealHash) {
            answers[_requestID].oracleAverageReveals[msg.sender] = _result;
            answers[_requestID].oracleHasSubmittedReveal[msg.sender] = true;
            answers[_requestID].correctRevealOracles.push(msg.sender);
            // return answers[_requestID].oracleHasSubmittedReveal[msg.sender];
        } else {
            answers[_requestID].incorrectRevealOracles.push(msg.sender);
            answers[_requestID].oracleHasSubmittedReveal[msg.sender] = true;
            //return false;
        }
        answers[_requestID].oracleCounter++;
        if (answers[_requestID].oracleCounter == (orc.getNumberOfOracles(_requestID) - 1)){
            averageAggregation(_requestID);
        }
        return true;
    }
    // voting aggregation
    function voteAggregation(uint256 _requestID) internal returns(bool) {
        address[] memory trueOracleResults;
        address[] memory falseOracleResults;
        bool _result;
        address orcAddress;
        for (uint i = 0; i < answers[_requestID].correctRevealOracles.length; i++) {
            orcAddress = answers[_requestID].correctRevealOracles[i];
            if (answers[_requestID].oracleVoteReveals[orcAddress] == true ) {
                trueOracleResults[answers[_requestID].t] = orcAddress;
                answers[_requestID].t++;
            } else {
                falseOracleResults[answers[_requestID].f] = orcAddress;
                answers[_requestID].f++;
            }
        }

        if (answers[_requestID].t > answers[_requestID].f) {
            // call
            require(!(answers[_requestID].f > (answers[_requestID].oracleCounter - 1) / 3), 'too many incorrect nodes');
            _result = true;
            orc.deliverVoteResponse(_requestID, _result, trueOracleResults,
                falseOracleResults,
                answers[_requestID].incorrectRevealOracles);
            return _result;
        } else if (answers[_requestID].f > answers[_requestID].t) {
            require(!(answers[_requestID].t > (answers[_requestID].oracleCounter - 1) / 3), 'too many incorrect nodes');
            _result = false;
            orc.deliverVoteResponse(_requestID, _result,
                falseOracleResults,
                trueOracleResults,
                answers[_requestID].incorrectRevealOracles);
            return _result;
        } else {
            orc.cancelRequestDueToDeadlock(_requestID, answers[_requestID].incorrectRevealOracles);
            answers[_requestID].cancelFlag = 1;
        }
        return false;
    }
    // averaging aggregation
    function averageAggregation(uint256 _requestID)
    internal
    returns(bool)
    {
        address orcAddress;
        int256 total = 0;
        int256 average;
        for (uint i=0; i<answers[_requestID].correctRevealOracles.length; i++){
            orcAddress = answers[_requestID].correctRevealOracles[i];
            total += answers[_requestID].oracleAverageReveals[orcAddress];
        }
        average = total / int(answers[_requestID].correctRevealOracles.length);
        orc.deliverAverageResponse(
                _requestID,
                average,
                answers[_requestID].correctRevealOracles,
                    answers[_requestID].incorrectRevealOracles);
        return true;
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
    modifier onlyAuthorisedOraclesYetToCommit(
        uint256 _requestID)
    {
        // checks that the oracle is in request list
        require(orc.getOracleForRequest(_requestID, msg.sender));
        // checks that the oracle hasn't yet committed.
        require(answers[_requestID].oracleHasSubmittedCommit[msg.sender] == false);
        _;
    }
    modifier onlyAuthorisedOraclesYetToReveal(
        uint256 _requestID)
    {
        // checks that the oracle is in request list
        require(orc.getOracleForRequest(_requestID, msg.sender));
        // checks that the oracle hasn't yet revealed.
        require(answers[_requestID].oracleHasSubmittedReveal[msg.sender] == false);
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
    modifier commitFlagCheck(uint256 _requestID) {
        require(answers[_requestID].commitsFlag == 0 );
        _;
    }
    modifier commitCompleteFlagCheck(uint256 _requestID) {
        require(answers[_requestID].commitsFlag == 1);
        _;
    }
    modifier revealsFlagCheck(uint256 _requestID){
        require(answers[_requestID].revealsFlag == 0);
        _;
    }
    modifier onlyORC()
    {
        require(msg.sender == address(orc));
        _;
    }
    modifier onlyVoteAggregationType(uint256 _requestID) {
        require(orc.getAggregationType(_requestID) == 1);
        _;
    }
    modifier onlyAverageAggregationType(uint256 _requestID) {
        require(orc.getAggregationType(_requestID) == 2);
        _;
    }
}










//    // not needed
//    // @notice each oracle sends their response to through this function
//    // filling up the Answer struct at index requestID.
//    // @dev
//    // @param _requestID ID of request
//    // @param _actualResult the value the off-chain oracle has collected
//    // @return bool, true to say answer has been submitted
//    function receiveResponse(
//        uint256 _requestID,
//        bytes memory _actualResult
//    )
//    public
//    onlyAuthorisedOraclesYetToVote(_requestID) // makes sure only oracles that are authorised call call this function,
//    // also only if it has not yet submitted an answer
//    cancelFlagCheck(_requestID) // requires cancelFlag == 0
//    returns(bool)
//    {
//        // only logs answers for any number of oracles that are under the number required
//        if (answers[_requestID].oracleCounter < orc.getNumberOfOracles(_requestID)) {
//            answers[_requestID].requestID = _requestID; // assign correct id
//            answers[_requestID].oracleCommits[msg.sender] = _actualResult; // orcale nodes answer mapped to its address
//            answers[_requestID].oracleHasSubmitted[msg.sender] = true; // it has now submitted.
//            answers[_requestID].oracleAddresses[answers[_requestID].oracleCounter] = msg.sender; // map counter to oracle addr
//            answers[_requestID].oracleCounter++; // once this is equal we know to enter next if statement
//            answers[_requestID].cancelFlag = 0; // sets cancelFlag to 0
//            emit ResponseReceived(msg.sender, "response logged"); // emits to listeners that a response has been logged
//        }
//        // if the counter has reached the number of oracles needed, on the last oracle,
//        // we can call the deliverResponse with the aggregation() method with returns a single result
//        if(answers[_requestID].oracleCounter == orc.getNumberOfOracles(_requestID)) {
//            answers[_requestID].lastOracle = true; // flag to know all oracles have submitted a vote
//            // call aggregator function, returns a final result and an array of correct oracles
//            (bytes memory finalResult, address[] memory correctOracles) = aggregation(_requestID);
//            orc.deliverResponse(_requestID, finalResult, correctOracles); // call the deliverResponse function
//        }
//        return answers[_requestID].oracleHasSubmitted[msg.sender]; // returns true
//    }

//    // @notice aggregates results on a voting system, also requires that false votes are not larger than a threshold
//    // @dev
//    // @param _requestID pass in requestID to index the answers
//    // @return final result
//    function aggregation(
//        uint256 _requestID)
//    internal
//    cancelFlagCheck(_requestID) // requires cancelFlag == 0
//    returns(
//        bytes memory,
//        address[] memory)
//    {
//        // local finalResult initialised
//        bytes memory finalResult;
//        //gets dataType from ORC request struct which is then used for a hash of the params to match the results
//        answers[_requestID].dataType = orc.getDataType(_requestID);
//        // loop through the oracle counter to get oracle address and determine their result
//        for (uint i=0; i<answers[_requestID].oracleCounter; i++){
//            address oracleAddress = answers[_requestID].oracleAddresses[i]; // assign oracleAddress
//            // ensures actual result and required result are the same
//            bytes32 responseHash = keccak256(abi.encodePacked(
//                    answers[_requestID].dataType, answers[_requestID].oracleResults[oracleAddress]
//                )); // response hash from answer
//            emit LogHashes(orc.getPHash(_requestID), responseHash); // emits to listeners to 2 hashes, these should equal
//            if (responseHash == orc.getPHash(_requestID)){
//                answers[_requestID].t++; // if correct increment the true count
//                answers[_requestID].correctOracles.push(oracleAddress); // adds correct oracles to an array
//                if (finalResult.length <= 0) {
//                    // first response will assign its answer to the actual answer
//                    finalResult = answers[_requestID].oracleResults[oracleAddress]; // as these are all the same
//                }
//            } else {
//                //otherwise increment the false count
//                answers[_requestID].f++;
//            }
//        }
//        // requires a threshold for false answers
//        require(!(answers[_requestID].f > (answers[_requestID].oracleCounter - 1) / 3), 'too many incorrect nodes');
//        emit AggregationCompleted(_requestID, "aggregation complete");
//        return (finalResult, answers[_requestID].correctOracles);
//    }






