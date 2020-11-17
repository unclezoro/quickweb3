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

// MulticallCall is an auto generated low-level Go binding around an user-defined struct.
type MulticallCall struct {
	Target   common.Address
	CallData []byte
}

// TimelockABI is the input ABI used to generate the binding from.
const TimelockABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Timelock is an auto generated Go binding around an Ethereum contract.
type Timelock struct {
	TimelockCaller     // Read-only binding to the contract
	TimelockTransactor // Write-only binding to the contract
	TimelockFilterer   // Log filterer for contract events
}

// TimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimelockSession struct {
	Contract     *Timelock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimelockCallerSession struct {
	Contract *TimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimelockTransactorSession struct {
	Contract     *TimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimelockRaw struct {
	Contract *Timelock // Generic contract binding to access the raw methods on
}

// TimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimelockCallerRaw struct {
	Contract *TimelockCaller // Generic read-only contract binding to access the raw methods on
}

// TimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimelockTransactorRaw struct {
	Contract *TimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimelock creates a new instance of Timelock, bound to a specific deployed contract.
func NewTimelock(address common.Address, backend bind.ContractBackend) (*Timelock, error) {
	contract, err := bindTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timelock{TimelockCaller: TimelockCaller{contract: contract}, TimelockTransactor: TimelockTransactor{contract: contract}, TimelockFilterer: TimelockFilterer{contract: contract}}, nil
}

// NewTimelockCaller creates a new read-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockCaller(address common.Address, caller bind.ContractCaller) (*TimelockCaller, error) {
	contract, err := bindTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockCaller{contract: contract}, nil
}

// NewTimelockTransactor creates a new write-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*TimelockTransactor, error) {
	contract, err := bindTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockTransactor{contract: contract}, nil
}

// NewTimelockFilterer creates a new log filterer instance of Timelock, bound to a specific deployed contract.
func NewTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*TimelockFilterer, error) {
	contract, err := bindTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimelockFilterer{contract: contract}, nil
}

// bindTimelock binds a generic wrapper to an already deployed contract.
func bindTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.TimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) constant returns(bytes32 blockHash)
func (_Timelock *TimelockCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getBlockHash", blockNumber)
	return *ret0, err
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) constant returns(bytes32 blockHash)
func (_Timelock *TimelockSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Timelock.Contract.GetBlockHash(&_Timelock.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) constant returns(bytes32 blockHash)
func (_Timelock *TimelockCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _Timelock.Contract.GetBlockHash(&_Timelock.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() constant returns(address coinbase)
func (_Timelock *TimelockCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getCurrentBlockCoinbase")
	return *ret0, err
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() constant returns(address coinbase)
func (_Timelock *TimelockSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Timelock.Contract.GetCurrentBlockCoinbase(&_Timelock.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() constant returns(address coinbase)
func (_Timelock *TimelockCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _Timelock.Contract.GetCurrentBlockCoinbase(&_Timelock.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() constant returns(uint256 difficulty)
func (_Timelock *TimelockCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getCurrentBlockDifficulty")
	return *ret0, err
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() constant returns(uint256 difficulty)
func (_Timelock *TimelockSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockDifficulty(&_Timelock.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() constant returns(uint256 difficulty)
func (_Timelock *TimelockCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockDifficulty(&_Timelock.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() constant returns(uint256 gaslimit)
func (_Timelock *TimelockCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getCurrentBlockGasLimit")
	return *ret0, err
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() constant returns(uint256 gaslimit)
func (_Timelock *TimelockSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockGasLimit(&_Timelock.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() constant returns(uint256 gaslimit)
func (_Timelock *TimelockCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockGasLimit(&_Timelock.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() constant returns(uint256 timestamp)
func (_Timelock *TimelockCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getCurrentBlockTimestamp")
	return *ret0, err
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() constant returns(uint256 timestamp)
func (_Timelock *TimelockSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockTimestamp(&_Timelock.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() constant returns(uint256 timestamp)
func (_Timelock *TimelockCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _Timelock.Contract.GetCurrentBlockTimestamp(&_Timelock.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) constant returns(uint256 balance)
func (_Timelock *TimelockCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getEthBalance", addr)
	return *ret0, err
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) constant returns(uint256 balance)
func (_Timelock *TimelockSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Timelock.Contract.GetEthBalance(&_Timelock.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) constant returns(uint256 balance)
func (_Timelock *TimelockCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _Timelock.Contract.GetEthBalance(&_Timelock.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() constant returns(bytes32 blockHash)
func (_Timelock *TimelockCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Timelock.contract.Call(opts, out, "getLastBlockHash")
	return *ret0, err
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() constant returns(bytes32 blockHash)
func (_Timelock *TimelockSession) GetLastBlockHash() ([32]byte, error) {
	return _Timelock.Contract.GetLastBlockHash(&_Timelock.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() constant returns(bytes32 blockHash)
func (_Timelock *TimelockCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _Timelock.Contract.GetLastBlockHash(&_Timelock.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate([]MulticallCall calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Timelock *TimelockTransactor) Aggregate(opts *bind.TransactOpts, calls []MulticallCall) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate([]MulticallCall calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Timelock *TimelockSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _Timelock.Contract.Aggregate(&_Timelock.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate([]MulticallCall calls) returns(uint256 blockNumber, bytes[] returnData)
func (_Timelock *TimelockTransactorSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _Timelock.Contract.Aggregate(&_Timelock.TransactOpts, calls)
}
