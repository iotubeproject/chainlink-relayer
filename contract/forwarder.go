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

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"OwnershipTransferRequestedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"forward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"multiForward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ownerForward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"setAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"transferOwnershipWithMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Forwarder *ForwarderCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Forwarder *ForwarderSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Forwarder.Contract.GetAuthorizedSenders(&_Forwarder.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Forwarder *ForwarderCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Forwarder.Contract.GetAuthorizedSenders(&_Forwarder.CallOpts)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Forwarder *ForwarderCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Forwarder *ForwarderSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Forwarder.Contract.IsAuthorizedSender(&_Forwarder.CallOpts, sender)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Forwarder *ForwarderCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Forwarder.Contract.IsAuthorizedSender(&_Forwarder.CallOpts, sender)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Forwarder *ForwarderCaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "linkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Forwarder *ForwarderSession) LinkToken() (common.Address, error) {
	return _Forwarder.Contract.LinkToken(&_Forwarder.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Forwarder *ForwarderCallerSession) LinkToken() (common.Address, error) {
	return _Forwarder.Contract.LinkToken(&_Forwarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Forwarder *ForwarderCallerSession) Owner() (common.Address, error) {
	return _Forwarder.Contract.Owner(&_Forwarder.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_Forwarder *ForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_Forwarder *ForwarderSession) TypeAndVersion() (string, error) {
	return _Forwarder.Contract.TypeAndVersion(&_Forwarder.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_Forwarder *ForwarderCallerSession) TypeAndVersion() (string, error) {
	return _Forwarder.Contract.TypeAndVersion(&_Forwarder.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Forwarder *ForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Forwarder *ForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.AcceptOwnership(&_Forwarder.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Forwarder *ForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Forwarder.Contract.AcceptOwnership(&_Forwarder.TransactOpts)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Forwarder *ForwarderTransactor) Forward(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "forward", to, data)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Forwarder *ForwarderSession) Forward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Forward(&_Forwarder.TransactOpts, to, data)
}

// Forward is a paid mutator transaction binding the contract method 0x6fadcf72.
//
// Solidity: function forward(address to, bytes data) returns()
func (_Forwarder *ForwarderTransactorSession) Forward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.Forward(&_Forwarder.TransactOpts, to, data)
}

// MultiForward is a paid mutator transaction binding the contract method 0xb64fa9e6.
//
// Solidity: function multiForward(address[] tos, bytes[] datas) returns()
func (_Forwarder *ForwarderTransactor) MultiForward(opts *bind.TransactOpts, tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "multiForward", tos, datas)
}

// MultiForward is a paid mutator transaction binding the contract method 0xb64fa9e6.
//
// Solidity: function multiForward(address[] tos, bytes[] datas) returns()
func (_Forwarder *ForwarderSession) MultiForward(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Forwarder.Contract.MultiForward(&_Forwarder.TransactOpts, tos, datas)
}

// MultiForward is a paid mutator transaction binding the contract method 0xb64fa9e6.
//
// Solidity: function multiForward(address[] tos, bytes[] datas) returns()
func (_Forwarder *ForwarderTransactorSession) MultiForward(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Forwarder.Contract.MultiForward(&_Forwarder.TransactOpts, tos, datas)
}

// OwnerForward is a paid mutator transaction binding the contract method 0x033f49f7.
//
// Solidity: function ownerForward(address to, bytes data) returns()
func (_Forwarder *ForwarderTransactor) OwnerForward(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "ownerForward", to, data)
}

// OwnerForward is a paid mutator transaction binding the contract method 0x033f49f7.
//
// Solidity: function ownerForward(address to, bytes data) returns()
func (_Forwarder *ForwarderSession) OwnerForward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.OwnerForward(&_Forwarder.TransactOpts, to, data)
}

// OwnerForward is a paid mutator transaction binding the contract method 0x033f49f7.
//
// Solidity: function ownerForward(address to, bytes data) returns()
func (_Forwarder *ForwarderTransactorSession) OwnerForward(to common.Address, data []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.OwnerForward(&_Forwarder.TransactOpts, to, data)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_Forwarder *ForwarderTransactor) SetAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "setAuthorizedSenders", senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_Forwarder *ForwarderSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.SetAuthorizedSenders(&_Forwarder.TransactOpts, senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_Forwarder *ForwarderTransactorSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.SetAuthorizedSenders(&_Forwarder.TransactOpts, senders)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Forwarder *ForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Forwarder *ForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Forwarder *ForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnership(&_Forwarder.TransactOpts, to)
}

// TransferOwnershipWithMessage is a paid mutator transaction binding the contract method 0x4d3e2323.
//
// Solidity: function transferOwnershipWithMessage(address to, bytes message) returns()
func (_Forwarder *ForwarderTransactor) TransferOwnershipWithMessage(opts *bind.TransactOpts, to common.Address, message []byte) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "transferOwnershipWithMessage", to, message)
}

// TransferOwnershipWithMessage is a paid mutator transaction binding the contract method 0x4d3e2323.
//
// Solidity: function transferOwnershipWithMessage(address to, bytes message) returns()
func (_Forwarder *ForwarderSession) TransferOwnershipWithMessage(to common.Address, message []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnershipWithMessage(&_Forwarder.TransactOpts, to, message)
}

// TransferOwnershipWithMessage is a paid mutator transaction binding the contract method 0x4d3e2323.
//
// Solidity: function transferOwnershipWithMessage(address to, bytes message) returns()
func (_Forwarder *ForwarderTransactorSession) TransferOwnershipWithMessage(to common.Address, message []byte) (*types.Transaction, error) {
	return _Forwarder.Contract.TransferOwnershipWithMessage(&_Forwarder.TransactOpts, to, message)
}

// ForwarderAuthorizedSendersChangedIterator is returned from FilterAuthorizedSendersChanged and is used to iterate over the raw logs and unpacked data for AuthorizedSendersChanged events raised by the Forwarder contract.
type ForwarderAuthorizedSendersChangedIterator struct {
	Event *ForwarderAuthorizedSendersChanged // Event containing the contract specifics and raw log

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
func (it *ForwarderAuthorizedSendersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderAuthorizedSendersChanged)
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
		it.Event = new(ForwarderAuthorizedSendersChanged)
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
func (it *ForwarderAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderAuthorizedSendersChanged represents a AuthorizedSendersChanged event raised by the Forwarder contract.
type ForwarderAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersChanged is a free log retrieval operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Forwarder *ForwarderFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*ForwarderAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &ForwarderAuthorizedSendersChangedIterator{contract: _Forwarder.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersChanged is a free log subscription operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Forwarder *ForwarderFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *ForwarderAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderAuthorizedSendersChanged)
				if err := _Forwarder.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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

// ParseAuthorizedSendersChanged is a log parse operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Forwarder *ForwarderFilterer) ParseAuthorizedSendersChanged(log types.Log) (*ForwarderAuthorizedSendersChanged, error) {
	event := new(ForwarderAuthorizedSendersChanged)
	if err := _Forwarder.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwarderOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Forwarder contract.
type ForwarderOwnershipTransferRequestedIterator struct {
	Event *ForwarderOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ForwarderOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferRequested)
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
		it.Event = new(ForwarderOwnershipTransferRequested)
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
func (it *ForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Forwarder contract.
type ForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferRequestedIterator{contract: _Forwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderOwnershipTransferRequested)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*ForwarderOwnershipTransferRequested, error) {
	event := new(ForwarderOwnershipTransferRequested)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwarderOwnershipTransferRequestedWithMessageIterator is returned from FilterOwnershipTransferRequestedWithMessage and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequestedWithMessage events raised by the Forwarder contract.
type ForwarderOwnershipTransferRequestedWithMessageIterator struct {
	Event *ForwarderOwnershipTransferRequestedWithMessage // Event containing the contract specifics and raw log

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
func (it *ForwarderOwnershipTransferRequestedWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferRequestedWithMessage)
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
		it.Event = new(ForwarderOwnershipTransferRequestedWithMessage)
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
func (it *ForwarderOwnershipTransferRequestedWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderOwnershipTransferRequestedWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderOwnershipTransferRequestedWithMessage represents a OwnershipTransferRequestedWithMessage event raised by the Forwarder contract.
type ForwarderOwnershipTransferRequestedWithMessage struct {
	From    common.Address
	To      common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequestedWithMessage is a free log retrieval operation binding the contract event 0x4e1e878dc28d5f040db5969163ff1acd75c44c3f655da2dde9c70bbd8e56dc7e.
//
// Solidity: event OwnershipTransferRequestedWithMessage(address indexed from, address indexed to, bytes message)
func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferRequestedWithMessage(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferRequestedWithMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferRequestedWithMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferRequestedWithMessageIterator{contract: _Forwarder.contract, event: "OwnershipTransferRequestedWithMessage", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequestedWithMessage is a free log subscription operation binding the contract event 0x4e1e878dc28d5f040db5969163ff1acd75c44c3f655da2dde9c70bbd8e56dc7e.
//
// Solidity: event OwnershipTransferRequestedWithMessage(address indexed from, address indexed to, bytes message)
func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferRequestedWithMessage(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferRequestedWithMessage, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferRequestedWithMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderOwnershipTransferRequestedWithMessage)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequestedWithMessage", log); err != nil {
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

// ParseOwnershipTransferRequestedWithMessage is a log parse operation binding the contract event 0x4e1e878dc28d5f040db5969163ff1acd75c44c3f655da2dde9c70bbd8e56dc7e.
//
// Solidity: event OwnershipTransferRequestedWithMessage(address indexed from, address indexed to, bytes message)
func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferRequestedWithMessage(log types.Log) (*ForwarderOwnershipTransferRequestedWithMessage, error) {
	event := new(ForwarderOwnershipTransferRequestedWithMessage)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferRequestedWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForwarderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Forwarder contract.
type ForwarderOwnershipTransferredIterator struct {
	Event *ForwarderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForwarderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderOwnershipTransferred)
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
		it.Event = new(ForwarderOwnershipTransferred)
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
func (it *ForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderOwnershipTransferred represents a OwnershipTransferred event raised by the Forwarder contract.
type ForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Forwarder *ForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ForwarderOwnershipTransferredIterator{contract: _Forwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Forwarder *ForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderOwnershipTransferred)
				if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Forwarder *ForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*ForwarderOwnershipTransferred, error) {
	event := new(ForwarderOwnershipTransferred)
	if err := _Forwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
