// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chef

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

// ChefABI is the input ABI used to generate the binding from.
const ChefABI = "[{\"inputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"_cake\",\"type\":\"address\"},{\"internalType\":\"contractSyrupBar\",\"name\":\"_syrup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cakePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cake\",\"outputs\":[{\"internalType\":\"contractCakeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enterStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"leaveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syrup\",\"outputs\":[{\"internalType\":\"contractSyrupBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"multiplierNumber\",\"type\":\"uint256\"}],\"name\":\"updateMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Chef is an auto generated Go binding around an Ethereum contract.
type Chef struct {
	ChefCaller     // Read-only binding to the contract
	ChefTransactor // Write-only binding to the contract
	ChefFilterer   // Log filterer for contract events
}

// ChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChefSession struct {
	Contract     *Chef             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChefCallerSession struct {
	Contract *ChefCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChefTransactorSession struct {
	Contract     *ChefTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChefRaw struct {
	Contract *Chef // Generic contract binding to access the raw methods on
}

// ChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChefCallerRaw struct {
	Contract *ChefCaller // Generic read-only contract binding to access the raw methods on
}

// ChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChefTransactorRaw struct {
	Contract *ChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChef creates a new instance of Chef, bound to a specific deployed contract.
func NewChef(address common.Address, backend bind.ContractBackend) (*Chef, error) {
	contract, err := bindChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chef{ChefCaller: ChefCaller{contract: contract}, ChefTransactor: ChefTransactor{contract: contract}, ChefFilterer: ChefFilterer{contract: contract}}, nil
}

// NewChefCaller creates a new read-only instance of Chef, bound to a specific deployed contract.
func NewChefCaller(address common.Address, caller bind.ContractCaller) (*ChefCaller, error) {
	contract, err := bindChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChefCaller{contract: contract}, nil
}

// NewChefTransactor creates a new write-only instance of Chef, bound to a specific deployed contract.
func NewChefTransactor(address common.Address, transactor bind.ContractTransactor) (*ChefTransactor, error) {
	contract, err := bindChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChefTransactor{contract: contract}, nil
}

// NewChefFilterer creates a new log filterer instance of Chef, bound to a specific deployed contract.
func NewChefFilterer(address common.Address, filterer bind.ContractFilterer) (*ChefFilterer, error) {
	contract, err := bindChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChefFilterer{contract: contract}, nil
}

// bindChef binds a generic wrapper to an already deployed contract.
func bindChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chef *ChefRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Chef.Contract.ChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chef *ChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chef.Contract.ChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chef *ChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chef.Contract.ChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chef *ChefCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Chef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chef *ChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chef *ChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chef.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() constant returns(uint256)
func (_Chef *ChefCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "BONUS_MULTIPLIER")
	return *ret0, err
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() constant returns(uint256)
func (_Chef *ChefSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Chef.Contract.BONUSMULTIPLIER(&_Chef.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() constant returns(uint256)
func (_Chef *ChefCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Chef.Contract.BONUSMULTIPLIER(&_Chef.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() constant returns(address)
func (_Chef *ChefCaller) Cake(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "cake")
	return *ret0, err
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() constant returns(address)
func (_Chef *ChefSession) Cake() (common.Address, error) {
	return _Chef.Contract.Cake(&_Chef.CallOpts)
}

// Cake is a free data retrieval call binding the contract method 0xdce17484.
//
// Solidity: function cake() constant returns(address)
func (_Chef *ChefCallerSession) Cake() (common.Address, error) {
	return _Chef.Contract.Cake(&_Chef.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() constant returns(uint256)
func (_Chef *ChefCaller) CakePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "cakePerBlock")
	return *ret0, err
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() constant returns(uint256)
func (_Chef *ChefSession) CakePerBlock() (*big.Int, error) {
	return _Chef.Contract.CakePerBlock(&_Chef.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x0755e0b6.
//
// Solidity: function cakePerBlock() constant returns(uint256)
func (_Chef *ChefCallerSession) CakePerBlock() (*big.Int, error) {
	return _Chef.Contract.CakePerBlock(&_Chef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() constant returns(address)
func (_Chef *ChefCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "devaddr")
	return *ret0, err
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() constant returns(address)
func (_Chef *ChefSession) Devaddr() (common.Address, error) {
	return _Chef.Contract.Devaddr(&_Chef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() constant returns(address)
func (_Chef *ChefCallerSession) Devaddr() (common.Address, error) {
	return _Chef.Contract.Devaddr(&_Chef.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Chef *ChefCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "getMultiplier", _from, _to)
	return *ret0, err
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Chef *ChefSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Chef.Contract.GetMultiplier(&_Chef.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) constant returns(uint256)
func (_Chef *ChefCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Chef.Contract.GetMultiplier(&_Chef.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() constant returns(address)
func (_Chef *ChefCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "migrator")
	return *ret0, err
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() constant returns(address)
func (_Chef *ChefSession) Migrator() (common.Address, error) {
	return _Chef.Contract.Migrator(&_Chef.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() constant returns(address)
func (_Chef *ChefCallerSession) Migrator() (common.Address, error) {
	return _Chef.Contract.Migrator(&_Chef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Chef *ChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Chef *ChefSession) Owner() (common.Address, error) {
	return _Chef.Contract.Owner(&_Chef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Chef *ChefCallerSession) Owner() (common.Address, error) {
	return _Chef.Contract.Owner(&_Chef.CallOpts)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) constant returns(uint256)
func (_Chef *ChefCaller) PendingCake(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "pendingCake", _pid, _user)
	return *ret0, err
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) constant returns(uint256)
func (_Chef *ChefSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Chef.Contract.PendingCake(&_Chef.CallOpts, _pid, _user)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) constant returns(uint256)
func (_Chef *ChefCallerSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Chef.Contract.PendingCake(&_Chef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Chef *ChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	ret := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccCakePerShare *big.Int
	})
	out := ret
	err := _Chef.contract.Call(opts, out, "poolInfo", arg0)
	return *ret, err
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Chef *ChefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Chef.Contract.PoolInfo(&_Chef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) constant returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accCakePerShare)
func (_Chef *ChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccCakePerShare *big.Int
}, error) {
	return _Chef.Contract.PoolInfo(&_Chef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() constant returns(uint256)
func (_Chef *ChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "poolLength")
	return *ret0, err
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() constant returns(uint256)
func (_Chef *ChefSession) PoolLength() (*big.Int, error) {
	return _Chef.Contract.PoolLength(&_Chef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() constant returns(uint256)
func (_Chef *ChefCallerSession) PoolLength() (*big.Int, error) {
	return _Chef.Contract.PoolLength(&_Chef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Chef *ChefCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "startBlock")
	return *ret0, err
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Chef *ChefSession) StartBlock() (*big.Int, error) {
	return _Chef.Contract.StartBlock(&_Chef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Chef *ChefCallerSession) StartBlock() (*big.Int, error) {
	return _Chef.Contract.StartBlock(&_Chef.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Chef *ChefCaller) Syrup(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "syrup")
	return *ret0, err
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Chef *ChefSession) Syrup() (common.Address, error) {
	return _Chef.Contract.Syrup(&_Chef.CallOpts)
}

// Syrup is a free data retrieval call binding the contract method 0x86a952c4.
//
// Solidity: function syrup() constant returns(address)
func (_Chef *ChefCallerSession) Syrup() (common.Address, error) {
	return _Chef.Contract.Syrup(&_Chef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() constant returns(uint256)
func (_Chef *ChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Chef.contract.Call(opts, out, "totalAllocPoint")
	return *ret0, err
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() constant returns(uint256)
func (_Chef *ChefSession) TotalAllocPoint() (*big.Int, error) {
	return _Chef.Contract.TotalAllocPoint(&_Chef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() constant returns(uint256)
func (_Chef *ChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Chef.Contract.TotalAllocPoint(&_Chef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Chef *ChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	ret := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	out := ret
	err := _Chef.contract.Call(opts, out, "userInfo", arg0, arg1)
	return *ret, err
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Chef *ChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Chef.Contract.UserInfo(&_Chef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) constant returns(uint256 amount, uint256 rewardDebt)
func (_Chef *ChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Chef.Contract.UserInfo(&_Chef.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Chef *ChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Chef *ChefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.Contract.Add(&_Chef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Chef *ChefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.Contract.Add(&_Chef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Deposit(&_Chef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Deposit(&_Chef.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Chef *ChefTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Chef *ChefSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Chef.Contract.Dev(&_Chef.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Chef *ChefTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Chef.Contract.Dev(&_Chef.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Chef *ChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Chef *ChefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.EmergencyWithdraw(&_Chef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Chef *ChefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.EmergencyWithdraw(&_Chef.TransactOpts, _pid)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Chef *ChefTransactor) EnterStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "enterStaking", _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Chef *ChefSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.EnterStaking(&_Chef.TransactOpts, _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Chef *ChefTransactorSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.EnterStaking(&_Chef.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Chef *ChefTransactor) LeaveStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "leaveStaking", _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Chef *ChefSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.LeaveStaking(&_Chef.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Chef *ChefTransactorSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.LeaveStaking(&_Chef.TransactOpts, _amount)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Chef *ChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Chef *ChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _Chef.Contract.MassUpdatePools(&_Chef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Chef *ChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Chef.Contract.MassUpdatePools(&_Chef.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Chef *ChefTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Chef *ChefSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Migrate(&_Chef.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Chef *ChefTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Migrate(&_Chef.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chef *ChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chef *ChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chef.Contract.RenounceOwnership(&_Chef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chef *ChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chef.Contract.RenounceOwnership(&_Chef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Chef *ChefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Chef *ChefSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.Contract.Set(&_Chef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Chef *ChefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Chef.Contract.Set(&_Chef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Chef *ChefTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Chef *ChefSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Chef.Contract.SetMigrator(&_Chef.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Chef *ChefTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Chef.Contract.SetMigrator(&_Chef.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chef *ChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chef *ChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chef.Contract.TransferOwnership(&_Chef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chef *ChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chef.Contract.TransferOwnership(&_Chef.TransactOpts, newOwner)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Chef *ChefTransactor) UpdateMultiplier(opts *bind.TransactOpts, multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "updateMultiplier", multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Chef *ChefSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.UpdateMultiplier(&_Chef.TransactOpts, multiplierNumber)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 multiplierNumber) returns()
func (_Chef *ChefTransactorSession) UpdateMultiplier(multiplierNumber *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.UpdateMultiplier(&_Chef.TransactOpts, multiplierNumber)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Chef *ChefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Chef *ChefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.UpdatePool(&_Chef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Chef *ChefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.UpdatePool(&_Chef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Withdraw(&_Chef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Chef *ChefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Chef.Contract.Withdraw(&_Chef.TransactOpts, _pid, _amount)
}

// ChefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Chef contract.
type ChefDepositIterator struct {
	Event *ChefDeposit // Event containing the contract specifics and raw log

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
func (it *ChefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChefDeposit)
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
		it.Event = new(ChefDeposit)
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
func (it *ChefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChefDeposit represents a Deposit event raised by the Chef contract.
type ChefDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ChefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ChefDepositIterator{contract: _Chef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ChefDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChefDeposit)
				if err := _Chef.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) ParseDeposit(log types.Log) (*ChefDeposit, error) {
	event := new(ChefDeposit)
	if err := _Chef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Chef contract.
type ChefEmergencyWithdrawIterator struct {
	Event *ChefEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *ChefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChefEmergencyWithdraw)
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
		it.Event = new(ChefEmergencyWithdraw)
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
func (it *ChefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChefEmergencyWithdraw represents a EmergencyWithdraw event raised by the Chef contract.
type ChefEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ChefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ChefEmergencyWithdrawIterator{contract: _Chef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *ChefEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChefEmergencyWithdraw)
				if err := _Chef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) ParseEmergencyWithdraw(log types.Log) (*ChefEmergencyWithdraw, error) {
	event := new(ChefEmergencyWithdraw)
	if err := _Chef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Chef contract.
type ChefOwnershipTransferredIterator struct {
	Event *ChefOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChefOwnershipTransferred)
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
		it.Event = new(ChefOwnershipTransferred)
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
func (it *ChefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChefOwnershipTransferred represents a OwnershipTransferred event raised by the Chef contract.
type ChefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chef *ChefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChefOwnershipTransferredIterator{contract: _Chef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chef *ChefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChefOwnershipTransferred)
				if err := _Chef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chef *ChefFilterer) ParseOwnershipTransferred(log types.Log) (*ChefOwnershipTransferred, error) {
	event := new(ChefOwnershipTransferred)
	if err := _Chef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Chef contract.
type ChefWithdrawIterator struct {
	Event *ChefWithdraw // Event containing the contract specifics and raw log

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
func (it *ChefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChefWithdraw)
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
		it.Event = new(ChefWithdraw)
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
func (it *ChefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChefWithdraw represents a Withdraw event raised by the Chef contract.
type ChefWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ChefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ChefWithdrawIterator{contract: _Chef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ChefWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Chef.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChefWithdraw)
				if err := _Chef.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Chef *ChefFilterer) ParseWithdraw(log types.Log) (*ChefWithdraw, error) {
	event := new(ChefWithdraw)
	if err := _Chef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
