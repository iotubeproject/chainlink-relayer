// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DummyAggregatorMetaData contains all meta data concerning the DummyAggregator contract.
var DummyAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int192\",\"name\":\"minAnswer_\",\"type\":\"int192\"},{\"internalType\":\"int192\",\"name\":\"maxAnswer_\",\"type\":\"int192\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"epochAndRound\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"roundId\",\"type\":\"uint32\"},{\"internalType\":\"int192\",\"name\":\"answer\",\"type\":\"int192\"},{\"internalType\":\"uint32\",\"name\":\"observationsTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"transmissionTimestamp\",\"type\":\"uint32\"}],\"name\":\"addTransmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId_\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAnswer\",\"outputs\":[{\"internalType\":\"int192\",\"name\":\"\",\"type\":\"int192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetEpochAndRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DummyAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyAggregatorMetaData.ABI instead.
var DummyAggregatorABI = DummyAggregatorMetaData.ABI

// DummyAggregator is an auto generated Go binding around an Ethereum contract.
type DummyAggregator struct {
	DummyAggregatorCaller     // Read-only binding to the contract
	DummyAggregatorTransactor // Write-only binding to the contract
	DummyAggregatorFilterer   // Log filterer for contract events
}

// DummyAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyAggregatorSession struct {
	Contract     *DummyAggregator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyAggregatorCallerSession struct {
	Contract *DummyAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DummyAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyAggregatorTransactorSession struct {
	Contract     *DummyAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DummyAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyAggregatorRaw struct {
	Contract *DummyAggregator // Generic contract binding to access the raw methods on
}

// DummyAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyAggregatorCallerRaw struct {
	Contract *DummyAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// DummyAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyAggregatorTransactorRaw struct {
	Contract *DummyAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyAggregator creates a new instance of DummyAggregator, bound to a specific deployed contract.
func NewDummyAggregator(address common.Address, backend bind.ContractBackend) (*DummyAggregator, error) {
	contract, err := bindDummyAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyAggregator{DummyAggregatorCaller: DummyAggregatorCaller{contract: contract}, DummyAggregatorTransactor: DummyAggregatorTransactor{contract: contract}, DummyAggregatorFilterer: DummyAggregatorFilterer{contract: contract}}, nil
}

// NewDummyAggregatorCaller creates a new read-only instance of DummyAggregator, bound to a specific deployed contract.
func NewDummyAggregatorCaller(address common.Address, caller bind.ContractCaller) (*DummyAggregatorCaller, error) {
	contract, err := bindDummyAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorCaller{contract: contract}, nil
}

// NewDummyAggregatorTransactor creates a new write-only instance of DummyAggregator, bound to a specific deployed contract.
func NewDummyAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyAggregatorTransactor, error) {
	contract, err := bindDummyAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorTransactor{contract: contract}, nil
}

// NewDummyAggregatorFilterer creates a new log filterer instance of DummyAggregator, bound to a specific deployed contract.
func NewDummyAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyAggregatorFilterer, error) {
	contract, err := bindDummyAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorFilterer{contract: contract}, nil
}

// bindDummyAggregator binds a generic wrapper to an already deployed contract.
func bindDummyAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DummyAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyAggregator *DummyAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyAggregator.Contract.DummyAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyAggregator *DummyAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAggregator.Contract.DummyAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyAggregator *DummyAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyAggregator.Contract.DummyAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyAggregator *DummyAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyAggregator *DummyAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyAggregator *DummyAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyAggregator.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DummyAggregator *DummyAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DummyAggregator *DummyAggregatorSession) Decimals() (uint8, error) {
	return _DummyAggregator.Contract.Decimals(&_DummyAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DummyAggregator *DummyAggregatorCallerSession) Decimals() (uint8, error) {
	return _DummyAggregator.Contract.Decimals(&_DummyAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_DummyAggregator *DummyAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_DummyAggregator *DummyAggregatorSession) Description() (string, error) {
	return _DummyAggregator.Contract.Description(&_DummyAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_DummyAggregator *DummyAggregatorCallerSession) Description() (string, error) {
	return _DummyAggregator.Contract.Description(&_DummyAggregator.CallOpts)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_DummyAggregator *DummyAggregatorCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_DummyAggregator *DummyAggregatorSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _DummyAggregator.Contract.GetAnswer(&_DummyAggregator.CallOpts, roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_DummyAggregator *DummyAggregatorCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _DummyAggregator.Contract.GetAnswer(&_DummyAggregator.CallOpts, roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 roundId) view returns(uint80 roundId_, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorCaller) GetRoundData(opts *bind.CallOpts, roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "getRoundData", roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 roundId) view returns(uint80 roundId_, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorSession) GetRoundData(roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _DummyAggregator.Contract.GetRoundData(&_DummyAggregator.CallOpts, roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 roundId) view returns(uint80 roundId_, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorCallerSession) GetRoundData(roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _DummyAggregator.Contract.GetRoundData(&_DummyAggregator.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_DummyAggregator *DummyAggregatorCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_DummyAggregator *DummyAggregatorSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _DummyAggregator.Contract.GetTimestamp(&_DummyAggregator.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_DummyAggregator *DummyAggregatorCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _DummyAggregator.Contract.GetTimestamp(&_DummyAggregator.CallOpts, roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_DummyAggregator *DummyAggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_DummyAggregator *DummyAggregatorSession) LatestAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestAnswer(&_DummyAggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_DummyAggregator *DummyAggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestAnswer(&_DummyAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_DummyAggregator *DummyAggregatorSession) LatestRound() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestRound(&_DummyAggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestRound(&_DummyAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _DummyAggregator.Contract.LatestRoundData(&_DummyAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_DummyAggregator *DummyAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _DummyAggregator.Contract.LatestRoundData(&_DummyAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_DummyAggregator *DummyAggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestTimestamp(&_DummyAggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _DummyAggregator.Contract.LatestTimestamp(&_DummyAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorCaller) MaxAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "maxAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorSession) MaxAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.MaxAnswer(&_DummyAggregator.CallOpts)
}

// MaxAnswer is a free data retrieval call binding the contract method 0x70da2f67.
//
// Solidity: function maxAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorCallerSession) MaxAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.MaxAnswer(&_DummyAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorCaller) MinAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "minAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorSession) MinAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.MinAnswer(&_DummyAggregator.CallOpts)
}

// MinAnswer is a free data retrieval call binding the contract method 0x22adbc78.
//
// Solidity: function minAnswer() view returns(int192)
func (_DummyAggregator *DummyAggregatorCallerSession) MinAnswer() (*big.Int, error) {
	return _DummyAggregator.Contract.MinAnswer(&_DummyAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DummyAggregator *DummyAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DummyAggregator *DummyAggregatorSession) Owner() (common.Address, error) {
	return _DummyAggregator.Contract.Owner(&_DummyAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DummyAggregator *DummyAggregatorCallerSession) Owner() (common.Address, error) {
	return _DummyAggregator.Contract.Owner(&_DummyAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DummyAggregator *DummyAggregatorSession) Version() (*big.Int, error) {
	return _DummyAggregator.Contract.Version(&_DummyAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DummyAggregator *DummyAggregatorCallerSession) Version() (*big.Int, error) {
	return _DummyAggregator.Contract.Version(&_DummyAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DummyAggregator *DummyAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DummyAggregator *DummyAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DummyAggregator.Contract.AcceptOwnership(&_DummyAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DummyAggregator *DummyAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DummyAggregator.Contract.AcceptOwnership(&_DummyAggregator.TransactOpts)
}

// AddTransmission is a paid mutator transaction binding the contract method 0x99fafa4e.
//
// Solidity: function addTransmission(uint40 epochAndRound, uint32 roundId, int192 answer, uint32 observationsTimestamp, uint32 transmissionTimestamp) returns()
func (_DummyAggregator *DummyAggregatorTransactor) AddTransmission(opts *bind.TransactOpts, epochAndRound *big.Int, roundId uint32, answer *big.Int, observationsTimestamp uint32, transmissionTimestamp uint32) (*types.Transaction, error) {
	return _DummyAggregator.contract.Transact(opts, "addTransmission", epochAndRound, roundId, answer, observationsTimestamp, transmissionTimestamp)
}

// AddTransmission is a paid mutator transaction binding the contract method 0x99fafa4e.
//
// Solidity: function addTransmission(uint40 epochAndRound, uint32 roundId, int192 answer, uint32 observationsTimestamp, uint32 transmissionTimestamp) returns()
func (_DummyAggregator *DummyAggregatorSession) AddTransmission(epochAndRound *big.Int, roundId uint32, answer *big.Int, observationsTimestamp uint32, transmissionTimestamp uint32) (*types.Transaction, error) {
	return _DummyAggregator.Contract.AddTransmission(&_DummyAggregator.TransactOpts, epochAndRound, roundId, answer, observationsTimestamp, transmissionTimestamp)
}

// AddTransmission is a paid mutator transaction binding the contract method 0x99fafa4e.
//
// Solidity: function addTransmission(uint40 epochAndRound, uint32 roundId, int192 answer, uint32 observationsTimestamp, uint32 transmissionTimestamp) returns()
func (_DummyAggregator *DummyAggregatorTransactorSession) AddTransmission(epochAndRound *big.Int, roundId uint32, answer *big.Int, observationsTimestamp uint32, transmissionTimestamp uint32) (*types.Transaction, error) {
	return _DummyAggregator.Contract.AddTransmission(&_DummyAggregator.TransactOpts, epochAndRound, roundId, answer, observationsTimestamp, transmissionTimestamp)
}

// ResetEpochAndRound is a paid mutator transaction binding the contract method 0x83f5e969.
//
// Solidity: function resetEpochAndRound() returns()
func (_DummyAggregator *DummyAggregatorTransactor) ResetEpochAndRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyAggregator.contract.Transact(opts, "resetEpochAndRound")
}

// ResetEpochAndRound is a paid mutator transaction binding the contract method 0x83f5e969.
//
// Solidity: function resetEpochAndRound() returns()
func (_DummyAggregator *DummyAggregatorSession) ResetEpochAndRound() (*types.Transaction, error) {
	return _DummyAggregator.Contract.ResetEpochAndRound(&_DummyAggregator.TransactOpts)
}

// ResetEpochAndRound is a paid mutator transaction binding the contract method 0x83f5e969.
//
// Solidity: function resetEpochAndRound() returns()
func (_DummyAggregator *DummyAggregatorTransactorSession) ResetEpochAndRound() (*types.Transaction, error) {
	return _DummyAggregator.Contract.ResetEpochAndRound(&_DummyAggregator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DummyAggregator *DummyAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DummyAggregator.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DummyAggregator *DummyAggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DummyAggregator.Contract.TransferOwnership(&_DummyAggregator.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_DummyAggregator *DummyAggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _DummyAggregator.Contract.TransferOwnership(&_DummyAggregator.TransactOpts, to)
}

// DummyAggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the DummyAggregator contract.
type DummyAggregatorAnswerUpdatedIterator struct {
	Event *DummyAggregatorAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DummyAggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyAggregatorAnswerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DummyAggregatorAnswerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DummyAggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyAggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyAggregatorAnswerUpdated represents a AnswerUpdated event raised by the DummyAggregator contract.
type DummyAggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_DummyAggregator *DummyAggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*DummyAggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DummyAggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorAnswerUpdatedIterator{contract: _DummyAggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_DummyAggregator *DummyAggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *DummyAggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _DummyAggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyAggregatorAnswerUpdated)
				if err := _DummyAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_DummyAggregator *DummyAggregatorFilterer) ParseAnswerUpdated(log types.Log) (*DummyAggregatorAnswerUpdated, error) {
	event := new(DummyAggregatorAnswerUpdated)
	if err := _DummyAggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DummyAggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the DummyAggregator contract.
type DummyAggregatorNewRoundIterator struct {
	Event *DummyAggregatorNewRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DummyAggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyAggregatorNewRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DummyAggregatorNewRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DummyAggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyAggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyAggregatorNewRound represents a NewRound event raised by the DummyAggregator contract.
type DummyAggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_DummyAggregator *DummyAggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*DummyAggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _DummyAggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorNewRoundIterator{contract: _DummyAggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_DummyAggregator *DummyAggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *DummyAggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _DummyAggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyAggregatorNewRound)
				if err := _DummyAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_DummyAggregator *DummyAggregatorFilterer) ParseNewRound(log types.Log) (*DummyAggregatorNewRound, error) {
	event := new(DummyAggregatorNewRound)
	if err := _DummyAggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DummyAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the DummyAggregator contract.
type DummyAggregatorOwnershipTransferRequestedIterator struct {
	Event *DummyAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DummyAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyAggregatorOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DummyAggregatorOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DummyAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the DummyAggregator contract.
type DummyAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DummyAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DummyAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorOwnershipTransferRequestedIterator{contract: _DummyAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *DummyAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DummyAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyAggregatorOwnershipTransferRequested)
				if err := _DummyAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*DummyAggregatorOwnershipTransferRequested, error) {
	event := new(DummyAggregatorOwnershipTransferRequested)
	if err := _DummyAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DummyAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DummyAggregator contract.
type DummyAggregatorOwnershipTransferredIterator struct {
	Event *DummyAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DummyAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DummyAggregatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DummyAggregatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DummyAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DummyAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DummyAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the DummyAggregator contract.
type DummyAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DummyAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DummyAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DummyAggregatorOwnershipTransferredIterator{contract: _DummyAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DummyAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DummyAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DummyAggregatorOwnershipTransferred)
				if err := _DummyAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_DummyAggregator *DummyAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*DummyAggregatorOwnershipTransferred, error) {
	event := new(DummyAggregatorOwnershipTransferred)
	if err := _DummyAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
