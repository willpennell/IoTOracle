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

// ReputationContractMetaData contains all meta data concerning the ReputationContract contract.
var ReputationContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getOracleRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPenaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"incrementRating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracleRatings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReputationContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationContractMetaData.ABI instead.
var ReputationContractABI = ReputationContractMetaData.ABI

// ReputationContract is an auto generated Go binding around an Ethereum contract.
type ReputationContract struct {
	ReputationContractCaller     // Read-only binding to the contract
	ReputationContractTransactor // Write-only binding to the contract
	ReputationContractFilterer   // Log filterer for contract events
}

// ReputationContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationContractSession struct {
	Contract     *ReputationContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReputationContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationContractCallerSession struct {
	Contract *ReputationContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ReputationContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationContractTransactorSession struct {
	Contract     *ReputationContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ReputationContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationContractRaw struct {
	Contract *ReputationContract // Generic contract binding to access the raw methods on
}

// ReputationContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationContractCallerRaw struct {
	Contract *ReputationContractCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationContractTransactorRaw struct {
	Contract *ReputationContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputationContract creates a new instance of ReputationContract, bound to a specific deployed contract.
func NewReputationContract(address common.Address, backend bind.ContractBackend) (*ReputationContract, error) {
	contract, err := bindReputationContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReputationContract{ReputationContractCaller: ReputationContractCaller{contract: contract}, ReputationContractTransactor: ReputationContractTransactor{contract: contract}, ReputationContractFilterer: ReputationContractFilterer{contract: contract}}, nil
}

// NewReputationContractCaller creates a new read-only instance of ReputationContract, bound to a specific deployed contract.
func NewReputationContractCaller(address common.Address, caller bind.ContractCaller) (*ReputationContractCaller, error) {
	contract, err := bindReputationContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationContractCaller{contract: contract}, nil
}

// NewReputationContractTransactor creates a new write-only instance of ReputationContract, bound to a specific deployed contract.
func NewReputationContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationContractTransactor, error) {
	contract, err := bindReputationContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationContractTransactor{contract: contract}, nil
}

// NewReputationContractFilterer creates a new log filterer instance of ReputationContract, bound to a specific deployed contract.
func NewReputationContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationContractFilterer, error) {
	contract, err := bindReputationContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationContractFilterer{contract: contract}, nil
}

// bindReputationContract binds a generic wrapper to an already deployed contract.
func bindReputationContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReputationContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationContract *ReputationContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationContract.Contract.ReputationContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationContract *ReputationContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationContract.Contract.ReputationContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationContract *ReputationContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationContract.Contract.ReputationContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReputationContract *ReputationContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReputationContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReputationContract *ReputationContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReputationContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReputationContract *ReputationContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReputationContract.Contract.contract.Transact(opts, method, params...)
}

// GetOracleRating is a free data retrieval call binding the contract method 0x0855078b.
//
// Solidity: function getOracleRating(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractCaller) GetOracleRating(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReputationContract.contract.Call(opts, &out, "getOracleRating", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOracleRating is a free data retrieval call binding the contract method 0x0855078b.
//
// Solidity: function getOracleRating(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractSession) GetOracleRating(_addr common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.GetOracleRating(&_ReputationContract.CallOpts, _addr)
}

// GetOracleRating is a free data retrieval call binding the contract method 0x0855078b.
//
// Solidity: function getOracleRating(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractCallerSession) GetOracleRating(_addr common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.GetOracleRating(&_ReputationContract.CallOpts, _addr)
}

// GetPenaltyFee is a free data retrieval call binding the contract method 0xb8f9c2a1.
//
// Solidity: function getPenaltyFee(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractCaller) GetPenaltyFee(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReputationContract.contract.Call(opts, &out, "getPenaltyFee", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPenaltyFee is a free data retrieval call binding the contract method 0xb8f9c2a1.
//
// Solidity: function getPenaltyFee(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractSession) GetPenaltyFee(_addr common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.GetPenaltyFee(&_ReputationContract.CallOpts, _addr)
}

// GetPenaltyFee is a free data retrieval call binding the contract method 0xb8f9c2a1.
//
// Solidity: function getPenaltyFee(address _addr) view returns(uint256)
func (_ReputationContract *ReputationContractCallerSession) GetPenaltyFee(_addr common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.GetPenaltyFee(&_ReputationContract.CallOpts, _addr)
}

// OracleRatings is a free data retrieval call binding the contract method 0x1d803317.
//
// Solidity: function oracleRatings(address ) view returns(uint256)
func (_ReputationContract *ReputationContractCaller) OracleRatings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReputationContract.contract.Call(opts, &out, "oracleRatings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleRatings is a free data retrieval call binding the contract method 0x1d803317.
//
// Solidity: function oracleRatings(address ) view returns(uint256)
func (_ReputationContract *ReputationContractSession) OracleRatings(arg0 common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.OracleRatings(&_ReputationContract.CallOpts, arg0)
}

// OracleRatings is a free data retrieval call binding the contract method 0x1d803317.
//
// Solidity: function oracleRatings(address ) view returns(uint256)
func (_ReputationContract *ReputationContractCallerSession) OracleRatings(arg0 common.Address) (*big.Int, error) {
	return _ReputationContract.Contract.OracleRatings(&_ReputationContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationContract *ReputationContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReputationContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationContract *ReputationContractSession) Owner() (common.Address, error) {
	return _ReputationContract.Contract.Owner(&_ReputationContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ReputationContract *ReputationContractCallerSession) Owner() (common.Address, error) {
	return _ReputationContract.Contract.Owner(&_ReputationContract.CallOpts)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _addr) returns(bool)
func (_ReputationContract *ReputationContractTransactor) AddOracle(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.contract.Transact(opts, "addOracle", _addr)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _addr) returns(bool)
func (_ReputationContract *ReputationContractSession) AddOracle(_addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.Contract.AddOracle(&_ReputationContract.TransactOpts, _addr)
}

// AddOracle is a paid mutator transaction binding the contract method 0xdf5dd1a5.
//
// Solidity: function addOracle(address _addr) returns(bool)
func (_ReputationContract *ReputationContractTransactorSession) AddOracle(_addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.Contract.AddOracle(&_ReputationContract.TransactOpts, _addr)
}

// IncrementRating is a paid mutator transaction binding the contract method 0x18b529c3.
//
// Solidity: function incrementRating(address _addr) returns(uint256)
func (_ReputationContract *ReputationContractTransactor) IncrementRating(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.contract.Transact(opts, "incrementRating", _addr)
}

// IncrementRating is a paid mutator transaction binding the contract method 0x18b529c3.
//
// Solidity: function incrementRating(address _addr) returns(uint256)
func (_ReputationContract *ReputationContractSession) IncrementRating(_addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.Contract.IncrementRating(&_ReputationContract.TransactOpts, _addr)
}

// IncrementRating is a paid mutator transaction binding the contract method 0x18b529c3.
//
// Solidity: function incrementRating(address _addr) returns(uint256)
func (_ReputationContract *ReputationContractTransactorSession) IncrementRating(_addr common.Address) (*types.Transaction, error) {
	return _ReputationContract.Contract.IncrementRating(&_ReputationContract.TransactOpts, _addr)
}

// ReputationContractOracleBlacklistedIterator is returned from FilterOracleBlacklisted and is used to iterate over the raw logs and unpacked data for OracleBlacklisted events raised by the ReputationContract contract.
type ReputationContractOracleBlacklistedIterator struct {
	Event *ReputationContractOracleBlacklisted // Event containing the contract specifics and raw log

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
func (it *ReputationContractOracleBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationContractOracleBlacklisted)
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
		it.Event = new(ReputationContractOracleBlacklisted)
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
func (it *ReputationContractOracleBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationContractOracleBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationContractOracleBlacklisted represents a OracleBlacklisted event raised by the ReputationContract contract.
type ReputationContractOracleBlacklisted struct {
	Arg0 common.Address
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOracleBlacklisted is a free log retrieval operation binding the contract event 0xdffbba745a852c1b2b282e3aa033195de3b054875aab03a94fdc488e7bed956c.
//
// Solidity: event OracleBlacklisted(address arg0, string arg1)
func (_ReputationContract *ReputationContractFilterer) FilterOracleBlacklisted(opts *bind.FilterOpts) (*ReputationContractOracleBlacklistedIterator, error) {

	logs, sub, err := _ReputationContract.contract.FilterLogs(opts, "OracleBlacklisted")
	if err != nil {
		return nil, err
	}
	return &ReputationContractOracleBlacklistedIterator{contract: _ReputationContract.contract, event: "OracleBlacklisted", logs: logs, sub: sub}, nil
}

// WatchOracleBlacklisted is a free log subscription operation binding the contract event 0xdffbba745a852c1b2b282e3aa033195de3b054875aab03a94fdc488e7bed956c.
//
// Solidity: event OracleBlacklisted(address arg0, string arg1)
func (_ReputationContract *ReputationContractFilterer) WatchOracleBlacklisted(opts *bind.WatchOpts, sink chan<- *ReputationContractOracleBlacklisted) (event.Subscription, error) {

	logs, sub, err := _ReputationContract.contract.WatchLogs(opts, "OracleBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationContractOracleBlacklisted)
				if err := _ReputationContract.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
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

// ParseOracleBlacklisted is a log parse operation binding the contract event 0xdffbba745a852c1b2b282e3aa033195de3b054875aab03a94fdc488e7bed956c.
//
// Solidity: event OracleBlacklisted(address arg0, string arg1)
func (_ReputationContract *ReputationContractFilterer) ParseOracleBlacklisted(log types.Log) (*ReputationContractOracleBlacklisted, error) {
	event := new(ReputationContractOracleBlacklisted)
	if err := _ReputationContract.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
