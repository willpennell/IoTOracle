const Web3 = require('web3');
const util = require('util')
const MyContract = require('../build/contracts/OracleRequesterContract.json');


const init = async () => {
    const web3 = new Web3('http://localhost:8546');
    const addr1 = "0x9e9E756A9E1605cD9824438fc13D87022670E3E6"
    const id = await web3.eth.net.getId();
    const deployedNetwork = MyContract.networks[id];
    const contract = new web3.eth.Contract(
        MyContract.abi,
        deployedNetwork.address
    )
    let oracle = await contract.methods.oracles(addr1).call(
    );
    console.log(oracle, addr1)
}


init();