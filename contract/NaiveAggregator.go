// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NaiveAggregatorABI is the input ABI used to generate the binding from.
const NaiveAggregatorABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_answer\",\"type\":\"int256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NaiveAggregator is an auto generated Go binding around an Ethereum contract.
type NaiveAggregator struct {
	NaiveAggregatorCaller     // Read-only binding to the contract
	NaiveAggregatorTransactor // Write-only binding to the contract
	NaiveAggregatorFilterer   // Log filterer for contract events
}

// NaiveAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type NaiveAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NaiveAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NaiveAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NaiveAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NaiveAggregatorSession struct {
	Contract     *NaiveAggregator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NaiveAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NaiveAggregatorCallerSession struct {
	Contract *NaiveAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NaiveAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NaiveAggregatorTransactorSession struct {
	Contract     *NaiveAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NaiveAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type NaiveAggregatorRaw struct {
	Contract *NaiveAggregator // Generic contract binding to access the raw methods on
}

// NaiveAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NaiveAggregatorCallerRaw struct {
	Contract *NaiveAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// NaiveAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NaiveAggregatorTransactorRaw struct {
	Contract *NaiveAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNaiveAggregator creates a new instance of NaiveAggregator, bound to a specific deployed contract.
func NewNaiveAggregator(address common.Address, backend bind.ContractBackend) (*NaiveAggregator, error) {
	contract, err := bindNaiveAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NaiveAggregator{NaiveAggregatorCaller: NaiveAggregatorCaller{contract: contract}, NaiveAggregatorTransactor: NaiveAggregatorTransactor{contract: contract}, NaiveAggregatorFilterer: NaiveAggregatorFilterer{contract: contract}}, nil
}

// NewNaiveAggregatorCaller creates a new read-only instance of NaiveAggregator, bound to a specific deployed contract.
func NewNaiveAggregatorCaller(address common.Address, caller bind.ContractCaller) (*NaiveAggregatorCaller, error) {
	contract, err := bindNaiveAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NaiveAggregatorCaller{contract: contract}, nil
}

// NewNaiveAggregatorTransactor creates a new write-only instance of NaiveAggregator, bound to a specific deployed contract.
func NewNaiveAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*NaiveAggregatorTransactor, error) {
	contract, err := bindNaiveAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NaiveAggregatorTransactor{contract: contract}, nil
}

// NewNaiveAggregatorFilterer creates a new log filterer instance of NaiveAggregator, bound to a specific deployed contract.
func NewNaiveAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*NaiveAggregatorFilterer, error) {
	contract, err := bindNaiveAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NaiveAggregatorFilterer{contract: contract}, nil
}

// bindNaiveAggregator binds a generic wrapper to an already deployed contract.
func bindNaiveAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NaiveAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NaiveAggregator *NaiveAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NaiveAggregator.Contract.NaiveAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NaiveAggregator *NaiveAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.NaiveAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NaiveAggregator *NaiveAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.NaiveAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NaiveAggregator *NaiveAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NaiveAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NaiveAggregator *NaiveAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NaiveAggregator *NaiveAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_NaiveAggregator *NaiveAggregatorCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_NaiveAggregator *NaiveAggregatorSession) Aggregator() (common.Address, error) {
	return _NaiveAggregator.Contract.Aggregator(&_NaiveAggregator.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_NaiveAggregator *NaiveAggregatorCallerSession) Aggregator() (common.Address, error) {
	return _NaiveAggregator.Contract.Aggregator(&_NaiveAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NaiveAggregator *NaiveAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NaiveAggregator *NaiveAggregatorSession) Decimals() (uint8, error) {
	return _NaiveAggregator.Contract.Decimals(&_NaiveAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NaiveAggregator *NaiveAggregatorCallerSession) Decimals() (uint8, error) {
	return _NaiveAggregator.Contract.Decimals(&_NaiveAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_NaiveAggregator *NaiveAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_NaiveAggregator *NaiveAggregatorSession) Description() (string, error) {
	return _NaiveAggregator.Contract.Description(&_NaiveAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_NaiveAggregator *NaiveAggregatorCallerSession) Description() (string, error) {
	return _NaiveAggregator.Contract.Description(&_NaiveAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_NaiveAggregator *NaiveAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "getRoundData", _roundId)

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
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_NaiveAggregator *NaiveAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _NaiveAggregator.Contract.GetRoundData(&_NaiveAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_NaiveAggregator *NaiveAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _NaiveAggregator.Contract.GetRoundData(&_NaiveAggregator.CallOpts, _roundId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NaiveAggregator *NaiveAggregatorCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NaiveAggregator *NaiveAggregatorSession) IsOwner() (bool, error) {
	return _NaiveAggregator.Contract.IsOwner(&_NaiveAggregator.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NaiveAggregator *NaiveAggregatorCallerSession) IsOwner() (bool, error) {
	return _NaiveAggregator.Contract.IsOwner(&_NaiveAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_NaiveAggregator *NaiveAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "latestRoundData")

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
func (_NaiveAggregator *NaiveAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _NaiveAggregator.Contract.LatestRoundData(&_NaiveAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_NaiveAggregator *NaiveAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _NaiveAggregator.Contract.LatestRoundData(&_NaiveAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_NaiveAggregator *NaiveAggregatorCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_NaiveAggregator *NaiveAggregatorSession) LatestRoundId() (*big.Int, error) {
	return _NaiveAggregator.Contract.LatestRoundId(&_NaiveAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_NaiveAggregator *NaiveAggregatorCallerSession) LatestRoundId() (*big.Int, error) {
	return _NaiveAggregator.Contract.LatestRoundId(&_NaiveAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NaiveAggregator *NaiveAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NaiveAggregator *NaiveAggregatorSession) Owner() (common.Address, error) {
	return _NaiveAggregator.Contract.Owner(&_NaiveAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NaiveAggregator *NaiveAggregatorCallerSession) Owner() (common.Address, error) {
	return _NaiveAggregator.Contract.Owner(&_NaiveAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_NaiveAggregator *NaiveAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NaiveAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_NaiveAggregator *NaiveAggregatorSession) Version() (*big.Int, error) {
	return _NaiveAggregator.Contract.Version(&_NaiveAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_NaiveAggregator *NaiveAggregatorCallerSession) Version() (*big.Int, error) {
	return _NaiveAggregator.Contract.Version(&_NaiveAggregator.CallOpts)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_NaiveAggregator *NaiveAggregatorTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_NaiveAggregator *NaiveAggregatorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.SetAggregator(&_NaiveAggregator.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_NaiveAggregator *NaiveAggregatorTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.SetAggregator(&_NaiveAggregator.TransactOpts, _aggregator)
}

// Submit is a paid mutator transaction binding the contract method 0x9b25df0b.
//
// Solidity: function submit(int256 _answer) returns()
func (_NaiveAggregator *NaiveAggregatorTransactor) Submit(opts *bind.TransactOpts, _answer *big.Int) (*types.Transaction, error) {
	return _NaiveAggregator.contract.Transact(opts, "submit", _answer)
}

// Submit is a paid mutator transaction binding the contract method 0x9b25df0b.
//
// Solidity: function submit(int256 _answer) returns()
func (_NaiveAggregator *NaiveAggregatorSession) Submit(_answer *big.Int) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.Submit(&_NaiveAggregator.TransactOpts, _answer)
}

// Submit is a paid mutator transaction binding the contract method 0x9b25df0b.
//
// Solidity: function submit(int256 _answer) returns()
func (_NaiveAggregator *NaiveAggregatorTransactorSession) Submit(_answer *big.Int) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.Submit(&_NaiveAggregator.TransactOpts, _answer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NaiveAggregator *NaiveAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NaiveAggregator *NaiveAggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.TransferOwnership(&_NaiveAggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NaiveAggregator *NaiveAggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NaiveAggregator.Contract.TransferOwnership(&_NaiveAggregator.TransactOpts, newOwner)
}

// NaiveAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NaiveAggregator contract.
type NaiveAggregatorOwnershipTransferredIterator struct {
	Event *NaiveAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NaiveAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NaiveAggregatorOwnershipTransferred)
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
		it.Event = new(NaiveAggregatorOwnershipTransferred)
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
func (it *NaiveAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NaiveAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NaiveAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the NaiveAggregator contract.
type NaiveAggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NaiveAggregator *NaiveAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NaiveAggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NaiveAggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NaiveAggregatorOwnershipTransferredIterator{contract: _NaiveAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NaiveAggregator *NaiveAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NaiveAggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NaiveAggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NaiveAggregatorOwnershipTransferred)
				if err := _NaiveAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NaiveAggregator *NaiveAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*NaiveAggregatorOwnershipTransferred, error) {
	event := new(NaiveAggregatorOwnershipTransferred)
	if err := _NaiveAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
