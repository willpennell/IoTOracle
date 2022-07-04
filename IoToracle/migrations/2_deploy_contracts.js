var IoTOracle = artifacts.require("IoTOracleContract");

module.exports = function(deployer) {
    deployer.deploy(IoTOracle);
}