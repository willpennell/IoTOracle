const Web3 = require('web3');
const util = require('util')
const MyContract = require('../build/contracts/IoTOracleContract.json');
// iot id
const iotid = "rp4200938820xx"
const iotidEncode = Web3.utils.asciiToHex(iotid)

console.log(iotidEncode)
const init = async () => {
    const web3 = new Web3('http://localhost:8546');
    const addr1 = "0xa4F40B53B37eA6471df700E34F0dd5eb9253b98D"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];
    const contract = new web3.eth.Contract(
        MyContract.abi,
        deployedNetwork.address
    )
    contract.methods.makeRequest(
        addr1,
        iotidEncode,
        true,
    ).send({from: addr1, gas: 1000000})
    console.log("Pushed to Blockchain")
}


init();