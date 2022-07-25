const Web3 = require('web3');
const util = require('util')
const MyContract = require('../build/contracts/OracleRequesterContract.json');
const AggCon = require('../build/contracts/AggregatorContract.json')


const init = async () => {
    const web3 = new Web3('http://localhost:8546');
    const addr1 = "0x0db4b258DfEC24FDAC43b75ee2fEAe1E214F7667"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];

    const contract = new web3.eth.Contract(
        AggCon.abi,
        deployedNetwork.address
    )
    contract.methods.cancelRequest(
        1
    ).send({from: addr1, gas: 3000000})
}


init();