const Web3 = require('web3');
const util = require('util')
const MyContract = require('../build/contracts/OracleRequesterContract.json');

// callbackFID
const cbfd = '{"function":"Test1"}'
const callbackFID = Web3.utils.asciiToHex(cbfd)
// iot id
const iotid = 'rp41992'
const iotidEncode = Web3.utils.asciiToHex(iotid)
// dataType
const dt = '{"type": bool, "topic": "PIRSensor", "tAfter": null, "tBefore": null}'
const dataTypeEncode = Web3.utils.asciiToHex(dt)

//const dataTypeEncode = Web3.utils.asciiToHex()

// requiredResult
let rr = '{"result": true}'
const requiredResultEncode = Web3.utils.asciiToHex(rr)
console.log(requiredResultEncode)


const init = async () => {
    const web3 = new Web3('http://localhost:8546');
    const addr1 = "0x0db4b258DfEC24FDAC43b75ee2fEAe1E214F7667"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];
    const contract = new web3.eth.Contract(
        MyContract.abi,
        deployedNetwork.address
    )
    contract.methods.createRequest(
        addr1,
        callbackFID,
        iotidEncode,
        dataTypeEncode,
        requiredResultEncode,
        1
    ).send({from: addr1, gas: 3000000, value: 10**11})
    console.log("Pushed to Blockchain")
}


init();