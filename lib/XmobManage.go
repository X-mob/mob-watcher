// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// XmobManageMetaData contains all meta data concerning the XmobManage contract.
var XmobManageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchangeCore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"DepositEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"feeRate\",\"type\":\"uint8\"}],\"name\":\"FeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseTarget\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takeProfitPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stopLossPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"targetMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"MobCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"XMobProxySet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_raiseTarget\",\"type\":\"uint256\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_raiseTarget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_takeProfitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stopLossPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_raiseDeadline\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_deadline\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_targetMode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"createMob\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeCore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mobsById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mobsTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_feeRate\",\"type\":\"uint8\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchangeCore\",\"type\":\"address\"}],\"name\":\"setXMobProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// XmobManageABI is the input ABI used to generate the binding from.
// Deprecated: Use XmobManageMetaData.ABI instead.
var XmobManageABI = XmobManageMetaData.ABI

// XmobManage is an auto generated Go binding around an Ethereum contract.
type XmobManage struct {
	XmobManageCaller     // Read-only binding to the contract
	XmobManageTransactor // Write-only binding to the contract
	XmobManageFilterer   // Log filterer for contract events
}

// XmobManageCaller is an auto generated read-only Go binding around an Ethereum contract.
type XmobManageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobManageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XmobManageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobManageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XmobManageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XmobManageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XmobManageSession struct {
	Contract     *XmobManage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XmobManageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XmobManageCallerSession struct {
	Contract *XmobManageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// XmobManageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XmobManageTransactorSession struct {
	Contract     *XmobManageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// XmobManageRaw is an auto generated low-level Go binding around an Ethereum contract.
type XmobManageRaw struct {
	Contract *XmobManage // Generic contract binding to access the raw methods on
}

// XmobManageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XmobManageCallerRaw struct {
	Contract *XmobManageCaller // Generic read-only contract binding to access the raw methods on
}

// XmobManageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XmobManageTransactorRaw struct {
	Contract *XmobManageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXmobManage creates a new instance of XmobManage, bound to a specific deployed contract.
func NewXmobManage(address common.Address, backend bind.ContractBackend) (*XmobManage, error) {
	contract, err := bindXmobManage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XmobManage{XmobManageCaller: XmobManageCaller{contract: contract}, XmobManageTransactor: XmobManageTransactor{contract: contract}, XmobManageFilterer: XmobManageFilterer{contract: contract}}, nil
}

// NewXmobManageCaller creates a new read-only instance of XmobManage, bound to a specific deployed contract.
func NewXmobManageCaller(address common.Address, caller bind.ContractCaller) (*XmobManageCaller, error) {
	contract, err := bindXmobManage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XmobManageCaller{contract: contract}, nil
}

// NewXmobManageTransactor creates a new write-only instance of XmobManage, bound to a specific deployed contract.
func NewXmobManageTransactor(address common.Address, transactor bind.ContractTransactor) (*XmobManageTransactor, error) {
	contract, err := bindXmobManage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XmobManageTransactor{contract: contract}, nil
}

// NewXmobManageFilterer creates a new log filterer instance of XmobManage, bound to a specific deployed contract.
func NewXmobManageFilterer(address common.Address, filterer bind.ContractFilterer) (*XmobManageFilterer, error) {
	contract, err := bindXmobManage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XmobManageFilterer{contract: contract}, nil
}

// bindXmobManage binds a generic wrapper to an already deployed contract.
func bindXmobManage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XmobManageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmobManage *XmobManageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XmobManage.Contract.XmobManageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmobManage *XmobManageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobManage.Contract.XmobManageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmobManage *XmobManageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmobManage.Contract.XmobManageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XmobManage *XmobManageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _XmobManage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XmobManage *XmobManageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobManage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XmobManage *XmobManageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XmobManage.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _raiseTarget) view returns(uint256 fee)
func (_XmobManage *XmobManageCaller) CalcFee(opts *bind.CallOpts, _raiseTarget *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "calcFee", _raiseTarget)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _raiseTarget) view returns(uint256 fee)
func (_XmobManage *XmobManageSession) CalcFee(_raiseTarget *big.Int) (*big.Int, error) {
	return _XmobManage.Contract.CalcFee(&_XmobManage.CallOpts, _raiseTarget)
}

// CalcFee is a free data retrieval call binding the contract method 0x75dc7d8c.
//
// Solidity: function calcFee(uint256 _raiseTarget) view returns(uint256 fee)
func (_XmobManage *XmobManageCallerSession) CalcFee(_raiseTarget *big.Int) (*big.Int, error) {
	return _XmobManage.Contract.CalcFee(&_XmobManage.CallOpts, _raiseTarget)
}

// ExchangeCore is a free data retrieval call binding the contract method 0x84a344b1.
//
// Solidity: function exchangeCore() view returns(address)
func (_XmobManage *XmobManageCaller) ExchangeCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "exchangeCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeCore is a free data retrieval call binding the contract method 0x84a344b1.
//
// Solidity: function exchangeCore() view returns(address)
func (_XmobManage *XmobManageSession) ExchangeCore() (common.Address, error) {
	return _XmobManage.Contract.ExchangeCore(&_XmobManage.CallOpts)
}

// ExchangeCore is a free data retrieval call binding the contract method 0x84a344b1.
//
// Solidity: function exchangeCore() view returns(address)
func (_XmobManage *XmobManageCallerSession) ExchangeCore() (common.Address, error) {
	return _XmobManage.Contract.ExchangeCore(&_XmobManage.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint8)
func (_XmobManage *XmobManageCaller) FeeRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint8)
func (_XmobManage *XmobManageSession) FeeRate() (uint8, error) {
	return _XmobManage.Contract.FeeRate(&_XmobManage.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint8)
func (_XmobManage *XmobManageCallerSession) FeeRate() (uint8, error) {
	return _XmobManage.Contract.FeeRate(&_XmobManage.CallOpts)
}

// MobsById is a free data retrieval call binding the contract method 0xdd790e6f.
//
// Solidity: function mobsById(uint256 ) view returns(address)
func (_XmobManage *XmobManageCaller) MobsById(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "mobsById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MobsById is a free data retrieval call binding the contract method 0xdd790e6f.
//
// Solidity: function mobsById(uint256 ) view returns(address)
func (_XmobManage *XmobManageSession) MobsById(arg0 *big.Int) (common.Address, error) {
	return _XmobManage.Contract.MobsById(&_XmobManage.CallOpts, arg0)
}

// MobsById is a free data retrieval call binding the contract method 0xdd790e6f.
//
// Solidity: function mobsById(uint256 ) view returns(address)
func (_XmobManage *XmobManageCallerSession) MobsById(arg0 *big.Int) (common.Address, error) {
	return _XmobManage.Contract.MobsById(&_XmobManage.CallOpts, arg0)
}

// MobsTotal is a free data retrieval call binding the contract method 0x633236f2.
//
// Solidity: function mobsTotal() view returns(uint256)
func (_XmobManage *XmobManageCaller) MobsTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "mobsTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MobsTotal is a free data retrieval call binding the contract method 0x633236f2.
//
// Solidity: function mobsTotal() view returns(uint256)
func (_XmobManage *XmobManageSession) MobsTotal() (*big.Int, error) {
	return _XmobManage.Contract.MobsTotal(&_XmobManage.CallOpts)
}

// MobsTotal is a free data retrieval call binding the contract method 0x633236f2.
//
// Solidity: function mobsTotal() view returns(uint256)
func (_XmobManage *XmobManageCallerSession) MobsTotal() (*big.Int, error) {
	return _XmobManage.Contract.MobsTotal(&_XmobManage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobManage *XmobManageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _XmobManage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobManage *XmobManageSession) Owner() (common.Address, error) {
	return _XmobManage.Contract.Owner(&_XmobManage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_XmobManage *XmobManageCallerSession) Owner() (common.Address, error) {
	return _XmobManage.Contract.Owner(&_XmobManage.CallOpts)
}

// CreateMob is a paid mutator transaction binding the contract method 0x2f944487.
//
// Solidity: function createMob(address _token, uint256 _tokenId, uint256 _raiseTarget, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint64 _raiseDeadline, uint64 _deadline, uint8 _targetMode, string _name) payable returns(address)
func (_XmobManage *XmobManageTransactor) CreateMob(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _raiseTarget *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raiseDeadline uint64, _deadline uint64, _targetMode uint8, _name string) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "createMob", _token, _tokenId, _raiseTarget, _takeProfitPrice, _stopLossPrice, _raiseDeadline, _deadline, _targetMode, _name)
}

// CreateMob is a paid mutator transaction binding the contract method 0x2f944487.
//
// Solidity: function createMob(address _token, uint256 _tokenId, uint256 _raiseTarget, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint64 _raiseDeadline, uint64 _deadline, uint8 _targetMode, string _name) payable returns(address)
func (_XmobManage *XmobManageSession) CreateMob(_token common.Address, _tokenId *big.Int, _raiseTarget *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raiseDeadline uint64, _deadline uint64, _targetMode uint8, _name string) (*types.Transaction, error) {
	return _XmobManage.Contract.CreateMob(&_XmobManage.TransactOpts, _token, _tokenId, _raiseTarget, _takeProfitPrice, _stopLossPrice, _raiseDeadline, _deadline, _targetMode, _name)
}

// CreateMob is a paid mutator transaction binding the contract method 0x2f944487.
//
// Solidity: function createMob(address _token, uint256 _tokenId, uint256 _raiseTarget, uint256 _takeProfitPrice, uint256 _stopLossPrice, uint64 _raiseDeadline, uint64 _deadline, uint8 _targetMode, string _name) payable returns(address)
func (_XmobManage *XmobManageTransactorSession) CreateMob(_token common.Address, _tokenId *big.Int, _raiseTarget *big.Int, _takeProfitPrice *big.Int, _stopLossPrice *big.Int, _raiseDeadline uint64, _deadline uint64, _targetMode uint8, _name string) (*types.Transaction, error) {
	return _XmobManage.Contract.CreateMob(&_XmobManage.TransactOpts, _token, _tokenId, _raiseTarget, _takeProfitPrice, _stopLossPrice, _raiseDeadline, _deadline, _targetMode, _name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobManage *XmobManageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobManage *XmobManageSession) RenounceOwnership() (*types.Transaction, error) {
	return _XmobManage.Contract.RenounceOwnership(&_XmobManage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_XmobManage *XmobManageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _XmobManage.Contract.RenounceOwnership(&_XmobManage.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x5177942a.
//
// Solidity: function setFeeRate(uint8 _feeRate) returns()
func (_XmobManage *XmobManageTransactor) SetFeeRate(opts *bind.TransactOpts, _feeRate uint8) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "setFeeRate", _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x5177942a.
//
// Solidity: function setFeeRate(uint8 _feeRate) returns()
func (_XmobManage *XmobManageSession) SetFeeRate(_feeRate uint8) (*types.Transaction, error) {
	return _XmobManage.Contract.SetFeeRate(&_XmobManage.TransactOpts, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x5177942a.
//
// Solidity: function setFeeRate(uint8 _feeRate) returns()
func (_XmobManage *XmobManageTransactorSession) SetFeeRate(_feeRate uint8) (*types.Transaction, error) {
	return _XmobManage.Contract.SetFeeRate(&_XmobManage.TransactOpts, _feeRate)
}

// SetXMobProxy is a paid mutator transaction binding the contract method 0x5aab6465.
//
// Solidity: function setXMobProxy(address _exchangeCore) returns()
func (_XmobManage *XmobManageTransactor) SetXMobProxy(opts *bind.TransactOpts, _exchangeCore common.Address) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "setXMobProxy", _exchangeCore)
}

// SetXMobProxy is a paid mutator transaction binding the contract method 0x5aab6465.
//
// Solidity: function setXMobProxy(address _exchangeCore) returns()
func (_XmobManage *XmobManageSession) SetXMobProxy(_exchangeCore common.Address) (*types.Transaction, error) {
	return _XmobManage.Contract.SetXMobProxy(&_XmobManage.TransactOpts, _exchangeCore)
}

// SetXMobProxy is a paid mutator transaction binding the contract method 0x5aab6465.
//
// Solidity: function setXMobProxy(address _exchangeCore) returns()
func (_XmobManage *XmobManageTransactorSession) SetXMobProxy(_exchangeCore common.Address) (*types.Transaction, error) {
	return _XmobManage.Contract.SetXMobProxy(&_XmobManage.TransactOpts, _exchangeCore)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobManage *XmobManageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobManage *XmobManageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XmobManage.Contract.TransferOwnership(&_XmobManage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_XmobManage *XmobManageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _XmobManage.Contract.TransferOwnership(&_XmobManage.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address addr, uint256 amt) returns()
func (_XmobManage *XmobManageTransactor) Withdraw(opts *bind.TransactOpts, addr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _XmobManage.contract.Transact(opts, "withdraw", addr, amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address addr, uint256 amt) returns()
func (_XmobManage *XmobManageSession) Withdraw(addr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _XmobManage.Contract.Withdraw(&_XmobManage.TransactOpts, addr, amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address addr, uint256 amt) returns()
func (_XmobManage *XmobManageTransactorSession) Withdraw(addr common.Address, amt *big.Int) (*types.Transaction, error) {
	return _XmobManage.Contract.Withdraw(&_XmobManage.TransactOpts, addr, amt)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobManage *XmobManageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _XmobManage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobManage *XmobManageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _XmobManage.Contract.Fallback(&_XmobManage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_XmobManage *XmobManageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _XmobManage.Contract.Fallback(&_XmobManage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobManage *XmobManageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XmobManage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobManage *XmobManageSession) Receive() (*types.Transaction, error) {
	return _XmobManage.Contract.Receive(&_XmobManage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_XmobManage *XmobManageTransactorSession) Receive() (*types.Transaction, error) {
	return _XmobManage.Contract.Receive(&_XmobManage.TransactOpts)
}

// XmobManageDepositEthIterator is returned from FilterDepositEth and is used to iterate over the raw logs and unpacked data for DepositEth events raised by the XmobManage contract.
type XmobManageDepositEthIterator struct {
	Event *XmobManageDepositEth // Event containing the contract specifics and raw log

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
func (it *XmobManageDepositEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageDepositEth)
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
		it.Event = new(XmobManageDepositEth)
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
func (it *XmobManageDepositEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageDepositEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageDepositEth represents a DepositEth event raised by the XmobManage contract.
type XmobManageDepositEth struct {
	Sender common.Address
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositEth is a free log retrieval operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobManage *XmobManageFilterer) FilterDepositEth(opts *bind.FilterOpts) (*XmobManageDepositEthIterator, error) {

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "DepositEth")
	if err != nil {
		return nil, err
	}
	return &XmobManageDepositEthIterator{contract: _XmobManage.contract, event: "DepositEth", logs: logs, sub: sub}, nil
}

// WatchDepositEth is a free log subscription operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobManage *XmobManageFilterer) WatchDepositEth(opts *bind.WatchOpts, sink chan<- *XmobManageDepositEth) (event.Subscription, error) {

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "DepositEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageDepositEth)
				if err := _XmobManage.contract.UnpackLog(event, "DepositEth", log); err != nil {
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

// ParseDepositEth is a log parse operation binding the contract event 0x7034bb05cfe54b0d147fc0574ed166101e7f0313eb404e113974fbe2a998ca83.
//
// Solidity: event DepositEth(address sender, uint256 amt)
func (_XmobManage *XmobManageFilterer) ParseDepositEth(log types.Log) (*XmobManageDepositEth, error) {
	event := new(XmobManageDepositEth)
	if err := _XmobManage.contract.UnpackLog(event, "DepositEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobManageFeeRateSetIterator is returned from FilterFeeRateSet and is used to iterate over the raw logs and unpacked data for FeeRateSet events raised by the XmobManage contract.
type XmobManageFeeRateSetIterator struct {
	Event *XmobManageFeeRateSet // Event containing the contract specifics and raw log

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
func (it *XmobManageFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageFeeRateSet)
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
		it.Event = new(XmobManageFeeRateSet)
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
func (it *XmobManageFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageFeeRateSet represents a FeeRateSet event raised by the XmobManage contract.
type XmobManageFeeRateSet struct {
	FeeRate uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeRateSet is a free log retrieval operation binding the contract event 0x257cfd119c6475688ef5e646cc62f82c07f56b22b2682c7554874957b3e3c81a.
//
// Solidity: event FeeRateSet(uint8 feeRate)
func (_XmobManage *XmobManageFilterer) FilterFeeRateSet(opts *bind.FilterOpts) (*XmobManageFeeRateSetIterator, error) {

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "FeeRateSet")
	if err != nil {
		return nil, err
	}
	return &XmobManageFeeRateSetIterator{contract: _XmobManage.contract, event: "FeeRateSet", logs: logs, sub: sub}, nil
}

// WatchFeeRateSet is a free log subscription operation binding the contract event 0x257cfd119c6475688ef5e646cc62f82c07f56b22b2682c7554874957b3e3c81a.
//
// Solidity: event FeeRateSet(uint8 feeRate)
func (_XmobManage *XmobManageFilterer) WatchFeeRateSet(opts *bind.WatchOpts, sink chan<- *XmobManageFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "FeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageFeeRateSet)
				if err := _XmobManage.contract.UnpackLog(event, "FeeRateSet", log); err != nil {
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

// ParseFeeRateSet is a log parse operation binding the contract event 0x257cfd119c6475688ef5e646cc62f82c07f56b22b2682c7554874957b3e3c81a.
//
// Solidity: event FeeRateSet(uint8 feeRate)
func (_XmobManage *XmobManageFilterer) ParseFeeRateSet(log types.Log) (*XmobManageFeeRateSet, error) {
	event := new(XmobManageFeeRateSet)
	if err := _XmobManage.contract.UnpackLog(event, "FeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobManageMobCreateIterator is returned from FilterMobCreate and is used to iterate over the raw logs and unpacked data for MobCreate events raised by the XmobManage contract.
type XmobManageMobCreateIterator struct {
	Event *XmobManageMobCreate // Event containing the contract specifics and raw log

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
func (it *XmobManageMobCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageMobCreate)
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
		it.Event = new(XmobManageMobCreate)
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
func (it *XmobManageMobCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageMobCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageMobCreate represents a MobCreate event raised by the XmobManage contract.
type XmobManageMobCreate struct {
	Creator         common.Address
	Token           common.Address
	TokenId         *big.Int
	Proxy           common.Address
	Id              *big.Int
	RaiseTarget     *big.Int
	TakeProfitPrice *big.Int
	StopLossPrice   *big.Int
	Fee             *big.Int
	Deadline        *big.Int
	RaiseDeadline   *big.Int
	TargetMode      uint8
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMobCreate is a free log retrieval operation binding the contract event 0x54a02bfe6e6d9b8d670bd4f360dee1cebffb303b62559c89e97f258d71bcd2a9.
//
// Solidity: event MobCreate(address indexed creator, address indexed token, uint256 indexed tokenId, address proxy, uint256 id, uint256 raiseTarget, uint256 takeProfitPrice, uint256 stopLossPrice, uint256 fee, uint256 deadline, uint256 raiseDeadline, uint8 targetMode, string name)
func (_XmobManage *XmobManageFilterer) FilterMobCreate(opts *bind.FilterOpts, creator []common.Address, token []common.Address, tokenId []*big.Int) (*XmobManageMobCreateIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "MobCreate", creatorRule, tokenRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &XmobManageMobCreateIterator{contract: _XmobManage.contract, event: "MobCreate", logs: logs, sub: sub}, nil
}

// WatchMobCreate is a free log subscription operation binding the contract event 0x54a02bfe6e6d9b8d670bd4f360dee1cebffb303b62559c89e97f258d71bcd2a9.
//
// Solidity: event MobCreate(address indexed creator, address indexed token, uint256 indexed tokenId, address proxy, uint256 id, uint256 raiseTarget, uint256 takeProfitPrice, uint256 stopLossPrice, uint256 fee, uint256 deadline, uint256 raiseDeadline, uint8 targetMode, string name)
func (_XmobManage *XmobManageFilterer) WatchMobCreate(opts *bind.WatchOpts, sink chan<- *XmobManageMobCreate, creator []common.Address, token []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "MobCreate", creatorRule, tokenRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageMobCreate)
				if err := _XmobManage.contract.UnpackLog(event, "MobCreate", log); err != nil {
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

// ParseMobCreate is a log parse operation binding the contract event 0x54a02bfe6e6d9b8d670bd4f360dee1cebffb303b62559c89e97f258d71bcd2a9.
//
// Solidity: event MobCreate(address indexed creator, address indexed token, uint256 indexed tokenId, address proxy, uint256 id, uint256 raiseTarget, uint256 takeProfitPrice, uint256 stopLossPrice, uint256 fee, uint256 deadline, uint256 raiseDeadline, uint8 targetMode, string name)
func (_XmobManage *XmobManageFilterer) ParseMobCreate(log types.Log) (*XmobManageMobCreate, error) {
	event := new(XmobManageMobCreate)
	if err := _XmobManage.contract.UnpackLog(event, "MobCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobManageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the XmobManage contract.
type XmobManageOwnershipTransferredIterator struct {
	Event *XmobManageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *XmobManageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageOwnershipTransferred)
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
		it.Event = new(XmobManageOwnershipTransferred)
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
func (it *XmobManageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageOwnershipTransferred represents a OwnershipTransferred event raised by the XmobManage contract.
type XmobManageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XmobManage *XmobManageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*XmobManageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &XmobManageOwnershipTransferredIterator{contract: _XmobManage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_XmobManage *XmobManageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *XmobManageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageOwnershipTransferred)
				if err := _XmobManage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_XmobManage *XmobManageFilterer) ParseOwnershipTransferred(log types.Log) (*XmobManageOwnershipTransferred, error) {
	event := new(XmobManageOwnershipTransferred)
	if err := _XmobManage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobManageWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the XmobManage contract.
type XmobManageWithdrawIterator struct {
	Event *XmobManageWithdraw // Event containing the contract specifics and raw log

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
func (it *XmobManageWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageWithdraw)
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
		it.Event = new(XmobManageWithdraw)
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
func (it *XmobManageWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageWithdraw represents a Withdraw event raised by the XmobManage contract.
type XmobManageWithdraw struct {
	Addr common.Address
	Amt  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amt)
func (_XmobManage *XmobManageFilterer) FilterWithdraw(opts *bind.FilterOpts) (*XmobManageWithdrawIterator, error) {

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &XmobManageWithdrawIterator{contract: _XmobManage.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amt)
func (_XmobManage *XmobManageFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *XmobManageWithdraw) (event.Subscription, error) {

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageWithdraw)
				if err := _XmobManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address addr, uint256 amt)
func (_XmobManage *XmobManageFilterer) ParseWithdraw(log types.Log) (*XmobManageWithdraw, error) {
	event := new(XmobManageWithdraw)
	if err := _XmobManage.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XmobManageXMobProxySetIterator is returned from FilterXMobProxySet and is used to iterate over the raw logs and unpacked data for XMobProxySet events raised by the XmobManage contract.
type XmobManageXMobProxySetIterator struct {
	Event *XmobManageXMobProxySet // Event containing the contract specifics and raw log

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
func (it *XmobManageXMobProxySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XmobManageXMobProxySet)
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
		it.Event = new(XmobManageXMobProxySet)
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
func (it *XmobManageXMobProxySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XmobManageXMobProxySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XmobManageXMobProxySet represents a XMobProxySet event raised by the XmobManage contract.
type XmobManageXMobProxySet struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterXMobProxySet is a free log retrieval operation binding the contract event 0xc2c0b30f4a244bba60f3de8f42949dbc76dec7bfcab5a5aaf3554295c8c83957.
//
// Solidity: event XMobProxySet(address proxy)
func (_XmobManage *XmobManageFilterer) FilterXMobProxySet(opts *bind.FilterOpts) (*XmobManageXMobProxySetIterator, error) {

	logs, sub, err := _XmobManage.contract.FilterLogs(opts, "XMobProxySet")
	if err != nil {
		return nil, err
	}
	return &XmobManageXMobProxySetIterator{contract: _XmobManage.contract, event: "XMobProxySet", logs: logs, sub: sub}, nil
}

// WatchXMobProxySet is a free log subscription operation binding the contract event 0xc2c0b30f4a244bba60f3de8f42949dbc76dec7bfcab5a5aaf3554295c8c83957.
//
// Solidity: event XMobProxySet(address proxy)
func (_XmobManage *XmobManageFilterer) WatchXMobProxySet(opts *bind.WatchOpts, sink chan<- *XmobManageXMobProxySet) (event.Subscription, error) {

	logs, sub, err := _XmobManage.contract.WatchLogs(opts, "XMobProxySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XmobManageXMobProxySet)
				if err := _XmobManage.contract.UnpackLog(event, "XMobProxySet", log); err != nil {
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

// ParseXMobProxySet is a log parse operation binding the contract event 0xc2c0b30f4a244bba60f3de8f42949dbc76dec7bfcab5a5aaf3554295c8c83957.
//
// Solidity: event XMobProxySet(address proxy)
func (_XmobManage *XmobManageFilterer) ParseXMobProxySet(log types.Log) (*XmobManageXMobProxySet, error) {
	event := new(XmobManageXMobProxySet)
	if err := _XmobManage.contract.UnpackLog(event, "XMobProxySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
