# chainlink-relayer
Chainlink Relayer instantly relays Chainlink price feeds on Ethereum to IoTeX network, for enabling dApps where price feeds are needed
- The prices are relayed to a contract (called shadow aggregator) on IoTeX which has exactly the same interface as Chainlink's aggregator, and it makes dApps to migrate to use Chainlink effortltessly once their integration with IoTeX is done.
- The client waits for 20 blocks (around 5 minutes) to confirm a transaction. 
- Anyone can run a relayer to relay the information. So it is permissionless!

## Shadow Aggregators
For each aggregator on Ethereum, we will create a shadow aggregator on IoTeX. 

**IoTeX Testnet**
The table below lists the shadow aggregators deployed on IoTeX testnet:
|Pair |Dec|Aggregator|Shadow Aggregator|
|----|----|----|----|
|BTC/USD|8|[0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c](https://etherscan.io/address/0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c)|[0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2](https://testnet.iotexscan.io/address/0xc4A29a94f12be03033daa4e6Ce9b9678c26275a2)|
|ETH/USD|8|[0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419](https://etherscan.io/address/0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419)|[0x107DF34D3B2F471eEff880956957e5068A987b81](https://testnet.iotexscan.io/address/0x107DF34D3B2F471eEff880956957e5068A987b81)|
|BUSD/USD|8|[0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A](https://etherscan.io/address/0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A)|[0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5](https://testnet.iotexscan.io/address/0x8A6A4407c77F1e04C39bAd8C089D639cbda40Df5)|
|USDC/USD|8|[0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6](https://etherscan.io/address/0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6)|[0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54](https://testnet.iotexscan.io/address/0xB1aa8c29d96720A80AFe9e3F6CD48822D27C8d54)|
|USDT/USD|8|[0x3E7d1eAB13ad0104d2750B8863b489D65364e32D](https://etherscan.io/address/0x3E7d1eAB13ad0104d2750B8863b489D65364e32D)|[0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63](https://testnet.iotexscan.io/address/0x63Bd61A642d1f3dbf1f47006AC03CD7e7eb72f63)|
|DAI/USD|8|[0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9](https://etherscan.io/address/0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9)|[0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151](https://testnet.iotexscan.io/address/0x9673b1b3fbB96E24f1C1AB40421Db9465f0f1151)|

**IoTeX Mainnet**
|Pair |Dec|Aggregator|Shadow Aggregator|
|----|----|----|----|
|BTC/USD|8|[0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c](https://etherscan.io/address/0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c)|[0x631f185E832DfBC3aDfeFa37c83aA23f75d0c8c7](https://iotexscan.io/address/0x631f185E832DfBC3aDfeFa37c83aA23f75d0c8c7)|
|ETH/USD|8|[0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419](https://etherscan.io/address/0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419)|[0x0a1886890c0633e32746bc5021E3c1EfAD2bd662](https://iotexscan.io/address/0x0a1886890c0633e32746bc5021E3c1EfAD2bd662)|
|BUSD/USD|8|[0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A](https://etherscan.io/address/0x833D8Eb16D306ed1FbB5D7A2E019e106B960965A)|[0x071F9106A9957e530a4B48269e38640ebAfc0f34](https://iotexscan.io/address/0x071F9106A9957e530a4B48269e38640ebAfc0f34)|
|USDC/USD|8|[0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6](https://etherscan.io/address/0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6)|[0xC296E7e92B3Ce84e9bF5780a47eF231E14A4506d](https://iotexscan.io/address/0xC296E7e92B3Ce84e9bF5780a47eF231E14A4506d)|
|USDT/USD|8|[0x3E7d1eAB13ad0104d2750B8863b489D65364e32D](https://etherscan.io/address/0x3E7d1eAB13ad0104d2750B8863b489D65364e32D)|[0xa900b5eB48F5A1122F9bfA660dd0B61Ddc56C872](https://iotexscan.io/address/0xa900b5eB48F5A1122F9bfA660dd0B61Ddc56C872)|
|DAI/USD|8|[0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9](https://etherscan.io/address/0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9)|[0x95eBC95FF2b81866D7Bc1c3c1257533795CF96B7](https://iotexscan.io/address/0x95eBC95FF2b81866D7Bc1c3c1257533795CF96B7)|

|Pair |Dec|Aggregator|
|----|----|----|
|IOTX/USD|8|[0x267Ef702F3422cC55C617218a4fB84446F5Ec646](https://iotexscan.io/address/0x267Ef702F3422cC55C617218a4fB84446F5Ec646)|

## Exchange Aggregators
Exchange aggregator has the same interface as Chainlink aggregator. The prices are read from exchanges, and the medium value of these prices will be feed to exchange aggregator by permitted relayers if
- the value is of a 0.5% deviation from the last price
- or the time interval is more than 1 hour  

## Run a Relayer
### Prepare Config File
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 

### Start Relayer Service
go run main.go -config config.yaml
