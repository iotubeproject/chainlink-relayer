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
|ETH/USD|8|[0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419](https://etherscan.io/address/0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419)|[0x107DF34D3B2F471eEff880956957e5068A987b81](https://iotexscan.io/address/0x107DF34D3B2F471eEff880956957e5068A987b81)|
|BUSD/USD|8|[0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A](https://etherscan.io/address/0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A)|[0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5](https://iotexscan.io/address/0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5)|
|USDC/USD|8|[0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6](https://etherscan.io/address/0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6)|[0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54](https://iotexscan.io/address/0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54)|
|USDT/USD|8|[0x3E7d1eAB13ad0104d2750B8863b489D65364e32D](https://etherscan.io/address/0x3E7d1eAB13ad0104d2750B8863b489D65364e32D)|[0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63](https://iotexscan.io/address/0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63)|
|DAI/USD|8|[0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9](https://etherscan.io/address/0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9)|[0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151](https://iotexscan.io/address/0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151)|

## Run a Relayer
### Prepare Config File
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 

### Start Relayer Service
go run main.go -config config.yaml
