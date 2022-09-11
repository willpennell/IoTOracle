1_initial_migration.js \
2_deploy_contracts \
auto generated, however we add to add parts our self

```javascript
deployer.then(async () => {
        console.log(accounts);
        await deployer.deploy(OracleRequest, {from: accounts[0]});
        await deployer.deploy(Aggregator, OracleRequest.address, {from: accounts[0]});
        await deployer.deploy(Reputation, OracleRequest.address, {from: accounts[0]});
        let OracleRequestInstance = await OracleRequest.deployed();
        await OracleRequestInstance.setAggregatorContract(Aggregator.address);
        await OracleRequestInstance.setReputationContract(Reputation.address);
    });
```

We had to deploy the contracts in a certain order as 
Reputation and Aggregation contracts constructors required the ORC-SC address