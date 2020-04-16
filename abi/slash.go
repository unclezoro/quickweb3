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
const SlashABI = "[{\"anonymous\":false,\"inputs\":[],\"name\":\"crashResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"indicatorCleaned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"knownResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"unKnownResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorSlashed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BSC_RELAYER_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FELONY_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MISDEMEANOR_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"felonyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getSlashIndicator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"misdemeanorThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) BINDCHANNELID() (uint8, error) {
	return _Slash.Contract.BINDCHANNELID(&_Slash.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) BINDCHANNELID() (uint8, error) {
	return _Slash.Contract.BINDCHANNELID(&_Slash.CallOpts)
}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() constant returns(uint256)
func (_Slash *SlashCaller) BSCRELAYERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "BSC_RELAYER_REWARD")
	return *ret0, err
}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() constant returns(uint256)
func (_Slash *SlashSession) BSCRELAYERREWARD() (*big.Int, error) {
	return _Slash.Contract.BSCRELAYERREWARD(&_Slash.CallOpts)
}

// BSCRELAYERREWARD is a free data retrieval call binding the contract method 0x9bc8e4f2.
//
// Solidity: function BSC_RELAYER_REWARD() constant returns(uint256)
func (_Slash *SlashCallerSession) BSCRELAYERREWARD() (*big.Int, error) {
	return _Slash.Contract.BSCRELAYERREWARD(&_Slash.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Slash *SlashCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Slash *SlashSession) CODEOK() (uint32, error) {
	return _Slash.Contract.CODEOK(&_Slash.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Slash *SlashCallerSession) CODEOK() (uint32, error) {
	return _Slash.Contract.CODEOK(&_Slash.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.CROSSCHAINCONTRACTADDR(&_Slash.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.CROSSCHAINCONTRACTADDR(&_Slash.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Slash *SlashCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Slash *SlashSession) ERRORFAILDECODE() (uint32, error) {
	return _Slash.Contract.ERRORFAILDECODE(&_Slash.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Slash *SlashCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Slash.Contract.ERRORFAILDECODE(&_Slash.CallOpts)
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

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) GOVCHANNELID() (uint8, error) {
	return _Slash.Contract.GOVCHANNELID(&_Slash.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) GOVCHANNELID() (uint8, error) {
	return _Slash.Contract.GOVCHANNELID(&_Slash.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Slash *SlashCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Slash *SlashSession) GOVHUBADDR() (common.Address, error) {
	return _Slash.Contract.GOVHUBADDR(&_Slash.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Slash.Contract.GOVHUBADDR(&_Slash.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Slash *SlashCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Slash *SlashSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Slash.Contract.INCENTIVIZEADDR(&_Slash.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Slash.Contract.INCENTIVIZEADDR(&_Slash.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Slash *SlashCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Slash *SlashSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Slash.Contract.LIGHTCLIENTADDR(&_Slash.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Slash.Contract.LIGHTCLIENTADDR(&_Slash.CallOpts)
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

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.RELAYERHUBCONTRACTADDR(&_Slash.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.RELAYERHUBCONTRACTADDR(&_Slash.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) SLASHCHANNELID() (uint8, error) {
	return _Slash.Contract.SLASHCHANNELID(&_Slash.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Slash.Contract.SLASHCHANNELID(&_Slash.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.SLASHCONTRACTADDR(&_Slash.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Slash.Contract.SLASHCONTRACTADDR(&_Slash.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) STAKINGCHANNELID() (uint8, error) {
	return _Slash.Contract.STAKINGCHANNELID(&_Slash.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Slash.Contract.STAKINGCHANNELID(&_Slash.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Slash *SlashCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Slash *SlashSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Slash.Contract.SYSTEMREWARDADDR(&_Slash.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Slash.Contract.SYSTEMREWARDADDR(&_Slash.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Slash *SlashCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Slash *SlashSession) TOKENHUBADDR() (common.Address, error) {
	return _Slash.Contract.TOKENHUBADDR(&_Slash.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Slash *SlashCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Slash.Contract.TOKENHUBADDR(&_Slash.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Slash.Contract.TRANSFERINCHANNELID(&_Slash.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Slash.Contract.TRANSFERINCHANNELID(&_Slash.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Slash *SlashCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Slash *SlashSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Slash.Contract.TRANSFEROUTCHANNELID(&_Slash.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Slash *SlashCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Slash.Contract.TRANSFEROUTCHANNELID(&_Slash.CallOpts)
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

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Slash *SlashCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Slash *SlashSession) AlreadyInit() (bool, error) {
	return _Slash.Contract.AlreadyInit(&_Slash.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Slash *SlashCallerSession) AlreadyInit() (bool, error) {
	return _Slash.Contract.AlreadyInit(&_Slash.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Slash *SlashCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Slash *SlashSession) BscChainID() (uint16, error) {
	return _Slash.Contract.BscChainID(&_Slash.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Slash *SlashCallerSession) BscChainID() (uint16, error) {
	return _Slash.Contract.BscChainID(&_Slash.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() constant returns(uint256)
func (_Slash *SlashCaller) FelonyThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "felonyThreshold")
	return *ret0, err
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() constant returns(uint256)
func (_Slash *SlashSession) FelonyThreshold() (*big.Int, error) {
	return _Slash.Contract.FelonyThreshold(&_Slash.CallOpts)
}

// FelonyThreshold is a free data retrieval call binding the contract method 0x389f4f71.
//
// Solidity: function felonyThreshold() constant returns(uint256)
func (_Slash *SlashCallerSession) FelonyThreshold() (*big.Int, error) {
	return _Slash.Contract.FelonyThreshold(&_Slash.CallOpts)
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

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() constant returns(uint256)
func (_Slash *SlashCaller) MisdemeanorThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Slash.contract.Call(opts, out, "misdemeanorThreshold")
	return *ret0, err
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() constant returns(uint256)
func (_Slash *SlashSession) MisdemeanorThreshold() (*big.Int, error) {
	return _Slash.Contract.MisdemeanorThreshold(&_Slash.CallOpts)
}

// MisdemeanorThreshold is a free data retrieval call binding the contract method 0x567a372d.
//
// Solidity: function misdemeanorThreshold() constant returns(uint256)
func (_Slash *SlashCallerSession) MisdemeanorThreshold() (*big.Int, error) {
	return _Slash.Contract.MisdemeanorThreshold(&_Slash.CallOpts)
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

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_Slash *SlashTransactor) HandleAckPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "handleAckPackage", arg0, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_Slash *SlashSession) HandleAckPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleAckPackage(&_Slash.TransactOpts, arg0, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 , bytes msgBytes) returns()
func (_Slash *SlashTransactorSession) HandleAckPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleAckPackage(&_Slash.TransactOpts, arg0, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Slash *SlashTransactor) HandleFailAckPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "handleFailAckPackage", arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Slash *SlashSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleFailAckPackage(&_Slash.TransactOpts, arg0, arg1)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 , bytes ) returns()
func (_Slash *SlashTransactorSession) HandleFailAckPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleFailAckPackage(&_Slash.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_Slash *SlashTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "handleSynPackage", arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_Slash *SlashSession) HandleSynPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleSynPackage(&_Slash.TransactOpts, arg0, arg1)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes ) returns(bytes)
func (_Slash *SlashTransactorSession) HandleSynPackage(arg0 uint8, arg1 []byte) (*types.Transaction, error) {
	return _Slash.Contract.HandleSynPackage(&_Slash.TransactOpts, arg0, arg1)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Slash *SlashTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Slash *SlashSession) Init() (*types.Transaction, error) {
	return _Slash.Contract.Init(&_Slash.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Slash *SlashTransactorSession) Init() (*types.Transaction, error) {
	return _Slash.Contract.Init(&_Slash.TransactOpts)
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

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Slash *SlashTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Slash.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Slash *SlashSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Slash.Contract.UpdateParam(&_Slash.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Slash *SlashTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Slash.Contract.UpdateParam(&_Slash.TransactOpts, key, value)
}

// SlashCrashResponseIterator is returned from FilterCrashResponse and is used to iterate over the raw logs and unpacked data for CrashResponse events raised by the Slash contract.
type SlashCrashResponseIterator struct {
	Event *SlashCrashResponse // Event containing the contract specifics and raw log

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
func (it *SlashCrashResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashCrashResponse)
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
		it.Event = new(SlashCrashResponse)
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
func (it *SlashCrashResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashCrashResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashCrashResponse represents a CrashResponse event raised by the Slash contract.
type SlashCrashResponse struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrashResponse is a free log retrieval operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_Slash *SlashFilterer) FilterCrashResponse(opts *bind.FilterOpts) (*SlashCrashResponseIterator, error) {

	logs, sub, err := _Slash.contract.FilterLogs(opts, "crashResponse")
	if err != nil {
		return nil, err
	}
	return &SlashCrashResponseIterator{contract: _Slash.contract, event: "crashResponse", logs: logs, sub: sub}, nil
}

// WatchCrashResponse is a free log subscription operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_Slash *SlashFilterer) WatchCrashResponse(opts *bind.WatchOpts, sink chan<- *SlashCrashResponse) (event.Subscription, error) {

	logs, sub, err := _Slash.contract.WatchLogs(opts, "crashResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashCrashResponse)
				if err := _Slash.contract.UnpackLog(event, "crashResponse", log); err != nil {
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

// ParseCrashResponse is a log parse operation binding the contract event 0x07db600eebe2ac176be8dcebad61858c245a4961bb32ca2aa3d159b09aa0810e.
//
// Solidity: event crashResponse()
func (_Slash *SlashFilterer) ParseCrashResponse(log types.Log) (*SlashCrashResponse, error) {
	event := new(SlashCrashResponse)
	if err := _Slash.contract.UnpackLog(event, "crashResponse", log); err != nil {
		return nil, err
	}
	return event, nil
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

// SlashKnownResponseIterator is returned from FilterKnownResponse and is used to iterate over the raw logs and unpacked data for KnownResponse events raised by the Slash contract.
type SlashKnownResponseIterator struct {
	Event *SlashKnownResponse // Event containing the contract specifics and raw log

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
func (it *SlashKnownResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashKnownResponse)
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
		it.Event = new(SlashKnownResponse)
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
func (it *SlashKnownResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashKnownResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashKnownResponse represents a KnownResponse event raised by the Slash contract.
type SlashKnownResponse struct {
	Code uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKnownResponse is a free log retrieval operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_Slash *SlashFilterer) FilterKnownResponse(opts *bind.FilterOpts) (*SlashKnownResponseIterator, error) {

	logs, sub, err := _Slash.contract.FilterLogs(opts, "knownResponse")
	if err != nil {
		return nil, err
	}
	return &SlashKnownResponseIterator{contract: _Slash.contract, event: "knownResponse", logs: logs, sub: sub}, nil
}

// WatchKnownResponse is a free log subscription operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_Slash *SlashFilterer) WatchKnownResponse(opts *bind.WatchOpts, sink chan<- *SlashKnownResponse) (event.Subscription, error) {

	logs, sub, err := _Slash.contract.WatchLogs(opts, "knownResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashKnownResponse)
				if err := _Slash.contract.UnpackLog(event, "knownResponse", log); err != nil {
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

// ParseKnownResponse is a log parse operation binding the contract event 0x7f0956d47419b9525356e7111652b653b530ec6f5096dccc04589bc38e629967.
//
// Solidity: event knownResponse(uint32 code)
func (_Slash *SlashFilterer) ParseKnownResponse(log types.Log) (*SlashKnownResponse, error) {
	event := new(SlashKnownResponse)
	if err := _Slash.contract.UnpackLog(event, "knownResponse", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SlashParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Slash contract.
type SlashParamChangeIterator struct {
	Event *SlashParamChange // Event containing the contract specifics and raw log

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
func (it *SlashParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashParamChange)
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
		it.Event = new(SlashParamChange)
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
func (it *SlashParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashParamChange represents a ParamChange event raised by the Slash contract.
type SlashParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Slash *SlashFilterer) FilterParamChange(opts *bind.FilterOpts) (*SlashParamChangeIterator, error) {

	logs, sub, err := _Slash.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &SlashParamChangeIterator{contract: _Slash.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Slash *SlashFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *SlashParamChange) (event.Subscription, error) {

	logs, sub, err := _Slash.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashParamChange)
				if err := _Slash.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Slash *SlashFilterer) ParseParamChange(log types.Log) (*SlashParamChange, error) {
	event := new(SlashParamChange)
	if err := _Slash.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SlashUnKnownResponseIterator is returned from FilterUnKnownResponse and is used to iterate over the raw logs and unpacked data for UnKnownResponse events raised by the Slash contract.
type SlashUnKnownResponseIterator struct {
	Event *SlashUnKnownResponse // Event containing the contract specifics and raw log

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
func (it *SlashUnKnownResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashUnKnownResponse)
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
		it.Event = new(SlashUnKnownResponse)
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
func (it *SlashUnKnownResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashUnKnownResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashUnKnownResponse represents a UnKnownResponse event raised by the Slash contract.
type SlashUnKnownResponse struct {
	Code uint32
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnKnownResponse is a free log retrieval operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_Slash *SlashFilterer) FilterUnKnownResponse(opts *bind.FilterOpts) (*SlashUnKnownResponseIterator, error) {

	logs, sub, err := _Slash.contract.FilterLogs(opts, "unKnownResponse")
	if err != nil {
		return nil, err
	}
	return &SlashUnKnownResponseIterator{contract: _Slash.contract, event: "unKnownResponse", logs: logs, sub: sub}, nil
}

// WatchUnKnownResponse is a free log subscription operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_Slash *SlashFilterer) WatchUnKnownResponse(opts *bind.WatchOpts, sink chan<- *SlashUnKnownResponse) (event.Subscription, error) {

	logs, sub, err := _Slash.contract.WatchLogs(opts, "unKnownResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashUnKnownResponse)
				if err := _Slash.contract.UnpackLog(event, "unKnownResponse", log); err != nil {
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

// ParseUnKnownResponse is a log parse operation binding the contract event 0x7d45f62d17443dd4547bca8a8112c60e2385669318dc300ec61a5d2492f262e7.
//
// Solidity: event unKnownResponse(uint32 code)
func (_Slash *SlashFilterer) ParseUnKnownResponse(log types.Log) (*SlashUnKnownResponse, error) {
	event := new(SlashUnKnownResponse)
	if err := _Slash.contract.UnpackLog(event, "unKnownResponse", log); err != nil {
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
