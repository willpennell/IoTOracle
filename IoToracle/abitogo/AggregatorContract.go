// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abitogo

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

// AggregatorContractMetaData contains all meta data concerning the AggregatorContract contract.
var AggregatorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"AggregationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ResponseReceived\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_actualResult\",\"type\":\"bytes\"}],\"name\":\"receiveResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregatorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorContractMetaData.ABI instead.
var AggregatorContractABI = AggregatorContractMetaData.ABI

// AggregatorContract is an auto generated Go binding around an Ethereum contract.
type AggregatorContract struct {
	AggregatorContractCaller     // Read-only binding to the contract
	AggregatorContractTransactor // Write-only binding to the contract
	AggregatorContractFilterer   // Log filterer for contract events
}

// AggregatorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorContractSession struct {
	Contract     *AggregatorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AggregatorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorContractCallerSession struct {
	Contract *AggregatorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AggregatorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorContractTransactorSession struct {
	Contract     *AggregatorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AggregatorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorContractRaw struct {
	Contract *AggregatorContract // Generic contract binding to access the raw methods on
}

// AggregatorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorContractCallerRaw struct {
	Contract *AggregatorContractCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorContractTransactorRaw struct {
	Contract *AggregatorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorContract creates a new instance of AggregatorContract, bound to a specific deployed contract.
func NewAggregatorContract(address common.Address, backend bind.ContractBackend) (*AggregatorContract, error) {
	contract, err := bindAggregatorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorContract{AggregatorContractCaller: AggregatorContractCaller{contract: contract}, AggregatorContractTransactor: AggregatorContractTransactor{contract: contract}, AggregatorContractFilterer: AggregatorContractFilterer{contract: contract}}, nil
}

// NewAggregatorContractCaller creates a new read-only instance of AggregatorContract, bound to a specific deployed contract.
func NewAggregatorContractCaller(address common.Address, caller bind.ContractCaller) (*AggregatorContractCaller, error) {
	contract, err := bindAggregatorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorContractCaller{contract: contract}, nil
}

// NewAggregatorContractTransactor creates a new write-only instance of AggregatorContract, bound to a specific deployed contract.
func NewAggregatorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorContractTransactor, error) {
	contract, err := bindAggregatorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorContractTransactor{contract: contract}, nil
}

// NewAggregatorContractFilterer creates a new log filterer instance of AggregatorContract, bound to a specific deployed contract.
func NewAggregatorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorContractFilterer, error) {
	contract, err := bindAggregatorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorContractFilterer{contract: contract}, nil
}

// bindAggregatorContract binds a generic wrapper to an already deployed contract.
func bindAggregatorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorContract *AggregatorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorContract.Contract.AggregatorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorContract *AggregatorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorContract.Contract.AggregatorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorContract *AggregatorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorContract.Contract.AggregatorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorContract *AggregatorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorContract *AggregatorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorContract *AggregatorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorContract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregatorContract *AggregatorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregatorContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregatorContract *AggregatorContractSession) Owner() (common.Address, error) {
	return _AggregatorContract.Contract.Owner(&_AggregatorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregatorContract *AggregatorContractCallerSession) Owner() (common.Address, error) {
	return _AggregatorContract.Contract.Owner(&_AggregatorContract.CallOpts)
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0xf8f14f42.
//
// Solidity: function receiveResponse(uint256 _requestID, bytes _actualResult) returns(bool)
func (_AggregatorContract *AggregatorContractTransactor) ReceiveResponse(opts *bind.TransactOpts, _requestID *big.Int, _actualResult []byte) (*types.Transaction, error) {
	return _AggregatorContract.contract.Transact(opts, "receiveResponse", _requestID, _actualResult)
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0xf8f14f42.
//
// Solidity: function receiveResponse(uint256 _requestID, bytes _actualResult) returns(bool)
func (_AggregatorContract *AggregatorContractSession) ReceiveResponse(_requestID *big.Int, _actualResult []byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.ReceiveResponse(&_AggregatorContract.TransactOpts, _requestID, _actualResult)
}

// ReceiveResponse is a paid mutator transaction binding the contract method 0xf8f14f42.
//
// Solidity: function receiveResponse(uint256 _requestID, bytes _actualResult) returns(bool)
func (_AggregatorContract *AggregatorContractTransactorSession) ReceiveResponse(_requestID *big.Int, _actualResult []byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.ReceiveResponse(&_AggregatorContract.TransactOpts, _requestID, _actualResult)
}

// AggregatorContractAggregationCompletedIterator is returned from FilterAggregationCompleted and is used to iterate over the raw logs and unpacked data for AggregationCompleted events raised by the AggregatorContract contract.
type AggregatorContractAggregationCompletedIterator struct {
	Event *AggregatorContractAggregationCompleted // Event containing the contract specifics and raw log

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
func (it *AggregatorContractAggregationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractAggregationCompleted)
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
		it.Event = new(AggregatorContractAggregationCompleted)
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
func (it *AggregatorContractAggregationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractAggregationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractAggregationCompleted represents a AggregationCompleted event raised by the AggregatorContract contract.
type AggregatorContractAggregationCompleted struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAggregationCompleted is a free log retrieval operation binding the contract event 0xab982296b91a9b70c6739c96524c481cf0a1714be617e9bb39538f46ad5c8800.
//
// Solidity: event AggregationCompleted(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) FilterAggregationCompleted(opts *bind.FilterOpts) (*AggregatorContractAggregationCompletedIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "AggregationCompleted")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractAggregationCompletedIterator{contract: _AggregatorContract.contract, event: "AggregationCompleted", logs: logs, sub: sub}, nil
}

// WatchAggregationCompleted is a free log subscription operation binding the contract event 0xab982296b91a9b70c6739c96524c481cf0a1714be617e9bb39538f46ad5c8800.
//
// Solidity: event AggregationCompleted(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) WatchAggregationCompleted(opts *bind.WatchOpts, sink chan<- *AggregatorContractAggregationCompleted) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "AggregationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractAggregationCompleted)
				if err := _AggregatorContract.contract.UnpackLog(event, "AggregationCompleted", log); err != nil {
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

// ParseAggregationCompleted is a log parse operation binding the contract event 0xab982296b91a9b70c6739c96524c481cf0a1714be617e9bb39538f46ad5c8800.
//
// Solidity: event AggregationCompleted(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) ParseAggregationCompleted(log types.Log) (*AggregatorContractAggregationCompleted, error) {
	event := new(AggregatorContractAggregationCompleted)
	if err := _AggregatorContract.contract.UnpackLog(event, "AggregationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorContractResponseReceivedIterator is returned from FilterResponseReceived and is used to iterate over the raw logs and unpacked data for ResponseReceived events raised by the AggregatorContract contract.
type AggregatorContractResponseReceivedIterator struct {
	Event *AggregatorContractResponseReceived // Event containing the contract specifics and raw log

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
func (it *AggregatorContractResponseReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractResponseReceived)
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
		it.Event = new(AggregatorContractResponseReceived)
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
func (it *AggregatorContractResponseReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractResponseReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractResponseReceived represents a ResponseReceived event raised by the AggregatorContract contract.
type AggregatorContractResponseReceived struct {
	Arg0 common.Address
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResponseReceived is a free log retrieval operation binding the contract event 0xe4cba15a82b5f91e40e6fcd30c425c90aa715d1d6d73eb1fedb1baf58c28d6dd.
//
// Solidity: event ResponseReceived(address arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) FilterResponseReceived(opts *bind.FilterOpts) (*AggregatorContractResponseReceivedIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "ResponseReceived")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractResponseReceivedIterator{contract: _AggregatorContract.contract, event: "ResponseReceived", logs: logs, sub: sub}, nil
}

// WatchResponseReceived is a free log subscription operation binding the contract event 0xe4cba15a82b5f91e40e6fcd30c425c90aa715d1d6d73eb1fedb1baf58c28d6dd.
//
// Solidity: event ResponseReceived(address arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) WatchResponseReceived(opts *bind.WatchOpts, sink chan<- *AggregatorContractResponseReceived) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "ResponseReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractResponseReceived)
				if err := _AggregatorContract.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
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

// ParseResponseReceived is a log parse operation binding the contract event 0xe4cba15a82b5f91e40e6fcd30c425c90aa715d1d6d73eb1fedb1baf58c28d6dd.
//
// Solidity: event ResponseReceived(address arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) ParseResponseReceived(log types.Log) (*AggregatorContractResponseReceived, error) {
	event := new(AggregatorContractResponseReceived)
	if err := _AggregatorContract.contract.UnpackLog(event, "ResponseReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
