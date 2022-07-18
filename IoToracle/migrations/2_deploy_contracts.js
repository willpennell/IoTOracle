var OracleRequester = artifacts.require("OracleRequesterContract")

module.exports = function(deployer) {
    deployer.deploy(OracleRequester);
}