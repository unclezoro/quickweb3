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

// RelayerABI is the input ABI used to generate the binding from.
const RelayerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerUnRegister\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_DUES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REQUIRED_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Relayer is an auto generated Go binding around an Ethereum contract.
type Relayer struct {
	RelayerCaller     // Read-only binding to the contract
	RelayerTransactor // Write-only binding to the contract
	RelayerFilterer   // Log filterer for contract events
}

// RelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerSession struct {
	Contract     *Relayer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerCallerSession struct {
	Contract *RelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerTransactorSession struct {
	Contract     *RelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerRaw struct {
	Contract *Relayer // Generic contract binding to access the raw methods on
}

// RelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerCallerRaw struct {
	Contract *RelayerCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerTransactorRaw struct {
	Contract *RelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayer creates a new instance of Relayer, bound to a specific deployed contract.
func NewRelayer(address common.Address, backend bind.ContractBackend) (*Relayer, error) {
	contract, err := bindRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayer{RelayerCaller: RelayerCaller{contract: contract}, RelayerTransactor: RelayerTransactor{contract: contract}, RelayerFilterer: RelayerFilterer{contract: contract}}, nil
}

// NewRelayerCaller creates a new read-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerCaller(address common.Address, caller bind.ContractCaller) (*RelayerCaller, error) {
	contract, err := bindRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerCaller{contract: contract}, nil
}

// NewRelayerTransactor creates a new write-only instance of Relayer, bound to a specific deployed contract.
func NewRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerTransactor, error) {
	contract, err := bindRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerTransactor{contract: contract}, nil
}

// NewRelayerFilterer creates a new log filterer instance of Relayer, bound to a specific deployed contract.
func NewRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerFilterer, error) {
	contract, err := bindRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerFilterer{contract: contract}, nil
}

// bindRelayer binds a generic wrapper to an already deployed contract.
func bindRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.RelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.RelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayer *RelayerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayer *RelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayer *RelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayer.Contract.contract.Transact(opts, method, params...)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayer *RelayerCaller) INITDUES(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "INIT_DUES")
	return *ret0, err
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayer *RelayerSession) INITDUES() (*big.Int, error) {
	return _Relayer.Contract.INITDUES(&_Relayer.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayer *RelayerCallerSession) INITDUES() (*big.Int, error) {
	return _Relayer.Contract.INITDUES(&_Relayer.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayer *RelayerCaller) INITREQUIREDDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "INIT_REQUIRED_DEPOSIT")
	return *ret0, err
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayer *RelayerSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _Relayer.Contract.INITREQUIREDDEPOSIT(&_Relayer.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayer *RelayerCallerSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _Relayer.Contract.INITREQUIREDDEPOSIT(&_Relayer.CallOpts)
}

// INITSYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xbaa3f053.
//
// Solidity: function INIT_SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayer *RelayerCaller) INITSYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "INIT_SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// INITSYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xbaa3f053.
//
// Solidity: function INIT_SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayer *RelayerSession) INITSYSTEMREWARDADDR() (common.Address, error) {
	return _Relayer.Contract.INITSYSTEMREWARDADDR(&_Relayer.CallOpts)
}

// INITSYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xbaa3f053.
//
// Solidity: function INIT_SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayer *RelayerCallerSession) INITSYSTEMREWARDADDR() (common.Address, error) {
	return _Relayer.Contract.INITSYSTEMREWARDADDR(&_Relayer.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayer *RelayerCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayer *RelayerSession) AlreadyInit() (bool, error) {
	return _Relayer.Contract.AlreadyInit(&_Relayer.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayer *RelayerCallerSession) AlreadyInit() (bool, error) {
	return _Relayer.Contract.AlreadyInit(&_Relayer.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayer *RelayerCaller) Dues(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "dues")
	return *ret0, err
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayer *RelayerSession) Dues() (*big.Int, error) {
	return _Relayer.Contract.Dues(&_Relayer.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayer *RelayerCallerSession) Dues() (*big.Int, error) {
	return _Relayer.Contract.Dues(&_Relayer.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayer *RelayerCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "isRelayer", sender)
	return *ret0, err
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayer *RelayerSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayer.Contract.IsRelayer(&_Relayer.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayer *RelayerCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayer.Contract.IsRelayer(&_Relayer.CallOpts, sender)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayer *RelayerCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayer.contract.Call(opts, out, "requiredDeposit")
	return *ret0, err
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayer *RelayerSession) RequiredDeposit() (*big.Int, error) {
	return _Relayer.Contract.RequiredDeposit(&_Relayer.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayer *RelayerCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Relayer.Contract.RequiredDeposit(&_Relayer.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayer *RelayerTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayer *RelayerSession) Init() (*types.Transaction, error) {
	return _Relayer.Contract.Init(&_Relayer.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayer *RelayerTransactorSession) Init() (*types.Transaction, error) {
	return _Relayer.Contract.Init(&_Relayer.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayer *RelayerTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayer *RelayerSession) Register() (*types.Transaction, error) {
	return _Relayer.Contract.Register(&_Relayer.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayer *RelayerTransactorSession) Register() (*types.Transaction, error) {
	return _Relayer.Contract.Register(&_Relayer.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayer *RelayerTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayer.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayer *RelayerSession) Unregister() (*types.Transaction, error) {
	return _Relayer.Contract.Unregister(&_Relayer.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayer *RelayerTransactorSession) Unregister() (*types.Transaction, error) {
	return _Relayer.Contract.Unregister(&_Relayer.TransactOpts)
}

// RelayerRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the Relayer contract.
type RelayerRelayerRegisterIterator struct {
	Event *RelayerRelayerRegister // Event containing the contract specifics and raw log

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
func (it *RelayerRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerRegister)
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
		it.Event = new(RelayerRelayerRegister)
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
func (it *RelayerRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerRegister represents a RelayerRegister event raised by the Relayer contract.
type RelayerRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayer *RelayerFilterer) FilterRelayerRegister(opts *bind.FilterOpts) (*RelayerRelayerRegisterIterator, error) {

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerRegisterIterator{contract: _Relayer.contract, event: "relayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayer *RelayerFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *RelayerRelayerRegister) (event.Subscription, error) {

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerRegister)
				if err := _Relayer.contract.UnpackLog(event, "relayerRegister", log); err != nil {
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

// ParseRelayerRegister is a log parse operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayer *RelayerFilterer) ParseRelayerRegister(log types.Log) (*RelayerRelayerRegister, error) {
	event := new(RelayerRelayerRegister)
	if err := _Relayer.contract.UnpackLog(event, "relayerRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the Relayer contract.
type RelayerRelayerUnRegisterIterator struct {
	Event *RelayerRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *RelayerRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerRelayerUnRegister)
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
		it.Event = new(RelayerRelayerUnRegister)
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
func (it *RelayerRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerRelayerUnRegister represents a RelayerUnRegister event raised by the Relayer contract.
type RelayerRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayer *RelayerFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts) (*RelayerRelayerUnRegisterIterator, error) {

	logs, sub, err := _Relayer.contract.FilterLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerRelayerUnRegisterIterator{contract: _Relayer.contract, event: "relayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayer *RelayerFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *RelayerRelayerUnRegister) (event.Subscription, error) {

	logs, sub, err := _Relayer.contract.WatchLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerRelayerUnRegister)
				if err := _Relayer.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
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

// ParseRelayerUnRegister is a log parse operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayer *RelayerFilterer) ParseRelayerUnRegister(log types.Log) (*RelayerRelayerUnRegister, error) {
	event := new(RelayerRelayerUnRegister)
	if err := _Relayer.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}
