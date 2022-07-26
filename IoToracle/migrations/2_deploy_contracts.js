var OracleRequest = artifacts.require("OracleRequestContract")
var Aggregator = artifacts.require("AggregatorContract")

module.exports = async function (deployer) {
    deployer.then(async () => {
        await deployer.deploy(OracleRequest);
        await deployer.deploy(Aggregator, OracleRequest.address);
        let OracleRequestInstance = await OracleRequest.deployed();
        await OracleRequestInstance.setAggregatorContract(Aggregator.address);
    });

}