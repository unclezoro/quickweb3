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

// CrossABI is the input ABI used to generate the binding from.
const CrossABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"addChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oracleSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"crossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"unexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"unexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"unsupportedPackage\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STORE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"batchSizeForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"handlePackage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isRelayRewardFromSystemReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleSequence\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"packageCounterForEachHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredContractMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Cross is an auto generated Go binding around an Ethereum contract.
type Cross struct {
	CrossCaller     // Read-only binding to the contract
	CrossTransactor // Write-only binding to the contract
	CrossFilterer   // Log filterer for contract events
}

// CrossCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossSession struct {
	Contract     *Cross            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossCallerSession struct {
	Contract *CrossCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrossTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossTransactorSession struct {
	Contract     *CrossTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossRaw struct {
	Contract *Cross // Generic contract binding to access the raw methods on
}

// CrossCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossCallerRaw struct {
	Contract *CrossCaller // Generic read-only contract binding to access the raw methods on
}

// CrossTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossTransactorRaw struct {
	Contract *CrossTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCross creates a new instance of Cross, bound to a specific deployed contract.
func NewCross(address common.Address, backend bind.ContractBackend) (*Cross, error) {
	contract, err := bindCross(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cross{CrossCaller: CrossCaller{contract: contract}, CrossTransactor: CrossTransactor{contract: contract}, CrossFilterer: CrossFilterer{contract: contract}}, nil
}

// NewCrossCaller creates a new read-only instance of Cross, bound to a specific deployed contract.
func NewCrossCaller(address common.Address, caller bind.ContractCaller) (*CrossCaller, error) {
	contract, err := bindCross(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossCaller{contract: contract}, nil
}

// NewCrossTransactor creates a new write-only instance of Cross, bound to a specific deployed contract.
func NewCrossTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossTransactor, error) {
	contract, err := bindCross(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossTransactor{contract: contract}, nil
}

// NewCrossFilterer creates a new log filterer instance of Cross, bound to a specific deployed contract.
func NewCrossFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossFilterer, error) {
	contract, err := bindCross(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossFilterer{contract: contract}, nil
}

// bindCross binds a generic wrapper to an already deployed contract.
func bindCross(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cross *CrossRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cross.Contract.CrossCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cross *CrossRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cross.Contract.CrossTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cross *CrossRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cross.Contract.CrossTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cross *CrossCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cross.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cross *CrossTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cross.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cross *CrossTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cross.Contract.contract.Transact(opts, method, params...)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossCaller) ACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "ACK_PACKAGE")
	return *ret0, err
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossSession) ACKPACKAGE() (uint8, error) {
	return _Cross.Contract.ACKPACKAGE(&_Cross.CallOpts)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossCallerSession) ACKPACKAGE() (uint8, error) {
	return _Cross.Contract.ACKPACKAGE(&_Cross.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) BINDCHANNELID() (uint8, error) {
	return _Cross.Contract.BINDCHANNELID(&_Cross.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) BINDCHANNELID() (uint8, error) {
	return _Cross.Contract.BINDCHANNELID(&_Cross.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Cross *CrossCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Cross *CrossSession) CODEOK() (uint32, error) {
	return _Cross.Contract.CODEOK(&_Cross.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Cross *CrossCallerSession) CODEOK() (uint32, error) {
	return _Cross.Contract.CODEOK(&_Cross.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.CROSSCHAINCONTRACTADDR(&_Cross.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.CROSSCHAINCONTRACTADDR(&_Cross.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Cross *CrossCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Cross *CrossSession) ERRORFAILDECODE() (uint32, error) {
	return _Cross.Contract.ERRORFAILDECODE(&_Cross.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Cross *CrossCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Cross.Contract.ERRORFAILDECODE(&_Cross.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossCaller) FAILACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "FAIL_ACK_PACKAGE")
	return *ret0, err
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossSession) FAILACKPACKAGE() (uint8, error) {
	return _Cross.Contract.FAILACKPACKAGE(&_Cross.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Cross *CrossCallerSession) FAILACKPACKAGE() (uint8, error) {
	return _Cross.Contract.FAILACKPACKAGE(&_Cross.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) GOVCHANNELID() (uint8, error) {
	return _Cross.Contract.GOVCHANNELID(&_Cross.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) GOVCHANNELID() (uint8, error) {
	return _Cross.Contract.GOVCHANNELID(&_Cross.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Cross *CrossCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Cross *CrossSession) GOVHUBADDR() (common.Address, error) {
	return _Cross.Contract.GOVHUBADDR(&_Cross.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Cross.Contract.GOVHUBADDR(&_Cross.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Cross *CrossCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Cross *CrossSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Cross.Contract.INCENTIVIZEADDR(&_Cross.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Cross.Contract.INCENTIVIZEADDR(&_Cross.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Cross *CrossCaller) INITBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "INIT_BATCH_SIZE")
	return *ret0, err
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Cross *CrossSession) INITBATCHSIZE() (*big.Int, error) {
	return _Cross.Contract.INITBATCHSIZE(&_Cross.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Cross *CrossCallerSession) INITBATCHSIZE() (*big.Int, error) {
	return _Cross.Contract.INITBATCHSIZE(&_Cross.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Cross *CrossCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Cross *CrossSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Cross.Contract.LIGHTCLIENTADDR(&_Cross.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Cross.Contract.LIGHTCLIENTADDR(&_Cross.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.RELAYERHUBCONTRACTADDR(&_Cross.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.RELAYERHUBCONTRACTADDR(&_Cross.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) SLASHCHANNELID() (uint8, error) {
	return _Cross.Contract.SLASHCHANNELID(&_Cross.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Cross.Contract.SLASHCHANNELID(&_Cross.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.SLASHCONTRACTADDR(&_Cross.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.SLASHCONTRACTADDR(&_Cross.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) STAKINGCHANNELID() (uint8, error) {
	return _Cross.Contract.STAKINGCHANNELID(&_Cross.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Cross.Contract.STAKINGCHANNELID(&_Cross.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Cross *CrossCaller) STORENAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "STORE_NAME")
	return *ret0, err
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Cross *CrossSession) STORENAME() (string, error) {
	return _Cross.Contract.STORENAME(&_Cross.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Cross *CrossCallerSession) STORENAME() (string, error) {
	return _Cross.Contract.STORENAME(&_Cross.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Cross *CrossCaller) SYNPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "SYN_PACKAGE")
	return *ret0, err
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Cross *CrossSession) SYNPACKAGE() (uint8, error) {
	return _Cross.Contract.SYNPACKAGE(&_Cross.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Cross *CrossCallerSession) SYNPACKAGE() (uint8, error) {
	return _Cross.Contract.SYNPACKAGE(&_Cross.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Cross *CrossCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Cross *CrossSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Cross.Contract.SYSTEMREWARDADDR(&_Cross.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Cross.Contract.SYSTEMREWARDADDR(&_Cross.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Cross *CrossCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Cross *CrossSession) TOKENHUBADDR() (common.Address, error) {
	return _Cross.Contract.TOKENHUBADDR(&_Cross.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Cross.Contract.TOKENHUBADDR(&_Cross.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Cross.Contract.TRANSFERINCHANNELID(&_Cross.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Cross.Contract.TRANSFERINCHANNELID(&_Cross.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Cross *CrossCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Cross *CrossSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Cross.Contract.TRANSFEROUTCHANNELID(&_Cross.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Cross *CrossCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Cross.Contract.TRANSFEROUTCHANNELID(&_Cross.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.VALIDATORCONTRACTADDR(&_Cross.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Cross *CrossCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Cross.Contract.VALIDATORCONTRACTADDR(&_Cross.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Cross *CrossCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Cross *CrossSession) AlreadyInit() (bool, error) {
	return _Cross.Contract.AlreadyInit(&_Cross.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Cross *CrossCallerSession) AlreadyInit() (bool, error) {
	return _Cross.Contract.AlreadyInit(&_Cross.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Cross *CrossCaller) BatchSizeForOracle(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "batchSizeForOracle")
	return *ret0, err
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Cross *CrossSession) BatchSizeForOracle() (*big.Int, error) {
	return _Cross.Contract.BatchSizeForOracle(&_Cross.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Cross *CrossCallerSession) BatchSizeForOracle() (*big.Int, error) {
	return _Cross.Contract.BatchSizeForOracle(&_Cross.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Cross *CrossCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Cross *CrossSession) BscChainID() (uint16, error) {
	return _Cross.Contract.BscChainID(&_Cross.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Cross *CrossCallerSession) BscChainID() (uint16, error) {
	return _Cross.Contract.BscChainID(&_Cross.CallOpts)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Cross *CrossCaller) ChannelHandlerContractMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "channelHandlerContractMap", arg0)
	return *ret0, err
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Cross *CrossSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Cross.Contract.ChannelHandlerContractMap(&_Cross.CallOpts, arg0)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Cross *CrossCallerSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Cross.Contract.ChannelHandlerContractMap(&_Cross.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossCaller) ChannelReceiveSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "channelReceiveSequenceMap", arg0)
	return *ret0, err
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Cross.Contract.ChannelReceiveSequenceMap(&_Cross.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossCallerSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Cross.Contract.ChannelReceiveSequenceMap(&_Cross.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossCaller) ChannelSendSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "channelSendSequenceMap", arg0)
	return *ret0, err
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Cross.Contract.ChannelSendSequenceMap(&_Cross.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Cross *CrossCallerSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Cross.Contract.ChannelSendSequenceMap(&_Cross.CallOpts, arg0)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Cross *CrossCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "encodePayload", packageType, relayFee, msgBytes)
	return *ret0, err
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Cross *CrossSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Cross.Contract.EncodePayload(&_Cross.CallOpts, packageType, relayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Cross *CrossCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Cross.Contract.EncodePayload(&_Cross.CallOpts, packageType, relayFee, msgBytes)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Cross *CrossCaller) IsRelayRewardFromSystemReward(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "isRelayRewardFromSystemReward", arg0)
	return *ret0, err
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Cross *CrossSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Cross.Contract.IsRelayRewardFromSystemReward(&_Cross.CallOpts, arg0)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Cross *CrossCallerSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Cross.Contract.IsRelayRewardFromSystemReward(&_Cross.CallOpts, arg0)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(uint64)
func (_Cross *CrossCaller) OracleSequence(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "oracleSequence")
	return *ret0, err
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(uint64)
func (_Cross *CrossSession) OracleSequence() (uint64, error) {
	return _Cross.Contract.OracleSequence(&_Cross.CallOpts)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(uint64)
func (_Cross *CrossCallerSession) OracleSequence() (uint64, error) {
	return _Cross.Contract.OracleSequence(&_Cross.CallOpts)
}

// PackageCounterForEachHeight is a free data retrieval call binding the contract method 0x37c9b8ba.
//
// Solidity: function packageCounterForEachHeight(uint256 ) constant returns(uint256)
func (_Cross *CrossCaller) PackageCounterForEachHeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "packageCounterForEachHeight", arg0)
	return *ret0, err
}

// PackageCounterForEachHeight is a free data retrieval call binding the contract method 0x37c9b8ba.
//
// Solidity: function packageCounterForEachHeight(uint256 ) constant returns(uint256)
func (_Cross *CrossSession) PackageCounterForEachHeight(arg0 *big.Int) (*big.Int, error) {
	return _Cross.Contract.PackageCounterForEachHeight(&_Cross.CallOpts, arg0)
}

// PackageCounterForEachHeight is a free data retrieval call binding the contract method 0x37c9b8ba.
//
// Solidity: function packageCounterForEachHeight(uint256 ) constant returns(uint256)
func (_Cross *CrossCallerSession) PackageCounterForEachHeight(arg0 *big.Int) (*big.Int, error) {
	return _Cross.Contract.PackageCounterForEachHeight(&_Cross.CallOpts, arg0)
}

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Cross *CrossCaller) RegisteredContractMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cross.contract.Call(opts, out, "registeredContractMap", arg0)
	return *ret0, err
}

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Cross *CrossSession) RegisteredContractMap(arg0 common.Address) (bool, error) {
	return _Cross.Contract.RegisteredContractMap(&_Cross.CallOpts, arg0)
}

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Cross *CrossCallerSession) RegisteredContractMap(arg0 common.Address) (bool, error) {
	return _Cross.Contract.RegisteredContractMap(&_Cross.CallOpts, arg0)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Cross *CrossTransactor) HandlePackage(opts *bind.TransactOpts, payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "handlePackage", payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Cross *CrossSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Cross.Contract.HandlePackage(&_Cross.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Cross *CrossTransactorSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Cross.Contract.HandlePackage(&_Cross.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Cross *CrossTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Cross *CrossSession) Init() (*types.Transaction, error) {
	return _Cross.Contract.Init(&_Cross.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Cross *CrossTransactorSession) Init() (*types.Transaction, error) {
	return _Cross.Contract.Init(&_Cross.TransactOpts)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns(bool)
func (_Cross *CrossTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns(bool)
func (_Cross *CrossSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.SendSynPackage(&_Cross.TransactOpts, channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns(bool)
func (_Cross *CrossTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Cross.Contract.SendSynPackage(&_Cross.TransactOpts, channelId, msgBytes, relayFee)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Cross *CrossTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Cross.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Cross *CrossSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Cross.Contract.UpdateParam(&_Cross.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Cross *CrossTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Cross.Contract.UpdateParam(&_Cross.TransactOpts, key, value)
}

// CrossAddChannelIterator is returned from FilterAddChannel and is used to iterate over the raw logs and unpacked data for AddChannel events raised by the Cross contract.
type CrossAddChannelIterator struct {
	Event *CrossAddChannel // Event containing the contract specifics and raw log

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
func (it *CrossAddChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossAddChannel)
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
		it.Event = new(CrossAddChannel)
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
func (it *CrossAddChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossAddChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossAddChannel represents a AddChannel event raised by the Cross contract.
type CrossAddChannel struct {
	ChannelId    uint8
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddChannel is a free log retrieval operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Cross *CrossFilterer) FilterAddChannel(opts *bind.FilterOpts, channelId []uint8, contractAddr []common.Address) (*CrossAddChannelIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "addChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrossAddChannelIterator{contract: _Cross.contract, event: "addChannel", logs: logs, sub: sub}, nil
}

// WatchAddChannel is a free log subscription operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Cross *CrossFilterer) WatchAddChannel(opts *bind.WatchOpts, sink chan<- *CrossAddChannel, channelId []uint8, contractAddr []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "addChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossAddChannel)
				if err := _Cross.contract.UnpackLog(event, "addChannel", log); err != nil {
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

// ParseAddChannel is a log parse operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Cross *CrossFilterer) ParseAddChannel(log types.Log) (*CrossAddChannel, error) {
	event := new(CrossAddChannel)
	if err := _Cross.contract.UnpackLog(event, "addChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossCrossChainPackageIterator is returned from FilterCrossChainPackage and is used to iterate over the raw logs and unpacked data for CrossChainPackage events raised by the Cross contract.
type CrossCrossChainPackageIterator struct {
	Event *CrossCrossChainPackage // Event containing the contract specifics and raw log

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
func (it *CrossCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossCrossChainPackage)
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
		it.Event = new(CrossCrossChainPackage)
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
func (it *CrossCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossCrossChainPackage represents a CrossChainPackage event raised by the Cross contract.
type CrossCrossChainPackage struct {
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (*CrossCrossChainPackageIterator, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "crossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossCrossChainPackageIterator{contract: _Cross.contract, event: "crossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrossCrossChainPackage, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "crossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossCrossChainPackage)
				if err := _Cross.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
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

// ParseCrossChainPackage is a log parse operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) ParseCrossChainPackage(log types.Log) (*CrossCrossChainPackage, error) {
	event := new(CrossCrossChainPackage)
	if err := _Cross.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Cross contract.
type CrossParamChangeIterator struct {
	Event *CrossParamChange // Event containing the contract specifics and raw log

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
func (it *CrossParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossParamChange)
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
		it.Event = new(CrossParamChange)
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
func (it *CrossParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossParamChange represents a ParamChange event raised by the Cross contract.
type CrossParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Cross *CrossFilterer) FilterParamChange(opts *bind.FilterOpts) (*CrossParamChangeIterator, error) {

	logs, sub, err := _Cross.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &CrossParamChangeIterator{contract: _Cross.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Cross *CrossFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *CrossParamChange) (event.Subscription, error) {

	logs, sub, err := _Cross.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossParamChange)
				if err := _Cross.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Cross *CrossFilterer) ParseParamChange(log types.Log) (*CrossParamChange, error) {
	event := new(CrossParamChange)
	if err := _Cross.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossUnexpectedFailureAssertionInPackageHandlerIterator is returned from FilterUnexpectedFailureAssertionInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedFailureAssertionInPackageHandler events raised by the Cross contract.
type CrossUnexpectedFailureAssertionInPackageHandlerIterator struct {
	Event *CrossUnexpectedFailureAssertionInPackageHandler // Event containing the contract specifics and raw log

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
func (it *CrossUnexpectedFailureAssertionInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossUnexpectedFailureAssertionInPackageHandler)
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
		it.Event = new(CrossUnexpectedFailureAssertionInPackageHandler)
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
func (it *CrossUnexpectedFailureAssertionInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossUnexpectedFailureAssertionInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossUnexpectedFailureAssertionInPackageHandler represents a UnexpectedFailureAssertionInPackageHandler event raised by the Cross contract.
type CrossUnexpectedFailureAssertionInPackageHandler struct {
	ContractAddr common.Address
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedFailureAssertionInPackageHandler is a free log retrieval operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Cross *CrossFilterer) FilterUnexpectedFailureAssertionInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrossUnexpectedFailureAssertionInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "unexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrossUnexpectedFailureAssertionInPackageHandlerIterator{contract: _Cross.contract, event: "unexpectedFailureAssertionInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedFailureAssertionInPackageHandler is a free log subscription operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Cross *CrossFilterer) WatchUnexpectedFailureAssertionInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrossUnexpectedFailureAssertionInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "unexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossUnexpectedFailureAssertionInPackageHandler)
				if err := _Cross.contract.UnpackLog(event, "unexpectedFailureAssertionInPackageHandler", log); err != nil {
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

// ParseUnexpectedFailureAssertionInPackageHandler is a log parse operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Cross *CrossFilterer) ParseUnexpectedFailureAssertionInPackageHandler(log types.Log) (*CrossUnexpectedFailureAssertionInPackageHandler, error) {
	event := new(CrossUnexpectedFailureAssertionInPackageHandler)
	if err := _Cross.contract.UnpackLog(event, "unexpectedFailureAssertionInPackageHandler", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossUnexpectedRevertInPackageHandlerIterator is returned from FilterUnexpectedRevertInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedRevertInPackageHandler events raised by the Cross contract.
type CrossUnexpectedRevertInPackageHandlerIterator struct {
	Event *CrossUnexpectedRevertInPackageHandler // Event containing the contract specifics and raw log

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
func (it *CrossUnexpectedRevertInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossUnexpectedRevertInPackageHandler)
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
		it.Event = new(CrossUnexpectedRevertInPackageHandler)
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
func (it *CrossUnexpectedRevertInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossUnexpectedRevertInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossUnexpectedRevertInPackageHandler represents a UnexpectedRevertInPackageHandler event raised by the Cross contract.
type CrossUnexpectedRevertInPackageHandler struct {
	ContractAddr common.Address
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedRevertInPackageHandler is a free log retrieval operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Cross *CrossFilterer) FilterUnexpectedRevertInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrossUnexpectedRevertInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "unexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrossUnexpectedRevertInPackageHandlerIterator{contract: _Cross.contract, event: "unexpectedRevertInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedRevertInPackageHandler is a free log subscription operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Cross *CrossFilterer) WatchUnexpectedRevertInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrossUnexpectedRevertInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "unexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossUnexpectedRevertInPackageHandler)
				if err := _Cross.contract.UnpackLog(event, "unexpectedRevertInPackageHandler", log); err != nil {
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

// ParseUnexpectedRevertInPackageHandler is a log parse operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Cross *CrossFilterer) ParseUnexpectedRevertInPackageHandler(log types.Log) (*CrossUnexpectedRevertInPackageHandler, error) {
	event := new(CrossUnexpectedRevertInPackageHandler)
	if err := _Cross.contract.UnpackLog(event, "unexpectedRevertInPackageHandler", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrossUnsupportedPackageIterator is returned from FilterUnsupportedPackage and is used to iterate over the raw logs and unpacked data for UnsupportedPackage events raised by the Cross contract.
type CrossUnsupportedPackageIterator struct {
	Event *CrossUnsupportedPackage // Event containing the contract specifics and raw log

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
func (it *CrossUnsupportedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossUnsupportedPackage)
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
		it.Event = new(CrossUnsupportedPackage)
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
func (it *CrossUnsupportedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossUnsupportedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossUnsupportedPackage represents a UnsupportedPackage event raised by the Cross contract.
type CrossUnsupportedPackage struct {
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnsupportedPackage is a free log retrieval operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) FilterUnsupportedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrossUnsupportedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Cross.contract.FilterLogs(opts, "unsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrossUnsupportedPackageIterator{contract: _Cross.contract, event: "unsupportedPackage", logs: logs, sub: sub}, nil
}

// WatchUnsupportedPackage is a free log subscription operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) WatchUnsupportedPackage(opts *bind.WatchOpts, sink chan<- *CrossUnsupportedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Cross.contract.WatchLogs(opts, "unsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossUnsupportedPackage)
				if err := _Cross.contract.UnpackLog(event, "unsupportedPackage", log); err != nil {
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

// ParseUnsupportedPackage is a log parse operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Cross *CrossFilterer) ParseUnsupportedPackage(log types.Log) (*CrossUnsupportedPackage, error) {
	event := new(CrossUnsupportedPackage)
	if err := _Cross.contract.UnpackLog(event, "unsupportedPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}
