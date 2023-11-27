// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subscriptioncontract

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

// SubscriptionMetaData contains all meta data concerning the Subscription contract.
var SubscriptionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownadd\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeadd\",\"type\":\"address\"}],\"name\":\"SetFeeAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeamn\",\"type\":\"uint256\"}],\"name\":\"SetFeeAmn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subamn\",\"type\":\"uint256\"}],\"name\":\"SetSubAmn\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"getFeeAdd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeAmn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subadd\",\"type\":\"address\"}],\"name\":\"getSubAdd\",\"outputs\":[{\"internalType\":\"address[3]\",\"name\":\"\",\"type\":\"address[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubAmn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subadd\",\"type\":\"address\"}],\"name\":\"getSubUnx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeadd\",\"type\":\"address\"}],\"name\":\"setFeeAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeamn\",\"type\":\"uint256\"}],\"name\":\"setFeeAmn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subamn\",\"type\":\"uint256\"}],\"name\":\"setSubAmn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subadd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unixsec\",\"type\":\"uint256\"}],\"name\":\"subOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subadd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creaone\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amntone\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creatwo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amnttwo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creathr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amntthr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unixsec\",\"type\":\"uint256\"}],\"name\":\"subThr\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subadd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creaone\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amntone\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creatwo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amnttwo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unixsec\",\"type\":\"uint256\"}],\"name\":\"subTwo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SubscriptionABI is the input ABI used to generate the binding from.
// Deprecated: Use SubscriptionMetaData.ABI instead.
var SubscriptionABI = SubscriptionMetaData.ABI

// Subscription is an auto generated Go binding around an Ethereum contract.
type Subscription struct {
	SubscriptionCaller     // Read-only binding to the contract
	SubscriptionTransactor // Write-only binding to the contract
	SubscriptionFilterer   // Log filterer for contract events
}

// SubscriptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionSession struct {
	Contract     *Subscription     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubscriptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionCallerSession struct {
	Contract *SubscriptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SubscriptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionTransactorSession struct {
	Contract     *SubscriptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SubscriptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionRaw struct {
	Contract *Subscription // Generic contract binding to access the raw methods on
}

// SubscriptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionCallerRaw struct {
	Contract *SubscriptionCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionTransactorRaw struct {
	Contract *SubscriptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscription creates a new instance of Subscription, bound to a specific deployed contract.
func NewSubscription(address common.Address, backend bind.ContractBackend) (*Subscription, error) {
	contract, err := bindSubscription(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subscription{SubscriptionCaller: SubscriptionCaller{contract: contract}, SubscriptionTransactor: SubscriptionTransactor{contract: contract}, SubscriptionFilterer: SubscriptionFilterer{contract: contract}}, nil
}

// NewSubscriptionCaller creates a new read-only instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionCaller, error) {
	contract, err := bindSubscription(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionCaller{contract: contract}, nil
}

// NewSubscriptionTransactor creates a new write-only instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionTransactor, error) {
	contract, err := bindSubscription(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionTransactor{contract: contract}, nil
}

// NewSubscriptionFilterer creates a new log filterer instance of Subscription, bound to a specific deployed contract.
func NewSubscriptionFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionFilterer, error) {
	contract, err := bindSubscription(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionFilterer{contract: contract}, nil
}

// bindSubscription binds a generic wrapper to an already deployed contract.
func bindSubscription(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscription *SubscriptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscription.Contract.SubscriptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscription *SubscriptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.Contract.SubscriptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscription *SubscriptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscription.Contract.SubscriptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscription *SubscriptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subscription.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscription *SubscriptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscription *SubscriptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscription.Contract.contract.Transact(opts, method, params...)
}

// GetFeeAdd is a free data retrieval call binding the contract method 0x4da84b64.
//
// Solidity: function getFeeAdd() view returns(address)
func (_Subscription *SubscriptionCaller) GetFeeAdd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "getFeeAdd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeAdd is a free data retrieval call binding the contract method 0x4da84b64.
//
// Solidity: function getFeeAdd() view returns(address)
func (_Subscription *SubscriptionSession) GetFeeAdd() (common.Address, error) {
	return _Subscription.Contract.GetFeeAdd(&_Subscription.CallOpts)
}

// GetFeeAdd is a free data retrieval call binding the contract method 0x4da84b64.
//
// Solidity: function getFeeAdd() view returns(address)
func (_Subscription *SubscriptionCallerSession) GetFeeAdd() (common.Address, error) {
	return _Subscription.Contract.GetFeeAdd(&_Subscription.CallOpts)
}

// GetFeeAmn is a free data retrieval call binding the contract method 0xecf9ef64.
//
// Solidity: function getFeeAmn() view returns(uint256)
func (_Subscription *SubscriptionCaller) GetFeeAmn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "getFeeAmn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeAmn is a free data retrieval call binding the contract method 0xecf9ef64.
//
// Solidity: function getFeeAmn() view returns(uint256)
func (_Subscription *SubscriptionSession) GetFeeAmn() (*big.Int, error) {
	return _Subscription.Contract.GetFeeAmn(&_Subscription.CallOpts)
}

// GetFeeAmn is a free data retrieval call binding the contract method 0xecf9ef64.
//
// Solidity: function getFeeAmn() view returns(uint256)
func (_Subscription *SubscriptionCallerSession) GetFeeAmn() (*big.Int, error) {
	return _Subscription.Contract.GetFeeAmn(&_Subscription.CallOpts)
}

// GetSubAdd is a free data retrieval call binding the contract method 0x0d981627.
//
// Solidity: function getSubAdd(address subadd) view returns(address[3])
func (_Subscription *SubscriptionCaller) GetSubAdd(opts *bind.CallOpts, subadd common.Address) ([3]common.Address, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "getSubAdd", subadd)

	if err != nil {
		return *new([3]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([3]common.Address)).(*[3]common.Address)

	return out0, err

}

// GetSubAdd is a free data retrieval call binding the contract method 0x0d981627.
//
// Solidity: function getSubAdd(address subadd) view returns(address[3])
func (_Subscription *SubscriptionSession) GetSubAdd(subadd common.Address) ([3]common.Address, error) {
	return _Subscription.Contract.GetSubAdd(&_Subscription.CallOpts, subadd)
}

// GetSubAdd is a free data retrieval call binding the contract method 0x0d981627.
//
// Solidity: function getSubAdd(address subadd) view returns(address[3])
func (_Subscription *SubscriptionCallerSession) GetSubAdd(subadd common.Address) ([3]common.Address, error) {
	return _Subscription.Contract.GetSubAdd(&_Subscription.CallOpts, subadd)
}

// GetSubAmn is a free data retrieval call binding the contract method 0x0deacca6.
//
// Solidity: function getSubAmn() view returns(uint256)
func (_Subscription *SubscriptionCaller) GetSubAmn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "getSubAmn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubAmn is a free data retrieval call binding the contract method 0x0deacca6.
//
// Solidity: function getSubAmn() view returns(uint256)
func (_Subscription *SubscriptionSession) GetSubAmn() (*big.Int, error) {
	return _Subscription.Contract.GetSubAmn(&_Subscription.CallOpts)
}

// GetSubAmn is a free data retrieval call binding the contract method 0x0deacca6.
//
// Solidity: function getSubAmn() view returns(uint256)
func (_Subscription *SubscriptionCallerSession) GetSubAmn() (*big.Int, error) {
	return _Subscription.Contract.GetSubAmn(&_Subscription.CallOpts)
}

// GetSubUnx is a free data retrieval call binding the contract method 0xcdabe7b7.
//
// Solidity: function getSubUnx(address subadd) view returns(uint256)
func (_Subscription *SubscriptionCaller) GetSubUnx(opts *bind.CallOpts, subadd common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "getSubUnx", subadd)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubUnx is a free data retrieval call binding the contract method 0xcdabe7b7.
//
// Solidity: function getSubUnx(address subadd) view returns(uint256)
func (_Subscription *SubscriptionSession) GetSubUnx(subadd common.Address) (*big.Int, error) {
	return _Subscription.Contract.GetSubUnx(&_Subscription.CallOpts, subadd)
}

// GetSubUnx is a free data retrieval call binding the contract method 0xcdabe7b7.
//
// Solidity: function getSubUnx(address subadd) view returns(uint256)
func (_Subscription *SubscriptionCallerSession) GetSubUnx(subadd common.Address) (*big.Int, error) {
	return _Subscription.Contract.GetSubUnx(&_Subscription.CallOpts, subadd)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subscription.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionSession) Owner() (common.Address, error) {
	return _Subscription.Contract.Owner(&_Subscription.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subscription *SubscriptionCallerSession) Owner() (common.Address, error) {
	return _Subscription.Contract.Owner(&_Subscription.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscription.Contract.RenounceOwnership(&_Subscription.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscription *SubscriptionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscription.Contract.RenounceOwnership(&_Subscription.TransactOpts)
}

// SetFeeAdd is a paid mutator transaction binding the contract method 0xa02d6c6b.
//
// Solidity: function setFeeAdd(address feeadd) returns()
func (_Subscription *SubscriptionTransactor) SetFeeAdd(opts *bind.TransactOpts, feeadd common.Address) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "setFeeAdd", feeadd)
}

// SetFeeAdd is a paid mutator transaction binding the contract method 0xa02d6c6b.
//
// Solidity: function setFeeAdd(address feeadd) returns()
func (_Subscription *SubscriptionSession) SetFeeAdd(feeadd common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.SetFeeAdd(&_Subscription.TransactOpts, feeadd)
}

// SetFeeAdd is a paid mutator transaction binding the contract method 0xa02d6c6b.
//
// Solidity: function setFeeAdd(address feeadd) returns()
func (_Subscription *SubscriptionTransactorSession) SetFeeAdd(feeadd common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.SetFeeAdd(&_Subscription.TransactOpts, feeadd)
}

// SetFeeAmn is a paid mutator transaction binding the contract method 0xe14cfeac.
//
// Solidity: function setFeeAmn(uint256 feeamn) returns()
func (_Subscription *SubscriptionTransactor) SetFeeAmn(opts *bind.TransactOpts, feeamn *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "setFeeAmn", feeamn)
}

// SetFeeAmn is a paid mutator transaction binding the contract method 0xe14cfeac.
//
// Solidity: function setFeeAmn(uint256 feeamn) returns()
func (_Subscription *SubscriptionSession) SetFeeAmn(feeamn *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetFeeAmn(&_Subscription.TransactOpts, feeamn)
}

// SetFeeAmn is a paid mutator transaction binding the contract method 0xe14cfeac.
//
// Solidity: function setFeeAmn(uint256 feeamn) returns()
func (_Subscription *SubscriptionTransactorSession) SetFeeAmn(feeamn *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetFeeAmn(&_Subscription.TransactOpts, feeamn)
}

// SetSubAmn is a paid mutator transaction binding the contract method 0x6aad8cee.
//
// Solidity: function setSubAmn(uint256 subamn) returns()
func (_Subscription *SubscriptionTransactor) SetSubAmn(opts *bind.TransactOpts, subamn *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "setSubAmn", subamn)
}

// SetSubAmn is a paid mutator transaction binding the contract method 0x6aad8cee.
//
// Solidity: function setSubAmn(uint256 subamn) returns()
func (_Subscription *SubscriptionSession) SetSubAmn(subamn *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetSubAmn(&_Subscription.TransactOpts, subamn)
}

// SetSubAmn is a paid mutator transaction binding the contract method 0x6aad8cee.
//
// Solidity: function setSubAmn(uint256 subamn) returns()
func (_Subscription *SubscriptionTransactorSession) SetSubAmn(subamn *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SetSubAmn(&_Subscription.TransactOpts, subamn)
}

// SubOne is a paid mutator transaction binding the contract method 0x41565a92.
//
// Solidity: function subOne(address subadd, address creator, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactor) SubOne(opts *bind.TransactOpts, subadd common.Address, creator common.Address, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "subOne", subadd, creator, unixsec)
}

// SubOne is a paid mutator transaction binding the contract method 0x41565a92.
//
// Solidity: function subOne(address subadd, address creator, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionSession) SubOne(subadd common.Address, creator common.Address, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubOne(&_Subscription.TransactOpts, subadd, creator, unixsec)
}

// SubOne is a paid mutator transaction binding the contract method 0x41565a92.
//
// Solidity: function subOne(address subadd, address creator, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactorSession) SubOne(subadd common.Address, creator common.Address, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubOne(&_Subscription.TransactOpts, subadd, creator, unixsec)
}

// SubThr is a paid mutator transaction binding the contract method 0x00f24362.
//
// Solidity: function subThr(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, address creathr, uint256 amntthr, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactor) SubThr(opts *bind.TransactOpts, subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, creathr common.Address, amntthr *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "subThr", subadd, creaone, amntone, creatwo, amnttwo, creathr, amntthr, unixsec)
}

// SubThr is a paid mutator transaction binding the contract method 0x00f24362.
//
// Solidity: function subThr(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, address creathr, uint256 amntthr, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionSession) SubThr(subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, creathr common.Address, amntthr *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubThr(&_Subscription.TransactOpts, subadd, creaone, amntone, creatwo, amnttwo, creathr, amntthr, unixsec)
}

// SubThr is a paid mutator transaction binding the contract method 0x00f24362.
//
// Solidity: function subThr(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, address creathr, uint256 amntthr, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactorSession) SubThr(subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, creathr common.Address, amntthr *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubThr(&_Subscription.TransactOpts, subadd, creaone, amntone, creatwo, amnttwo, creathr, amntthr, unixsec)
}

// SubTwo is a paid mutator transaction binding the contract method 0xbb3a4b34.
//
// Solidity: function subTwo(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactor) SubTwo(opts *bind.TransactOpts, subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "subTwo", subadd, creaone, amntone, creatwo, amnttwo, unixsec)
}

// SubTwo is a paid mutator transaction binding the contract method 0xbb3a4b34.
//
// Solidity: function subTwo(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionSession) SubTwo(subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubTwo(&_Subscription.TransactOpts, subadd, creaone, amntone, creatwo, amnttwo, unixsec)
}

// SubTwo is a paid mutator transaction binding the contract method 0xbb3a4b34.
//
// Solidity: function subTwo(address subadd, address creaone, uint256 amntone, address creatwo, uint256 amnttwo, uint256 unixsec) payable returns()
func (_Subscription *SubscriptionTransactorSession) SubTwo(subadd common.Address, creaone common.Address, amntone *big.Int, creatwo common.Address, amnttwo *big.Int, unixsec *big.Int) (*types.Transaction, error) {
	return _Subscription.Contract.SubTwo(&_Subscription.TransactOpts, subadd, creaone, amntone, creatwo, amnttwo, unixsec)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.TransferOwnership(&_Subscription.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscription *SubscriptionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscription.Contract.TransferOwnership(&_Subscription.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Subscription *SubscriptionTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Subscription.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Subscription *SubscriptionSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Subscription.Contract.Fallback(&_Subscription.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Subscription *SubscriptionTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Subscription.Contract.Fallback(&_Subscription.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Subscription *SubscriptionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscription.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Subscription *SubscriptionSession) Receive() (*types.Transaction, error) {
	return _Subscription.Contract.Receive(&_Subscription.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Subscription *SubscriptionTransactorSession) Receive() (*types.Transaction, error) {
	return _Subscription.Contract.Receive(&_Subscription.TransactOpts)
}

// SubscriptionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Subscription contract.
type SubscriptionOwnershipTransferredIterator struct {
	Event *SubscriptionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubscriptionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionOwnershipTransferred)
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
		it.Event = new(SubscriptionOwnershipTransferred)
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
func (it *SubscriptionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionOwnershipTransferred represents a OwnershipTransferred event raised by the Subscription contract.
type SubscriptionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscription *SubscriptionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubscriptionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionOwnershipTransferredIterator{contract: _Subscription.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscription *SubscriptionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubscriptionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionOwnershipTransferred)
				if err := _Subscription.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Subscription *SubscriptionFilterer) ParseOwnershipTransferred(log types.Log) (*SubscriptionOwnershipTransferred, error) {
	event := new(SubscriptionOwnershipTransferred)
	if err := _Subscription.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscriptionSetFeeAddIterator is returned from FilterSetFeeAdd and is used to iterate over the raw logs and unpacked data for SetFeeAdd events raised by the Subscription contract.
type SubscriptionSetFeeAddIterator struct {
	Event *SubscriptionSetFeeAdd // Event containing the contract specifics and raw log

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
func (it *SubscriptionSetFeeAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionSetFeeAdd)
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
		it.Event = new(SubscriptionSetFeeAdd)
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
func (it *SubscriptionSetFeeAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionSetFeeAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionSetFeeAdd represents a SetFeeAdd event raised by the Subscription contract.
type SubscriptionSetFeeAdd struct {
	Feeadd common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAdd is a free log retrieval operation binding the contract event 0x4d4874814fed3dcf20b33af17e5ae81b9c693fd3991fb1745ba5b354faa66b12.
//
// Solidity: event SetFeeAdd(address feeadd)
func (_Subscription *SubscriptionFilterer) FilterSetFeeAdd(opts *bind.FilterOpts) (*SubscriptionSetFeeAddIterator, error) {

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "SetFeeAdd")
	if err != nil {
		return nil, err
	}
	return &SubscriptionSetFeeAddIterator{contract: _Subscription.contract, event: "SetFeeAdd", logs: logs, sub: sub}, nil
}

// WatchSetFeeAdd is a free log subscription operation binding the contract event 0x4d4874814fed3dcf20b33af17e5ae81b9c693fd3991fb1745ba5b354faa66b12.
//
// Solidity: event SetFeeAdd(address feeadd)
func (_Subscription *SubscriptionFilterer) WatchSetFeeAdd(opts *bind.WatchOpts, sink chan<- *SubscriptionSetFeeAdd) (event.Subscription, error) {

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "SetFeeAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionSetFeeAdd)
				if err := _Subscription.contract.UnpackLog(event, "SetFeeAdd", log); err != nil {
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

// ParseSetFeeAdd is a log parse operation binding the contract event 0x4d4874814fed3dcf20b33af17e5ae81b9c693fd3991fb1745ba5b354faa66b12.
//
// Solidity: event SetFeeAdd(address feeadd)
func (_Subscription *SubscriptionFilterer) ParseSetFeeAdd(log types.Log) (*SubscriptionSetFeeAdd, error) {
	event := new(SubscriptionSetFeeAdd)
	if err := _Subscription.contract.UnpackLog(event, "SetFeeAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscriptionSetFeeAmnIterator is returned from FilterSetFeeAmn and is used to iterate over the raw logs and unpacked data for SetFeeAmn events raised by the Subscription contract.
type SubscriptionSetFeeAmnIterator struct {
	Event *SubscriptionSetFeeAmn // Event containing the contract specifics and raw log

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
func (it *SubscriptionSetFeeAmnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionSetFeeAmn)
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
		it.Event = new(SubscriptionSetFeeAmn)
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
func (it *SubscriptionSetFeeAmnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionSetFeeAmnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionSetFeeAmn represents a SetFeeAmn event raised by the Subscription contract.
type SubscriptionSetFeeAmn struct {
	Feeamn *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFeeAmn is a free log retrieval operation binding the contract event 0x50d8b1ef22a705f331c824019061de8e1458c06f93414471c37eb8ef45967556.
//
// Solidity: event SetFeeAmn(uint256 feeamn)
func (_Subscription *SubscriptionFilterer) FilterSetFeeAmn(opts *bind.FilterOpts) (*SubscriptionSetFeeAmnIterator, error) {

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "SetFeeAmn")
	if err != nil {
		return nil, err
	}
	return &SubscriptionSetFeeAmnIterator{contract: _Subscription.contract, event: "SetFeeAmn", logs: logs, sub: sub}, nil
}

// WatchSetFeeAmn is a free log subscription operation binding the contract event 0x50d8b1ef22a705f331c824019061de8e1458c06f93414471c37eb8ef45967556.
//
// Solidity: event SetFeeAmn(uint256 feeamn)
func (_Subscription *SubscriptionFilterer) WatchSetFeeAmn(opts *bind.WatchOpts, sink chan<- *SubscriptionSetFeeAmn) (event.Subscription, error) {

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "SetFeeAmn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionSetFeeAmn)
				if err := _Subscription.contract.UnpackLog(event, "SetFeeAmn", log); err != nil {
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

// ParseSetFeeAmn is a log parse operation binding the contract event 0x50d8b1ef22a705f331c824019061de8e1458c06f93414471c37eb8ef45967556.
//
// Solidity: event SetFeeAmn(uint256 feeamn)
func (_Subscription *SubscriptionFilterer) ParseSetFeeAmn(log types.Log) (*SubscriptionSetFeeAmn, error) {
	event := new(SubscriptionSetFeeAmn)
	if err := _Subscription.contract.UnpackLog(event, "SetFeeAmn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubscriptionSetSubAmnIterator is returned from FilterSetSubAmn and is used to iterate over the raw logs and unpacked data for SetSubAmn events raised by the Subscription contract.
type SubscriptionSetSubAmnIterator struct {
	Event *SubscriptionSetSubAmn // Event containing the contract specifics and raw log

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
func (it *SubscriptionSetSubAmnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionSetSubAmn)
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
		it.Event = new(SubscriptionSetSubAmn)
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
func (it *SubscriptionSetSubAmnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionSetSubAmnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionSetSubAmn represents a SetSubAmn event raised by the Subscription contract.
type SubscriptionSetSubAmn struct {
	Subamn *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSubAmn is a free log retrieval operation binding the contract event 0xc34e816c188f4d4080cc4f5ecc5587b4425c34350a5e9632efd8a370449d9be9.
//
// Solidity: event SetSubAmn(uint256 subamn)
func (_Subscription *SubscriptionFilterer) FilterSetSubAmn(opts *bind.FilterOpts) (*SubscriptionSetSubAmnIterator, error) {

	logs, sub, err := _Subscription.contract.FilterLogs(opts, "SetSubAmn")
	if err != nil {
		return nil, err
	}
	return &SubscriptionSetSubAmnIterator{contract: _Subscription.contract, event: "SetSubAmn", logs: logs, sub: sub}, nil
}

// WatchSetSubAmn is a free log subscription operation binding the contract event 0xc34e816c188f4d4080cc4f5ecc5587b4425c34350a5e9632efd8a370449d9be9.
//
// Solidity: event SetSubAmn(uint256 subamn)
func (_Subscription *SubscriptionFilterer) WatchSetSubAmn(opts *bind.WatchOpts, sink chan<- *SubscriptionSetSubAmn) (event.Subscription, error) {

	logs, sub, err := _Subscription.contract.WatchLogs(opts, "SetSubAmn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionSetSubAmn)
				if err := _Subscription.contract.UnpackLog(event, "SetSubAmn", log); err != nil {
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

// ParseSetSubAmn is a log parse operation binding the contract event 0xc34e816c188f4d4080cc4f5ecc5587b4425c34350a5e9632efd8a370449d9be9.
//
// Solidity: event SetSubAmn(uint256 subamn)
func (_Subscription *SubscriptionFilterer) ParseSetSubAmn(log types.Log) (*SubscriptionSetSubAmn, error) {
	event := new(SubscriptionSetSubAmn)
	if err := _Subscription.contract.UnpackLog(event, "SetSubAmn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
