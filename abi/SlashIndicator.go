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

// SlashABI is the input ABI used to generate the binding from.
const SlashABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"indicatorCleaned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FELONY_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MISDEMEANOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getSlashIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Slash is an auto generated Go binding around an Ethereum contract.
type Slash struct {
	SlashCaller     // Read-only binding to the contract
	SlashTransactor // Write-only binding to the contract
	SlashFilterer   // Log filterer for contract events
}

// SlashCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashSession struct {
	Contract     *Slash            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashCallerSession struct {
	Contract *SlashCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SlashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashTransactorSession struct {
	Contract     *SlashTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashRaw struct {
	Contract *Slash // Generic contract binding to access the raw methods on
}

// SlashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashCallerRaw struct {
	Contract *SlashCaller // Generic read-only contract binding to access the raw methods on
}

// SlashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashTransactorRaw struct {
	Contract *SlashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlash creates a new instance of Slash, bound to a specific deployed contract.
func NewSlash(address common.Address, backend bind.ContractBackend) (*Slash, error) {
	contract, err := bindSlash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slash{SlashCaller: SlashCaller{contract: contract}, SlashTransactor: SlashTransactor{contract: contract}, SlashFilterer: SlashFilterer{contract: contract}}, nil
}

// NewSlashCaller creates a new read-only instance of Slash, bound to a specific deployed contract.
func NewSlashCaller(address common.Address, caller bind.ContractCaller) (*SlashCaller, error) {
	contract, err := bindSlash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashCaller{contract: contract}, nil
}

// NewSlashTransactor creates a new write-only instance of Slash, bound to a specific deployed contract.
func NewSlashTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashTransactor, error) {
	contract, err := bindSlash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashTransactor{contract: contract}, nil
}

// NewSlashFilterer creates a new log filterer instance of Slash, bound to a specific deployed contract.
func NewSlashFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashFilterer, error) {
	contract, err := bindSlash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashFilterer{contract: contract}, nil
}

// bindSlash binds a generic wrapper to an already deployed contract.
func bindSlash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slash *SlashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Slash.Contract.SlashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slash *SlashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slash.Contract.SlashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slash *SlashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slash.Contract.SlashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slash *SlashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Slash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slash *SlashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slash *SlashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slash.Contract.contract.Transact(opts, method, params...)
}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() constant returns(uint256)
func (_Slash *SlashCaller) FELONYTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "FELONY_THRESHOLD")
	return *ret0, err
}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() constant returns(uint256)
func (_Slash *SlashSession) FELONYTHRESHOLD() (*big.Int, error) {
	return _Slash.Contract.FELONYTHRESHOLD(&_Slash.CallOpts)
}

// FELONYTHRESHOLD is a free data retrieval call binding the contract method 0xc80d4b8f.
//
// Solidity: function FELONY_THRESHOLD() constant returns(uint256)
func (_Slash *SlashCallerSession) FELONYTHRESHOLD() (*big.Int, error) {
	return _Slash.Contract.FELONYTHRESHOLD(&_Slash.CallOpts)
}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() constant returns(uint256)
func (_Slash *SlashCaller) MISDEMEANORTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "MISDEMEANOR_THRESHOLD")
	return *ret0, err
}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() constant returns(uint256)
func (_Slash *SlashSession) MISDEMEANORTHRESHOLD() (*big.Int, error) {
	return _Slash.Contract.MISDEMEANORTHRESHOLD(&_Slash.CallOpts)
}

// MISDEMEANORTHRESHOLD is a free data retrieval call binding the contract method 0x7912a65d.
//
// Solidity: function MISDEMEANOR_THRESHOLD() constant returns(uint256)
func (_Slash *SlashCallerSession) MISDEMEANORTHRESHOLD() (*big.Int, error) {
	return _Slash.Contract.MISDEMEANORTHRESHOLD(&_Slash.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.VALIDATORCONTRACTADDR(&_Slash.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.VALIDATORCONTRACTADDR(&_Slash.CallOpts)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) constant returns(uint256, uint256)
func (_Slash *SlashCaller) GetSlashIndicator(opts *bind.CallOpts, validator common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Slash.contract.Call(opts, out, "getSlashIndicator", validator)
	return *ret0, *ret1, err
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) constant returns(uint256, uint256)
func (_Slash *SlashSession) GetSlashIndicator(validator common.Address) (*big.Int, *big.Int, error) {
	return _Slash.Contract.GetSlashIndicator(&_Slash.CallOpts, validator)
}

// GetSlashIndicator is a free data retrieval call binding the contract method 0x37c8dab9.
//
// Solidity: function getSlashIndicator(address validator) constant returns(uint256, uint256)
func (_Slash *SlashCallerSession) GetSlashIndicator(validator common.Address) (*big.Int, *big.Int, error) {
	return _Slash.Contract.GetSlashIndicator(&_Slash.CallOpts, validator)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() constant returns(uint256)
func (_Slash *SlashCaller) PreviousHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "previousHeight")
	return *ret0, err
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() constant returns(uint256)
func (_Slash *SlashSession) PreviousHeight() (*big.Int, error) {
	return _Slash.Contract.PreviousHeight(&_Slash.CallOpts)
}

// PreviousHeight is a free data retrieval call binding the contract method 0x62b72cf5.
//
// Solidity: function previousHeight() constant returns(uint256)
func (_Slash *SlashCallerSession) PreviousHeight() (*big.Int, error) {
	return _Slash.Contract.PreviousHeight(&_Slash.CallOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Slash *SlashTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Slash *SlashSession) Clean() (*types.Transaction, error) {
	return _Slash.Contract.Clean(&_Slash.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_Slash *SlashTransactorSession) Clean() (*types.Transaction, error) {
	return _Slash.Contract.Clean(&_Slash.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_Slash *SlashTransactor) Slash(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "slash", validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_Slash *SlashSession) Slash(validator common.Address) (*types.Transaction, error) {
	return _Slash.Contract.Slash(&_Slash.TransactOpts, validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address validator) returns()
func (_Slash *SlashTransactorSession) Slash(validator common.Address) (*types.Transaction, error) {
	return _Slash.Contract.Slash(&_Slash.TransactOpts, validator)
}

// SlashIndicatorCleanedIterator is returned from FilterIndicatorCleaned and is used to iterate over the raw logs and unpacked data for IndicatorCleaned events raised by the Slash contract.
type SlashIndicatorCleanedIterator struct {
	Event *SlashIndicatorCleaned // Event containing the contract specifics and raw log

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
func (it *SlashIndicatorCleanedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashIndicatorCleaned)
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
		it.Event = new(SlashIndicatorCleaned)
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
func (it *SlashIndicatorCleanedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashIndicatorCleanedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashIndicatorCleaned represents a IndicatorCleaned event raised by the Slash contract.
type SlashIndicatorCleaned struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIndicatorCleaned is a free log retrieval operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_Slash *SlashFilterer) FilterIndicatorCleaned(opts *bind.FilterOpts) (*SlashIndicatorCleanedIterator, error) {

	logs, sub, err := _Slash.contract.FilterLogs(opts, "indicatorCleaned")
	if err != nil {
		return nil, err
	}
	return &SlashIndicatorCleanedIterator{contract: _Slash.contract, event: "indicatorCleaned", logs: logs, sub: sub}, nil
}

// WatchIndicatorCleaned is a free log subscription operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_Slash *SlashFilterer) WatchIndicatorCleaned(opts *bind.WatchOpts, sink chan<- *SlashIndicatorCleaned) (event.Subscription, error) {

	logs, sub, err := _Slash.contract.WatchLogs(opts, "indicatorCleaned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashIndicatorCleaned)
				if err := _Slash.contract.UnpackLog(event, "indicatorCleaned", log); err != nil {
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

// ParseIndicatorCleaned is a log parse operation binding the contract event 0xcfdb3b6ccaeccbdc68be3c59c840e3b3c90f0a7c491f5fff1cf56cfda200dd9c.
//
// Solidity: event indicatorCleaned()
func (_Slash *SlashFilterer) ParseIndicatorCleaned(log types.Log) (*SlashIndicatorCleaned, error) {
	event := new(SlashIndicatorCleaned)
	if err := _Slash.contract.UnpackLog(event, "indicatorCleaned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SlashValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Slash contract.
type SlashValidatorSlashedIterator struct {
	Event *SlashValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *SlashValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashValidatorSlashed)
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
		it.Event = new(SlashValidatorSlashed)
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
func (it *SlashValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashValidatorSlashed represents a ValidatorSlashed event raised by the Slash contract.
type SlashValidatorSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_Slash *SlashFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*SlashValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Slash.contract.FilterLogs(opts, "validatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SlashValidatorSlashedIterator{contract: _Slash.contract, event: "validatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_Slash *SlashFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *SlashValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Slash.contract.WatchLogs(opts, "validatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashValidatorSlashed)
				if err := _Slash.contract.UnpackLog(event, "validatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0xddb6012116e51abf5436d956a4f0ebd927e92c576ff96d7918290c8782291e3e.
//
// Solidity: event validatorSlashed(address indexed validator)
func (_Slash *SlashFilterer) ParseValidatorSlashed(log types.Log) (*SlashValidatorSlashed, error) {
	event := new(SlashValidatorSlashed)
	if err := _Slash.contract.UnpackLog(event, "validatorSlashed", log); err != nil {
		return nil, err
	}
	return event, nil
}
