const Web3 = require("web3");
const OracleRequestContract = artifacts.require("OracleRequestContract")
const AggregatorContract = artifacts.require("AggregatorContract")
const ReputationContract = artifacts.require("ReputationContract")

// callbackFID
const cbfd = '{"function":"Test1"}'
const callbackFID = Web3.utils.asciiToHex(cbfd)
const iotid = 'rp41992'
const iotidEncode = Web3.utils.asciiToHex(iotid)
const dt = '{"topic": "PIRSensor", "tAfter": 1658686070, "tBefore": 1658858840}'
const dataTypeEncode = Web3.utils.asciiToHex(dt)
const aggregationType = 1;

contract("OracleRequestContract, AggregationContract, ReputationContract", (accounts) =>{
    let orcInstance;
    let aggInstance;
    let repInstance;
    let node1 = accounts[9];
    let node2 = accounts[1];
    let node3 = accounts[2];
    let usersc = accounts[4];




    before(async () => {
        orcInstance = await OracleRequestContract.deployed();
        aggInstance = await AggregatorContract.deployed(orcInstance.address);
        repInstance = await ReputationContract.deployed(orcInstance.address);
        await orcInstance.setAggregatorContract(aggInstance.address)
        await orcInstance.setReputationContract(repInstance.address)
    });
    it("oracles be able to join the network", async () => {

        await orcInstance.joinAsOracle.call({from: node1, value: 1**20})

    });
});