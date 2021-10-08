# chainlink-relayer
## prepare config file
1. fill "databaseURL" with your DB link
2. input ethereum rpc url in "sourceClientURL"
3. input your relayer private key, make sure the account has sufficient balance 
## start service
go run main.go -config config.yaml
