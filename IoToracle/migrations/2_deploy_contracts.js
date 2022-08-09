var OracleRequest = artifacts.require("OracleRequestContract")
var Aggregator = artifacts.require("AggregatorContract")
var Reputation = artifacts.require("ReputationContract")

module.exports = async function (deployer) {
    deployer.then(async () => {
        await deployer.deploy(OracleRequest);
        await deployer.deploy(Aggregator, OracleRequest.address);
        await deployer.deploy(Reputation, OracleRequest.address);
        let OracleRequestInstance = await OracleRequest.deployed();
        await OracleRequestInstance.setAggregatorContract(Aggregator.address);
        await OracleRequestInstance.setReputationContract(Reputation.address);
    });

}