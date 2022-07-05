const IoTOracleContract = artifacts.require("IoTOracleContract")

contract("IoTOracleContract", (accounts) =>{
    let IoTOracleInstance;
    let expectedRequestID = 1;
    let callbackAddr = accounts[1];
    let iotId = 'rp42022';
    let requiredResult = true;

    before(async () => {
        IoTOracleInstance = await IoTOracleContract.deployed();
    });
    it("should make call to IoTOracle.makeRequest() and return ID value: 1", async () => {

        const returnedRequestID = await IoTOracleInstance.makeRequest.call(
            callbackAddr, iotId, requiredResult, {from: callbackAddr}
        );
        assert.equal(returnedRequestID, expectedRequestID, "These should be equal");
    });
    it("Should emit Event RequestInfo and match expected outputs.", async () => {
        const event = await IoTOracleInstance.makeRequest.sendTransaction(
            callbackAddr, iotId, requiredResult, {from: callbackAddr}
        );
        assert.equal(event.logs.length, 1, "RequestInfo should have been emitted");
        assert.equal(event.logs[0].args.requester.valueOf(), callbackAddr);
        assert.equal(event.logs[0].args.requiredResult.valueOf(), true);
        assert.equal(event.logs[0].args.id.valueOf(), expectedRequestID);
    });







});