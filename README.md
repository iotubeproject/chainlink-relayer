# chainlink-relayer
## Aggregators
For each aggregator on source chain (e.g. Ethereum), we will create a shadow aggregator on IoTeX. The table below lists the shadow aggregators we created on IoTeX testnet:
|Pair |Dec|Aggregator|Shadow Aggregator|
|----|----|----|----|
|BTC/USD|8|[0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c](https://etherscan.io/address/0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c)|[0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2](https://iotexscan.io/address/0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2)|


## Want to ba a relayer
### prepare config file
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 
### start relayer service
go run main.go -config config.yaml
