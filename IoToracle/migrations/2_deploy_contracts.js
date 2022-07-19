var OracleRequester = artifacts.require("OracleRequesterContract")
var Aggregator = artifacts.require("AggregatorContract")

module.exports = function(deployer) {
    deployer.then(async () => {
        await deployer.deploy(OracleRequester);
        await deployer.deploy(Aggregator, OracleRequester.address);
    });
}