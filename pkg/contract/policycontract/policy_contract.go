// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package policycontract

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

// TripleRecord is an auto generated low-level Go binding around an user-defined struct.
type TripleRecord struct {
	Sys *big.Int
	Mem common.Address
	Acc *big.Int
}

// PolicyMetaData contains all meta data concerning the Policy contract.
var PolicyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amo\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"name\":\"CreateMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"name\":\"CreateSystem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"name\":\"DeleteMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"name\":\"DeleteSystem\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"internalType\":\"structTriple.Record\",\"name\":\"rec\",\"type\":\"tuple\"}],\"name\":\"createRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"internalType\":\"structTriple.Record\",\"name\":\"rec\",\"type\":\"tuple\"}],\"name\":\"deleteRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cur\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blo\",\"type\":\"uint256\"}],\"name\":\"searchRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"acc\",\"type\":\"uint256\"}],\"internalType\":\"structTriple.Record[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyMetaData.ABI instead.
var PolicyABI = PolicyMetaData.ABI

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// SearchAmount is a free data retrieval call binding the contract method 0x248ecc97.
//
// Solidity: function searchAmount() view returns(uint256)
func (_Policy *PolicyCaller) SearchAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "searchAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SearchAmount is a free data retrieval call binding the contract method 0x248ecc97.
//
// Solidity: function searchAmount() view returns(uint256)
func (_Policy *PolicySession) SearchAmount() (*big.Int, error) {
	return _Policy.Contract.SearchAmount(&_Policy.CallOpts)
}

// SearchAmount is a free data retrieval call binding the contract method 0x248ecc97.
//
// Solidity: function searchAmount() view returns(uint256)
func (_Policy *PolicyCallerSession) SearchAmount() (*big.Int, error) {
	return _Policy.Contract.SearchAmount(&_Policy.CallOpts)
}

// SearchBlocks is a free data retrieval call binding the contract method 0xa28155e7.
//
// Solidity: function searchBlocks() view returns(uint256)
func (_Policy *PolicyCaller) SearchBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "searchBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SearchBlocks is a free data retrieval call binding the contract method 0xa28155e7.
//
// Solidity: function searchBlocks() view returns(uint256)
func (_Policy *PolicySession) SearchBlocks() (*big.Int, error) {
	return _Policy.Contract.SearchBlocks(&_Policy.CallOpts)
}

// SearchBlocks is a free data retrieval call binding the contract method 0xa28155e7.
//
// Solidity: function searchBlocks() view returns(uint256)
func (_Policy *PolicyCallerSession) SearchBlocks() (*big.Int, error) {
	return _Policy.Contract.SearchBlocks(&_Policy.CallOpts)
}

// SearchRecord is a free data retrieval call binding the contract method 0xf0279ec9.
//
// Solidity: function searchRecord(uint256 cur, uint256 blo) view returns(uint256, (uint256,address,uint256)[])
func (_Policy *PolicyCaller) SearchRecord(opts *bind.CallOpts, cur *big.Int, blo *big.Int) (*big.Int, []TripleRecord, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "searchRecord", cur, blo)

	if err != nil {
		return *new(*big.Int), *new([]TripleRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]TripleRecord)).(*[]TripleRecord)

	return out0, out1, err

}

// SearchRecord is a free data retrieval call binding the contract method 0xf0279ec9.
//
// Solidity: function searchRecord(uint256 cur, uint256 blo) view returns(uint256, (uint256,address,uint256)[])
func (_Policy *PolicySession) SearchRecord(cur *big.Int, blo *big.Int) (*big.Int, []TripleRecord, error) {
	return _Policy.Contract.SearchRecord(&_Policy.CallOpts, cur, blo)
}

// SearchRecord is a free data retrieval call binding the contract method 0xf0279ec9.
//
// Solidity: function searchRecord(uint256 cur, uint256 blo) view returns(uint256, (uint256,address,uint256)[])
func (_Policy *PolicyCallerSession) SearchRecord(cur *big.Int, blo *big.Int) (*big.Int, []TripleRecord, error) {
	return _Policy.Contract.SearchRecord(&_Policy.CallOpts, cur, blo)
}

// CreateRecord is a paid mutator transaction binding the contract method 0x3ba76024.
//
// Solidity: function createRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicyTransactor) CreateRecord(opts *bind.TransactOpts, rec TripleRecord) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "createRecord", rec)
}

// CreateRecord is a paid mutator transaction binding the contract method 0x3ba76024.
//
// Solidity: function createRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicySession) CreateRecord(rec TripleRecord) (*types.Transaction, error) {
	return _Policy.Contract.CreateRecord(&_Policy.TransactOpts, rec)
}

// CreateRecord is a paid mutator transaction binding the contract method 0x3ba76024.
//
// Solidity: function createRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicyTransactorSession) CreateRecord(rec TripleRecord) (*types.Transaction, error) {
	return _Policy.Contract.CreateRecord(&_Policy.TransactOpts, rec)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x9a0bb02d.
//
// Solidity: function deleteRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicyTransactor) DeleteRecord(opts *bind.TransactOpts, rec TripleRecord) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "deleteRecord", rec)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x9a0bb02d.
//
// Solidity: function deleteRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicySession) DeleteRecord(rec TripleRecord) (*types.Transaction, error) {
	return _Policy.Contract.DeleteRecord(&_Policy.TransactOpts, rec)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x9a0bb02d.
//
// Solidity: function deleteRecord((uint256,address,uint256) rec) returns()
func (_Policy *PolicyTransactorSession) DeleteRecord(rec TripleRecord) (*types.Transaction, error) {
	return _Policy.Contract.DeleteRecord(&_Policy.TransactOpts, rec)
}

// PolicyCreateMemberIterator is returned from FilterCreateMember and is used to iterate over the raw logs and unpacked data for CreateMember events raised by the Policy contract.
type PolicyCreateMemberIterator struct {
	Event *PolicyCreateMember // Event containing the contract specifics and raw log

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
func (it *PolicyCreateMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyCreateMember)
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
		it.Event = new(PolicyCreateMember)
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
func (it *PolicyCreateMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyCreateMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyCreateMember represents a CreateMember event raised by the Policy contract.
type PolicyCreateMember struct {
	Sys *big.Int
	Mem common.Address
	Acc *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreateMember is a free log retrieval operation binding the contract event 0x597262d9c149019effb55207ba032436b598e66a3ec20876b7adde4c99fe3a51.
//
// Solidity: event CreateMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) FilterCreateMember(opts *bind.FilterOpts, sys []*big.Int, mem []common.Address, acc []*big.Int) (*PolicyCreateMemberIterator, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "CreateMember", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return &PolicyCreateMemberIterator{contract: _Policy.contract, event: "CreateMember", logs: logs, sub: sub}, nil
}

// WatchCreateMember is a free log subscription operation binding the contract event 0x597262d9c149019effb55207ba032436b598e66a3ec20876b7adde4c99fe3a51.
//
// Solidity: event CreateMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) WatchCreateMember(opts *bind.WatchOpts, sink chan<- *PolicyCreateMember, sys []*big.Int, mem []common.Address, acc []*big.Int) (event.Subscription, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "CreateMember", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyCreateMember)
				if err := _Policy.contract.UnpackLog(event, "CreateMember", log); err != nil {
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

// ParseCreateMember is a log parse operation binding the contract event 0x597262d9c149019effb55207ba032436b598e66a3ec20876b7adde4c99fe3a51.
//
// Solidity: event CreateMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) ParseCreateMember(log types.Log) (*PolicyCreateMember, error) {
	event := new(PolicyCreateMember)
	if err := _Policy.contract.UnpackLog(event, "CreateMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyCreateSystemIterator is returned from FilterCreateSystem and is used to iterate over the raw logs and unpacked data for CreateSystem events raised by the Policy contract.
type PolicyCreateSystemIterator struct {
	Event *PolicyCreateSystem // Event containing the contract specifics and raw log

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
func (it *PolicyCreateSystemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyCreateSystem)
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
		it.Event = new(PolicyCreateSystem)
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
func (it *PolicyCreateSystemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyCreateSystemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyCreateSystem represents a CreateSystem event raised by the Policy contract.
type PolicyCreateSystem struct {
	Sys *big.Int
	Mem common.Address
	Acc *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreateSystem is a free log retrieval operation binding the contract event 0x29a61a256d6b667ad24bd3de68182f43b9e500e07a8cc8eb1fc738923496a105.
//
// Solidity: event CreateSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) FilterCreateSystem(opts *bind.FilterOpts, sys []*big.Int, mem []common.Address, acc []*big.Int) (*PolicyCreateSystemIterator, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "CreateSystem", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return &PolicyCreateSystemIterator{contract: _Policy.contract, event: "CreateSystem", logs: logs, sub: sub}, nil
}

// WatchCreateSystem is a free log subscription operation binding the contract event 0x29a61a256d6b667ad24bd3de68182f43b9e500e07a8cc8eb1fc738923496a105.
//
// Solidity: event CreateSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) WatchCreateSystem(opts *bind.WatchOpts, sink chan<- *PolicyCreateSystem, sys []*big.Int, mem []common.Address, acc []*big.Int) (event.Subscription, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "CreateSystem", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyCreateSystem)
				if err := _Policy.contract.UnpackLog(event, "CreateSystem", log); err != nil {
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

// ParseCreateSystem is a log parse operation binding the contract event 0x29a61a256d6b667ad24bd3de68182f43b9e500e07a8cc8eb1fc738923496a105.
//
// Solidity: event CreateSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) ParseCreateSystem(log types.Log) (*PolicyCreateSystem, error) {
	event := new(PolicyCreateSystem)
	if err := _Policy.contract.UnpackLog(event, "CreateSystem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyDeleteMemberIterator is returned from FilterDeleteMember and is used to iterate over the raw logs and unpacked data for DeleteMember events raised by the Policy contract.
type PolicyDeleteMemberIterator struct {
	Event *PolicyDeleteMember // Event containing the contract specifics and raw log

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
func (it *PolicyDeleteMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyDeleteMember)
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
		it.Event = new(PolicyDeleteMember)
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
func (it *PolicyDeleteMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyDeleteMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyDeleteMember represents a DeleteMember event raised by the Policy contract.
type PolicyDeleteMember struct {
	Sys *big.Int
	Mem common.Address
	Acc *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeleteMember is a free log retrieval operation binding the contract event 0xf657bb80ad274cd01d7af02a864775c2c3f307137297c9812da2df0943f0d4a2.
//
// Solidity: event DeleteMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) FilterDeleteMember(opts *bind.FilterOpts, sys []*big.Int, mem []common.Address, acc []*big.Int) (*PolicyDeleteMemberIterator, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "DeleteMember", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return &PolicyDeleteMemberIterator{contract: _Policy.contract, event: "DeleteMember", logs: logs, sub: sub}, nil
}

// WatchDeleteMember is a free log subscription operation binding the contract event 0xf657bb80ad274cd01d7af02a864775c2c3f307137297c9812da2df0943f0d4a2.
//
// Solidity: event DeleteMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) WatchDeleteMember(opts *bind.WatchOpts, sink chan<- *PolicyDeleteMember, sys []*big.Int, mem []common.Address, acc []*big.Int) (event.Subscription, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "DeleteMember", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyDeleteMember)
				if err := _Policy.contract.UnpackLog(event, "DeleteMember", log); err != nil {
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

// ParseDeleteMember is a log parse operation binding the contract event 0xf657bb80ad274cd01d7af02a864775c2c3f307137297c9812da2df0943f0d4a2.
//
// Solidity: event DeleteMember(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) ParseDeleteMember(log types.Log) (*PolicyDeleteMember, error) {
	event := new(PolicyDeleteMember)
	if err := _Policy.contract.UnpackLog(event, "DeleteMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyDeleteSystemIterator is returned from FilterDeleteSystem and is used to iterate over the raw logs and unpacked data for DeleteSystem events raised by the Policy contract.
type PolicyDeleteSystemIterator struct {
	Event *PolicyDeleteSystem // Event containing the contract specifics and raw log

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
func (it *PolicyDeleteSystemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyDeleteSystem)
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
		it.Event = new(PolicyDeleteSystem)
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
func (it *PolicyDeleteSystemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyDeleteSystemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyDeleteSystem represents a DeleteSystem event raised by the Policy contract.
type PolicyDeleteSystem struct {
	Sys *big.Int
	Mem common.Address
	Acc *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeleteSystem is a free log retrieval operation binding the contract event 0x82b66663a2fb041ac966f141cf64abfda9d4e78f5cc742e4e0cb22706c872f3c.
//
// Solidity: event DeleteSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) FilterDeleteSystem(opts *bind.FilterOpts, sys []*big.Int, mem []common.Address, acc []*big.Int) (*PolicyDeleteSystemIterator, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "DeleteSystem", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return &PolicyDeleteSystemIterator{contract: _Policy.contract, event: "DeleteSystem", logs: logs, sub: sub}, nil
}

// WatchDeleteSystem is a free log subscription operation binding the contract event 0x82b66663a2fb041ac966f141cf64abfda9d4e78f5cc742e4e0cb22706c872f3c.
//
// Solidity: event DeleteSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) WatchDeleteSystem(opts *bind.WatchOpts, sink chan<- *PolicyDeleteSystem, sys []*big.Int, mem []common.Address, acc []*big.Int) (event.Subscription, error) {

	var sysRule []interface{}
	for _, sysItem := range sys {
		sysRule = append(sysRule, sysItem)
	}
	var memRule []interface{}
	for _, memItem := range mem {
		memRule = append(memRule, memItem)
	}
	var accRule []interface{}
	for _, accItem := range acc {
		accRule = append(accRule, accItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "DeleteSystem", sysRule, memRule, accRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyDeleteSystem)
				if err := _Policy.contract.UnpackLog(event, "DeleteSystem", log); err != nil {
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

// ParseDeleteSystem is a log parse operation binding the contract event 0x82b66663a2fb041ac966f141cf64abfda9d4e78f5cc742e4e0cb22706c872f3c.
//
// Solidity: event DeleteSystem(uint256 indexed sys, address indexed mem, uint256 indexed acc)
func (_Policy *PolicyFilterer) ParseDeleteSystem(log types.Log) (*PolicyDeleteSystem, error) {
	event := new(PolicyDeleteSystem)
	if err := _Policy.contract.UnpackLog(event, "DeleteSystem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
