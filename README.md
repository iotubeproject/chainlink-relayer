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
|ETH/USD|8|[0x37bc7498f4ff12c19678ee8fe19d713b87f6a9e6](https://etherscan.io/address/0x37bc7498f4ff12c19678ee8fe19d713b87f6a9e6)|[0x107DF34D3B2F471eEff880956957e5068A987b81](https://iotexscan.io/address/0x107DF34D3B2F471eEff880956957e5068A987b81)|
|BUSD/USD|8|[0x73dc1b226f7dfac353bdb41a27c4212213e6af07](https://etherscan.io/address/0x73dc1b226f7dfac353bdb41a27c4212213e6af07)|[0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5](https://iotexscan.io/address/0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5)|
|USDC/USD|8|[0x789190466e21a8b78b8027866cbbdc151542a26c](https://etherscan.io/address/0x789190466e21a8b78b8027866cbbdc151542a26c)|[0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54](https://iotexscan.io/address/0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54)|
|USDT/USD|8|[0xa964273552c1dba201f5f000215f5bd5576e8f93](https://etherscan.io/address/0xa964273552c1dba201f5f000215f5bd5576e8f93)|[0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63](https://iotexscan.io/address/0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63)|
|DAI/USD|8|[0xdec0a100ead1faa37407f0edc76033426cf90b82](https://etherscan.io/address/0xdec0a100ead1faa37407f0edc76033426cf90b82)|[0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151](https://iotexscan.io/address/0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151)|

## Run a Relayer
### Prepare Config File
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 

### Start Relayer Service
go run main.go -config config.yaml
