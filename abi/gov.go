// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovABI is the input ABI used to generate the binding from.
const GovABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"failReasonWithBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"failReasonWithStr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_TARGET_CONTRACT_FAIL\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_TARGET_NOT_CONTRACT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PARAM_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Gov is an auto generated Go binding around an Ethereum contract.
type Gov struct {
	GovCaller     // Read-only binding to the contract
	GovTransactor // Write-only binding to the contract
	GovFilterer   // Log filterer for contract events
}

// GovCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovSession struct {
	Contract     *Gov              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovCallerSession struct {
	Contract *GovCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovTransactorSession struct {
	Contract     *GovTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovRaw struct {
	Contract *Gov // Generic contract binding to access the raw methods on
}

// GovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovCallerRaw struct {
	Contract *GovCaller // Generic read-only contract binding to access the raw methods on
}

// GovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovTransactorRaw struct {
	Contract *GovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGov creates a new instance of Gov, bound to a specific deployed contract.
func NewGov(address common.Address, backend bind.ContractBackend) (*Gov, error) {
	contract, err := bindGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// NewGovCaller creates a new read-only instance of Gov, bound to a specific deployed contract.
func NewGovCaller(address common.Address, caller bind.ContractCaller) (*GovCaller, error) {
	contract, err := bindGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovCaller{contract: contract}, nil
}

// NewGovTransactor creates a new write-only instance of Gov, bound to a specific deployed contract.
func NewGovTransactor(address common.Address, transactor bind.ContractTransactor) (*GovTransactor, error) {
	contract, err := bindGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovTransactor{contract: contract}, nil
}

// NewGovFilterer creates a new log filterer instance of Gov, bound to a specific deployed contract.
func NewGovFilterer(address common.Address, filterer bind.ContractFilterer) (*GovFilterer, error) {
	contract, err := bindGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovFilterer{contract: contract}, nil
}

// bindGov binds a generic wrapper to an already deployed contract.
func bindGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.GovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) BINDCHANNELID() (uint8, error) {
	return _Gov.Contract.BINDCHANNELID(&_Gov.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) BINDCHANNELID() (uint8, error) {
	return _Gov.Contract.BINDCHANNELID(&_Gov.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Gov *GovCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Gov *GovSession) CODEOK() (uint32, error) {
	return _Gov.Contract.CODEOK(&_Gov.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Gov *GovCallerSession) CODEOK() (uint32, error) {
	return _Gov.Contract.CODEOK(&_Gov.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.CROSSCHAINCONTRACTADDR(&_Gov.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.CROSSCHAINCONTRACTADDR(&_Gov.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Gov *GovCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Gov *GovSession) ERRORFAILDECODE() (uint32, error) {
	return _Gov.Contract.ERRORFAILDECODE(&_Gov.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Gov *GovCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Gov.Contract.ERRORFAILDECODE(&_Gov.CallOpts)
}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() constant returns(uint32)
func (_Gov *GovCaller) ERRORTARGETCONTRACTFAIL(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "ERROR_TARGET_CONTRACT_FAIL")
	return *ret0, err
}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() constant returns(uint32)
func (_Gov *GovSession) ERRORTARGETCONTRACTFAIL() (uint32, error) {
	return _Gov.Contract.ERRORTARGETCONTRACTFAIL(&_Gov.CallOpts)
}

// ERRORTARGETCONTRACTFAIL is a free data retrieval call binding the contract method 0x3a21baae.
//
// Solidity: function ERROR_TARGET_CONTRACT_FAIL() constant returns(uint32)
func (_Gov *GovCallerSession) ERRORTARGETCONTRACTFAIL() (uint32, error) {
	return _Gov.Contract.ERRORTARGETCONTRACTFAIL(&_Gov.CallOpts)
}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() constant returns(uint32)
func (_Gov *GovCaller) ERRORTARGETNOTCONTRACT(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "ERROR_TARGET_NOT_CONTRACT")
	return *ret0, err
}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() constant returns(uint32)
func (_Gov *GovSession) ERRORTARGETNOTCONTRACT() (uint32, error) {
	return _Gov.Contract.ERRORTARGETNOTCONTRACT(&_Gov.CallOpts)
}

// ERRORTARGETNOTCONTRACT is a free data retrieval call binding the contract method 0x9ab1a373.
//
// Solidity: function ERROR_TARGET_NOT_CONTRACT() constant returns(uint32)
func (_Gov *GovCallerSession) ERRORTARGETNOTCONTRACT() (uint32, error) {
	return _Gov.Contract.ERRORTARGETNOTCONTRACT(&_Gov.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) GOVCHANNELID() (uint8, error) {
	return _Gov.Contract.GOVCHANNELID(&_Gov.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) GOVCHANNELID() (uint8, error) {
	return _Gov.Contract.GOVCHANNELID(&_Gov.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Gov *GovCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Gov *GovSession) GOVHUBADDR() (common.Address, error) {
	return _Gov.Contract.GOVHUBADDR(&_Gov.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Gov *GovCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Gov.Contract.GOVHUBADDR(&_Gov.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Gov *GovCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Gov *GovSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Gov.Contract.INCENTIVIZEADDR(&_Gov.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Gov *GovCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Gov.Contract.INCENTIVIZEADDR(&_Gov.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Gov *GovCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Gov *GovSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Gov.Contract.LIGHTCLIENTADDR(&_Gov.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Gov *GovCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Gov.Contract.LIGHTCLIENTADDR(&_Gov.CallOpts)
}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Gov *GovCaller) PARAMUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "PARAM_UPDATE_MESSAGE_TYPE")
	return *ret0, err
}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Gov *GovSession) PARAMUPDATEMESSAGETYPE() (uint8, error) {
	return _Gov.Contract.PARAMUPDATEMESSAGETYPE(&_Gov.CallOpts)
}

// PARAMUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x4900c4ea.
//
// Solidity: function PARAM_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Gov *GovCallerSession) PARAMUPDATEMESSAGETYPE() (uint8, error) {
	return _Gov.Contract.PARAMUPDATEMESSAGETYPE(&_Gov.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.RELAYERHUBCONTRACTADDR(&_Gov.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.RELAYERHUBCONTRACTADDR(&_Gov.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) SLASHCHANNELID() (uint8, error) {
	return _Gov.Contract.SLASHCHANNELID(&_Gov.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Gov.Contract.SLASHCHANNELID(&_Gov.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.SLASHCONTRACTADDR(&_Gov.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.SLASHCONTRACTADDR(&_Gov.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) STAKINGCHANNELID() (uint8, error) {
	return _Gov.Contract.STAKINGCHANNELID(&_Gov.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Gov.Contract.STAKINGCHANNELID(&_Gov.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Gov *GovCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Gov *GovSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Gov.Contract.SYSTEMREWARDADDR(&_Gov.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Gov *GovCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Gov.Contract.SYSTEMREWARDADDR(&_Gov.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Gov *GovCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Gov *GovSession) TOKENHUBADDR() (common.Address, error) {
	return _Gov.Contract.TOKENHUBADDR(&_Gov.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Gov *GovCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Gov.Contract.TOKENHUBADDR(&_Gov.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Gov.Contract.TRANSFERINCHANNELID(&_Gov.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Gov.Contract.TRANSFERINCHANNELID(&_Gov.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Gov *GovCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Gov *GovSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Gov.Contract.TRANSFEROUTCHANNELID(&_Gov.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Gov *GovCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Gov.Contract.TRANSFEROUTCHANNELID(&_Gov.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.VALIDATORCONTRACTADDR(&_Gov.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Gov *GovCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Gov.Contract.VALIDATORCONTRACTADDR(&_Gov.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Gov *GovCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Gov *GovSession) AlreadyInit() (bool, error) {
	return _Gov.Contract.AlreadyInit(&_Gov.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Gov *GovCallerSession) AlreadyInit() (bool, error) {
	return _Gov.Contract.AlreadyInit(&_Gov.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Gov *GovCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Gov *GovSession) BscChainID() (uint16, error) {
	return _Gov.Contract.BscChainID(&_Gov.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Gov *GovCallerSession) BscChainID() (uint16, error) {
	return _Gov.Contract.BscChainID(&_Gov.CallOpts)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Gov *GovTransactor) HandleAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "handleAckPackage", arg0, arg1)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Gov *GovSession) HandleAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleAckPackage(&_Gov.TransactOpts, arg0, arg1)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes ) returns()
func (_Gov *GovTransactorSession) HandleAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleAckPackage(&_Gov.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Gov *GovTransactor) HandleFailAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "handleFailAckPackage", arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Gov *GovSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleFailAckPackage(&_Gov.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Gov *GovTransactorSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleFailAckPackage(&_Gov.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Gov *GovTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Gov *GovSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleSynPackage(&_Gov.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_Gov *GovTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Gov.Contract.HandleSynPackage(&_Gov.TransactOpts, arg0, msgBytes)
}

// GovFailReasonWithBytesIterator is returned from FilterFailReasonWithBytes and is used to iterate over the raw logs and unpacked data for FailReasonWithBytes events raised by the Gov contract.
type GovFailReasonWithBytesIterator struct {
	Event *GovFailReasonWithBytes // Event containing the contract specifics and raw log

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
func (it *GovFailReasonWithBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovFailReasonWithBytes)
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
		it.Event = new(GovFailReasonWithBytes)
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
func (it *GovFailReasonWithBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovFailReasonWithBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovFailReasonWithBytes represents a FailReasonWithBytes event raised by the Gov contract.
type GovFailReasonWithBytes struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithBytes is a free log retrieval operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Gov *GovFilterer) FilterFailReasonWithBytes(opts *bind.FilterOpts) (*GovFailReasonWithBytesIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "failReasonWithBytes")
	if err != nil {
		return nil, err
	}
	return &GovFailReasonWithBytesIterator{contract: _Gov.contract, event: "failReasonWithBytes", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithBytes is a free log subscription operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Gov *GovFilterer) WatchFailReasonWithBytes(opts *bind.WatchOpts, sink chan<- *GovFailReasonWithBytes) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "failReasonWithBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovFailReasonWithBytes)
				if err := _Gov.contract.UnpackLog(event, "failReasonWithBytes", log); err != nil {
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

// ParseFailReasonWithBytes is a log parse operation binding the contract event 0x1279f84165b4fd69c35e1f338ff107231b036c655cd1688851e011ce617c4e8d.
//
// Solidity: event failReasonWithBytes(bytes message)
func (_Gov *GovFilterer) ParseFailReasonWithBytes(log types.Log) (*GovFailReasonWithBytes, error) {
	event := new(GovFailReasonWithBytes)
	if err := _Gov.contract.UnpackLog(event, "failReasonWithBytes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the Gov contract.
type GovFailReasonWithStrIterator struct {
	Event *GovFailReasonWithStr // Event containing the contract specifics and raw log

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
func (it *GovFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovFailReasonWithStr)
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
		it.Event = new(GovFailReasonWithStr)
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
func (it *GovFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovFailReasonWithStr represents a FailReasonWithStr event raised by the Gov contract.
type GovFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Gov *GovFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*GovFailReasonWithStrIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &GovFailReasonWithStrIterator{contract: _Gov.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Gov *GovFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *GovFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovFailReasonWithStr)
				if err := _Gov.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
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

// ParseFailReasonWithStr is a log parse operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_Gov *GovFilterer) ParseFailReasonWithStr(log types.Log) (*GovFailReasonWithStr, error) {
	event := new(GovFailReasonWithStr)
	if err := _Gov.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Gov contract.
type GovParamChangeIterator struct {
	Event *GovParamChange // Event containing the contract specifics and raw log

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
func (it *GovParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamChange)
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
		it.Event = new(GovParamChange)
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
func (it *GovParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamChange represents a ParamChange event raised by the Gov contract.
type GovParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Gov *GovFilterer) FilterParamChange(opts *bind.FilterOpts) (*GovParamChangeIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &GovParamChangeIterator{contract: _Gov.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Gov *GovFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *GovParamChange) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamChange)
				if err := _Gov.contract.UnpackLog(event, "paramChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Gov *GovFilterer) ParseParamChange(log types.Log) (*GovParamChange, error) {
	event := new(GovParamChange)
	if err := _Gov.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}
