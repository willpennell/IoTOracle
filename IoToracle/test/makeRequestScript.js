const Web3 = require('web3');
const MyContract = require('../build/contracts/IoTOracleContract.json');


const iotid = 'rp42022'
const init = async () => {
    const web3 = new Web3('http://localhost:7545');
    const addr1 = "0xC19D4b3B2F17a08f6324d15D80E6F4d1F45C392e"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];
    const contract = new web3.eth.Contract(
        MyContract.abi,
        deployedNetwork.address
    )
    contract.methods.makeRequest(
        addr1,
        iotid,
        true,
    ).send({from: addr1, gas: 1000000})

}
init();