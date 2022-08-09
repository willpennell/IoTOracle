// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./OracleRequestContract.sol";


contract ReputationContract {
    address public owner; // owner of contract
    OracleRequestContract orc; // instance of ORC contract
    // global vars
    mapping (address => uint256 ) public oracleRatings;
    // rating values
    uint constant MIN_RATING = 1;
    uint constant MAX_RATING = 10;
    // penalty fee
    uint constant PENALTY_FEE = 35000;

    event OracleBlacklisted(address, string);

    constructor(address _addr) {
        owner = msg.sender;
        // address of deployed OracleRequestContract
        orc = OracleRequestContract(_addr);
    }
    // @notice addOracle adds oracle to rating mapping to track behaviour
    function addOracle(address _addr)
    external
    onlyORC()
    returns(bool)
    {
        oracleRatings[_addr] = MIN_RATING;
        return true;
    }
    // @notice increments penalty rating as oracle address keeps being dishonest
    function incrementRating(address _addr)
    external
    onlyORC() // only oracle request contract can call this
    onlyNonBlacklistedOracles(_addr) // only non-blacklisted oracles are to be used
    returns(uint)
    {
        oracleRatings[_addr] ++; // increments the value in map at address
        // check if the oracle rating is greater than the max rating
        if (oracleRatings[_addr] > MAX_RATING) {
            // orc blacklist function
            orc.blacklistOracle(_addr);
            emit OracleBlacklisted(_addr, "oracle node has been blacklisted from network.");
        }
        return oracleRatings[_addr];
    }
    // ***getters***
    // @notice gets oracle rating
    function getOracleRating(address _addr)
    public
    view
    returns(uint256)
    {
        return oracleRatings[_addr];
    }
    // @notice gets oracle penalty fee
    function getPenaltyFee(address _addr)
    public
    view
    returns(uint256)
    {
        return PENALTY_FEE**oracleRatings[_addr];
    }
    // ***modifiers***
    // @notice only OracleRequestContract can call the function with modifier
    modifier onlyORC()
    {
        require(msg.sender == address(orc));
        _;
    }
    // @notice only non blacklisted oracles
    modifier onlyNonBlacklistedOracles(address _addr) {
        require(!orc.getBlacklistedOracle(_addr) == true);
        _;
    }

}
