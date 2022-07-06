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

// IoTOracleContractMetaData contains all meta data concerning the IoTOracleContract contract.
var IoTOracleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Delivery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"requiredResult\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"IotId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"}],\"name\":\"RequestIoTInfo\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_IoTId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_requiredResult\",\"type\":\"bool\"}],\"name\":\"makeRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"IoTId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"requiredResult\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"pHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IoTOracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use IoTOracleContractMetaData.ABI instead.
var IoTOracleContractABI = IoTOracleContractMetaData.ABI

// IoTOracleContract is an auto generated Go binding around an Ethereum contract.
type IoTOracleContract struct {
	IoTOracleContractCaller     // Read-only binding to the contract
	IoTOracleContractTransactor // Write-only binding to the contract
	IoTOracleContractFilterer   // Log filterer for contract events
}

// IoTOracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTOracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTOracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTOracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTOracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTOracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTOracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTOracleContractSession struct {
	Contract     *IoTOracleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IoTOracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTOracleContractCallerSession struct {
	Contract *IoTOracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IoTOracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTOracleContractTransactorSession struct {
	Contract     *IoTOracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IoTOracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTOracleContractRaw struct {
	Contract *IoTOracleContract // Generic contract binding to access the raw methods on
}

// IoTOracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTOracleContractCallerRaw struct {
	Contract *IoTOracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// IoTOracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTOracleContractTransactorRaw struct {
	Contract *IoTOracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTOracleContract creates a new instance of IoTOracleContract, bound to a specific deployed contract.
func NewIoTOracleContract(address common.Address, backend bind.ContractBackend) (*IoTOracleContract, error) {
	contract, err := bindIoTOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTOracleContract{IoTOracleContractCaller: IoTOracleContractCaller{contract: contract}, IoTOracleContractTransactor: IoTOracleContractTransactor{contract: contract}, IoTOracleContractFilterer: IoTOracleContractFilterer{contract: contract}}, nil
}

// NewIoTOracleContractCaller creates a new read-only instance of IoTOracleContract, bound to a specific deployed contract.
func NewIoTOracleContractCaller(address common.Address, caller bind.ContractCaller) (*IoTOracleContractCaller, error) {
	contract, err := bindIoTOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTOracleContractCaller{contract: contract}, nil
}

// NewIoTOracleContractTransactor creates a new write-only instance of IoTOracleContract, bound to a specific deployed contract.
func NewIoTOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTOracleContractTransactor, error) {
	contract, err := bindIoTOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTOracleContractTransactor{contract: contract}, nil
}

// NewIoTOracleContractFilterer creates a new log filterer instance of IoTOracleContract, bound to a specific deployed contract.
func NewIoTOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTOracleContractFilterer, error) {
	contract, err := bindIoTOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTOracleContractFilterer{contract: contract}, nil
}

// bindIoTOracleContract binds a generic wrapper to an already deployed contract.
func bindIoTOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTOracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTOracleContract *IoTOracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoTOracleContract.Contract.IoTOracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTOracleContract *IoTOracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.IoTOracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTOracleContract *IoTOracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.IoTOracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTOracleContract *IoTOracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoTOracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTOracleContract *IoTOracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTOracleContract *IoTOracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.contract.Transact(opts, method, params...)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address requester, address callbackAddress, bytes IoTId, bool requiredResult, bytes32 pHash)
func (_IoTOracleContract *IoTOracleContractCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestId       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	IoTId           []byte
	RequiredResult  bool
	PHash           [32]byte
}, error) {
	var out []interface{}
	err := _IoTOracleContract.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		RequestId       *big.Int
		Requester       common.Address
		CallbackAddress common.Address
		IoTId           []byte
		RequiredResult  bool
		PHash           [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CallbackAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IoTId = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.RequiredResult = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.PHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address requester, address callbackAddress, bytes IoTId, bool requiredResult, bytes32 pHash)
func (_IoTOracleContract *IoTOracleContractSession) Requests(arg0 *big.Int) (struct {
	RequestId       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	IoTId           []byte
	RequiredResult  bool
	PHash           [32]byte
}, error) {
	return _IoTOracleContract.Contract.Requests(&_IoTOracleContract.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address requester, address callbackAddress, bytes IoTId, bool requiredResult, bytes32 pHash)
func (_IoTOracleContract *IoTOracleContractCallerSession) Requests(arg0 *big.Int) (struct {
	RequestId       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	IoTId           []byte
	RequiredResult  bool
	PHash           [32]byte
}, error) {
	return _IoTOracleContract.Contract.Requests(&_IoTOracleContract.CallOpts, arg0)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x6f9e87d6.
//
// Solidity: function makeRequest(address _callbackAddress, bytes _IoTId, bool _requiredResult) returns(uint256)
func (_IoTOracleContract *IoTOracleContractTransactor) MakeRequest(opts *bind.TransactOpts, _callbackAddress common.Address, _IoTId []byte, _requiredResult bool) (*types.Transaction, error) {
	return _IoTOracleContract.contract.Transact(opts, "makeRequest", _callbackAddress, _IoTId, _requiredResult)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x6f9e87d6.
//
// Solidity: function makeRequest(address _callbackAddress, bytes _IoTId, bool _requiredResult) returns(uint256)
func (_IoTOracleContract *IoTOracleContractSession) MakeRequest(_callbackAddress common.Address, _IoTId []byte, _requiredResult bool) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.MakeRequest(&_IoTOracleContract.TransactOpts, _callbackAddress, _IoTId, _requiredResult)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x6f9e87d6.
//
// Solidity: function makeRequest(address _callbackAddress, bytes _IoTId, bool _requiredResult) returns(uint256)
func (_IoTOracleContract *IoTOracleContractTransactorSession) MakeRequest(_callbackAddress common.Address, _IoTId []byte, _requiredResult bool) (*types.Transaction, error) {
	return _IoTOracleContract.Contract.MakeRequest(&_IoTOracleContract.TransactOpts, _callbackAddress, _IoTId, _requiredResult)
}

// IoTOracleContractDeliveryIterator is returned from FilterDelivery and is used to iterate over the raw logs and unpacked data for Delivery events raised by the IoTOracleContract contract.
type IoTOracleContractDeliveryIterator struct {
	Event *IoTOracleContractDelivery // Event containing the contract specifics and raw log

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
func (it *IoTOracleContractDeliveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTOracleContractDelivery)
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
		it.Event = new(IoTOracleContractDelivery)
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
func (it *IoTOracleContractDeliveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTOracleContractDeliveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTOracleContractDelivery represents a Delivery event raised by the IoTOracleContract contract.
type IoTOracleContractDelivery struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelivery is a free log retrieval operation binding the contract event 0x5ef6ad57e4a13b3646f0527fbd8571dfd05d25fe8fdec201e9da21dbda3943fd.
//
// Solidity: event Delivery()
func (_IoTOracleContract *IoTOracleContractFilterer) FilterDelivery(opts *bind.FilterOpts) (*IoTOracleContractDeliveryIterator, error) {

	logs, sub, err := _IoTOracleContract.contract.FilterLogs(opts, "Delivery")
	if err != nil {
		return nil, err
	}
	return &IoTOracleContractDeliveryIterator{contract: _IoTOracleContract.contract, event: "Delivery", logs: logs, sub: sub}, nil
}

// WatchDelivery is a free log subscription operation binding the contract event 0x5ef6ad57e4a13b3646f0527fbd8571dfd05d25fe8fdec201e9da21dbda3943fd.
//
// Solidity: event Delivery()
func (_IoTOracleContract *IoTOracleContractFilterer) WatchDelivery(opts *bind.WatchOpts, sink chan<- *IoTOracleContractDelivery) (event.Subscription, error) {

	logs, sub, err := _IoTOracleContract.contract.WatchLogs(opts, "Delivery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTOracleContractDelivery)
				if err := _IoTOracleContract.contract.UnpackLog(event, "Delivery", log); err != nil {
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

// ParseDelivery is a log parse operation binding the contract event 0x5ef6ad57e4a13b3646f0527fbd8571dfd05d25fe8fdec201e9da21dbda3943fd.
//
// Solidity: event Delivery()
func (_IoTOracleContract *IoTOracleContractFilterer) ParseDelivery(log types.Log) (*IoTOracleContractDelivery, error) {
	event := new(IoTOracleContractDelivery)
	if err := _IoTOracleContract.contract.UnpackLog(event, "Delivery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IoTOracleContractRequestIoTInfoIterator is returned from FilterRequestIoTInfo and is used to iterate over the raw logs and unpacked data for RequestIoTInfo events raised by the IoTOracleContract contract.
type IoTOracleContractRequestIoTInfoIterator struct {
	Event *IoTOracleContractRequestIoTInfo // Event containing the contract specifics and raw log

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
func (it *IoTOracleContractRequestIoTInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTOracleContractRequestIoTInfo)
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
		it.Event = new(IoTOracleContractRequestIoTInfo)
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
func (it *IoTOracleContractRequestIoTInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTOracleContractRequestIoTInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTOracleContractRequestIoTInfo represents a RequestIoTInfo event raised by the IoTOracleContract contract.
type IoTOracleContractRequestIoTInfo struct {
	Id              *big.Int
	RequiredResult  bool
	IotId           []byte
	Requester       common.Address
	CallbackAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequestIoTInfo is a free log retrieval operation binding the contract event 0xbe45a1d403dfb2b6c846aec2e9c45905a86c684b22c0e031968e3b41edc39d3d.
//
// Solidity: event RequestIoTInfo(uint256 id, bool requiredResult, bytes IotId, address requester, address callbackAddress)
func (_IoTOracleContract *IoTOracleContractFilterer) FilterRequestIoTInfo(opts *bind.FilterOpts) (*IoTOracleContractRequestIoTInfoIterator, error) {

	logs, sub, err := _IoTOracleContract.contract.FilterLogs(opts, "RequestIoTInfo")
	if err != nil {
		return nil, err
	}
	return &IoTOracleContractRequestIoTInfoIterator{contract: _IoTOracleContract.contract, event: "RequestIoTInfo", logs: logs, sub: sub}, nil
}

// WatchRequestIoTInfo is a free log subscription operation binding the contract event 0xbe45a1d403dfb2b6c846aec2e9c45905a86c684b22c0e031968e3b41edc39d3d.
//
// Solidity: event RequestIoTInfo(uint256 id, bool requiredResult, bytes IotId, address requester, address callbackAddress)
func (_IoTOracleContract *IoTOracleContractFilterer) WatchRequestIoTInfo(opts *bind.WatchOpts, sink chan<- *IoTOracleContractRequestIoTInfo) (event.Subscription, error) {

	logs, sub, err := _IoTOracleContract.contract.WatchLogs(opts, "RequestIoTInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTOracleContractRequestIoTInfo)
				if err := _IoTOracleContract.contract.UnpackLog(event, "RequestIoTInfo", log); err != nil {
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

// ParseRequestIoTInfo is a log parse operation binding the contract event 0xbe45a1d403dfb2b6c846aec2e9c45905a86c684b22c0e031968e3b41edc39d3d.
//
// Solidity: event RequestIoTInfo(uint256 id, bool requiredResult, bytes IotId, address requester, address callbackAddress)
func (_IoTOracleContract *IoTOracleContractFilterer) ParseRequestIoTInfo(log types.Log) (*IoTOracleContractRequestIoTInfo, error) {
	event := new(IoTOracleContractRequestIoTInfo)
	if err := _IoTOracleContract.contract.UnpackLog(event, "RequestIoTInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
