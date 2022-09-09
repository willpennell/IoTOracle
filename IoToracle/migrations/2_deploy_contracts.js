var OracleRequest = artifacts.require("OracleRequestContract")
var Aggregator = artifacts.require("AggregatorContract")
var Reputation = artifacts.require("ReputationContract")

module.exports = async function (deployer, network, accounts) {
    deployer.then(async () => {
        console.log(accounts);
        await deployer.deploy(OracleRequest, {from: accounts[0]});
        await deployer.deploy(Aggregator, OracleRequest.address, {from: accounts[0]});
        await deployer.deploy(Reputation, OracleRequest.address, {from: accounts[0]});
        let OracleRequestInstance = await OracleRequest.deployed();
        await OracleRequestInstance.setAggregatorContract(Aggregator.address);
        await OracleRequestInstance.setReputationContract(Reputation.address);
    });

}