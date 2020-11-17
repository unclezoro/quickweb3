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

// WbnbABI is the input ABI used to generate the binding from.
const WbnbABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"}]"

// Wbnb is an auto generated Go binding around an Ethereum contract.
type Wbnb struct {
	WbnbCaller     // Read-only binding to the contract
	WbnbTransactor // Write-only binding to the contract
	WbnbFilterer   // Log filterer for contract events
}

// WbnbCaller is an auto generated read-only Go binding around an Ethereum contract.
type WbnbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WbnbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WbnbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WbnbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WbnbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WbnbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WbnbSession struct {
	Contract     *Wbnb             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WbnbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WbnbCallerSession struct {
	Contract *WbnbCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WbnbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WbnbTransactorSession struct {
	Contract     *WbnbTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WbnbRaw is an auto generated low-level Go binding around an Ethereum contract.
type WbnbRaw struct {
	Contract *Wbnb // Generic contract binding to access the raw methods on
}

// WbnbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WbnbCallerRaw struct {
	Contract *WbnbCaller // Generic read-only contract binding to access the raw methods on
}

// WbnbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WbnbTransactorRaw struct {
	Contract *WbnbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWbnb creates a new instance of Wbnb, bound to a specific deployed contract.
func NewWbnb(address common.Address, backend bind.ContractBackend) (*Wbnb, error) {
	contract, err := bindWbnb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wbnb{WbnbCaller: WbnbCaller{contract: contract}, WbnbTransactor: WbnbTransactor{contract: contract}, WbnbFilterer: WbnbFilterer{contract: contract}}, nil
}

// NewWbnbCaller creates a new read-only instance of Wbnb, bound to a specific deployed contract.
func NewWbnbCaller(address common.Address, caller bind.ContractCaller) (*WbnbCaller, error) {
	contract, err := bindWbnb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WbnbCaller{contract: contract}, nil
}

// NewWbnbTransactor creates a new write-only instance of Wbnb, bound to a specific deployed contract.
func NewWbnbTransactor(address common.Address, transactor bind.ContractTransactor) (*WbnbTransactor, error) {
	contract, err := bindWbnb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WbnbTransactor{contract: contract}, nil
}

// NewWbnbFilterer creates a new log filterer instance of Wbnb, bound to a specific deployed contract.
func NewWbnbFilterer(address common.Address, filterer bind.ContractFilterer) (*WbnbFilterer, error) {
	contract, err := bindWbnb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WbnbFilterer{contract: contract}, nil
}

// bindWbnb binds a generic wrapper to an already deployed contract.
func bindWbnb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WbnbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wbnb *WbnbRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wbnb.Contract.WbnbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wbnb *WbnbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wbnb.Contract.WbnbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wbnb *WbnbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wbnb.Contract.WbnbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wbnb *WbnbCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wbnb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wbnb *WbnbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wbnb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wbnb *WbnbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wbnb.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Wbnb *WbnbCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Wbnb *WbnbSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Wbnb.Contract.Allowance(&_Wbnb.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Wbnb *WbnbCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Wbnb.Contract.Allowance(&_Wbnb.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Wbnb *WbnbCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Wbnb *WbnbSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Wbnb.Contract.BalanceOf(&_Wbnb.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Wbnb *WbnbCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Wbnb.Contract.BalanceOf(&_Wbnb.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wbnb *WbnbCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wbnb *WbnbSession) Decimals() (uint8, error) {
	return _Wbnb.Contract.Decimals(&_Wbnb.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Wbnb *WbnbCallerSession) Decimals() (uint8, error) {
	return _Wbnb.Contract.Decimals(&_Wbnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wbnb *WbnbCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wbnb *WbnbSession) Name() (string, error) {
	return _Wbnb.Contract.Name(&_Wbnb.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Wbnb *WbnbCallerSession) Name() (string, error) {
	return _Wbnb.Contract.Name(&_Wbnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wbnb *WbnbCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wbnb *WbnbSession) Symbol() (string, error) {
	return _Wbnb.Contract.Symbol(&_Wbnb.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Wbnb *WbnbCallerSession) Symbol() (string, error) {
	return _Wbnb.Contract.Symbol(&_Wbnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wbnb *WbnbCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wbnb.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wbnb *WbnbSession) TotalSupply() (*big.Int, error) {
	return _Wbnb.Contract.TotalSupply(&_Wbnb.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Wbnb *WbnbCallerSession) TotalSupply() (*big.Int, error) {
	return _Wbnb.Contract.TotalSupply(&_Wbnb.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Wbnb *WbnbSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Approve(&_Wbnb.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Approve(&_Wbnb.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Wbnb *WbnbTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wbnb.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Wbnb *WbnbSession) Deposit() (*types.Transaction, error) {
	return _Wbnb.Contract.Deposit(&_Wbnb.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Wbnb *WbnbTransactorSession) Deposit() (*types.Transaction, error) {
	return _Wbnb.Contract.Deposit(&_Wbnb.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Transfer(&_Wbnb.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Transfer(&_Wbnb.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.TransferFrom(&_Wbnb.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Wbnb *WbnbTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.TransferFrom(&_Wbnb.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Wbnb *WbnbTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Wbnb *WbnbSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Withdraw(&_Wbnb.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Wbnb *WbnbTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _Wbnb.Contract.Withdraw(&_Wbnb.TransactOpts, wad)
}

// WbnbApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wbnb contract.
type WbnbApprovalIterator struct {
	Event *WbnbApproval // Event containing the contract specifics and raw log

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
func (it *WbnbApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WbnbApproval)
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
		it.Event = new(WbnbApproval)
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
func (it *WbnbApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WbnbApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WbnbApproval represents a Approval event raised by the Wbnb contract.
type WbnbApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Wbnb *WbnbFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WbnbApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Wbnb.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WbnbApprovalIterator{contract: _Wbnb.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Wbnb *WbnbFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WbnbApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Wbnb.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WbnbApproval)
				if err := _Wbnb.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Wbnb *WbnbFilterer) ParseApproval(log types.Log) (*WbnbApproval, error) {
	event := new(WbnbApproval)
	if err := _Wbnb.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WbnbDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Wbnb contract.
type WbnbDepositIterator struct {
	Event *WbnbDeposit // Event containing the contract specifics and raw log

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
func (it *WbnbDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WbnbDeposit)
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
		it.Event = new(WbnbDeposit)
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
func (it *WbnbDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WbnbDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WbnbDeposit represents a Deposit event raised by the Wbnb contract.
type WbnbDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WbnbDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Wbnb.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WbnbDepositIterator{contract: _Wbnb.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WbnbDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Wbnb.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WbnbDeposit)
				if err := _Wbnb.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) ParseDeposit(log types.Log) (*WbnbDeposit, error) {
	event := new(WbnbDeposit)
	if err := _Wbnb.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WbnbTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wbnb contract.
type WbnbTransferIterator struct {
	Event *WbnbTransfer // Event containing the contract specifics and raw log

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
func (it *WbnbTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WbnbTransfer)
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
		it.Event = new(WbnbTransfer)
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
func (it *WbnbTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WbnbTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WbnbTransfer represents a Transfer event raised by the Wbnb contract.
type WbnbTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WbnbTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Wbnb.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WbnbTransferIterator{contract: _Wbnb.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WbnbTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Wbnb.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WbnbTransfer)
				if err := _Wbnb.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Wbnb *WbnbFilterer) ParseTransfer(log types.Log) (*WbnbTransfer, error) {
	event := new(WbnbTransfer)
	if err := _Wbnb.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WbnbWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Wbnb contract.
type WbnbWithdrawalIterator struct {
	Event *WbnbWithdrawal // Event containing the contract specifics and raw log

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
func (it *WbnbWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WbnbWithdrawal)
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
		it.Event = new(WbnbWithdrawal)
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
func (it *WbnbWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WbnbWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WbnbWithdrawal represents a Withdrawal event raised by the Wbnb contract.
type WbnbWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_Wbnb *WbnbFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WbnbWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _Wbnb.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WbnbWithdrawalIterator{contract: _Wbnb.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_Wbnb *WbnbFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WbnbWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _Wbnb.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WbnbWithdrawal)
				if err := _Wbnb.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_Wbnb *WbnbFilterer) ParseWithdrawal(log types.Log) (*WbnbWithdrawal, error) {
	event := new(WbnbWithdrawal)
	if err := _Wbnb.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
