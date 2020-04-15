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

// SystemRewardABI is the input ABI used to generate the binding from.
const SystemRewardABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DeleteOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceiveDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RewardEmpty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardTo\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REWARDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_SET_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemReward is an auto generated Go binding around an Ethereum contract.
type SystemReward struct {
	SystemRewardCaller     // Read-only binding to the contract
	SystemRewardTransactor // Write-only binding to the contract
	SystemRewardFilterer   // Log filterer for contract events
}

// SystemRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRewardSession struct {
	Contract     *SystemReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRewardCallerSession struct {
	Contract *SystemRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRewardTransactorSession struct {
	Contract     *SystemRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRewardRaw struct {
	Contract *SystemReward // Generic contract binding to access the raw methods on
}

// SystemRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRewardCallerRaw struct {
	Contract *SystemRewardCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRewardTransactorRaw struct {
	Contract *SystemRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemReward creates a new instance of SystemReward, bound to a specific deployed contract.
func NewSystemReward(address common.Address, backend bind.ContractBackend) (*SystemReward, error) {
	contract, err := bindSystemReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemReward{SystemRewardCaller: SystemRewardCaller{contract: contract}, SystemRewardTransactor: SystemRewardTransactor{contract: contract}, SystemRewardFilterer: SystemRewardFilterer{contract: contract}}, nil
}

// NewSystemRewardCaller creates a new read-only instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardCaller(address common.Address, caller bind.ContractCaller) (*SystemRewardCaller, error) {
	contract, err := bindSystemReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRewardCaller{contract: contract}, nil
}

// NewSystemRewardTransactor creates a new write-only instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRewardTransactor, error) {
	contract, err := bindSystemReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRewardTransactor{contract: contract}, nil
}

// NewSystemRewardFilterer creates a new log filterer instance of SystemReward, bound to a specific deployed contract.
func NewSystemRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRewardFilterer, error) {
	contract, err := bindSystemReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRewardFilterer{contract: contract}, nil
}

// bindSystemReward binds a generic wrapper to an already deployed contract.
func bindSystemReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReward *SystemRewardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemReward.Contract.SystemRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReward *SystemRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.Contract.SystemRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReward *SystemRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReward.Contract.SystemRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemReward *SystemRewardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemReward *SystemRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemReward *SystemRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemReward.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCaller) LIGHTCLIENTCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "LIGHT_CLIENT_CONTRACT")
	return *ret0, err
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.LIGHTCLIENTCONTRACT(&_SystemReward.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCallerSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.LIGHTCLIENTCONTRACT(&_SystemReward.CallOpts)
}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() constant returns(uint256)
func (_SystemReward *SystemRewardCaller) MAXREWARDS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "MAX_REWARDS")
	return *ret0, err
}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() constant returns(uint256)
func (_SystemReward *SystemRewardSession) MAXREWARDS() (*big.Int, error) {
	return _SystemReward.Contract.MAXREWARDS(&_SystemReward.CallOpts)
}

// MAXREWARDS is a free data retrieval call binding the contract method 0xfb5478b3.
//
// Solidity: function MAX_REWARDS() constant returns(uint256)
func (_SystemReward *SystemRewardCallerSession) MAXREWARDS() (*big.Int, error) {
	return _SystemReward.Contract.MAXREWARDS(&_SystemReward.CallOpts)
}

// TOKENHUBCONTRACT is a free data retrieval call binding the contract method 0x1a93b73f.
//
// Solidity: function TOKEN_HUB_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCaller) TOKENHUBCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "TOKEN_HUB_CONTRACT")
	return *ret0, err
}

// TOKENHUBCONTRACT is a free data retrieval call binding the contract method 0x1a93b73f.
//
// Solidity: function TOKEN_HUB_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardSession) TOKENHUBCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.TOKENHUBCONTRACT(&_SystemReward.CallOpts)
}

// TOKENHUBCONTRACT is a free data retrieval call binding the contract method 0x1a93b73f.
//
// Solidity: function TOKEN_HUB_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCallerSession) TOKENHUBCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.TOKENHUBCONTRACT(&_SystemReward.CallOpts)
}

// VALIDATORSETCONTRACT is a free data retrieval call binding the contract method 0x76137dd9.
//
// Solidity: function VALIDATOR_SET_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCaller) VALIDATORSETCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "VALIDATOR_SET_CONTRACT")
	return *ret0, err
}

// VALIDATORSETCONTRACT is a free data retrieval call binding the contract method 0x76137dd9.
//
// Solidity: function VALIDATOR_SET_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardSession) VALIDATORSETCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.VALIDATORSETCONTRACT(&_SystemReward.CallOpts)
}

// VALIDATORSETCONTRACT is a free data retrieval call binding the contract method 0x76137dd9.
//
// Solidity: function VALIDATOR_SET_CONTRACT() constant returns(address)
func (_SystemReward *SystemRewardCallerSession) VALIDATORSETCONTRACT() (common.Address, error) {
	return _SystemReward.Contract.VALIDATORSETCONTRACT(&_SystemReward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_SystemReward *SystemRewardCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_SystemReward *SystemRewardSession) AlreadyInit() (bool, error) {
	return _SystemReward.Contract.AlreadyInit(&_SystemReward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_SystemReward *SystemRewardCallerSession) AlreadyInit() (bool, error) {
	return _SystemReward.Contract.AlreadyInit(&_SystemReward.CallOpts)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) constant returns(bool)
func (_SystemReward *SystemRewardCaller) IsOperator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "isOperator", addr)
	return *ret0, err
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) constant returns(bool)
func (_SystemReward *SystemRewardSession) IsOperator(addr common.Address) (bool, error) {
	return _SystemReward.Contract.IsOperator(&_SystemReward.CallOpts, addr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address addr) constant returns(bool)
func (_SystemReward *SystemRewardCallerSession) IsOperator(addr common.Address) (bool, error) {
	return _SystemReward.Contract.IsOperator(&_SystemReward.CallOpts, addr)
}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() constant returns(uint256)
func (_SystemReward *SystemRewardCaller) NumOperator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemReward.contract.Call(opts, out, "numOperator")
	return *ret0, err
}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() constant returns(uint256)
func (_SystemReward *SystemRewardSession) NumOperator() (*big.Int, error) {
	return _SystemReward.Contract.NumOperator(&_SystemReward.CallOpts)
}

// NumOperator is a free data retrieval call binding the contract method 0x3a0b0eff.
//
// Solidity: function numOperator() constant returns(uint256)
func (_SystemReward *SystemRewardCallerSession) NumOperator() (*big.Int, error) {
	return _SystemReward.Contract.NumOperator(&_SystemReward.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_SystemReward *SystemRewardTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "addOperator", operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_SystemReward *SystemRewardSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.AddOperator(&_SystemReward.TransactOpts, operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator) returns()
func (_SystemReward *SystemRewardTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.AddOperator(&_SystemReward.TransactOpts, operator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_SystemReward *SystemRewardTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_SystemReward *SystemRewardSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SystemReward.Contract.ClaimRewards(&_SystemReward.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_SystemReward *SystemRewardTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SystemReward.Contract.ClaimRewards(&_SystemReward.TransactOpts, to, amount)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SystemReward *SystemRewardTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _SystemReward.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SystemReward *SystemRewardSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.RemoveOperator(&_SystemReward.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_SystemReward *SystemRewardTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _SystemReward.Contract.RemoveOperator(&_SystemReward.TransactOpts, operator)
}

// SystemRewardDeleteOperatorIterator is returned from FilterDeleteOperator and is used to iterate over the raw logs and unpacked data for DeleteOperator events raised by the SystemReward contract.
type SystemRewardDeleteOperatorIterator struct {
	Event *SystemRewardDeleteOperator // Event containing the contract specifics and raw log

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
func (it *SystemRewardDeleteOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardDeleteOperator)
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
		it.Event = new(SystemRewardDeleteOperator)
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
func (it *SystemRewardDeleteOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardDeleteOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardDeleteOperator represents a DeleteOperator event raised by the SystemReward contract.
type SystemRewardDeleteOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteOperator is a free log retrieval operation binding the contract event 0x97058961cf85cc43abe97e01e175f56263c18351b7c0040d9534805f3e2b75e3.
//
// Solidity: event DeleteOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) FilterDeleteOperator(opts *bind.FilterOpts, operator []common.Address) (*SystemRewardDeleteOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "DeleteOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemRewardDeleteOperatorIterator{contract: _SystemReward.contract, event: "DeleteOperator", logs: logs, sub: sub}, nil
}

// WatchDeleteOperator is a free log subscription operation binding the contract event 0x97058961cf85cc43abe97e01e175f56263c18351b7c0040d9534805f3e2b75e3.
//
// Solidity: event DeleteOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) WatchDeleteOperator(opts *bind.WatchOpts, sink chan<- *SystemRewardDeleteOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "DeleteOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardDeleteOperator)
				if err := _SystemReward.contract.UnpackLog(event, "DeleteOperator", log); err != nil {
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

// ParseDeleteOperator is a log parse operation binding the contract event 0x97058961cf85cc43abe97e01e175f56263c18351b7c0040d9534805f3e2b75e3.
//
// Solidity: event DeleteOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) ParseDeleteOperator(log types.Log) (*SystemRewardDeleteOperator, error) {
	event := new(SystemRewardDeleteOperator)
	if err := _SystemReward.contract.UnpackLog(event, "DeleteOperator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemRewardNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the SystemReward contract.
type SystemRewardNewOperatorIterator struct {
	Event *SystemRewardNewOperator // Event containing the contract specifics and raw log

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
func (it *SystemRewardNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardNewOperator)
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
		it.Event = new(SystemRewardNewOperator)
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
func (it *SystemRewardNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardNewOperator represents a NewOperator event raised by the SystemReward contract.
type SystemRewardNewOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) FilterNewOperator(opts *bind.FilterOpts, operator []common.Address) (*SystemRewardNewOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "NewOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemRewardNewOperatorIterator{contract: _SystemReward.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *SystemRewardNewOperator, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "NewOperator", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardNewOperator)
				if err := _SystemReward.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address indexed operator)
func (_SystemReward *SystemRewardFilterer) ParseNewOperator(log types.Log) (*SystemRewardNewOperator, error) {
	event := new(SystemRewardNewOperator)
	if err := _SystemReward.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemRewardReceiveDepositIterator is returned from FilterReceiveDeposit and is used to iterate over the raw logs and unpacked data for ReceiveDeposit events raised by the SystemReward contract.
type SystemRewardReceiveDepositIterator struct {
	Event *SystemRewardReceiveDeposit // Event containing the contract specifics and raw log

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
func (it *SystemRewardReceiveDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardReceiveDeposit)
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
		it.Event = new(SystemRewardReceiveDeposit)
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
func (it *SystemRewardReceiveDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardReceiveDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardReceiveDeposit represents a ReceiveDeposit event raised by the SystemReward contract.
type SystemRewardReceiveDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveDeposit is a free log retrieval operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address indexed from, uint256 amount)
func (_SystemReward *SystemRewardFilterer) FilterReceiveDeposit(opts *bind.FilterOpts, from []common.Address) (*SystemRewardReceiveDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "ReceiveDeposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &SystemRewardReceiveDepositIterator{contract: _SystemReward.contract, event: "ReceiveDeposit", logs: logs, sub: sub}, nil
}

// WatchReceiveDeposit is a free log subscription operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address indexed from, uint256 amount)
func (_SystemReward *SystemRewardFilterer) WatchReceiveDeposit(opts *bind.WatchOpts, sink chan<- *SystemRewardReceiveDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "ReceiveDeposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardReceiveDeposit)
				if err := _SystemReward.contract.UnpackLog(event, "ReceiveDeposit", log); err != nil {
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

// ParseReceiveDeposit is a log parse operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address indexed from, uint256 amount)
func (_SystemReward *SystemRewardFilterer) ParseReceiveDeposit(log types.Log) (*SystemRewardReceiveDeposit, error) {
	event := new(SystemRewardReceiveDeposit)
	if err := _SystemReward.contract.UnpackLog(event, "ReceiveDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemRewardRewardEmptyIterator is returned from FilterRewardEmpty and is used to iterate over the raw logs and unpacked data for RewardEmpty events raised by the SystemReward contract.
type SystemRewardRewardEmptyIterator struct {
	Event *SystemRewardRewardEmpty // Event containing the contract specifics and raw log

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
func (it *SystemRewardRewardEmptyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardRewardEmpty)
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
		it.Event = new(SystemRewardRewardEmpty)
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
func (it *SystemRewardRewardEmptyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardRewardEmptyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardRewardEmpty represents a RewardEmpty event raised by the SystemReward contract.
type SystemRewardRewardEmpty struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRewardEmpty is a free log retrieval operation binding the contract event 0x9bd2da9e978ed6606b09c156a1bd98da2b52fe189aa7632b66e358095cef70eb.
//
// Solidity: event RewardEmpty()
func (_SystemReward *SystemRewardFilterer) FilterRewardEmpty(opts *bind.FilterOpts) (*SystemRewardRewardEmptyIterator, error) {

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "RewardEmpty")
	if err != nil {
		return nil, err
	}
	return &SystemRewardRewardEmptyIterator{contract: _SystemReward.contract, event: "RewardEmpty", logs: logs, sub: sub}, nil
}

// WatchRewardEmpty is a free log subscription operation binding the contract event 0x9bd2da9e978ed6606b09c156a1bd98da2b52fe189aa7632b66e358095cef70eb.
//
// Solidity: event RewardEmpty()
func (_SystemReward *SystemRewardFilterer) WatchRewardEmpty(opts *bind.WatchOpts, sink chan<- *SystemRewardRewardEmpty) (event.Subscription, error) {

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "RewardEmpty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardRewardEmpty)
				if err := _SystemReward.contract.UnpackLog(event, "RewardEmpty", log); err != nil {
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

// ParseRewardEmpty is a log parse operation binding the contract event 0x9bd2da9e978ed6606b09c156a1bd98da2b52fe189aa7632b66e358095cef70eb.
//
// Solidity: event RewardEmpty()
func (_SystemReward *SystemRewardFilterer) ParseRewardEmpty(log types.Log) (*SystemRewardRewardEmpty, error) {
	event := new(SystemRewardRewardEmpty)
	if err := _SystemReward.contract.UnpackLog(event, "RewardEmpty", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemRewardRewardToIterator is returned from FilterRewardTo and is used to iterate over the raw logs and unpacked data for RewardTo events raised by the SystemReward contract.
type SystemRewardRewardToIterator struct {
	Event *SystemRewardRewardTo // Event containing the contract specifics and raw log

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
func (it *SystemRewardRewardToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRewardRewardTo)
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
		it.Event = new(SystemRewardRewardTo)
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
func (it *SystemRewardRewardToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRewardRewardToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRewardRewardTo represents a RewardTo event raised by the SystemReward contract.
type SystemRewardRewardTo struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTo is a free log retrieval operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address indexed to, uint256 amount)
func (_SystemReward *SystemRewardFilterer) FilterRewardTo(opts *bind.FilterOpts, to []common.Address) (*SystemRewardRewardToIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SystemReward.contract.FilterLogs(opts, "RewardTo", toRule)
	if err != nil {
		return nil, err
	}
	return &SystemRewardRewardToIterator{contract: _SystemReward.contract, event: "RewardTo", logs: logs, sub: sub}, nil
}

// WatchRewardTo is a free log subscription operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address indexed to, uint256 amount)
func (_SystemReward *SystemRewardFilterer) WatchRewardTo(opts *bind.WatchOpts, sink chan<- *SystemRewardRewardTo, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SystemReward.contract.WatchLogs(opts, "RewardTo", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRewardRewardTo)
				if err := _SystemReward.contract.UnpackLog(event, "RewardTo", log); err != nil {
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

// ParseRewardTo is a log parse operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address indexed to, uint256 amount)
func (_SystemReward *SystemRewardFilterer) ParseRewardTo(log types.Log) (*SystemRewardRewardTo, error) {
	event := new(SystemRewardRewardTo)
	if err := _SystemReward.contract.UnpackLog(event, "RewardTo", log); err != nil {
		return nil, err
	}
	return event, nil
}
