# chainlink-relayer
Chainlink Relayer instantly relays Chainlink price feeds on Ethereum to IoTeX network, for enabling dApps where price feeds are needed
- The prices are relayed to a contract (called shadow aggregator) on IoTeX which has exactly the same interface as Chainlink's aggregator, and it makes dApps to migrate to use Chainlink effortltessly once their integration with IoTeX is done.
- Anyone can run a relayer to relay the information. So it is permissionless!


## Aggregators
For each aggregator on Ethereum, we will create a shadow aggregator on IoTeX. 

The table below lists the shadow aggregators deployed on IoTeX testnet:
|Pair |Dec|Aggregator|Shadow Aggregator|
|----|----|----|----|
|BTC/USD|8|[0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c](https://etherscan.io/address/0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c)|[0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2](https://iotexscan.io/address/0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2)|


## Run a Relayer
### Prepare Config File
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 

### Start Relayer Service
go run main.go -config config.yaml
