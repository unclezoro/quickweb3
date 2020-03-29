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

// ValidatorABI is the input ABI used to generate the binding from.
const ValidatorABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"batchTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deprecatedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"systemTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorFelony\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorMisdemeanor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"validatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DUSTY_INCOMING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTRA_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JAIL_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BBCFeeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"incoming\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fromChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initLightClientAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initSlashContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initSystemRewardAddr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initTokenHubAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initValidatorSetBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyPrefix\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousDepositHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInComing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getIncoming\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// CHANNELID is a free data retrieval call binding the contract method 0xb0c0c514.
//
// Solidity: function CHANNEL_ID() constant returns(uint8)
func (_Validator *ValidatorCaller) CHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "CHANNEL_ID")
	return *ret0, err
}

// CHANNELID is a free data retrieval call binding the contract method 0xb0c0c514.
//
// Solidity: function CHANNEL_ID() constant returns(uint8)
func (_Validator *ValidatorSession) CHANNELID() (uint8, error) {
	return _Validator.Contract.CHANNELID(&_Validator.CallOpts)
}

// CHANNELID is a free data retrieval call binding the contract method 0xb0c0c514.
//
// Solidity: function CHANNEL_ID() constant returns(uint8)
func (_Validator *ValidatorCallerSession) CHANNELID() (uint8, error) {
	return _Validator.Contract.CHANNELID(&_Validator.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Validator *ValidatorCaller) DUSTYINCOMING(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "DUSTY_INCOMING")
	return *ret0, err
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Validator *ValidatorSession) DUSTYINCOMING() (*big.Int, error) {
	return _Validator.Contract.DUSTYINCOMING(&_Validator.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() constant returns(uint256)
func (_Validator *ValidatorCallerSession) DUSTYINCOMING() (*big.Int, error) {
	return _Validator.Contract.DUSTYINCOMING(&_Validator.CallOpts)
}

// EXTRAFEE is a free data retrieval call binding the contract method 0x6a65855e.
//
// Solidity: function EXTRA_FEE() constant returns(uint256)
func (_Validator *ValidatorCaller) EXTRAFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "EXTRA_FEE")
	return *ret0, err
}

// EXTRAFEE is a free data retrieval call binding the contract method 0x6a65855e.
//
// Solidity: function EXTRA_FEE() constant returns(uint256)
func (_Validator *ValidatorSession) EXTRAFEE() (*big.Int, error) {
	return _Validator.Contract.EXTRAFEE(&_Validator.CallOpts)
}

// EXTRAFEE is a free data retrieval call binding the contract method 0x6a65855e.
//
// Solidity: function EXTRA_FEE() constant returns(uint256)
func (_Validator *ValidatorCallerSession) EXTRAFEE() (*big.Int, error) {
	return _Validator.Contract.EXTRAFEE(&_Validator.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorCaller) JAILMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "JAIL_MESSAGE_TYPE")
	return *ret0, err
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorSession) JAILMESSAGETYPE() (uint8, error) {
	return _Validator.Contract.JAILMESSAGETYPE(&_Validator.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorCallerSession) JAILMESSAGETYPE() (uint8, error) {
	return _Validator.Contract.JAILMESSAGETYPE(&_Validator.CallOpts)
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Validator *ValidatorCaller) RELAYERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "RELAYER_REWARD")
	return *ret0, err
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Validator *ValidatorSession) RELAYERREWARD() (*big.Int, error) {
	return _Validator.Contract.RELAYERREWARD(&_Validator.CallOpts)
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Validator *ValidatorCallerSession) RELAYERREWARD() (*big.Int, error) {
	return _Validator.Contract.RELAYERREWARD(&_Validator.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() constant returns(address)
func (_Validator *ValidatorCaller) SYSTEMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "SYSTEM_ADDRESS")
	return *ret0, err
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() constant returns(address)
func (_Validator *ValidatorSession) SYSTEMADDRESS() (common.Address, error) {
	return _Validator.Contract.SYSTEMADDRESS(&_Validator.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() constant returns(address)
func (_Validator *ValidatorCallerSession) SYSTEMADDRESS() (common.Address, error) {
	return _Validator.Contract.SYSTEMADDRESS(&_Validator.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "VALIDATORS_UPDATE_MESSAGE_TYPE")
	return *ret0, err
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Validator.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Validator.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() constant returns(uint8)
func (_Validator *ValidatorCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _Validator.Contract.VALIDATORSUPDATEMESSAGETYPE(&_Validator.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Validator *ValidatorCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Validator *ValidatorSession) AlreadyInit() (bool, error) {
	return _Validator.Contract.AlreadyInit(&_Validator.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Validator *ValidatorCallerSession) AlreadyInit() (bool, error) {
	return _Validator.Contract.AlreadyInit(&_Validator.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validator *ValidatorCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	ret := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	out := ret
	err := _Validator.contract.Call(opts, out, "currentValidatorSet", arg0)
	return *ret, err
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validator *ValidatorSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Validator.Contract.CurrentValidatorSet(&_Validator.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) constant returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_Validator *ValidatorCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _Validator.Contract.CurrentValidatorSet(&_Validator.CallOpts, arg0)
}

// FromChainId is a free data retrieval call binding the contract method 0x3e2ae2f8.
//
// Solidity: function fromChainId() constant returns(uint16)
func (_Validator *ValidatorCaller) FromChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "fromChainId")
	return *ret0, err
}

// FromChainId is a free data retrieval call binding the contract method 0x3e2ae2f8.
//
// Solidity: function fromChainId() constant returns(uint16)
func (_Validator *ValidatorSession) FromChainId() (uint16, error) {
	return _Validator.Contract.FromChainId(&_Validator.CallOpts)
}

// FromChainId is a free data retrieval call binding the contract method 0x3e2ae2f8.
//
// Solidity: function fromChainId() constant returns(uint16)
func (_Validator *ValidatorCallerSession) FromChainId() (uint16, error) {
	return _Validator.Contract.FromChainId(&_Validator.CallOpts)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Validator *ValidatorCaller) GetIncoming(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getIncoming", validator)
	return *ret0, err
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Validator *ValidatorSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Validator.Contract.GetIncoming(&_Validator.CallOpts, validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) constant returns(uint256)
func (_Validator *ValidatorCallerSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _Validator.Contract.GetIncoming(&_Validator.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_Validator *ValidatorCallerSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// InitLightClientAddr is a free data retrieval call binding the contract method 0x355fd362.
//
// Solidity: function initLightClientAddr() constant returns(address)
func (_Validator *ValidatorCaller) InitLightClientAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "initLightClientAddr")
	return *ret0, err
}

// InitLightClientAddr is a free data retrieval call binding the contract method 0x355fd362.
//
// Solidity: function initLightClientAddr() constant returns(address)
func (_Validator *ValidatorSession) InitLightClientAddr() (common.Address, error) {
	return _Validator.Contract.InitLightClientAddr(&_Validator.CallOpts)
}

// InitLightClientAddr is a free data retrieval call binding the contract method 0x355fd362.
//
// Solidity: function initLightClientAddr() constant returns(address)
func (_Validator *ValidatorCallerSession) InitLightClientAddr() (common.Address, error) {
	return _Validator.Contract.InitLightClientAddr(&_Validator.CallOpts)
}

// InitSlashContract is a free data retrieval call binding the contract method 0x2bd3c7aa.
//
// Solidity: function initSlashContract() constant returns(address)
func (_Validator *ValidatorCaller) InitSlashContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "initSlashContract")
	return *ret0, err
}

// InitSlashContract is a free data retrieval call binding the contract method 0x2bd3c7aa.
//
// Solidity: function initSlashContract() constant returns(address)
func (_Validator *ValidatorSession) InitSlashContract() (common.Address, error) {
	return _Validator.Contract.InitSlashContract(&_Validator.CallOpts)
}

// InitSlashContract is a free data retrieval call binding the contract method 0x2bd3c7aa.
//
// Solidity: function initSlashContract() constant returns(address)
func (_Validator *ValidatorCallerSession) InitSlashContract() (common.Address, error) {
	return _Validator.Contract.InitSlashContract(&_Validator.CallOpts)
}

// InitSystemRewardAddr is a free data retrieval call binding the contract method 0xf4154338.
//
// Solidity: function initSystemRewardAddr() constant returns(address)
func (_Validator *ValidatorCaller) InitSystemRewardAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "initSystemRewardAddr")
	return *ret0, err
}

// InitSystemRewardAddr is a free data retrieval call binding the contract method 0xf4154338.
//
// Solidity: function initSystemRewardAddr() constant returns(address)
func (_Validator *ValidatorSession) InitSystemRewardAddr() (common.Address, error) {
	return _Validator.Contract.InitSystemRewardAddr(&_Validator.CallOpts)
}

// InitSystemRewardAddr is a free data retrieval call binding the contract method 0xf4154338.
//
// Solidity: function initSystemRewardAddr() constant returns(address)
func (_Validator *ValidatorCallerSession) InitSystemRewardAddr() (common.Address, error) {
	return _Validator.Contract.InitSystemRewardAddr(&_Validator.CallOpts)
}

// InitTokenHubAddr is a free data retrieval call binding the contract method 0x83b27d06.
//
// Solidity: function initTokenHubAddr() constant returns(address)
func (_Validator *ValidatorCaller) InitTokenHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "initTokenHubAddr")
	return *ret0, err
}

// InitTokenHubAddr is a free data retrieval call binding the contract method 0x83b27d06.
//
// Solidity: function initTokenHubAddr() constant returns(address)
func (_Validator *ValidatorSession) InitTokenHubAddr() (common.Address, error) {
	return _Validator.Contract.InitTokenHubAddr(&_Validator.CallOpts)
}

// InitTokenHubAddr is a free data retrieval call binding the contract method 0x83b27d06.
//
// Solidity: function initTokenHubAddr() constant returns(address)
func (_Validator *ValidatorCallerSession) InitTokenHubAddr() (common.Address, error) {
	return _Validator.Contract.InitTokenHubAddr(&_Validator.CallOpts)
}

// InitValidatorSetBytes is a free data retrieval call binding the contract method 0x32c55e89.
//
// Solidity: function initValidatorSetBytes() constant returns(bytes)
func (_Validator *ValidatorCaller) InitValidatorSetBytes(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "initValidatorSetBytes")
	return *ret0, err
}

// InitValidatorSetBytes is a free data retrieval call binding the contract method 0x32c55e89.
//
// Solidity: function initValidatorSetBytes() constant returns(bytes)
func (_Validator *ValidatorSession) InitValidatorSetBytes() ([]byte, error) {
	return _Validator.Contract.InitValidatorSetBytes(&_Validator.CallOpts)
}

// InitValidatorSetBytes is a free data retrieval call binding the contract method 0x32c55e89.
//
// Solidity: function initValidatorSetBytes() constant returns(bytes)
func (_Validator *ValidatorCallerSession) InitValidatorSetBytes() ([]byte, error) {
	return _Validator.Contract.InitValidatorSetBytes(&_Validator.CallOpts)
}

// KeyPrefix is a free data retrieval call binding the contract method 0x5fe1ef8b.
//
// Solidity: function keyPrefix() constant returns(bytes)
func (_Validator *ValidatorCaller) KeyPrefix(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "keyPrefix")
	return *ret0, err
}

// KeyPrefix is a free data retrieval call binding the contract method 0x5fe1ef8b.
//
// Solidity: function keyPrefix() constant returns(bytes)
func (_Validator *ValidatorSession) KeyPrefix() ([]byte, error) {
	return _Validator.Contract.KeyPrefix(&_Validator.CallOpts)
}

// KeyPrefix is a free data retrieval call binding the contract method 0x5fe1ef8b.
//
// Solidity: function keyPrefix() constant returns(bytes)
func (_Validator *ValidatorCallerSession) KeyPrefix() ([]byte, error) {
	return _Validator.Contract.KeyPrefix(&_Validator.CallOpts)
}

// PreviousDepositHeight is a free data retrieval call binding the contract method 0x4e3e3ca9.
//
// Solidity: function previousDepositHeight() constant returns(uint64)
func (_Validator *ValidatorCaller) PreviousDepositHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "previousDepositHeight")
	return *ret0, err
}

// PreviousDepositHeight is a free data retrieval call binding the contract method 0x4e3e3ca9.
//
// Solidity: function previousDepositHeight() constant returns(uint64)
func (_Validator *ValidatorSession) PreviousDepositHeight() (uint64, error) {
	return _Validator.Contract.PreviousDepositHeight(&_Validator.CallOpts)
}

// PreviousDepositHeight is a free data retrieval call binding the contract method 0x4e3e3ca9.
//
// Solidity: function previousDepositHeight() constant returns(uint64)
func (_Validator *ValidatorCallerSession) PreviousDepositHeight() (uint64, error) {
	return _Validator.Contract.PreviousDepositHeight(&_Validator.CallOpts)
}

// Sequence is a free data retrieval call binding the contract method 0x529d15cc.
//
// Solidity: function sequence() constant returns(uint64)
func (_Validator *ValidatorCaller) Sequence(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "sequence")
	return *ret0, err
}

// Sequence is a free data retrieval call binding the contract method 0x529d15cc.
//
// Solidity: function sequence() constant returns(uint64)
func (_Validator *ValidatorSession) Sequence() (uint64, error) {
	return _Validator.Contract.Sequence(&_Validator.CallOpts)
}

// Sequence is a free data retrieval call binding the contract method 0x529d15cc.
//
// Solidity: function sequence() constant returns(uint64)
func (_Validator *ValidatorCallerSession) Sequence() (uint64, error) {
	return _Validator.Contract.Sequence(&_Validator.CallOpts)
}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() constant returns(uint16)
func (_Validator *ValidatorCaller) ToChainId(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "toChainId")
	return *ret0, err
}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() constant returns(uint16)
func (_Validator *ValidatorSession) ToChainId() (uint16, error) {
	return _Validator.Contract.ToChainId(&_Validator.CallOpts)
}

// ToChainId is a free data retrieval call binding the contract method 0x07bb7655.
//
// Solidity: function toChainId() constant returns(uint16)
func (_Validator *ValidatorCallerSession) ToChainId() (uint16, error) {
	return _Validator.Contract.ToChainId(&_Validator.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Validator *ValidatorCaller) TotalInComing(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validator.contract.Call(opts, out, "totalInComing")
	return *ret0, err
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Validator *ValidatorSession) TotalInComing() (*big.Int, error) {
	return _Validator.Contract.TotalInComing(&_Validator.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() constant returns(uint256)
func (_Validator *ValidatorCallerSession) TotalInComing() (*big.Int, error) {
	return _Validator.Contract.TotalInComing(&_Validator.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Validator *ValidatorTransactor) Deposit(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "deposit", valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Validator *ValidatorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Deposit(&_Validator.TransactOpts, valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) returns()
func (_Validator *ValidatorTransactorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Deposit(&_Validator.TransactOpts, valAddr)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validator *ValidatorTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validator *ValidatorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Felony(&_Validator.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_Validator *ValidatorTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Felony(&_Validator.TransactOpts, validator)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validator *ValidatorTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validator *ValidatorSession) Init() (*types.Transaction, error) {
	return _Validator.Contract.Init(&_Validator.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Validator *ValidatorTransactorSession) Init() (*types.Transaction, error) {
	return _Validator.Contract.Init(&_Validator.TransactOpts)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validator *ValidatorTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validator *ValidatorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Misdemeanor(&_Validator.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_Validator *ValidatorTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _Validator.Contract.Misdemeanor(&_Validator.TransactOpts, validator)
}

// Update is a paid mutator transaction binding the contract method 0xe695c4ae.
//
// Solidity: function update(bytes msgBytes, bytes proof, uint64 height, uint64 packageSequence) returns()
func (_Validator *ValidatorTransactor) Update(opts *bind.TransactOpts, msgBytes []byte, proof []byte, height uint64, packageSequence uint64) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "update", msgBytes, proof, height, packageSequence)
}

// Update is a paid mutator transaction binding the contract method 0xe695c4ae.
//
// Solidity: function update(bytes msgBytes, bytes proof, uint64 height, uint64 packageSequence) returns()
func (_Validator *ValidatorSession) Update(msgBytes []byte, proof []byte, height uint64, packageSequence uint64) (*types.Transaction, error) {
	return _Validator.Contract.Update(&_Validator.TransactOpts, msgBytes, proof, height, packageSequence)
}

// Update is a paid mutator transaction binding the contract method 0xe695c4ae.
//
// Solidity: function update(bytes msgBytes, bytes proof, uint64 height, uint64 packageSequence) returns()
func (_Validator *ValidatorTransactorSession) Update(msgBytes []byte, proof []byte, height uint64, packageSequence uint64) (*types.Transaction, error) {
	return _Validator.Contract.Update(&_Validator.TransactOpts, msgBytes, proof, height, packageSequence)
}

// ValidatorBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the Validator contract.
type ValidatorBatchTransferIterator struct {
	Event *ValidatorBatchTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBatchTransfer)
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
		it.Event = new(ValidatorBatchTransfer)
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
func (it *ValidatorBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBatchTransfer represents a BatchTransfer event raised by the Validator contract.
type ValidatorBatchTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterBatchTransfer(opts *bind.FilterOpts, amount []*big.Int) (*ValidatorBatchTransferIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "batchTransfer", amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBatchTransferIterator{contract: _Validator.contract, event: "batchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorBatchTransfer, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "batchTransfer", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBatchTransfer)
				if err := _Validator.contract.UnpackLog(event, "batchTransfer", log); err != nil {
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

// ParseBatchTransfer is a log parse operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseBatchTransfer(log types.Log) (*ValidatorBatchTransfer, error) {
	event := new(ValidatorBatchTransfer)
	if err := _Validator.contract.UnpackLog(event, "batchTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorDeprecatedDepositIterator is returned from FilterDeprecatedDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedDeposit events raised by the Validator contract.
type ValidatorDeprecatedDepositIterator struct {
	Event *ValidatorDeprecatedDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorDeprecatedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDeprecatedDeposit)
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
		it.Event = new(ValidatorDeprecatedDeposit)
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
func (it *ValidatorDeprecatedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDeprecatedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDeprecatedDeposit represents a DeprecatedDeposit event raised by the Validator contract.
type ValidatorDeprecatedDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedDeposit is a free log retrieval operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterDeprecatedDeposit(opts *bind.FilterOpts, validator []common.Address, amount []*big.Int) (*ValidatorDeprecatedDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "deprecatedDeposit", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDeprecatedDepositIterator{contract: _Validator.contract, event: "deprecatedDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedDeposit is a free log subscription operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchDeprecatedDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorDeprecatedDeposit, validator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "deprecatedDeposit", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDeprecatedDeposit)
				if err := _Validator.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
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

// ParseDeprecatedDeposit is a log parse operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseDeprecatedDeposit(log types.Log) (*ValidatorDeprecatedDeposit, error) {
	event := new(ValidatorDeprecatedDeposit)
	if err := _Validator.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorDirectTransferIterator is returned from FilterDirectTransfer and is used to iterate over the raw logs and unpacked data for DirectTransfer events raised by the Validator contract.
type ValidatorDirectTransferIterator struct {
	Event *ValidatorDirectTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorDirectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDirectTransfer)
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
		it.Event = new(ValidatorDirectTransfer)
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
func (it *ValidatorDirectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDirectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDirectTransfer represents a DirectTransfer event raised by the Validator contract.
type ValidatorDirectTransfer struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransfer is a free log retrieval operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterDirectTransfer(opts *bind.FilterOpts, validator []common.Address, amount []*big.Int) (*ValidatorDirectTransferIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "directTransfer", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDirectTransferIterator{contract: _Validator.contract, event: "directTransfer", logs: logs, sub: sub}, nil
}

// WatchDirectTransfer is a free log subscription operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchDirectTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorDirectTransfer, validator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "directTransfer", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDirectTransfer)
				if err := _Validator.contract.UnpackLog(event, "directTransfer", log); err != nil {
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

// ParseDirectTransfer is a log parse operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseDirectTransfer(log types.Log) (*ValidatorDirectTransfer, error) {
	event := new(ValidatorDirectTransfer)
	if err := _Validator.contract.UnpackLog(event, "directTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorSystemTransferIterator is returned from FilterSystemTransfer and is used to iterate over the raw logs and unpacked data for SystemTransfer events raised by the Validator contract.
type ValidatorSystemTransferIterator struct {
	Event *ValidatorSystemTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorSystemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSystemTransfer)
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
		it.Event = new(ValidatorSystemTransfer)
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
func (it *ValidatorSystemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSystemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSystemTransfer represents a SystemTransfer event raised by the Validator contract.
type ValidatorSystemTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemTransfer is a free log retrieval operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterSystemTransfer(opts *bind.FilterOpts, amount []*big.Int) (*ValidatorSystemTransferIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "systemTransfer", amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSystemTransferIterator{contract: _Validator.contract, event: "systemTransfer", logs: logs, sub: sub}, nil
}

// WatchSystemTransfer is a free log subscription operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchSystemTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorSystemTransfer, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "systemTransfer", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSystemTransfer)
				if err := _Validator.contract.UnpackLog(event, "systemTransfer", log); err != nil {
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

// ParseSystemTransfer is a log parse operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseSystemTransfer(log types.Log) (*ValidatorSystemTransfer, error) {
	event := new(ValidatorSystemTransfer)
	if err := _Validator.contract.UnpackLog(event, "systemTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorDepositIterator is returned from FilterValidatorDeposit and is used to iterate over the raw logs and unpacked data for ValidatorDeposit events raised by the Validator contract.
type ValidatorValidatorDepositIterator struct {
	Event *ValidatorValidatorDeposit // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorDeposit)
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
		it.Event = new(ValidatorValidatorDeposit)
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
func (it *ValidatorValidatorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorDeposit represents a ValidatorDeposit event raised by the Validator contract.
type ValidatorValidatorDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposit is a free log retrieval operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterValidatorDeposit(opts *bind.FilterOpts, validator []common.Address, amount []*big.Int) (*ValidatorValidatorDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "validatorDeposit", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorDepositIterator{contract: _Validator.contract, event: "validatorDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposit is a free log subscription operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchValidatorDeposit(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorDeposit, validator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "validatorDeposit", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorDeposit)
				if err := _Validator.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
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

// ParseValidatorDeposit is a log parse operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseValidatorDeposit(log types.Log) (*ValidatorValidatorDeposit, error) {
	event := new(ValidatorValidatorDeposit)
	if err := _Validator.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorFelonyIterator is returned from FilterValidatorFelony and is used to iterate over the raw logs and unpacked data for ValidatorFelony events raised by the Validator contract.
type ValidatorValidatorFelonyIterator struct {
	Event *ValidatorValidatorFelony // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorFelonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorFelony)
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
		it.Event = new(ValidatorValidatorFelony)
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
func (it *ValidatorValidatorFelonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorFelonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorFelony represents a ValidatorFelony event raised by the Validator contract.
type ValidatorValidatorFelony struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorFelony is a free log retrieval operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterValidatorFelony(opts *bind.FilterOpts, validator []common.Address, amount []*big.Int) (*ValidatorValidatorFelonyIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "validatorFelony", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorFelonyIterator{contract: _Validator.contract, event: "validatorFelony", logs: logs, sub: sub}, nil
}

// WatchValidatorFelony is a free log subscription operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchValidatorFelony(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorFelony, validator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "validatorFelony", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorFelony)
				if err := _Validator.contract.UnpackLog(event, "validatorFelony", log); err != nil {
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

// ParseValidatorFelony is a log parse operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseValidatorFelony(log types.Log) (*ValidatorValidatorFelony, error) {
	event := new(ValidatorValidatorFelony)
	if err := _Validator.contract.UnpackLog(event, "validatorFelony", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Validator contract.
type ValidatorValidatorJailedIterator struct {
	Event *ValidatorValidatorJailed // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorJailed)
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
		it.Event = new(ValidatorValidatorJailed)
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
func (it *ValidatorValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorJailed represents a ValidatorJailed event raised by the Validator contract.
type ValidatorValidatorJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validator *ValidatorFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorJailedIterator{contract: _Validator.contract, event: "validatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validator *ValidatorFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorJailed)
				if err := _Validator.contract.UnpackLog(event, "validatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_Validator *ValidatorFilterer) ParseValidatorJailed(log types.Log) (*ValidatorValidatorJailed, error) {
	event := new(ValidatorValidatorJailed)
	if err := _Validator.contract.UnpackLog(event, "validatorJailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorMisdemeanorIterator is returned from FilterValidatorMisdemeanor and is used to iterate over the raw logs and unpacked data for ValidatorMisdemeanor events raised by the Validator contract.
type ValidatorValidatorMisdemeanorIterator struct {
	Event *ValidatorValidatorMisdemeanor // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorMisdemeanorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorMisdemeanor)
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
		it.Event = new(ValidatorValidatorMisdemeanor)
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
func (it *ValidatorValidatorMisdemeanorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorMisdemeanorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorMisdemeanor represents a ValidatorMisdemeanor event raised by the Validator contract.
type ValidatorValidatorMisdemeanor struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorMisdemeanor is a free log retrieval operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) FilterValidatorMisdemeanor(opts *bind.FilterOpts, validator []common.Address, amount []*big.Int) (*ValidatorValidatorMisdemeanorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "validatorMisdemeanor", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorMisdemeanorIterator{contract: _Validator.contract, event: "validatorMisdemeanor", logs: logs, sub: sub}, nil
}

// WatchValidatorMisdemeanor is a free log subscription operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) WatchValidatorMisdemeanor(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorMisdemeanor, validator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "validatorMisdemeanor", validatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorMisdemeanor)
				if err := _Validator.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
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

// ParseValidatorMisdemeanor is a log parse operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 indexed amount)
func (_Validator *ValidatorFilterer) ParseValidatorMisdemeanor(log types.Log) (*ValidatorValidatorMisdemeanor, error) {
	event := new(ValidatorValidatorMisdemeanor)
	if err := _Validator.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Validator contract.
type ValidatorValidatorSetUpdatedIterator struct {
	Event *ValidatorValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorSetUpdated)
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
		it.Event = new(ValidatorValidatorSetUpdated)
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
func (it *ValidatorValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Validator contract.
type ValidatorValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validator *ValidatorFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*ValidatorValidatorSetUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorSetUpdatedIterator{contract: _Validator.contract, event: "validatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validator *ValidatorFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_Validator *ValidatorFilterer) ParseValidatorSetUpdated(log types.Log) (*ValidatorValidatorSetUpdated, error) {
	event := new(ValidatorValidatorSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
