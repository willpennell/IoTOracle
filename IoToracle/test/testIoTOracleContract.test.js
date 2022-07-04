const IoTOracleContract = artifacts.require("IoTOracleContract")

contract("IoTOracleContract", (accounts) =>{
    let IoTOracleInstance;
    let expectedRequestID = 1;

    before(async () => {
        IoTOracleInstance = await IoTOracleContract.deployed();
    });

    it()


});