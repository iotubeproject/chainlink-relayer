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

// ShadowAggregatorABI is the input ABI used to generate the binding from.
const ShadowAggregatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"submission\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint80\",\"name\":\"round\",\"type\":\"uint80\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SubmissionReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configDigest\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"internalType\":\"enumShadowAggregator.Role\",\"name\":\"role\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_threshold\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_encodedConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_encoded\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_count\",\"type\":\"uint32\"}],\"name\":\"setConfigCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rawVs\",\"type\":\"bytes32\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ShadowAggregator is an auto generated Go binding around an Ethereum contract.
type ShadowAggregator struct {
	ShadowAggregatorCaller     // Read-only binding to the contract
	ShadowAggregatorTransactor // Write-only binding to the contract
	ShadowAggregatorFilterer   // Log filterer for contract events
}

// ShadowAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShadowAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShadowAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShadowAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShadowAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShadowAggregatorSession struct {
	Contract     *ShadowAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShadowAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShadowAggregatorCallerSession struct {
	Contract *ShadowAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ShadowAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShadowAggregatorTransactorSession struct {
	Contract     *ShadowAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ShadowAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShadowAggregatorRaw struct {
	Contract *ShadowAggregator // Generic contract binding to access the raw methods on
}

// ShadowAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShadowAggregatorCallerRaw struct {
	Contract *ShadowAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// ShadowAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShadowAggregatorTransactorRaw struct {
	Contract *ShadowAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShadowAggregator creates a new instance of ShadowAggregator, bound to a specific deployed contract.
func NewShadowAggregator(address common.Address, backend bind.ContractBackend) (*ShadowAggregator, error) {
	contract, err := bindShadowAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregator{ShadowAggregatorCaller: ShadowAggregatorCaller{contract: contract}, ShadowAggregatorTransactor: ShadowAggregatorTransactor{contract: contract}, ShadowAggregatorFilterer: ShadowAggregatorFilterer{contract: contract}}, nil
}

// NewShadowAggregatorCaller creates a new read-only instance of ShadowAggregator, bound to a specific deployed contract.
func NewShadowAggregatorCaller(address common.Address, caller bind.ContractCaller) (*ShadowAggregatorCaller, error) {
	contract, err := bindShadowAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorCaller{contract: contract}, nil
}

// NewShadowAggregatorTransactor creates a new write-only instance of ShadowAggregator, bound to a specific deployed contract.
func NewShadowAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ShadowAggregatorTransactor, error) {
	contract, err := bindShadowAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorTransactor{contract: contract}, nil
}

// NewShadowAggregatorFilterer creates a new log filterer instance of ShadowAggregator, bound to a specific deployed contract.
func NewShadowAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ShadowAggregatorFilterer, error) {
	contract, err := bindShadowAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorFilterer{contract: contract}, nil
}

// bindShadowAggregator binds a generic wrapper to an already deployed contract.
func bindShadowAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShadowAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowAggregator *ShadowAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowAggregator.Contract.ShadowAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowAggregator *ShadowAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.ShadowAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowAggregator *ShadowAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.ShadowAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShadowAggregator *ShadowAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShadowAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShadowAggregator *ShadowAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShadowAggregator *ShadowAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorSession) Aggregator() (common.Address, error) {
	return _ShadowAggregator.Contract.Aggregator(&_ShadowAggregator.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Aggregator() (common.Address, error) {
	return _ShadowAggregator.Contract.Aggregator(&_ShadowAggregator.CallOpts)
}

// ConfigCount is a free data retrieval call binding the contract method 0xa0620634.
//
// Solidity: function configCount() view returns(uint32)
func (_ShadowAggregator *ShadowAggregatorCaller) ConfigCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "configCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigCount is a free data retrieval call binding the contract method 0xa0620634.
//
// Solidity: function configCount() view returns(uint32)
func (_ShadowAggregator *ShadowAggregatorSession) ConfigCount() (uint32, error) {
	return _ShadowAggregator.Contract.ConfigCount(&_ShadowAggregator.CallOpts)
}

// ConfigCount is a free data retrieval call binding the contract method 0xa0620634.
//
// Solidity: function configCount() view returns(uint32)
func (_ShadowAggregator *ShadowAggregatorCallerSession) ConfigCount() (uint32, error) {
	return _ShadowAggregator.Contract.ConfigCount(&_ShadowAggregator.CallOpts)
}

// ConfigDigest is a free data retrieval call binding the contract method 0xac540d40.
//
// Solidity: function configDigest() view returns(bytes16)
func (_ShadowAggregator *ShadowAggregatorCaller) ConfigDigest(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "configDigest")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// ConfigDigest is a free data retrieval call binding the contract method 0xac540d40.
//
// Solidity: function configDigest() view returns(bytes16)
func (_ShadowAggregator *ShadowAggregatorSession) ConfigDigest() ([16]byte, error) {
	return _ShadowAggregator.Contract.ConfigDigest(&_ShadowAggregator.CallOpts)
}

// ConfigDigest is a free data retrieval call binding the contract method 0xac540d40.
//
// Solidity: function configDigest() view returns(bytes16)
func (_ShadowAggregator *ShadowAggregatorCallerSession) ConfigDigest() ([16]byte, error) {
	return _ShadowAggregator.Contract.ConfigDigest(&_ShadowAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorSession) Decimals() (uint8, error) {
	return _ShadowAggregator.Contract.Decimals(&_ShadowAggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Decimals() (uint8, error) {
	return _ShadowAggregator.Contract.Decimals(&_ShadowAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ShadowAggregator *ShadowAggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ShadowAggregator *ShadowAggregatorSession) Description() (string, error) {
	return _ShadowAggregator.Contract.Description(&_ShadowAggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Description() (string, error) {
	return _ShadowAggregator.Contract.Description(&_ShadowAggregator.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ShadowAggregator *ShadowAggregatorCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ShadowAggregator *ShadowAggregatorSession) GetOracles() ([]common.Address, error) {
	return _ShadowAggregator.Contract.GetOracles(&_ShadowAggregator.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_ShadowAggregator *ShadowAggregatorCallerSession) GetOracles() ([]common.Address, error) {
	return _ShadowAggregator.Contract.GetOracles(&_ShadowAggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ShadowAggregator *ShadowAggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "getRoundData", _roundId)

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
func (_ShadowAggregator *ShadowAggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ShadowAggregator.Contract.GetRoundData(&_ShadowAggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ShadowAggregator *ShadowAggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ShadowAggregator.Contract.GetRoundData(&_ShadowAggregator.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ShadowAggregator *ShadowAggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "latestRoundData")

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
func (_ShadowAggregator *ShadowAggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ShadowAggregator.Contract.LatestRoundData(&_ShadowAggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ShadowAggregator *ShadowAggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ShadowAggregator.Contract.LatestRoundData(&_ShadowAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_ShadowAggregator *ShadowAggregatorCaller) LatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "latestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_ShadowAggregator *ShadowAggregatorSession) LatestRoundId() (*big.Int, error) {
	return _ShadowAggregator.Contract.LatestRoundId(&_ShadowAggregator.CallOpts)
}

// LatestRoundId is a free data retrieval call binding the contract method 0x11a8f413.
//
// Solidity: function latestRoundId() view returns(uint80)
func (_ShadowAggregator *ShadowAggregatorCallerSession) LatestRoundId() (*big.Int, error) {
	return _ShadowAggregator.Contract.LatestRoundId(&_ShadowAggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorCaller) OracleCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "oracleCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorSession) OracleCount() (uint8, error) {
	return _ShadowAggregator.Contract.OracleCount(&_ShadowAggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_ShadowAggregator *ShadowAggregatorCallerSession) OracleCount() (uint8, error) {
	return _ShadowAggregator.Contract.OracleCount(&_ShadowAggregator.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint16 index, uint8 role)
func (_ShadowAggregator *ShadowAggregatorCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index uint16
	Role  uint8
}, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "oracles", arg0)

	outstruct := new(struct {
		Index uint16
		Role  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Role = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint16 index, uint8 role)
func (_ShadowAggregator *ShadowAggregatorSession) Oracles(arg0 common.Address) (struct {
	Index uint16
	Role  uint8
}, error) {
	return _ShadowAggregator.Contract.Oracles(&_ShadowAggregator.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(uint16 index, uint8 role)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Oracles(arg0 common.Address) (struct {
	Index uint16
	Role  uint8
}, error) {
	return _ShadowAggregator.Contract.Oracles(&_ShadowAggregator.CallOpts, arg0)
}

// OriginAggregator is a free data retrieval call binding the contract method 0x1378256a.
//
// Solidity: function originAggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCaller) OriginAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "originAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginAggregator is a free data retrieval call binding the contract method 0x1378256a.
//
// Solidity: function originAggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorSession) OriginAggregator() (common.Address, error) {
	return _ShadowAggregator.Contract.OriginAggregator(&_ShadowAggregator.CallOpts)
}

// OriginAggregator is a free data retrieval call binding the contract method 0x1378256a.
//
// Solidity: function originAggregator() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCallerSession) OriginAggregator() (common.Address, error) {
	return _ShadowAggregator.Contract.OriginAggregator(&_ShadowAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorSession) Owner() (common.Address, error) {
	return _ShadowAggregator.Contract.Owner(&_ShadowAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Owner() (common.Address, error) {
	return _ShadowAggregator.Contract.Owner(&_ShadowAggregator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorSession) PendingOwner() (common.Address, error) {
	return _ShadowAggregator.Contract.PendingOwner(&_ShadowAggregator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ShadowAggregator *ShadowAggregatorCallerSession) PendingOwner() (common.Address, error) {
	return _ShadowAggregator.Contract.PendingOwner(&_ShadowAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ShadowAggregator *ShadowAggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShadowAggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ShadowAggregator *ShadowAggregatorSession) Version() (*big.Int, error) {
	return _ShadowAggregator.Contract.Version(&_ShadowAggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ShadowAggregator *ShadowAggregatorCallerSession) Version() (*big.Int, error) {
	return _ShadowAggregator.Contract.Version(&_ShadowAggregator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ShadowAggregator *ShadowAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ShadowAggregator.Contract.AcceptOwnership(&_ShadowAggregator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ShadowAggregator.Contract.AcceptOwnership(&_ShadowAggregator.TransactOpts)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ShadowAggregator *ShadowAggregatorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetAggregator(&_ShadowAggregator.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetAggregator(&_ShadowAggregator.TransactOpts, _aggregator)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "setConfig", _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_ShadowAggregator *ShadowAggregatorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetConfig(&_ShadowAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfig is a paid mutator transaction binding the contract method 0x585aa7de.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _threshold, uint64 _encodedConfigVersion, bytes _encoded) returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _threshold uint8, _encodedConfigVersion uint64, _encoded []byte) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetConfig(&_ShadowAggregator.TransactOpts, _signers, _transmitters, _threshold, _encodedConfigVersion, _encoded)
}

// SetConfigCount is a paid mutator transaction binding the contract method 0xa483bd99.
//
// Solidity: function setConfigCount(uint32 _count) returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) SetConfigCount(opts *bind.TransactOpts, _count uint32) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "setConfigCount", _count)
}

// SetConfigCount is a paid mutator transaction binding the contract method 0xa483bd99.
//
// Solidity: function setConfigCount(uint32 _count) returns()
func (_ShadowAggregator *ShadowAggregatorSession) SetConfigCount(_count uint32) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetConfigCount(&_ShadowAggregator.TransactOpts, _count)
}

// SetConfigCount is a paid mutator transaction binding the contract method 0xa483bd99.
//
// Solidity: function setConfigCount(uint32 _count) returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) SetConfigCount(_count uint32) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.SetConfigCount(&_ShadowAggregator.TransactOpts, _count)
}

// Submit is a paid mutator transaction binding the contract method 0x9e787f0b.
//
// Solidity: function submit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) Submit(opts *bind.TransactOpts, _report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "submit", _report, _rs, _ss, _rawVs)
}

// Submit is a paid mutator transaction binding the contract method 0x9e787f0b.
//
// Solidity: function submit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_ShadowAggregator *ShadowAggregatorSession) Submit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.Submit(&_ShadowAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// Submit is a paid mutator transaction binding the contract method 0x9e787f0b.
//
// Solidity: function submit(bytes _report, bytes32[] _rs, bytes32[] _ss, bytes32 _rawVs) returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) Submit(_report []byte, _rs [][32]byte, _ss [][32]byte, _rawVs [32]byte) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.Submit(&_ShadowAggregator.TransactOpts, _report, _rs, _ss, _rawVs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ShadowAggregator *ShadowAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ShadowAggregator *ShadowAggregatorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.TransferOwnership(&_ShadowAggregator.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_ShadowAggregator *ShadowAggregatorTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _ShadowAggregator.Contract.TransferOwnership(&_ShadowAggregator.TransactOpts, _to)
}

// ShadowAggregatorOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the ShadowAggregator contract.
type ShadowAggregatorOwnershipTransferRequestedIterator struct {
	Event *ShadowAggregatorOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ShadowAggregatorOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowAggregatorOwnershipTransferRequested)
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
		it.Event = new(ShadowAggregatorOwnershipTransferRequested)
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
func (it *ShadowAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowAggregatorOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the ShadowAggregator contract.
type ShadowAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ShadowAggregator *ShadowAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShadowAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorOwnershipTransferRequestedIterator{contract: _ShadowAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_ShadowAggregator *ShadowAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ShadowAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowAggregatorOwnershipTransferRequested)
				if err := _ShadowAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_ShadowAggregator *ShadowAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ShadowAggregatorOwnershipTransferRequested, error) {
	event := new(ShadowAggregatorOwnershipTransferRequested)
	if err := _ShadowAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShadowAggregator contract.
type ShadowAggregatorOwnershipTransferredIterator struct {
	Event *ShadowAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShadowAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowAggregatorOwnershipTransferred)
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
		it.Event = new(ShadowAggregatorOwnershipTransferred)
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
func (it *ShadowAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the ShadowAggregator contract.
type ShadowAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ShadowAggregator *ShadowAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShadowAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorOwnershipTransferredIterator{contract: _ShadowAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_ShadowAggregator *ShadowAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShadowAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShadowAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowAggregatorOwnershipTransferred)
				if err := _ShadowAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShadowAggregator *ShadowAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*ShadowAggregatorOwnershipTransferred, error) {
	event := new(ShadowAggregatorOwnershipTransferred)
	if err := _ShadowAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShadowAggregatorSubmissionReceivedIterator is returned from FilterSubmissionReceived and is used to iterate over the raw logs and unpacked data for SubmissionReceived events raised by the ShadowAggregator contract.
type ShadowAggregatorSubmissionReceivedIterator struct {
	Event *ShadowAggregatorSubmissionReceived // Event containing the contract specifics and raw log

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
func (it *ShadowAggregatorSubmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShadowAggregatorSubmissionReceived)
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
		it.Event = new(ShadowAggregatorSubmissionReceived)
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
func (it *ShadowAggregatorSubmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShadowAggregatorSubmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShadowAggregatorSubmissionReceived represents a SubmissionReceived event raised by the ShadowAggregator contract.
type ShadowAggregatorSubmissionReceived struct {
	Submission *big.Int
	Round      *big.Int
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmissionReceived is a free log retrieval operation binding the contract event 0x22158772010459843096453be6465f97c8e8f17a50c4aa07ef3849c729010ab4.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint80 indexed round, address indexed oracle)
func (_ShadowAggregator *ShadowAggregatorFilterer) FilterSubmissionReceived(opts *bind.FilterOpts, submission []*big.Int, round []*big.Int, oracle []common.Address) (*ShadowAggregatorSubmissionReceivedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ShadowAggregator.contract.FilterLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &ShadowAggregatorSubmissionReceivedIterator{contract: _ShadowAggregator.contract, event: "SubmissionReceived", logs: logs, sub: sub}, nil
}

// WatchSubmissionReceived is a free log subscription operation binding the contract event 0x22158772010459843096453be6465f97c8e8f17a50c4aa07ef3849c729010ab4.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint80 indexed round, address indexed oracle)
func (_ShadowAggregator *ShadowAggregatorFilterer) WatchSubmissionReceived(opts *bind.WatchOpts, sink chan<- *ShadowAggregatorSubmissionReceived, submission []*big.Int, round []*big.Int, oracle []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ShadowAggregator.contract.WatchLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShadowAggregatorSubmissionReceived)
				if err := _ShadowAggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
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

// ParseSubmissionReceived is a log parse operation binding the contract event 0x22158772010459843096453be6465f97c8e8f17a50c4aa07ef3849c729010ab4.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint80 indexed round, address indexed oracle)
func (_ShadowAggregator *ShadowAggregatorFilterer) ParseSubmissionReceived(log types.Log) (*ShadowAggregatorSubmissionReceived, error) {
	event := new(ShadowAggregatorSubmissionReceived)
	if err := _ShadowAggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
