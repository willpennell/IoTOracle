const Web3 = require('web3');
const util = require('util')
const MyContract = require('../build/contracts/OracleRequestContract.json');

const init = async () => {
    const web3 = new Web3('http://localhost:8546');
    const aggAddr = "0xAF1CD1650f8c07f688901D444BE0305D35D39BFf"
    const orcAddr = "0x0772d4c59f2D6B17Eaf26251EeEFf4873dBd6C56"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];
    const contract = new web3.eth.Contract(
        MyContract.abi,
        deployedNetwork.address
    )
    contract.methods.setAggregatorContract(
        aggAddr,
    ).send({from: orcAddr, gas: 3000000})
    console.log("Pushed to Blockchain")
}


init();