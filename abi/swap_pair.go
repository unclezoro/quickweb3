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

// SwapPairABI is the input ABI used to generate the binding from.
const SwapPairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"GovernorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldImpl\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"ImplChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"InterestRatePerBlockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProductivityDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProductivityIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"changeInterestRatePerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseProductivity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getProductivity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"impl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incNounce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseProductivity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestsPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCumulation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nounce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"take\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"takeWithAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takeWithBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"upgradeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"upgradeImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardEarn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwapPair is an auto generated Go binding around an Ethereum contract.
type SwapPair struct {
	SwapPairCaller     // Read-only binding to the contract
	SwapPairTransactor // Write-only binding to the contract
	SwapPairFilterer   // Log filterer for contract events
}

// SwapPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapPairSession struct {
	Contract     *SwapPair         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapPairCallerSession struct {
	Contract *SwapPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SwapPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapPairTransactorSession struct {
	Contract     *SwapPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SwapPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapPairRaw struct {
	Contract *SwapPair // Generic contract binding to access the raw methods on
}

// SwapPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapPairCallerRaw struct {
	Contract *SwapPairCaller // Generic read-only contract binding to access the raw methods on
}

// SwapPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapPairTransactorRaw struct {
	Contract *SwapPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapPair creates a new instance of SwapPair, bound to a specific deployed contract.
func NewSwapPair(address common.Address, backend bind.ContractBackend) (*SwapPair, error) {
	contract, err := bindSwapPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapPair{SwapPairCaller: SwapPairCaller{contract: contract}, SwapPairTransactor: SwapPairTransactor{contract: contract}, SwapPairFilterer: SwapPairFilterer{contract: contract}}, nil
}

// NewSwapPairCaller creates a new read-only instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairCaller(address common.Address, caller bind.ContractCaller) (*SwapPairCaller, error) {
	contract, err := bindSwapPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapPairCaller{contract: contract}, nil
}

// NewSwapPairTransactor creates a new write-only instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapPairTransactor, error) {
	contract, err := bindSwapPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapPairTransactor{contract: contract}, nil
}

// NewSwapPairFilterer creates a new log filterer instance of SwapPair, bound to a specific deployed contract.
func NewSwapPairFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapPairFilterer, error) {
	contract, err := bindSwapPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapPairFilterer{contract: contract}, nil
}

// bindSwapPair binds a generic wrapper to an already deployed contract.
func bindSwapPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapPair *SwapPairRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapPair.Contract.SwapPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapPair *SwapPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.Contract.SwapPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapPair *SwapPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapPair.Contract.SwapPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapPair *SwapPairCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapPair *SwapPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapPair *SwapPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapPair.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SwapPair *SwapPairCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SwapPair *SwapPairSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Allowance(&_SwapPair.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.Allowance(&_SwapPair.CallOpts, arg0, arg1)
}

// AmountPerBlock is a free data retrieval call binding the contract method 0x4c61f047.
//
// Solidity: function amountPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairCaller) AmountPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "amountPerBlock")
	return *ret0, err
}

// AmountPerBlock is a free data retrieval call binding the contract method 0x4c61f047.
//
// Solidity: function amountPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairSession) AmountPerBlock() (*big.Int, error) {
	return _SwapPair.Contract.AmountPerBlock(&_SwapPair.CallOpts)
}

// AmountPerBlock is a free data retrieval call binding the contract method 0x4c61f047.
//
// Solidity: function amountPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) AmountPerBlock() (*big.Int, error) {
	return _SwapPair.Contract.AmountPerBlock(&_SwapPair.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_SwapPair *SwapPairCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_SwapPair *SwapPairSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.BalanceOf(&_SwapPair.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SwapPair.Contract.BalanceOf(&_SwapPair.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SwapPair *SwapPairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SwapPair *SwapPairSession) Decimals() (uint8, error) {
	return _SwapPair.Contract.Decimals(&_SwapPair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SwapPair *SwapPairCallerSession) Decimals() (uint8, error) {
	return _SwapPair.Contract.Decimals(&_SwapPair.CallOpts)
}

// GetProductivity is a free data retrieval call binding the contract method 0x28e964e9.
//
// Solidity: function getProductivity(address user) constant returns(uint256, uint256)
func (_SwapPair *SwapPairCaller) GetProductivity(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SwapPair.contract.Call(opts, out, "getProductivity", user)
	return *ret0, *ret1, err
}

// GetProductivity is a free data retrieval call binding the contract method 0x28e964e9.
//
// Solidity: function getProductivity(address user) constant returns(uint256, uint256)
func (_SwapPair *SwapPairSession) GetProductivity(user common.Address) (*big.Int, *big.Int, error) {
	return _SwapPair.Contract.GetProductivity(&_SwapPair.CallOpts, user)
}

// GetProductivity is a free data retrieval call binding the contract method 0x28e964e9.
//
// Solidity: function getProductivity(address user) constant returns(uint256, uint256)
func (_SwapPair *SwapPairCallerSession) GetProductivity(user common.Address) (*big.Int, *big.Int, error) {
	return _SwapPair.Contract.GetProductivity(&_SwapPair.CallOpts, user)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_SwapPair *SwapPairCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "governor")
	return *ret0, err
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_SwapPair *SwapPairSession) Governor() (common.Address, error) {
	return _SwapPair.Contract.Governor(&_SwapPair.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() constant returns(address)
func (_SwapPair *SwapPairCallerSession) Governor() (common.Address, error) {
	return _SwapPair.Contract.Governor(&_SwapPair.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_SwapPair *SwapPairCaller) Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "impl")
	return *ret0, err
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_SwapPair *SwapPairSession) Impl() (common.Address, error) {
	return _SwapPair.Contract.Impl(&_SwapPair.CallOpts)
}

// Impl is a free data retrieval call binding the contract method 0x8abf6077.
//
// Solidity: function impl() constant returns(address)
func (_SwapPair *SwapPairCallerSession) Impl() (common.Address, error) {
	return _SwapPair.Contract.Impl(&_SwapPair.CallOpts)
}

// InterestsPerBlock is a free data retrieval call binding the contract method 0xbfc8b208.
//
// Solidity: function interestsPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairCaller) InterestsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "interestsPerBlock")
	return *ret0, err
}

// InterestsPerBlock is a free data retrieval call binding the contract method 0xbfc8b208.
//
// Solidity: function interestsPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairSession) InterestsPerBlock() (*big.Int, error) {
	return _SwapPair.Contract.InterestsPerBlock(&_SwapPair.CallOpts)
}

// InterestsPerBlock is a free data retrieval call binding the contract method 0xbfc8b208.
//
// Solidity: function interestsPerBlock() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) InterestsPerBlock() (*big.Int, error) {
	return _SwapPair.Contract.InterestsPerBlock(&_SwapPair.CallOpts)
}

// MintCumulation is a free data retrieval call binding the contract method 0xcf675365.
//
// Solidity: function mintCumulation() constant returns(uint256)
func (_SwapPair *SwapPairCaller) MintCumulation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "mintCumulation")
	return *ret0, err
}

// MintCumulation is a free data retrieval call binding the contract method 0xcf675365.
//
// Solidity: function mintCumulation() constant returns(uint256)
func (_SwapPair *SwapPairSession) MintCumulation() (*big.Int, error) {
	return _SwapPair.Contract.MintCumulation(&_SwapPair.CallOpts)
}

// MintCumulation is a free data retrieval call binding the contract method 0xcf675365.
//
// Solidity: function mintCumulation() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) MintCumulation() (*big.Int, error) {
	return _SwapPair.Contract.MintCumulation(&_SwapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SwapPair *SwapPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SwapPair *SwapPairSession) Name() (string, error) {
	return _SwapPair.Contract.Name(&_SwapPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SwapPair *SwapPairCallerSession) Name() (string, error) {
	return _SwapPair.Contract.Name(&_SwapPair.CallOpts)
}

// Nounce is a free data retrieval call binding the contract method 0x616d2463.
//
// Solidity: function nounce() constant returns(uint256)
func (_SwapPair *SwapPairCaller) Nounce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "nounce")
	return *ret0, err
}

// Nounce is a free data retrieval call binding the contract method 0x616d2463.
//
// Solidity: function nounce() constant returns(uint256)
func (_SwapPair *SwapPairSession) Nounce() (*big.Int, error) {
	return _SwapPair.Contract.Nounce(&_SwapPair.CallOpts)
}

// Nounce is a free data retrieval call binding the contract method 0x616d2463.
//
// Solidity: function nounce() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) Nounce() (*big.Int, error) {
	return _SwapPair.Contract.Nounce(&_SwapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SwapPair *SwapPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SwapPair *SwapPairSession) Symbol() (string, error) {
	return _SwapPair.Contract.Symbol(&_SwapPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SwapPair *SwapPairCallerSession) Symbol() (string, error) {
	return _SwapPair.Contract.Symbol(&_SwapPair.CallOpts)
}

// Take is a free data retrieval call binding the contract method 0x159090bd.
//
// Solidity: function take() constant returns(uint256)
func (_SwapPair *SwapPairCaller) Take(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "take")
	return *ret0, err
}

// Take is a free data retrieval call binding the contract method 0x159090bd.
//
// Solidity: function take() constant returns(uint256)
func (_SwapPair *SwapPairSession) Take() (*big.Int, error) {
	return _SwapPair.Contract.Take(&_SwapPair.CallOpts)
}

// Take is a free data retrieval call binding the contract method 0x159090bd.
//
// Solidity: function take() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) Take() (*big.Int, error) {
	return _SwapPair.Contract.Take(&_SwapPair.CallOpts)
}

// TakeWithAddress is a free data retrieval call binding the contract method 0x6622c838.
//
// Solidity: function takeWithAddress(address user) constant returns(uint256)
func (_SwapPair *SwapPairCaller) TakeWithAddress(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "takeWithAddress", user)
	return *ret0, err
}

// TakeWithAddress is a free data retrieval call binding the contract method 0x6622c838.
//
// Solidity: function takeWithAddress(address user) constant returns(uint256)
func (_SwapPair *SwapPairSession) TakeWithAddress(user common.Address) (*big.Int, error) {
	return _SwapPair.Contract.TakeWithAddress(&_SwapPair.CallOpts, user)
}

// TakeWithAddress is a free data retrieval call binding the contract method 0x6622c838.
//
// Solidity: function takeWithAddress(address user) constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) TakeWithAddress(user common.Address) (*big.Int, error) {
	return _SwapPair.Contract.TakeWithAddress(&_SwapPair.CallOpts, user)
}

// TakeWithBlock is a free data retrieval call binding the contract method 0x0e0b6eb5.
//
// Solidity: function takeWithBlock() constant returns(uint256, uint256)
func (_SwapPair *SwapPairCaller) TakeWithBlock(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SwapPair.contract.Call(opts, out, "takeWithBlock")
	return *ret0, *ret1, err
}

// TakeWithBlock is a free data retrieval call binding the contract method 0x0e0b6eb5.
//
// Solidity: function takeWithBlock() constant returns(uint256, uint256)
func (_SwapPair *SwapPairSession) TakeWithBlock() (*big.Int, *big.Int, error) {
	return _SwapPair.Contract.TakeWithBlock(&_SwapPair.CallOpts)
}

// TakeWithBlock is a free data retrieval call binding the contract method 0x0e0b6eb5.
//
// Solidity: function takeWithBlock() constant returns(uint256, uint256)
func (_SwapPair *SwapPairCallerSession) TakeWithBlock() (*big.Int, *big.Int, error) {
	return _SwapPair.Contract.TakeWithBlock(&_SwapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SwapPair *SwapPairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapPair.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SwapPair *SwapPairSession) TotalSupply() (*big.Int, error) {
	return _SwapPair.Contract.TotalSupply(&_SwapPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SwapPair *SwapPairCallerSession) TotalSupply() (*big.Int, error) {
	return _SwapPair.Contract.TotalSupply(&_SwapPair.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(uint256 amount, uint256 rewardDebt, uint256 rewardEarn)
func (_SwapPair *SwapPairCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	RewardEarn *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
		RewardEarn *big.Int
	})
	out := ret
	err := _SwapPair.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(uint256 amount, uint256 rewardDebt, uint256 rewardEarn)
func (_SwapPair *SwapPairSession) Users(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	RewardEarn *big.Int
}, error) {
	return _SwapPair.Contract.Users(&_SwapPair.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(uint256 amount, uint256 rewardDebt, uint256 rewardEarn)
func (_SwapPair *SwapPairCallerSession) Users(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
	RewardEarn *big.Int
}, error) {
	return _SwapPair.Contract.Users(&_SwapPair.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Approve(&_SwapPair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Approve(&_SwapPair.TransactOpts, spender, value)
}

// ChangeInterestRatePerBlock is a paid mutator transaction binding the contract method 0x7b381b35.
//
// Solidity: function changeInterestRatePerBlock(uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) ChangeInterestRatePerBlock(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "changeInterestRatePerBlock", value)
}

// ChangeInterestRatePerBlock is a paid mutator transaction binding the contract method 0x7b381b35.
//
// Solidity: function changeInterestRatePerBlock(uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) ChangeInterestRatePerBlock(value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.ChangeInterestRatePerBlock(&_SwapPair.TransactOpts, value)
}

// ChangeInterestRatePerBlock is a paid mutator transaction binding the contract method 0x7b381b35.
//
// Solidity: function changeInterestRatePerBlock(uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) ChangeInterestRatePerBlock(value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.ChangeInterestRatePerBlock(&_SwapPair.TransactOpts, value)
}

// DecreaseProductivity is a paid mutator transaction binding the contract method 0x4afa66d6.
//
// Solidity: function decreaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) DecreaseProductivity(opts *bind.TransactOpts, user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "decreaseProductivity", user, value)
}

// DecreaseProductivity is a paid mutator transaction binding the contract method 0x4afa66d6.
//
// Solidity: function decreaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) DecreaseProductivity(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.DecreaseProductivity(&_SwapPair.TransactOpts, user, value)
}

// DecreaseProductivity is a paid mutator transaction binding the contract method 0x4afa66d6.
//
// Solidity: function decreaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) DecreaseProductivity(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.DecreaseProductivity(&_SwapPair.TransactOpts, user, value)
}

// IncNounce is a paid mutator transaction binding the contract method 0x5808b75b.
//
// Solidity: function incNounce() returns()
func (_SwapPair *SwapPairTransactor) IncNounce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "incNounce")
}

// IncNounce is a paid mutator transaction binding the contract method 0x5808b75b.
//
// Solidity: function incNounce() returns()
func (_SwapPair *SwapPairSession) IncNounce() (*types.Transaction, error) {
	return _SwapPair.Contract.IncNounce(&_SwapPair.TransactOpts)
}

// IncNounce is a paid mutator transaction binding the contract method 0x5808b75b.
//
// Solidity: function incNounce() returns()
func (_SwapPair *SwapPairTransactorSession) IncNounce() (*types.Transaction, error) {
	return _SwapPair.Contract.IncNounce(&_SwapPair.TransactOpts)
}

// IncreaseProductivity is a paid mutator transaction binding the contract method 0x36f04e45.
//
// Solidity: function increaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) IncreaseProductivity(opts *bind.TransactOpts, user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "increaseProductivity", user, value)
}

// IncreaseProductivity is a paid mutator transaction binding the contract method 0x36f04e45.
//
// Solidity: function increaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) IncreaseProductivity(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.IncreaseProductivity(&_SwapPair.TransactOpts, user, value)
}

// IncreaseProductivity is a paid mutator transaction binding the contract method 0x36f04e45.
//
// Solidity: function increaseProductivity(address user, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) IncreaseProductivity(user common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.IncreaseProductivity(&_SwapPair.TransactOpts, user, value)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_SwapPair *SwapPairTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_SwapPair *SwapPairSession) Mint() (*types.Transaction, error) {
	return _SwapPair.Contract.Mint(&_SwapPair.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_SwapPair *SwapPairTransactorSession) Mint() (*types.Transaction, error) {
	return _SwapPair.Contract.Mint(&_SwapPair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Transfer(&_SwapPair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.Transfer(&_SwapPair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.TransferFrom(&_SwapPair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SwapPair *SwapPairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SwapPair.Contract.TransferFrom(&_SwapPair.TransactOpts, from, to, value)
}

// UpgradeGovernance is a paid mutator transaction binding the contract method 0x1fedded5.
//
// Solidity: function upgradeGovernance(address _newGovernor) returns()
func (_SwapPair *SwapPairTransactor) UpgradeGovernance(opts *bind.TransactOpts, _newGovernor common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "upgradeGovernance", _newGovernor)
}

// UpgradeGovernance is a paid mutator transaction binding the contract method 0x1fedded5.
//
// Solidity: function upgradeGovernance(address _newGovernor) returns()
func (_SwapPair *SwapPairSession) UpgradeGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.UpgradeGovernance(&_SwapPair.TransactOpts, _newGovernor)
}

// UpgradeGovernance is a paid mutator transaction binding the contract method 0x1fedded5.
//
// Solidity: function upgradeGovernance(address _newGovernor) returns()
func (_SwapPair *SwapPairTransactorSession) UpgradeGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.UpgradeGovernance(&_SwapPair.TransactOpts, _newGovernor)
}

// UpgradeImpl is a paid mutator transaction binding the contract method 0x1a2f1363.
//
// Solidity: function upgradeImpl(address _newImpl) returns()
func (_SwapPair *SwapPairTransactor) UpgradeImpl(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _SwapPair.contract.Transact(opts, "upgradeImpl", _newImpl)
}

// UpgradeImpl is a paid mutator transaction binding the contract method 0x1a2f1363.
//
// Solidity: function upgradeImpl(address _newImpl) returns()
func (_SwapPair *SwapPairSession) UpgradeImpl(_newImpl common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.UpgradeImpl(&_SwapPair.TransactOpts, _newImpl)
}

// UpgradeImpl is a paid mutator transaction binding the contract method 0x1a2f1363.
//
// Solidity: function upgradeImpl(address _newImpl) returns()
func (_SwapPair *SwapPairTransactorSession) UpgradeImpl(_newImpl common.Address) (*types.Transaction, error) {
	return _SwapPair.Contract.UpgradeImpl(&_SwapPair.TransactOpts, _newImpl)
}

// SwapPairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SwapPair contract.
type SwapPairApprovalIterator struct {
	Event *SwapPairApproval // Event containing the contract specifics and raw log

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
func (it *SwapPairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairApproval)
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
		it.Event = new(SwapPairApproval)
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
func (it *SwapPairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairApproval represents a Approval event raised by the SwapPair contract.
type SwapPairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SwapPairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairApprovalIterator{contract: _SwapPair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SwapPairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairApproval)
				if err := _SwapPair.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SwapPair *SwapPairFilterer) ParseApproval(log types.Log) (*SwapPairApproval, error) {
	event := new(SwapPairApproval)
	if err := _SwapPair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairGovernorChangedIterator is returned from FilterGovernorChanged and is used to iterate over the raw logs and unpacked data for GovernorChanged events raised by the SwapPair contract.
type SwapPairGovernorChangedIterator struct {
	Event *SwapPairGovernorChanged // Event containing the contract specifics and raw log

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
func (it *SwapPairGovernorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairGovernorChanged)
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
		it.Event = new(SwapPairGovernorChanged)
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
func (it *SwapPairGovernorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairGovernorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairGovernorChanged represents a GovernorChanged event raised by the SwapPair contract.
type SwapPairGovernorChanged struct {
	OldGovernor common.Address
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovernorChanged is a free log retrieval operation binding the contract event 0xde4b3f61490b74c0ed6237523974fe299126bbbf8a8a7482fd220104c59b0c84.
//
// Solidity: event GovernorChanged(address indexed _oldGovernor, address indexed _newGovernor)
func (_SwapPair *SwapPairFilterer) FilterGovernorChanged(opts *bind.FilterOpts, _oldGovernor []common.Address, _newGovernor []common.Address) (*SwapPairGovernorChangedIterator, error) {

	var _oldGovernorRule []interface{}
	for _, _oldGovernorItem := range _oldGovernor {
		_oldGovernorRule = append(_oldGovernorRule, _oldGovernorItem)
	}
	var _newGovernorRule []interface{}
	for _, _newGovernorItem := range _newGovernor {
		_newGovernorRule = append(_newGovernorRule, _newGovernorItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "GovernorChanged", _oldGovernorRule, _newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairGovernorChangedIterator{contract: _SwapPair.contract, event: "GovernorChanged", logs: logs, sub: sub}, nil
}

// WatchGovernorChanged is a free log subscription operation binding the contract event 0xde4b3f61490b74c0ed6237523974fe299126bbbf8a8a7482fd220104c59b0c84.
//
// Solidity: event GovernorChanged(address indexed _oldGovernor, address indexed _newGovernor)
func (_SwapPair *SwapPairFilterer) WatchGovernorChanged(opts *bind.WatchOpts, sink chan<- *SwapPairGovernorChanged, _oldGovernor []common.Address, _newGovernor []common.Address) (event.Subscription, error) {

	var _oldGovernorRule []interface{}
	for _, _oldGovernorItem := range _oldGovernor {
		_oldGovernorRule = append(_oldGovernorRule, _oldGovernorItem)
	}
	var _newGovernorRule []interface{}
	for _, _newGovernorItem := range _newGovernor {
		_newGovernorRule = append(_newGovernorRule, _newGovernorItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "GovernorChanged", _oldGovernorRule, _newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairGovernorChanged)
				if err := _SwapPair.contract.UnpackLog(event, "GovernorChanged", log); err != nil {
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

// ParseGovernorChanged is a log parse operation binding the contract event 0xde4b3f61490b74c0ed6237523974fe299126bbbf8a8a7482fd220104c59b0c84.
//
// Solidity: event GovernorChanged(address indexed _oldGovernor, address indexed _newGovernor)
func (_SwapPair *SwapPairFilterer) ParseGovernorChanged(log types.Log) (*SwapPairGovernorChanged, error) {
	event := new(SwapPairGovernorChanged)
	if err := _SwapPair.contract.UnpackLog(event, "GovernorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairImplChangedIterator is returned from FilterImplChanged and is used to iterate over the raw logs and unpacked data for ImplChanged events raised by the SwapPair contract.
type SwapPairImplChangedIterator struct {
	Event *SwapPairImplChanged // Event containing the contract specifics and raw log

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
func (it *SwapPairImplChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairImplChanged)
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
		it.Event = new(SwapPairImplChanged)
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
func (it *SwapPairImplChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairImplChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairImplChanged represents a ImplChanged event raised by the SwapPair contract.
type SwapPairImplChanged struct {
	OldImpl common.Address
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplChanged is a free log retrieval operation binding the contract event 0xaad46b89531ed10d02d926f4d6bfe234a6126e3fffc02d3b07167575f9c14337.
//
// Solidity: event ImplChanged(address indexed _oldImpl, address indexed _newImpl)
func (_SwapPair *SwapPairFilterer) FilterImplChanged(opts *bind.FilterOpts, _oldImpl []common.Address, _newImpl []common.Address) (*SwapPairImplChangedIterator, error) {

	var _oldImplRule []interface{}
	for _, _oldImplItem := range _oldImpl {
		_oldImplRule = append(_oldImplRule, _oldImplItem)
	}
	var _newImplRule []interface{}
	for _, _newImplItem := range _newImpl {
		_newImplRule = append(_newImplRule, _newImplItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "ImplChanged", _oldImplRule, _newImplRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairImplChangedIterator{contract: _SwapPair.contract, event: "ImplChanged", logs: logs, sub: sub}, nil
}

// WatchImplChanged is a free log subscription operation binding the contract event 0xaad46b89531ed10d02d926f4d6bfe234a6126e3fffc02d3b07167575f9c14337.
//
// Solidity: event ImplChanged(address indexed _oldImpl, address indexed _newImpl)
func (_SwapPair *SwapPairFilterer) WatchImplChanged(opts *bind.WatchOpts, sink chan<- *SwapPairImplChanged, _oldImpl []common.Address, _newImpl []common.Address) (event.Subscription, error) {

	var _oldImplRule []interface{}
	for _, _oldImplItem := range _oldImpl {
		_oldImplRule = append(_oldImplRule, _oldImplItem)
	}
	var _newImplRule []interface{}
	for _, _newImplItem := range _newImpl {
		_newImplRule = append(_newImplRule, _newImplItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "ImplChanged", _oldImplRule, _newImplRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairImplChanged)
				if err := _SwapPair.contract.UnpackLog(event, "ImplChanged", log); err != nil {
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

// ParseImplChanged is a log parse operation binding the contract event 0xaad46b89531ed10d02d926f4d6bfe234a6126e3fffc02d3b07167575f9c14337.
//
// Solidity: event ImplChanged(address indexed _oldImpl, address indexed _newImpl)
func (_SwapPair *SwapPairFilterer) ParseImplChanged(log types.Log) (*SwapPairImplChanged, error) {
	event := new(SwapPairImplChanged)
	if err := _SwapPair.contract.UnpackLog(event, "ImplChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairInterestRatePerBlockChangedIterator is returned from FilterInterestRatePerBlockChanged and is used to iterate over the raw logs and unpacked data for InterestRatePerBlockChanged events raised by the SwapPair contract.
type SwapPairInterestRatePerBlockChangedIterator struct {
	Event *SwapPairInterestRatePerBlockChanged // Event containing the contract specifics and raw log

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
func (it *SwapPairInterestRatePerBlockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairInterestRatePerBlockChanged)
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
		it.Event = new(SwapPairInterestRatePerBlockChanged)
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
func (it *SwapPairInterestRatePerBlockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairInterestRatePerBlockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairInterestRatePerBlockChanged represents a InterestRatePerBlockChanged event raised by the SwapPair contract.
type SwapPairInterestRatePerBlockChanged struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInterestRatePerBlockChanged is a free log retrieval operation binding the contract event 0xc44c80d7c6ac73ee13b2c62e50522dadb6c646b4509ca93e52a45abb38a78e33.
//
// Solidity: event InterestRatePerBlockChanged(uint256 oldValue, uint256 newValue)
func (_SwapPair *SwapPairFilterer) FilterInterestRatePerBlockChanged(opts *bind.FilterOpts) (*SwapPairInterestRatePerBlockChangedIterator, error) {

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "InterestRatePerBlockChanged")
	if err != nil {
		return nil, err
	}
	return &SwapPairInterestRatePerBlockChangedIterator{contract: _SwapPair.contract, event: "InterestRatePerBlockChanged", logs: logs, sub: sub}, nil
}

// WatchInterestRatePerBlockChanged is a free log subscription operation binding the contract event 0xc44c80d7c6ac73ee13b2c62e50522dadb6c646b4509ca93e52a45abb38a78e33.
//
// Solidity: event InterestRatePerBlockChanged(uint256 oldValue, uint256 newValue)
func (_SwapPair *SwapPairFilterer) WatchInterestRatePerBlockChanged(opts *bind.WatchOpts, sink chan<- *SwapPairInterestRatePerBlockChanged) (event.Subscription, error) {

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "InterestRatePerBlockChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairInterestRatePerBlockChanged)
				if err := _SwapPair.contract.UnpackLog(event, "InterestRatePerBlockChanged", log); err != nil {
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

// ParseInterestRatePerBlockChanged is a log parse operation binding the contract event 0xc44c80d7c6ac73ee13b2c62e50522dadb6c646b4509ca93e52a45abb38a78e33.
//
// Solidity: event InterestRatePerBlockChanged(uint256 oldValue, uint256 newValue)
func (_SwapPair *SwapPairFilterer) ParseInterestRatePerBlockChanged(log types.Log) (*SwapPairInterestRatePerBlockChanged, error) {
	event := new(SwapPairInterestRatePerBlockChanged)
	if err := _SwapPair.contract.UnpackLog(event, "InterestRatePerBlockChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairProductivityDecreasedIterator is returned from FilterProductivityDecreased and is used to iterate over the raw logs and unpacked data for ProductivityDecreased events raised by the SwapPair contract.
type SwapPairProductivityDecreasedIterator struct {
	Event *SwapPairProductivityDecreased // Event containing the contract specifics and raw log

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
func (it *SwapPairProductivityDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairProductivityDecreased)
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
		it.Event = new(SwapPairProductivityDecreased)
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
func (it *SwapPairProductivityDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairProductivityDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairProductivityDecreased represents a ProductivityDecreased event raised by the SwapPair contract.
type SwapPairProductivityDecreased struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProductivityDecreased is a free log retrieval operation binding the contract event 0xddb757202feefdd10c1666f1bb8f9744309ab811b6ef3ec0404a20c36e426697.
//
// Solidity: event ProductivityDecreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterProductivityDecreased(opts *bind.FilterOpts, user []common.Address) (*SwapPairProductivityDecreasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "ProductivityDecreased", userRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairProductivityDecreasedIterator{contract: _SwapPair.contract, event: "ProductivityDecreased", logs: logs, sub: sub}, nil
}

// WatchProductivityDecreased is a free log subscription operation binding the contract event 0xddb757202feefdd10c1666f1bb8f9744309ab811b6ef3ec0404a20c36e426697.
//
// Solidity: event ProductivityDecreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchProductivityDecreased(opts *bind.WatchOpts, sink chan<- *SwapPairProductivityDecreased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "ProductivityDecreased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairProductivityDecreased)
				if err := _SwapPair.contract.UnpackLog(event, "ProductivityDecreased", log); err != nil {
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

// ParseProductivityDecreased is a log parse operation binding the contract event 0xddb757202feefdd10c1666f1bb8f9744309ab811b6ef3ec0404a20c36e426697.
//
// Solidity: event ProductivityDecreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) ParseProductivityDecreased(log types.Log) (*SwapPairProductivityDecreased, error) {
	event := new(SwapPairProductivityDecreased)
	if err := _SwapPair.contract.UnpackLog(event, "ProductivityDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairProductivityIncreasedIterator is returned from FilterProductivityIncreased and is used to iterate over the raw logs and unpacked data for ProductivityIncreased events raised by the SwapPair contract.
type SwapPairProductivityIncreasedIterator struct {
	Event *SwapPairProductivityIncreased // Event containing the contract specifics and raw log

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
func (it *SwapPairProductivityIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairProductivityIncreased)
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
		it.Event = new(SwapPairProductivityIncreased)
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
func (it *SwapPairProductivityIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairProductivityIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairProductivityIncreased represents a ProductivityIncreased event raised by the SwapPair contract.
type SwapPairProductivityIncreased struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProductivityIncreased is a free log retrieval operation binding the contract event 0x6c0a24b06ee7f1d9c3c6680c7328531c32bc59cd665436fc686f973cb8c01be7.
//
// Solidity: event ProductivityIncreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterProductivityIncreased(opts *bind.FilterOpts, user []common.Address) (*SwapPairProductivityIncreasedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "ProductivityIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairProductivityIncreasedIterator{contract: _SwapPair.contract, event: "ProductivityIncreased", logs: logs, sub: sub}, nil
}

// WatchProductivityIncreased is a free log subscription operation binding the contract event 0x6c0a24b06ee7f1d9c3c6680c7328531c32bc59cd665436fc686f973cb8c01be7.
//
// Solidity: event ProductivityIncreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchProductivityIncreased(opts *bind.WatchOpts, sink chan<- *SwapPairProductivityIncreased, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "ProductivityIncreased", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairProductivityIncreased)
				if err := _SwapPair.contract.UnpackLog(event, "ProductivityIncreased", log); err != nil {
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

// ParseProductivityIncreased is a log parse operation binding the contract event 0x6c0a24b06ee7f1d9c3c6680c7328531c32bc59cd665436fc686f973cb8c01be7.
//
// Solidity: event ProductivityIncreased(address indexed user, uint256 value)
func (_SwapPair *SwapPairFilterer) ParseProductivityIncreased(log types.Log) (*SwapPairProductivityIncreased, error) {
	event := new(SwapPairProductivityIncreased)
	if err := _SwapPair.contract.UnpackLog(event, "ProductivityIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SwapPairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SwapPair contract.
type SwapPairTransferIterator struct {
	Event *SwapPairTransfer // Event containing the contract specifics and raw log

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
func (it *SwapPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPairTransfer)
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
		it.Event = new(SwapPairTransfer)
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
func (it *SwapPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPairTransfer represents a Transfer event raised by the SwapPair contract.
type SwapPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SwapPair *SwapPairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SwapPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SwapPairTransferIterator{contract: _SwapPair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SwapPair *SwapPairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SwapPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SwapPair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPairTransfer)
				if err := _SwapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SwapPair *SwapPairFilterer) ParseTransfer(log types.Log) (*SwapPairTransfer, error) {
	event := new(SwapPairTransfer)
	if err := _SwapPair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
