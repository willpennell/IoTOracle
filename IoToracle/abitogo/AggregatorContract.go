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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"AggregationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"CommitsPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LogBool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogHashes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Logging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"RevealsPlaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"answers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataType\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"oracleCounter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"f\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancelFlag\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitsFlag\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealsFlag\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"lastOracle\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitHash\",\"type\":\"bytes32\"}],\"name\":\"commitResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_result\",\"type\":\"int256\"},{\"internalType\":\"bytes[]\",\"name\":\"_secret\",\"type\":\"bytes[]\"}],\"name\":\"revealAverageResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_secret\",\"type\":\"bytes\"}],\"name\":\"revealVoteResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(uint256 requestID, bytes dataType, uint256 oracleCounter, uint256 t, uint256 f, uint256 cancelFlag, uint256 commitsFlag, uint256 revealsFlag, bool lastOracle)
func (_AggregatorContract *AggregatorContractCaller) Answers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestID     *big.Int
	DataType      []byte
	OracleCounter *big.Int
	T             *big.Int
	F             *big.Int
	CancelFlag    *big.Int
	CommitsFlag   *big.Int
	RevealsFlag   *big.Int
	LastOracle    bool
}, error) {
	var out []interface{}
	err := _AggregatorContract.contract.Call(opts, &out, "answers", arg0)

	outstruct := new(struct {
		RequestID     *big.Int
		DataType      []byte
		OracleCounter *big.Int
		T             *big.Int
		F             *big.Int
		CancelFlag    *big.Int
		CommitsFlag   *big.Int
		RevealsFlag   *big.Int
		LastOracle    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataType = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.OracleCounter = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.T = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.F = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CancelFlag = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CommitsFlag = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RevealsFlag = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastOracle = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(uint256 requestID, bytes dataType, uint256 oracleCounter, uint256 t, uint256 f, uint256 cancelFlag, uint256 commitsFlag, uint256 revealsFlag, bool lastOracle)
func (_AggregatorContract *AggregatorContractSession) Answers(arg0 *big.Int) (struct {
	RequestID     *big.Int
	DataType      []byte
	OracleCounter *big.Int
	T             *big.Int
	F             *big.Int
	CancelFlag    *big.Int
	CommitsFlag   *big.Int
	RevealsFlag   *big.Int
	LastOracle    bool
}, error) {
	return _AggregatorContract.Contract.Answers(&_AggregatorContract.CallOpts, arg0)
}

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(uint256 requestID, bytes dataType, uint256 oracleCounter, uint256 t, uint256 f, uint256 cancelFlag, uint256 commitsFlag, uint256 revealsFlag, bool lastOracle)
func (_AggregatorContract *AggregatorContractCallerSession) Answers(arg0 *big.Int) (struct {
	RequestID     *big.Int
	DataType      []byte
	OracleCounter *big.Int
	T             *big.Int
	F             *big.Int
	CancelFlag    *big.Int
	CommitsFlag   *big.Int
	RevealsFlag   *big.Int
	LastOracle    bool
}, error) {
	return _AggregatorContract.Contract.Answers(&_AggregatorContract.CallOpts, arg0)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(bool b)
func (_AggregatorContract *AggregatorContractCaller) Decode(opts *bind.CallOpts, data []byte) (bool, error) {
	var out []interface{}
	err := _AggregatorContract.contract.Call(opts, &out, "decode", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(bool b)
func (_AggregatorContract *AggregatorContractSession) Decode(data []byte) (bool, error) {
	return _AggregatorContract.Contract.Decode(&_AggregatorContract.CallOpts, data)
}

// Decode is a free data retrieval call binding the contract method 0xe5c5e9a3.
//
// Solidity: function decode(bytes data) pure returns(bool b)
func (_AggregatorContract *AggregatorContractCallerSession) Decode(data []byte) (bool, error) {
	return _AggregatorContract.Contract.Decode(&_AggregatorContract.CallOpts, data)
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

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_AggregatorContract *AggregatorContractTransactor) CancelRequest(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _AggregatorContract.contract.Transact(opts, "cancelRequest", _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_AggregatorContract *AggregatorContractSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _AggregatorContract.Contract.CancelRequest(&_AggregatorContract.TransactOpts, _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_AggregatorContract *AggregatorContractTransactorSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _AggregatorContract.Contract.CancelRequest(&_AggregatorContract.TransactOpts, _requestID)
}

// CommitResponse is a paid mutator transaction binding the contract method 0x074e324d.
//
// Solidity: function commitResponse(uint256 _requestID, bytes32 _commitHash) returns(bool)
func (_AggregatorContract *AggregatorContractTransactor) CommitResponse(opts *bind.TransactOpts, _requestID *big.Int, _commitHash [32]byte) (*types.Transaction, error) {
	return _AggregatorContract.contract.Transact(opts, "commitResponse", _requestID, _commitHash)
}

// CommitResponse is a paid mutator transaction binding the contract method 0x074e324d.
//
// Solidity: function commitResponse(uint256 _requestID, bytes32 _commitHash) returns(bool)
func (_AggregatorContract *AggregatorContractSession) CommitResponse(_requestID *big.Int, _commitHash [32]byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.CommitResponse(&_AggregatorContract.TransactOpts, _requestID, _commitHash)
}

// CommitResponse is a paid mutator transaction binding the contract method 0x074e324d.
//
// Solidity: function commitResponse(uint256 _requestID, bytes32 _commitHash) returns(bool)
func (_AggregatorContract *AggregatorContractTransactorSession) CommitResponse(_requestID *big.Int, _commitHash [32]byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.CommitResponse(&_AggregatorContract.TransactOpts, _requestID, _commitHash)
}

// RevealAverageResponse is a paid mutator transaction binding the contract method 0x8259a8c1.
//
// Solidity: function revealAverageResponse(uint256 _requestID, int256 _result, bytes[] _secret) returns(bool)
func (_AggregatorContract *AggregatorContractTransactor) RevealAverageResponse(opts *bind.TransactOpts, _requestID *big.Int, _result *big.Int, _secret [][]byte) (*types.Transaction, error) {
	return _AggregatorContract.contract.Transact(opts, "revealAverageResponse", _requestID, _result, _secret)
}

// RevealAverageResponse is a paid mutator transaction binding the contract method 0x8259a8c1.
//
// Solidity: function revealAverageResponse(uint256 _requestID, int256 _result, bytes[] _secret) returns(bool)
func (_AggregatorContract *AggregatorContractSession) RevealAverageResponse(_requestID *big.Int, _result *big.Int, _secret [][]byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.RevealAverageResponse(&_AggregatorContract.TransactOpts, _requestID, _result, _secret)
}

// RevealAverageResponse is a paid mutator transaction binding the contract method 0x8259a8c1.
//
// Solidity: function revealAverageResponse(uint256 _requestID, int256 _result, bytes[] _secret) returns(bool)
func (_AggregatorContract *AggregatorContractTransactorSession) RevealAverageResponse(_requestID *big.Int, _result *big.Int, _secret [][]byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.RevealAverageResponse(&_AggregatorContract.TransactOpts, _requestID, _result, _secret)
}

// RevealVoteResponse is a paid mutator transaction binding the contract method 0x35605407.
//
// Solidity: function revealVoteResponse(uint256 _requestID, bytes _result, bytes _secret) returns()
func (_AggregatorContract *AggregatorContractTransactor) RevealVoteResponse(opts *bind.TransactOpts, _requestID *big.Int, _result []byte, _secret []byte) (*types.Transaction, error) {
	return _AggregatorContract.contract.Transact(opts, "revealVoteResponse", _requestID, _result, _secret)
}

// RevealVoteResponse is a paid mutator transaction binding the contract method 0x35605407.
//
// Solidity: function revealVoteResponse(uint256 _requestID, bytes _result, bytes _secret) returns()
func (_AggregatorContract *AggregatorContractSession) RevealVoteResponse(_requestID *big.Int, _result []byte, _secret []byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.RevealVoteResponse(&_AggregatorContract.TransactOpts, _requestID, _result, _secret)
}

// RevealVoteResponse is a paid mutator transaction binding the contract method 0x35605407.
//
// Solidity: function revealVoteResponse(uint256 _requestID, bytes _result, bytes _secret) returns()
func (_AggregatorContract *AggregatorContractTransactorSession) RevealVoteResponse(_requestID *big.Int, _result []byte, _secret []byte) (*types.Transaction, error) {
	return _AggregatorContract.Contract.RevealVoteResponse(&_AggregatorContract.TransactOpts, _requestID, _result, _secret)
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

// AggregatorContractCommitsPlacedIterator is returned from FilterCommitsPlaced and is used to iterate over the raw logs and unpacked data for CommitsPlaced events raised by the AggregatorContract contract.
type AggregatorContractCommitsPlacedIterator struct {
	Event *AggregatorContractCommitsPlaced // Event containing the contract specifics and raw log

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
func (it *AggregatorContractCommitsPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractCommitsPlaced)
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
		it.Event = new(AggregatorContractCommitsPlaced)
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
func (it *AggregatorContractCommitsPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractCommitsPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractCommitsPlaced represents a CommitsPlaced event raised by the AggregatorContract contract.
type AggregatorContractCommitsPlaced struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCommitsPlaced is a free log retrieval operation binding the contract event 0x94654b3250060d23222c9a81fdafe2f23accfb1dc46fb56b8e4aa75c0776e0fd.
//
// Solidity: event CommitsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) FilterCommitsPlaced(opts *bind.FilterOpts) (*AggregatorContractCommitsPlacedIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "CommitsPlaced")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractCommitsPlacedIterator{contract: _AggregatorContract.contract, event: "CommitsPlaced", logs: logs, sub: sub}, nil
}

// WatchCommitsPlaced is a free log subscription operation binding the contract event 0x94654b3250060d23222c9a81fdafe2f23accfb1dc46fb56b8e4aa75c0776e0fd.
//
// Solidity: event CommitsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) WatchCommitsPlaced(opts *bind.WatchOpts, sink chan<- *AggregatorContractCommitsPlaced) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "CommitsPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractCommitsPlaced)
				if err := _AggregatorContract.contract.UnpackLog(event, "CommitsPlaced", log); err != nil {
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

// ParseCommitsPlaced is a log parse operation binding the contract event 0x94654b3250060d23222c9a81fdafe2f23accfb1dc46fb56b8e4aa75c0776e0fd.
//
// Solidity: event CommitsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) ParseCommitsPlaced(log types.Log) (*AggregatorContractCommitsPlaced, error) {
	event := new(AggregatorContractCommitsPlaced)
	if err := _AggregatorContract.contract.UnpackLog(event, "CommitsPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorContractLogBoolIterator is returned from FilterLogBool and is used to iterate over the raw logs and unpacked data for LogBool events raised by the AggregatorContract contract.
type AggregatorContractLogBoolIterator struct {
	Event *AggregatorContractLogBool // Event containing the contract specifics and raw log

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
func (it *AggregatorContractLogBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractLogBool)
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
		it.Event = new(AggregatorContractLogBool)
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
func (it *AggregatorContractLogBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractLogBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractLogBool represents a LogBool event raised by the AggregatorContract contract.
type AggregatorContractLogBool struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBool is a free log retrieval operation binding the contract event 0xc33356bc2bad2ce263b056da5d061d4e89c336823d5e77f14c1383aedb7a1b3a.
//
// Solidity: event LogBool(bool arg0)
func (_AggregatorContract *AggregatorContractFilterer) FilterLogBool(opts *bind.FilterOpts) (*AggregatorContractLogBoolIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractLogBoolIterator{contract: _AggregatorContract.contract, event: "LogBool", logs: logs, sub: sub}, nil
}

// WatchLogBool is a free log subscription operation binding the contract event 0xc33356bc2bad2ce263b056da5d061d4e89c336823d5e77f14c1383aedb7a1b3a.
//
// Solidity: event LogBool(bool arg0)
func (_AggregatorContract *AggregatorContractFilterer) WatchLogBool(opts *bind.WatchOpts, sink chan<- *AggregatorContractLogBool) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractLogBool)
				if err := _AggregatorContract.contract.UnpackLog(event, "LogBool", log); err != nil {
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

// ParseLogBool is a log parse operation binding the contract event 0xc33356bc2bad2ce263b056da5d061d4e89c336823d5e77f14c1383aedb7a1b3a.
//
// Solidity: event LogBool(bool arg0)
func (_AggregatorContract *AggregatorContractFilterer) ParseLogBool(log types.Log) (*AggregatorContractLogBool, error) {
	event := new(AggregatorContractLogBool)
	if err := _AggregatorContract.contract.UnpackLog(event, "LogBool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorContractLogHashesIterator is returned from FilterLogHashes and is used to iterate over the raw logs and unpacked data for LogHashes events raised by the AggregatorContract contract.
type AggregatorContractLogHashesIterator struct {
	Event *AggregatorContractLogHashes // Event containing the contract specifics and raw log

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
func (it *AggregatorContractLogHashesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractLogHashes)
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
		it.Event = new(AggregatorContractLogHashes)
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
func (it *AggregatorContractLogHashesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractLogHashesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractLogHashes represents a LogHashes event raised by the AggregatorContract contract.
type AggregatorContractLogHashes struct {
	Arg0 [32]byte
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogHashes is a free log retrieval operation binding the contract event 0x1034809265ff4258719a9aebc8a98d3ae11c847f3db81864425e812fd387c188.
//
// Solidity: event LogHashes(bytes32 arg0, bytes32 arg1)
func (_AggregatorContract *AggregatorContractFilterer) FilterLogHashes(opts *bind.FilterOpts) (*AggregatorContractLogHashesIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "LogHashes")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractLogHashesIterator{contract: _AggregatorContract.contract, event: "LogHashes", logs: logs, sub: sub}, nil
}

// WatchLogHashes is a free log subscription operation binding the contract event 0x1034809265ff4258719a9aebc8a98d3ae11c847f3db81864425e812fd387c188.
//
// Solidity: event LogHashes(bytes32 arg0, bytes32 arg1)
func (_AggregatorContract *AggregatorContractFilterer) WatchLogHashes(opts *bind.WatchOpts, sink chan<- *AggregatorContractLogHashes) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "LogHashes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractLogHashes)
				if err := _AggregatorContract.contract.UnpackLog(event, "LogHashes", log); err != nil {
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

// ParseLogHashes is a log parse operation binding the contract event 0x1034809265ff4258719a9aebc8a98d3ae11c847f3db81864425e812fd387c188.
//
// Solidity: event LogHashes(bytes32 arg0, bytes32 arg1)
func (_AggregatorContract *AggregatorContractFilterer) ParseLogHashes(log types.Log) (*AggregatorContractLogHashes, error) {
	event := new(AggregatorContractLogHashes)
	if err := _AggregatorContract.contract.UnpackLog(event, "LogHashes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorContractLoggingIterator is returned from FilterLogging and is used to iterate over the raw logs and unpacked data for Logging events raised by the AggregatorContract contract.
type AggregatorContractLoggingIterator struct {
	Event *AggregatorContractLogging // Event containing the contract specifics and raw log

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
func (it *AggregatorContractLoggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractLogging)
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
		it.Event = new(AggregatorContractLogging)
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
func (it *AggregatorContractLoggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractLoggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractLogging represents a Logging event raised by the AggregatorContract contract.
type AggregatorContractLogging struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogging is a free log retrieval operation binding the contract event 0xb1ed47ced4562a888b3c47e728f82f7863e472cad724018fe1585bc6853ec82d.
//
// Solidity: event Logging(string arg0)
func (_AggregatorContract *AggregatorContractFilterer) FilterLogging(opts *bind.FilterOpts) (*AggregatorContractLoggingIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "Logging")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractLoggingIterator{contract: _AggregatorContract.contract, event: "Logging", logs: logs, sub: sub}, nil
}

// WatchLogging is a free log subscription operation binding the contract event 0xb1ed47ced4562a888b3c47e728f82f7863e472cad724018fe1585bc6853ec82d.
//
// Solidity: event Logging(string arg0)
func (_AggregatorContract *AggregatorContractFilterer) WatchLogging(opts *bind.WatchOpts, sink chan<- *AggregatorContractLogging) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "Logging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractLogging)
				if err := _AggregatorContract.contract.UnpackLog(event, "Logging", log); err != nil {
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

// ParseLogging is a log parse operation binding the contract event 0xb1ed47ced4562a888b3c47e728f82f7863e472cad724018fe1585bc6853ec82d.
//
// Solidity: event Logging(string arg0)
func (_AggregatorContract *AggregatorContractFilterer) ParseLogging(log types.Log) (*AggregatorContractLogging, error) {
	event := new(AggregatorContractLogging)
	if err := _AggregatorContract.contract.UnpackLog(event, "Logging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorContractRevealsPlacedIterator is returned from FilterRevealsPlaced and is used to iterate over the raw logs and unpacked data for RevealsPlaced events raised by the AggregatorContract contract.
type AggregatorContractRevealsPlacedIterator struct {
	Event *AggregatorContractRevealsPlaced // Event containing the contract specifics and raw log

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
func (it *AggregatorContractRevealsPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorContractRevealsPlaced)
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
		it.Event = new(AggregatorContractRevealsPlaced)
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
func (it *AggregatorContractRevealsPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorContractRevealsPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorContractRevealsPlaced represents a RevealsPlaced event raised by the AggregatorContract contract.
type AggregatorContractRevealsPlaced struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRevealsPlaced is a free log retrieval operation binding the contract event 0xcd9014e59ca398764ef85e5719d2151f53668c6d08f5ea98e064a65a6cccf06e.
//
// Solidity: event RevealsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) FilterRevealsPlaced(opts *bind.FilterOpts) (*AggregatorContractRevealsPlacedIterator, error) {

	logs, sub, err := _AggregatorContract.contract.FilterLogs(opts, "RevealsPlaced")
	if err != nil {
		return nil, err
	}
	return &AggregatorContractRevealsPlacedIterator{contract: _AggregatorContract.contract, event: "RevealsPlaced", logs: logs, sub: sub}, nil
}

// WatchRevealsPlaced is a free log subscription operation binding the contract event 0xcd9014e59ca398764ef85e5719d2151f53668c6d08f5ea98e064a65a6cccf06e.
//
// Solidity: event RevealsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) WatchRevealsPlaced(opts *bind.WatchOpts, sink chan<- *AggregatorContractRevealsPlaced) (event.Subscription, error) {

	logs, sub, err := _AggregatorContract.contract.WatchLogs(opts, "RevealsPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorContractRevealsPlaced)
				if err := _AggregatorContract.contract.UnpackLog(event, "RevealsPlaced", log); err != nil {
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

// ParseRevealsPlaced is a log parse operation binding the contract event 0xcd9014e59ca398764ef85e5719d2151f53668c6d08f5ea98e064a65a6cccf06e.
//
// Solidity: event RevealsPlaced(uint256 arg0, string arg1)
func (_AggregatorContract *AggregatorContractFilterer) ParseRevealsPlaced(log types.Log) (*AggregatorContractRevealsPlaced, error) {
	event := new(AggregatorContractRevealsPlaced)
	if err := _AggregatorContract.contract.UnpackLog(event, "RevealsPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
