// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./OracleRequestContract.sol";


contract ReputationContract {
    address public owner; // owner of contract
    // global vars
    mapping (address => uint256 ) public oracleRatings;
    // rating values
    uint constant MIN_RATING = 1;
    uint constant MAX_RATING = 10;
    // penalty fee
    uint constant PENALTY_FEE = 350000;

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
    oracleNotJoined(_addr)
    returns(bool)
    {
        oracleRatings[_addr] = MIN_RATING;
        return true;
    }

    function incrementRating(address _addr)
    external
    onlyORC()
    onlyNonBlacklistedOracles(_addr)
    returns(uint)
    {
        oracleRatings[_addr] ++;
        if (oracleRatings[_addr] > MIN_RATING) {
            // orc blacklist function
            orc.blacklistOracle(_addr);
            emit OracleBlacklisted(_addr, "oracle node has been blacklisted from network.");
            return false;
        }
        return oracleRatings[_addr];
    }
    // @notice


    // ***getters***
    function getOracleRating(address _addr)
    public
    view
    returns(uint256)
    {
        return oracleRatings[_addr];
    }

    function getPenaltyFee(address _addr)
    public
    external
    returns(uint256)
    {
        return PENALTY_FEE**oracleRatings[_addr];
    }

    // ***modifiers***

    // @notice modifier to confirm calling address is not yet an oracle node address acknowledged
    modifier oracleNotJoined(address _addr)
    {
        require(oracleRatings[_addr] != true);
        _;
    }
    // @notice only OracleRequestContract can call the function with modifier
    modifier onlyORC()
    {
        require(msg.sender == address(orc));
        _;
    }
    // @notice only non blacklisted oracles
    modifier onlyNonBlacklistedOracles(address _addr) {
        require(!orc.getBlacklistedOracle[_addr] == true);
        _;
    }

}
