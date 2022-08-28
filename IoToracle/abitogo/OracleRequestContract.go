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

// OracleRequestContractMetaData contains all meta data concerning the OracleRequestContract contract.
var OracleRequestContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"FinalInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"FinalResultLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"LogNumberOfOracles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"LogTimeOutLen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Logging\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OpenForBids\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OracleJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OracleLeft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OraclePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"PrintAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"PrintAddresses\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ReleaseRequestDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumOracleRequestContract.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"StatusChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UpdateTimestamp\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"aggregatorAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"blacklistOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklistedOracles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_incorrectOracleReveals\",\"type\":\"address[]\"}],\"name\":\"cancelRequestDueToDeadlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"completedRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackFID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_IoTID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_dataType\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_aggregationType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_numberOfOracles\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_elapsedTime\",\"type\":\"uint256\"}],\"name\":\"createRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_finalResult\",\"type\":\"int256\"},{\"internalType\":\"address[]\",\"name\":\"_correctOracles\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_incorrectOracles\",\"type\":\"address[]\"}],\"name\":\"deliverAverageResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_finalResult\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_correctOracles\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_incorrectOracles\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_incorrectRevealOracles\",\"type\":\"address[]\"}],\"name\":\"deliverVoteResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getAggregationType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getBlacklistedOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getDataType\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getNumberOfOracles\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getOracleAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_orcAddr\",\"type\":\"address\"}],\"name\":\"getOracleForRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getTimeoutOracleLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinAsOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveOracleNetwork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputationAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackFID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"IoTID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"dataType\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"numberOfOracles\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"aggregationType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"oracleCounter\",\"type\":\"uint32\"},{\"internalType\":\"enumOracleRequestContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"cancelFlag\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggAddr\",\"type\":\"address\"}],\"name\":\"setAggregatorContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_repAddr\",\"type\":\"address\"}],\"name\":\"setReputationContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"timeoutBiddingAppeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"timeoutCommitsAppeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"timeoutRevealsAppeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"appealFlag\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"elapsedTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OracleRequestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleRequestContractMetaData.ABI instead.
var OracleRequestContractABI = OracleRequestContractMetaData.ABI

// OracleRequestContract is an auto generated Go binding around an Ethereum contract.
type OracleRequestContract struct {
	OracleRequestContractCaller     // Read-only binding to the contract
	OracleRequestContractTransactor // Write-only binding to the contract
	OracleRequestContractFilterer   // Log filterer for contract events
}

// OracleRequestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleRequestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleRequestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleRequestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleRequestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleRequestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleRequestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleRequestContractSession struct {
	Contract     *OracleRequestContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OracleRequestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleRequestContractCallerSession struct {
	Contract *OracleRequestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OracleRequestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleRequestContractTransactorSession struct {
	Contract     *OracleRequestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OracleRequestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRequestContractRaw struct {
	Contract *OracleRequestContract // Generic contract binding to access the raw methods on
}

// OracleRequestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleRequestContractCallerRaw struct {
	Contract *OracleRequestContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleRequestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleRequestContractTransactorRaw struct {
	Contract *OracleRequestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleRequestContract creates a new instance of OracleRequestContract, bound to a specific deployed contract.
func NewOracleRequestContract(address common.Address, backend bind.ContractBackend) (*OracleRequestContract, error) {
	contract, err := bindOracleRequestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleRequestContract{OracleRequestContractCaller: OracleRequestContractCaller{contract: contract}, OracleRequestContractTransactor: OracleRequestContractTransactor{contract: contract}, OracleRequestContractFilterer: OracleRequestContractFilterer{contract: contract}}, nil
}

// NewOracleRequestContractCaller creates a new read-only instance of OracleRequestContract, bound to a specific deployed contract.
func NewOracleRequestContractCaller(address common.Address, caller bind.ContractCaller) (*OracleRequestContractCaller, error) {
	contract, err := bindOracleRequestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractCaller{contract: contract}, nil
}

// NewOracleRequestContractTransactor creates a new write-only instance of OracleRequestContract, bound to a specific deployed contract.
func NewOracleRequestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleRequestContractTransactor, error) {
	contract, err := bindOracleRequestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractTransactor{contract: contract}, nil
}

// NewOracleRequestContractFilterer creates a new log filterer instance of OracleRequestContract, bound to a specific deployed contract.
func NewOracleRequestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleRequestContractFilterer, error) {
	contract, err := bindOracleRequestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractFilterer{contract: contract}, nil
}

// bindOracleRequestContract binds a generic wrapper to an already deployed contract.
func bindOracleRequestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleRequestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleRequestContract *OracleRequestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleRequestContract.Contract.OracleRequestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleRequestContract *OracleRequestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.OracleRequestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleRequestContract *OracleRequestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.OracleRequestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleRequestContract *OracleRequestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleRequestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleRequestContract *OracleRequestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleRequestContract *OracleRequestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.contract.Transact(opts, method, params...)
}

// AggregatorAddr is a free data retrieval call binding the contract method 0x82762600.
//
// Solidity: function aggregatorAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractCaller) AggregatorAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "aggregatorAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggregatorAddr is a free data retrieval call binding the contract method 0x82762600.
//
// Solidity: function aggregatorAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractSession) AggregatorAddr() (common.Address, error) {
	return _OracleRequestContract.Contract.AggregatorAddr(&_OracleRequestContract.CallOpts)
}

// AggregatorAddr is a free data retrieval call binding the contract method 0x82762600.
//
// Solidity: function aggregatorAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractCallerSession) AggregatorAddr() (common.Address, error) {
	return _OracleRequestContract.Contract.AggregatorAddr(&_OracleRequestContract.CallOpts)
}

// BlacklistedOracles is a free data retrieval call binding the contract method 0x7d02da09.
//
// Solidity: function blacklistedOracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCaller) BlacklistedOracles(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "blacklistedOracles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlacklistedOracles is a free data retrieval call binding the contract method 0x7d02da09.
//
// Solidity: function blacklistedOracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) BlacklistedOracles(arg0 common.Address) (bool, error) {
	return _OracleRequestContract.Contract.BlacklistedOracles(&_OracleRequestContract.CallOpts, arg0)
}

// BlacklistedOracles is a free data retrieval call binding the contract method 0x7d02da09.
//
// Solidity: function blacklistedOracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCallerSession) BlacklistedOracles(arg0 common.Address) (bool, error) {
	return _OracleRequestContract.Contract.BlacklistedOracles(&_OracleRequestContract.CallOpts, arg0)
}

// CompletedRequests is a free data retrieval call binding the contract method 0xf0fbfbf3.
//
// Solidity: function completedRequests(uint256 ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCaller) CompletedRequests(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "completedRequests", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompletedRequests is a free data retrieval call binding the contract method 0xf0fbfbf3.
//
// Solidity: function completedRequests(uint256 ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) CompletedRequests(arg0 *big.Int) (bool, error) {
	return _OracleRequestContract.Contract.CompletedRequests(&_OracleRequestContract.CallOpts, arg0)
}

// CompletedRequests is a free data retrieval call binding the contract method 0xf0fbfbf3.
//
// Solidity: function completedRequests(uint256 ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCallerSession) CompletedRequests(arg0 *big.Int) (bool, error) {
	return _OracleRequestContract.Contract.CompletedRequests(&_OracleRequestContract.CallOpts, arg0)
}

// GetAggregationType is a free data retrieval call binding the contract method 0xe4888222.
//
// Solidity: function getAggregationType(uint256 _requestID) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractCaller) GetAggregationType(opts *bind.CallOpts, _requestID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getAggregationType", _requestID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggregationType is a free data retrieval call binding the contract method 0xe4888222.
//
// Solidity: function getAggregationType(uint256 _requestID) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractSession) GetAggregationType(_requestID *big.Int) (*big.Int, error) {
	return _OracleRequestContract.Contract.GetAggregationType(&_OracleRequestContract.CallOpts, _requestID)
}

// GetAggregationType is a free data retrieval call binding the contract method 0xe4888222.
//
// Solidity: function getAggregationType(uint256 _requestID) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetAggregationType(_requestID *big.Int) (*big.Int, error) {
	return _OracleRequestContract.Contract.GetAggregationType(&_OracleRequestContract.CallOpts, _requestID)
}

// GetBlacklistedOracle is a free data retrieval call binding the contract method 0x6fbf7973.
//
// Solidity: function getBlacklistedOracle(address _addr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCaller) GetBlacklistedOracle(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getBlacklistedOracle", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlacklistedOracle is a free data retrieval call binding the contract method 0x6fbf7973.
//
// Solidity: function getBlacklistedOracle(address _addr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) GetBlacklistedOracle(_addr common.Address) (bool, error) {
	return _OracleRequestContract.Contract.GetBlacklistedOracle(&_OracleRequestContract.CallOpts, _addr)
}

// GetBlacklistedOracle is a free data retrieval call binding the contract method 0x6fbf7973.
//
// Solidity: function getBlacklistedOracle(address _addr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetBlacklistedOracle(_addr common.Address) (bool, error) {
	return _OracleRequestContract.Contract.GetBlacklistedOracle(&_OracleRequestContract.CallOpts, _addr)
}

// GetDataType is a free data retrieval call binding the contract method 0x57442e79.
//
// Solidity: function getDataType(uint256 _requestID) view returns(bytes)
func (_OracleRequestContract *OracleRequestContractCaller) GetDataType(opts *bind.CallOpts, _requestID *big.Int) ([]byte, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getDataType", _requestID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDataType is a free data retrieval call binding the contract method 0x57442e79.
//
// Solidity: function getDataType(uint256 _requestID) view returns(bytes)
func (_OracleRequestContract *OracleRequestContractSession) GetDataType(_requestID *big.Int) ([]byte, error) {
	return _OracleRequestContract.Contract.GetDataType(&_OracleRequestContract.CallOpts, _requestID)
}

// GetDataType is a free data retrieval call binding the contract method 0x57442e79.
//
// Solidity: function getDataType(uint256 _requestID) view returns(bytes)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetDataType(_requestID *big.Int) ([]byte, error) {
	return _OracleRequestContract.Contract.GetDataType(&_OracleRequestContract.CallOpts, _requestID)
}

// GetNumberOfOracles is a free data retrieval call binding the contract method 0x91a3632a.
//
// Solidity: function getNumberOfOracles(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractCaller) GetNumberOfOracles(opts *bind.CallOpts, _requestID *big.Int) (uint32, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getNumberOfOracles", _requestID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetNumberOfOracles is a free data retrieval call binding the contract method 0x91a3632a.
//
// Solidity: function getNumberOfOracles(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractSession) GetNumberOfOracles(_requestID *big.Int) (uint32, error) {
	return _OracleRequestContract.Contract.GetNumberOfOracles(&_OracleRequestContract.CallOpts, _requestID)
}

// GetNumberOfOracles is a free data retrieval call binding the contract method 0x91a3632a.
//
// Solidity: function getNumberOfOracles(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetNumberOfOracles(_requestID *big.Int) (uint32, error) {
	return _OracleRequestContract.Contract.GetNumberOfOracles(&_OracleRequestContract.CallOpts, _requestID)
}

// GetOracleAddresses is a free data retrieval call binding the contract method 0xc26cc94f.
//
// Solidity: function getOracleAddresses(uint256 _requestID) view returns(address[])
func (_OracleRequestContract *OracleRequestContractCaller) GetOracleAddresses(opts *bind.CallOpts, _requestID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getOracleAddresses", _requestID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracleAddresses is a free data retrieval call binding the contract method 0xc26cc94f.
//
// Solidity: function getOracleAddresses(uint256 _requestID) view returns(address[])
func (_OracleRequestContract *OracleRequestContractSession) GetOracleAddresses(_requestID *big.Int) ([]common.Address, error) {
	return _OracleRequestContract.Contract.GetOracleAddresses(&_OracleRequestContract.CallOpts, _requestID)
}

// GetOracleAddresses is a free data retrieval call binding the contract method 0xc26cc94f.
//
// Solidity: function getOracleAddresses(uint256 _requestID) view returns(address[])
func (_OracleRequestContract *OracleRequestContractCallerSession) GetOracleAddresses(_requestID *big.Int) ([]common.Address, error) {
	return _OracleRequestContract.Contract.GetOracleAddresses(&_OracleRequestContract.CallOpts, _requestID)
}

// GetOracleForRequest is a free data retrieval call binding the contract method 0xa37b354d.
//
// Solidity: function getOracleForRequest(uint256 _requestID, address _orcAddr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCaller) GetOracleForRequest(opts *bind.CallOpts, _requestID *big.Int, _orcAddr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getOracleForRequest", _requestID, _orcAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOracleForRequest is a free data retrieval call binding the contract method 0xa37b354d.
//
// Solidity: function getOracleForRequest(uint256 _requestID, address _orcAddr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) GetOracleForRequest(_requestID *big.Int, _orcAddr common.Address) (bool, error) {
	return _OracleRequestContract.Contract.GetOracleForRequest(&_OracleRequestContract.CallOpts, _requestID, _orcAddr)
}

// GetOracleForRequest is a free data retrieval call binding the contract method 0xa37b354d.
//
// Solidity: function getOracleForRequest(uint256 _requestID, address _orcAddr) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetOracleForRequest(_requestID *big.Int, _orcAddr common.Address) (bool, error) {
	return _OracleRequestContract.Contract.GetOracleForRequest(&_OracleRequestContract.CallOpts, _requestID, _orcAddr)
}

// GetTimeoutOracleLength is a free data retrieval call binding the contract method 0x1c6fd0fc.
//
// Solidity: function getTimeoutOracleLength(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractCaller) GetTimeoutOracleLength(opts *bind.CallOpts, _requestID *big.Int) (uint32, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getTimeoutOracleLength", _requestID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTimeoutOracleLength is a free data retrieval call binding the contract method 0x1c6fd0fc.
//
// Solidity: function getTimeoutOracleLength(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractSession) GetTimeoutOracleLength(_requestID *big.Int) (uint32, error) {
	return _OracleRequestContract.Contract.GetTimeoutOracleLength(&_OracleRequestContract.CallOpts, _requestID)
}

// GetTimeoutOracleLength is a free data retrieval call binding the contract method 0x1c6fd0fc.
//
// Solidity: function getTimeoutOracleLength(uint256 _requestID) view returns(uint32)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetTimeoutOracleLength(_requestID *big.Int) (uint32, error) {
	return _OracleRequestContract.Contract.GetTimeoutOracleLength(&_OracleRequestContract.CallOpts, _requestID)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCaller) Oracles(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) Oracles(arg0 common.Address) (bool, error) {
	return _OracleRequestContract.Contract.Oracles(&_OracleRequestContract.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xaddd5099.
//
// Solidity: function oracles(address ) view returns(bool)
func (_OracleRequestContract *OracleRequestContractCallerSession) Oracles(arg0 common.Address) (bool, error) {
	return _OracleRequestContract.Contract.Oracles(&_OracleRequestContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleRequestContract *OracleRequestContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleRequestContract *OracleRequestContractSession) Owner() (common.Address, error) {
	return _OracleRequestContract.Contract.Owner(&_OracleRequestContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleRequestContract *OracleRequestContractCallerSession) Owner() (common.Address, error) {
	return _OracleRequestContract.Contract.Owner(&_OracleRequestContract.CallOpts)
}

// ReputationAddr is a free data retrieval call binding the contract method 0x053ebc9a.
//
// Solidity: function reputationAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractCaller) ReputationAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "reputationAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReputationAddr is a free data retrieval call binding the contract method 0x053ebc9a.
//
// Solidity: function reputationAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractSession) ReputationAddr() (common.Address, error) {
	return _OracleRequestContract.Contract.ReputationAddr(&_OracleRequestContract.CallOpts)
}

// ReputationAddr is a free data retrieval call binding the contract method 0x053ebc9a.
//
// Solidity: function reputationAddr() view returns(address)
func (_OracleRequestContract *OracleRequestContractCallerSession) ReputationAddr() (common.Address, error) {
	return _OracleRequestContract.Contract.ReputationAddr(&_OracleRequestContract.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, uint32 numberOfOracles, uint256 aggregationType, uint32 oracleCounter, uint8 status, uint8 cancelFlag, uint256 fee)
func (_OracleRequestContract *OracleRequestContractCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	NumberOfOracles uint32
	AggregationType *big.Int
	OracleCounter   uint32
	Status          uint8
	CancelFlag      uint8
	Fee             *big.Int
}, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		RequestID       *big.Int
		Requester       common.Address
		CallbackAddress common.Address
		CallbackFID     []byte
		IoTID           []byte
		DataType        []byte
		NumberOfOracles uint32
		AggregationType *big.Int
		OracleCounter   uint32
		Status          uint8
		CancelFlag      uint8
		Fee             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CallbackAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CallbackFID = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.IoTID = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.DataType = *abi.ConvertType(out[5], new([]byte)).(*[]byte)
	outstruct.NumberOfOracles = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.AggregationType = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OracleCounter = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.CancelFlag = *abi.ConvertType(out[10], new(uint8)).(*uint8)
	outstruct.Fee = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, uint32 numberOfOracles, uint256 aggregationType, uint32 oracleCounter, uint8 status, uint8 cancelFlag, uint256 fee)
func (_OracleRequestContract *OracleRequestContractSession) Requests(arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	NumberOfOracles uint32
	AggregationType *big.Int
	OracleCounter   uint32
	Status          uint8
	CancelFlag      uint8
	Fee             *big.Int
}, error) {
	return _OracleRequestContract.Contract.Requests(&_OracleRequestContract.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, uint32 numberOfOracles, uint256 aggregationType, uint32 oracleCounter, uint8 status, uint8 cancelFlag, uint256 fee)
func (_OracleRequestContract *OracleRequestContractCallerSession) Requests(arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	NumberOfOracles uint32
	AggregationType *big.Int
	OracleCounter   uint32
	Status          uint8
	CancelFlag      uint8
	Fee             *big.Int
}, error) {
	return _OracleRequestContract.Contract.Requests(&_OracleRequestContract.CallOpts, arg0)
}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractCaller) StakeBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "stakeBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractSession) StakeBalance(arg0 common.Address) (*big.Int, error) {
	return _OracleRequestContract.Contract.StakeBalance(&_OracleRequestContract.CallOpts, arg0)
}

// StakeBalance is a free data retrieval call binding the contract method 0x4e7c57a6.
//
// Solidity: function stakeBalance(address ) view returns(uint256)
func (_OracleRequestContract *OracleRequestContractCallerSession) StakeBalance(arg0 common.Address) (*big.Int, error) {
	return _OracleRequestContract.Contract.StakeBalance(&_OracleRequestContract.CallOpts, arg0)
}

// Timings is a free data retrieval call binding the contract method 0x9610f691.
//
// Solidity: function timings(uint256 ) view returns(uint256 appealFlag, uint256 timestamp, uint256 elapsedTime)
func (_OracleRequestContract *OracleRequestContractCaller) Timings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AppealFlag  *big.Int
	Timestamp   *big.Int
	ElapsedTime *big.Int
}, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "timings", arg0)

	outstruct := new(struct {
		AppealFlag  *big.Int
		Timestamp   *big.Int
		ElapsedTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppealFlag = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ElapsedTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Timings is a free data retrieval call binding the contract method 0x9610f691.
//
// Solidity: function timings(uint256 ) view returns(uint256 appealFlag, uint256 timestamp, uint256 elapsedTime)
func (_OracleRequestContract *OracleRequestContractSession) Timings(arg0 *big.Int) (struct {
	AppealFlag  *big.Int
	Timestamp   *big.Int
	ElapsedTime *big.Int
}, error) {
	return _OracleRequestContract.Contract.Timings(&_OracleRequestContract.CallOpts, arg0)
}

// Timings is a free data retrieval call binding the contract method 0x9610f691.
//
// Solidity: function timings(uint256 ) view returns(uint256 appealFlag, uint256 timestamp, uint256 elapsedTime)
func (_OracleRequestContract *OracleRequestContractCallerSession) Timings(arg0 *big.Int) (struct {
	AppealFlag  *big.Int
	Timestamp   *big.Int
	ElapsedTime *big.Int
}, error) {
	return _OracleRequestContract.Contract.Timings(&_OracleRequestContract.CallOpts, arg0)
}

// BlacklistOracle is a paid mutator transaction binding the contract method 0xeee5a596.
//
// Solidity: function blacklistOracle(address _addr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) BlacklistOracle(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "blacklistOracle", _addr)
}

// BlacklistOracle is a paid mutator transaction binding the contract method 0xeee5a596.
//
// Solidity: function blacklistOracle(address _addr) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) BlacklistOracle(_addr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.BlacklistOracle(&_OracleRequestContract.TransactOpts, _addr)
}

// BlacklistOracle is a paid mutator transaction binding the contract method 0xeee5a596.
//
// Solidity: function blacklistOracle(address _addr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) BlacklistOracle(_addr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.BlacklistOracle(&_OracleRequestContract.TransactOpts, _addr)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) CancelRequest(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "cancelRequest", _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequest(&_OracleRequestContract.TransactOpts, _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequest(&_OracleRequestContract.TransactOpts, _requestID)
}

// CancelRequestDueToDeadlock is a paid mutator transaction binding the contract method 0x498d7df0.
//
// Solidity: function cancelRequestDueToDeadlock(uint256 _requestID, address[] _incorrectOracleReveals) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) CancelRequestDueToDeadlock(opts *bind.TransactOpts, _requestID *big.Int, _incorrectOracleReveals []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "cancelRequestDueToDeadlock", _requestID, _incorrectOracleReveals)
}

// CancelRequestDueToDeadlock is a paid mutator transaction binding the contract method 0x498d7df0.
//
// Solidity: function cancelRequestDueToDeadlock(uint256 _requestID, address[] _incorrectOracleReveals) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) CancelRequestDueToDeadlock(_requestID *big.Int, _incorrectOracleReveals []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequestDueToDeadlock(&_OracleRequestContract.TransactOpts, _requestID, _incorrectOracleReveals)
}

// CancelRequestDueToDeadlock is a paid mutator transaction binding the contract method 0x498d7df0.
//
// Solidity: function cancelRequestDueToDeadlock(uint256 _requestID, address[] _incorrectOracleReveals) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) CancelRequestDueToDeadlock(_requestID *big.Int, _incorrectOracleReveals []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequestDueToDeadlock(&_OracleRequestContract.TransactOpts, _requestID, _incorrectOracleReveals)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x301f1dbd.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, uint256 _aggregationType, uint32 _numberOfOracles, uint256 _elapsedTime) payable returns(uint256)
func (_OracleRequestContract *OracleRequestContractTransactor) CreateRequest(opts *bind.TransactOpts, _callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _aggregationType *big.Int, _numberOfOracles uint32, _elapsedTime *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "createRequest", _callbackAddress, _callbackFID, _IoTID, _dataType, _aggregationType, _numberOfOracles, _elapsedTime)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x301f1dbd.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, uint256 _aggregationType, uint32 _numberOfOracles, uint256 _elapsedTime) payable returns(uint256)
func (_OracleRequestContract *OracleRequestContractSession) CreateRequest(_callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _aggregationType *big.Int, _numberOfOracles uint32, _elapsedTime *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CreateRequest(&_OracleRequestContract.TransactOpts, _callbackAddress, _callbackFID, _IoTID, _dataType, _aggregationType, _numberOfOracles, _elapsedTime)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x301f1dbd.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, uint256 _aggregationType, uint32 _numberOfOracles, uint256 _elapsedTime) payable returns(uint256)
func (_OracleRequestContract *OracleRequestContractTransactorSession) CreateRequest(_callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _aggregationType *big.Int, _numberOfOracles uint32, _elapsedTime *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CreateRequest(&_OracleRequestContract.TransactOpts, _callbackAddress, _callbackFID, _IoTID, _dataType, _aggregationType, _numberOfOracles, _elapsedTime)
}

// DeliverAverageResponse is a paid mutator transaction binding the contract method 0x58621b7a.
//
// Solidity: function deliverAverageResponse(uint256 _requestID, int256 _finalResult, address[] _correctOracles, address[] _incorrectOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) DeliverAverageResponse(opts *bind.TransactOpts, _requestID *big.Int, _finalResult *big.Int, _correctOracles []common.Address, _incorrectOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "deliverAverageResponse", _requestID, _finalResult, _correctOracles, _incorrectOracles)
}

// DeliverAverageResponse is a paid mutator transaction binding the contract method 0x58621b7a.
//
// Solidity: function deliverAverageResponse(uint256 _requestID, int256 _finalResult, address[] _correctOracles, address[] _incorrectOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) DeliverAverageResponse(_requestID *big.Int, _finalResult *big.Int, _correctOracles []common.Address, _incorrectOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverAverageResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult, _correctOracles, _incorrectOracles)
}

// DeliverAverageResponse is a paid mutator transaction binding the contract method 0x58621b7a.
//
// Solidity: function deliverAverageResponse(uint256 _requestID, int256 _finalResult, address[] _correctOracles, address[] _incorrectOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) DeliverAverageResponse(_requestID *big.Int, _finalResult *big.Int, _correctOracles []common.Address, _incorrectOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverAverageResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult, _correctOracles, _incorrectOracles)
}

// DeliverVoteResponse is a paid mutator transaction binding the contract method 0x30407ad9.
//
// Solidity: function deliverVoteResponse(uint256 _requestID, bool _finalResult, address[] _correctOracles, address[] _incorrectOracles, address[] _incorrectRevealOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) DeliverVoteResponse(opts *bind.TransactOpts, _requestID *big.Int, _finalResult bool, _correctOracles []common.Address, _incorrectOracles []common.Address, _incorrectRevealOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "deliverVoteResponse", _requestID, _finalResult, _correctOracles, _incorrectOracles, _incorrectRevealOracles)
}

// DeliverVoteResponse is a paid mutator transaction binding the contract method 0x30407ad9.
//
// Solidity: function deliverVoteResponse(uint256 _requestID, bool _finalResult, address[] _correctOracles, address[] _incorrectOracles, address[] _incorrectRevealOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) DeliverVoteResponse(_requestID *big.Int, _finalResult bool, _correctOracles []common.Address, _incorrectOracles []common.Address, _incorrectRevealOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverVoteResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult, _correctOracles, _incorrectOracles, _incorrectRevealOracles)
}

// DeliverVoteResponse is a paid mutator transaction binding the contract method 0x30407ad9.
//
// Solidity: function deliverVoteResponse(uint256 _requestID, bool _finalResult, address[] _correctOracles, address[] _incorrectOracles, address[] _incorrectRevealOracles) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) DeliverVoteResponse(_requestID *big.Int, _finalResult bool, _correctOracles []common.Address, _incorrectOracles []common.Address, _incorrectRevealOracles []common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverVoteResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult, _correctOracles, _incorrectOracles, _incorrectRevealOracles)
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() payable returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) JoinAsOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "joinAsOracle")
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() payable returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) JoinAsOracle() (*types.Transaction, error) {
	return _OracleRequestContract.Contract.JoinAsOracle(&_OracleRequestContract.TransactOpts)
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() payable returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) JoinAsOracle() (*types.Transaction, error) {
	return _OracleRequestContract.Contract.JoinAsOracle(&_OracleRequestContract.TransactOpts)
}

// LeaveOracleNetwork is a paid mutator transaction binding the contract method 0xa08702f3.
//
// Solidity: function leaveOracleNetwork() returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) LeaveOracleNetwork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "leaveOracleNetwork")
}

// LeaveOracleNetwork is a paid mutator transaction binding the contract method 0xa08702f3.
//
// Solidity: function leaveOracleNetwork() returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) LeaveOracleNetwork() (*types.Transaction, error) {
	return _OracleRequestContract.Contract.LeaveOracleNetwork(&_OracleRequestContract.TransactOpts)
}

// LeaveOracleNetwork is a paid mutator transaction binding the contract method 0xa08702f3.
//
// Solidity: function leaveOracleNetwork() returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) LeaveOracleNetwork() (*types.Transaction, error) {
	return _OracleRequestContract.Contract.LeaveOracleNetwork(&_OracleRequestContract.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) PlaceBid(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "placeBid", _requestID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) PlaceBid(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.PlaceBid(&_OracleRequestContract.TransactOpts, _requestID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _requestID) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) PlaceBid(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.PlaceBid(&_OracleRequestContract.TransactOpts, _requestID)
}

// SetAggregatorContract is a paid mutator transaction binding the contract method 0x9967cf60.
//
// Solidity: function setAggregatorContract(address _aggAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) SetAggregatorContract(opts *bind.TransactOpts, _aggAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "setAggregatorContract", _aggAddr)
}

// SetAggregatorContract is a paid mutator transaction binding the contract method 0x9967cf60.
//
// Solidity: function setAggregatorContract(address _aggAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) SetAggregatorContract(_aggAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.SetAggregatorContract(&_OracleRequestContract.TransactOpts, _aggAddr)
}

// SetAggregatorContract is a paid mutator transaction binding the contract method 0x9967cf60.
//
// Solidity: function setAggregatorContract(address _aggAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) SetAggregatorContract(_aggAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.SetAggregatorContract(&_OracleRequestContract.TransactOpts, _aggAddr)
}

// SetReputationContract is a paid mutator transaction binding the contract method 0x9584660f.
//
// Solidity: function setReputationContract(address _repAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) SetReputationContract(opts *bind.TransactOpts, _repAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "setReputationContract", _repAddr)
}

// SetReputationContract is a paid mutator transaction binding the contract method 0x9584660f.
//
// Solidity: function setReputationContract(address _repAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) SetReputationContract(_repAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.SetReputationContract(&_OracleRequestContract.TransactOpts, _repAddr)
}

// SetReputationContract is a paid mutator transaction binding the contract method 0x9584660f.
//
// Solidity: function setReputationContract(address _repAddr) returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactorSession) SetReputationContract(_repAddr common.Address) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.SetReputationContract(&_OracleRequestContract.TransactOpts, _repAddr)
}

// TimeoutBiddingAppeal is a paid mutator transaction binding the contract method 0xce11babd.
//
// Solidity: function timeoutBiddingAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactor) TimeoutBiddingAppeal(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "timeoutBiddingAppeal", _requestID)
}

// TimeoutBiddingAppeal is a paid mutator transaction binding the contract method 0xce11babd.
//
// Solidity: function timeoutBiddingAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractSession) TimeoutBiddingAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutBiddingAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// TimeoutBiddingAppeal is a paid mutator transaction binding the contract method 0xce11babd.
//
// Solidity: function timeoutBiddingAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) TimeoutBiddingAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutBiddingAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// TimeoutCommitsAppeal is a paid mutator transaction binding the contract method 0x8394d49e.
//
// Solidity: function timeoutCommitsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactor) TimeoutCommitsAppeal(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "timeoutCommitsAppeal", _requestID)
}

// TimeoutCommitsAppeal is a paid mutator transaction binding the contract method 0x8394d49e.
//
// Solidity: function timeoutCommitsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractSession) TimeoutCommitsAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutCommitsAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// TimeoutCommitsAppeal is a paid mutator transaction binding the contract method 0x8394d49e.
//
// Solidity: function timeoutCommitsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) TimeoutCommitsAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutCommitsAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// TimeoutRevealsAppeal is a paid mutator transaction binding the contract method 0xd175a47f.
//
// Solidity: function timeoutRevealsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactor) TimeoutRevealsAppeal(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "timeoutRevealsAppeal", _requestID)
}

// TimeoutRevealsAppeal is a paid mutator transaction binding the contract method 0xd175a47f.
//
// Solidity: function timeoutRevealsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractSession) TimeoutRevealsAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutRevealsAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// TimeoutRevealsAppeal is a paid mutator transaction binding the contract method 0xd175a47f.
//
// Solidity: function timeoutRevealsAppeal(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) TimeoutRevealsAppeal(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.TimeoutRevealsAppeal(&_OracleRequestContract.TransactOpts, _requestID)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_OracleRequestContract *OracleRequestContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OracleRequestContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_OracleRequestContract *OracleRequestContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.Fallback(&_OracleRequestContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.Fallback(&_OracleRequestContract.TransactOpts, calldata)
}

// OracleRequestContractBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the OracleRequestContract contract.
type OracleRequestContractBidPlacedIterator struct {
	Event *OracleRequestContractBidPlaced // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractBidPlaced)
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
		it.Event = new(OracleRequestContractBidPlaced)
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
func (it *OracleRequestContractBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractBidPlaced represents a BidPlaced event raised by the OracleRequestContract contract.
type OracleRequestContractBidPlaced struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0xe677a6b9a16a696674f50d2c98a1ce8c5cb51d6b48a0223412c512bbc93664d5.
//
// Solidity: event BidPlaced(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*OracleRequestContractBidPlacedIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractBidPlacedIterator{contract: _OracleRequestContract.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0xe677a6b9a16a696674f50d2c98a1ce8c5cb51d6b48a0223412c512bbc93664d5.
//
// Solidity: event BidPlaced(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *OracleRequestContractBidPlaced) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractBidPlaced)
				if err := _OracleRequestContract.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0xe677a6b9a16a696674f50d2c98a1ce8c5cb51d6b48a0223412c512bbc93664d5.
//
// Solidity: event BidPlaced(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseBidPlaced(log types.Log) (*OracleRequestContractBidPlaced, error) {
	event := new(OracleRequestContractBidPlaced)
	if err := _OracleRequestContract.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractFinalIntIterator is returned from FilterFinalInt and is used to iterate over the raw logs and unpacked data for FinalInt events raised by the OracleRequestContract contract.
type OracleRequestContractFinalIntIterator struct {
	Event *OracleRequestContractFinalInt // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractFinalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractFinalInt)
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
		it.Event = new(OracleRequestContractFinalInt)
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
func (it *OracleRequestContractFinalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractFinalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractFinalInt represents a FinalInt event raised by the OracleRequestContract contract.
type OracleRequestContractFinalInt struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFinalInt is a free log retrieval operation binding the contract event 0x3691000541ef974fd47bc9eebb3e5c3acc93be77b1a5e07d20477d2193be6019.
//
// Solidity: event FinalInt(int256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterFinalInt(opts *bind.FilterOpts) (*OracleRequestContractFinalIntIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "FinalInt")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractFinalIntIterator{contract: _OracleRequestContract.contract, event: "FinalInt", logs: logs, sub: sub}, nil
}

// WatchFinalInt is a free log subscription operation binding the contract event 0x3691000541ef974fd47bc9eebb3e5c3acc93be77b1a5e07d20477d2193be6019.
//
// Solidity: event FinalInt(int256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchFinalInt(opts *bind.WatchOpts, sink chan<- *OracleRequestContractFinalInt) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "FinalInt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractFinalInt)
				if err := _OracleRequestContract.contract.UnpackLog(event, "FinalInt", log); err != nil {
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

// ParseFinalInt is a log parse operation binding the contract event 0x3691000541ef974fd47bc9eebb3e5c3acc93be77b1a5e07d20477d2193be6019.
//
// Solidity: event FinalInt(int256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseFinalInt(log types.Log) (*OracleRequestContractFinalInt, error) {
	event := new(OracleRequestContractFinalInt)
	if err := _OracleRequestContract.contract.UnpackLog(event, "FinalInt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractFinalResultLogIterator is returned from FilterFinalResultLog and is used to iterate over the raw logs and unpacked data for FinalResultLog events raised by the OracleRequestContract contract.
type OracleRequestContractFinalResultLogIterator struct {
	Event *OracleRequestContractFinalResultLog // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractFinalResultLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractFinalResultLog)
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
		it.Event = new(OracleRequestContractFinalResultLog)
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
func (it *OracleRequestContractFinalResultLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractFinalResultLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractFinalResultLog represents a FinalResultLog event raised by the OracleRequestContract contract.
type OracleRequestContractFinalResultLog struct {
	Arg0 bool
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFinalResultLog is a free log retrieval operation binding the contract event 0x5286d5a6d3a87f8e130652b705a8d0ec84042e834444bd00313a08d13a384755.
//
// Solidity: event FinalResultLog(bool arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterFinalResultLog(opts *bind.FilterOpts) (*OracleRequestContractFinalResultLogIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "FinalResultLog")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractFinalResultLogIterator{contract: _OracleRequestContract.contract, event: "FinalResultLog", logs: logs, sub: sub}, nil
}

// WatchFinalResultLog is a free log subscription operation binding the contract event 0x5286d5a6d3a87f8e130652b705a8d0ec84042e834444bd00313a08d13a384755.
//
// Solidity: event FinalResultLog(bool arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchFinalResultLog(opts *bind.WatchOpts, sink chan<- *OracleRequestContractFinalResultLog) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "FinalResultLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractFinalResultLog)
				if err := _OracleRequestContract.contract.UnpackLog(event, "FinalResultLog", log); err != nil {
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

// ParseFinalResultLog is a log parse operation binding the contract event 0x5286d5a6d3a87f8e130652b705a8d0ec84042e834444bd00313a08d13a384755.
//
// Solidity: event FinalResultLog(bool arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseFinalResultLog(log types.Log) (*OracleRequestContractFinalResultLog, error) {
	event := new(OracleRequestContractFinalResultLog)
	if err := _OracleRequestContract.contract.UnpackLog(event, "FinalResultLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractLogNumberOfOraclesIterator is returned from FilterLogNumberOfOracles and is used to iterate over the raw logs and unpacked data for LogNumberOfOracles events raised by the OracleRequestContract contract.
type OracleRequestContractLogNumberOfOraclesIterator struct {
	Event *OracleRequestContractLogNumberOfOracles // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractLogNumberOfOraclesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractLogNumberOfOracles)
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
		it.Event = new(OracleRequestContractLogNumberOfOracles)
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
func (it *OracleRequestContractLogNumberOfOraclesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractLogNumberOfOraclesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractLogNumberOfOracles represents a LogNumberOfOracles event raised by the OracleRequestContract contract.
type OracleRequestContractLogNumberOfOracles struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogNumberOfOracles is a free log retrieval operation binding the contract event 0x226fa98be5b4c2f9d7cf2b650939b5adffe55f9156bbde7b6ed95694304f3a2a.
//
// Solidity: event LogNumberOfOracles(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterLogNumberOfOracles(opts *bind.FilterOpts) (*OracleRequestContractLogNumberOfOraclesIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "LogNumberOfOracles")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractLogNumberOfOraclesIterator{contract: _OracleRequestContract.contract, event: "LogNumberOfOracles", logs: logs, sub: sub}, nil
}

// WatchLogNumberOfOracles is a free log subscription operation binding the contract event 0x226fa98be5b4c2f9d7cf2b650939b5adffe55f9156bbde7b6ed95694304f3a2a.
//
// Solidity: event LogNumberOfOracles(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchLogNumberOfOracles(opts *bind.WatchOpts, sink chan<- *OracleRequestContractLogNumberOfOracles) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "LogNumberOfOracles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractLogNumberOfOracles)
				if err := _OracleRequestContract.contract.UnpackLog(event, "LogNumberOfOracles", log); err != nil {
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

// ParseLogNumberOfOracles is a log parse operation binding the contract event 0x226fa98be5b4c2f9d7cf2b650939b5adffe55f9156bbde7b6ed95694304f3a2a.
//
// Solidity: event LogNumberOfOracles(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseLogNumberOfOracles(log types.Log) (*OracleRequestContractLogNumberOfOracles, error) {
	event := new(OracleRequestContractLogNumberOfOracles)
	if err := _OracleRequestContract.contract.UnpackLog(event, "LogNumberOfOracles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractLogTimeOutLenIterator is returned from FilterLogTimeOutLen and is used to iterate over the raw logs and unpacked data for LogTimeOutLen events raised by the OracleRequestContract contract.
type OracleRequestContractLogTimeOutLenIterator struct {
	Event *OracleRequestContractLogTimeOutLen // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractLogTimeOutLenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractLogTimeOutLen)
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
		it.Event = new(OracleRequestContractLogTimeOutLen)
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
func (it *OracleRequestContractLogTimeOutLenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractLogTimeOutLenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractLogTimeOutLen represents a LogTimeOutLen event raised by the OracleRequestContract contract.
type OracleRequestContractLogTimeOutLen struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogTimeOutLen is a free log retrieval operation binding the contract event 0x484653374230165d70776055de5c30786c08673eb4740ee73dfc172c5e1b5cd0.
//
// Solidity: event LogTimeOutLen(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterLogTimeOutLen(opts *bind.FilterOpts) (*OracleRequestContractLogTimeOutLenIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "LogTimeOutLen")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractLogTimeOutLenIterator{contract: _OracleRequestContract.contract, event: "LogTimeOutLen", logs: logs, sub: sub}, nil
}

// WatchLogTimeOutLen is a free log subscription operation binding the contract event 0x484653374230165d70776055de5c30786c08673eb4740ee73dfc172c5e1b5cd0.
//
// Solidity: event LogTimeOutLen(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchLogTimeOutLen(opts *bind.WatchOpts, sink chan<- *OracleRequestContractLogTimeOutLen) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "LogTimeOutLen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractLogTimeOutLen)
				if err := _OracleRequestContract.contract.UnpackLog(event, "LogTimeOutLen", log); err != nil {
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

// ParseLogTimeOutLen is a log parse operation binding the contract event 0x484653374230165d70776055de5c30786c08673eb4740ee73dfc172c5e1b5cd0.
//
// Solidity: event LogTimeOutLen(uint256 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseLogTimeOutLen(log types.Log) (*OracleRequestContractLogTimeOutLen, error) {
	event := new(OracleRequestContractLogTimeOutLen)
	if err := _OracleRequestContract.contract.UnpackLog(event, "LogTimeOutLen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractLoggingIterator is returned from FilterLogging and is used to iterate over the raw logs and unpacked data for Logging events raised by the OracleRequestContract contract.
type OracleRequestContractLoggingIterator struct {
	Event *OracleRequestContractLogging // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractLoggingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractLogging)
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
		it.Event = new(OracleRequestContractLogging)
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
func (it *OracleRequestContractLoggingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractLoggingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractLogging represents a Logging event raised by the OracleRequestContract contract.
type OracleRequestContractLogging struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogging is a free log retrieval operation binding the contract event 0xb1ed47ced4562a888b3c47e728f82f7863e472cad724018fe1585bc6853ec82d.
//
// Solidity: event Logging(string arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterLogging(opts *bind.FilterOpts) (*OracleRequestContractLoggingIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "Logging")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractLoggingIterator{contract: _OracleRequestContract.contract, event: "Logging", logs: logs, sub: sub}, nil
}

// WatchLogging is a free log subscription operation binding the contract event 0xb1ed47ced4562a888b3c47e728f82f7863e472cad724018fe1585bc6853ec82d.
//
// Solidity: event Logging(string arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchLogging(opts *bind.WatchOpts, sink chan<- *OracleRequestContractLogging) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "Logging")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractLogging)
				if err := _OracleRequestContract.contract.UnpackLog(event, "Logging", log); err != nil {
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
func (_OracleRequestContract *OracleRequestContractFilterer) ParseLogging(log types.Log) (*OracleRequestContractLogging, error) {
	event := new(OracleRequestContractLogging)
	if err := _OracleRequestContract.contract.UnpackLog(event, "Logging", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractOpenForBidsIterator is returned from FilterOpenForBids and is used to iterate over the raw logs and unpacked data for OpenForBids events raised by the OracleRequestContract contract.
type OracleRequestContractOpenForBidsIterator struct {
	Event *OracleRequestContractOpenForBids // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractOpenForBidsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractOpenForBids)
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
		it.Event = new(OracleRequestContractOpenForBids)
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
func (it *OracleRequestContractOpenForBidsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractOpenForBidsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractOpenForBids represents a OpenForBids event raised by the OracleRequestContract contract.
type OracleRequestContractOpenForBids struct {
	Arg0 *big.Int
	Arg1 []byte
	Arg2 *big.Int
	Arg3 *big.Int
	Arg4 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOpenForBids is a free log retrieval operation binding the contract event 0x45467d61e6f35d1866a79e0828e55247a1cbdbd929532d73c56d73ae64cfe83a.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterOpenForBids(opts *bind.FilterOpts) (*OracleRequestContractOpenForBidsIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "OpenForBids")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractOpenForBidsIterator{contract: _OracleRequestContract.contract, event: "OpenForBids", logs: logs, sub: sub}, nil
}

// WatchOpenForBids is a free log subscription operation binding the contract event 0x45467d61e6f35d1866a79e0828e55247a1cbdbd929532d73c56d73ae64cfe83a.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchOpenForBids(opts *bind.WatchOpts, sink chan<- *OracleRequestContractOpenForBids) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "OpenForBids")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractOpenForBids)
				if err := _OracleRequestContract.contract.UnpackLog(event, "OpenForBids", log); err != nil {
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

// ParseOpenForBids is a log parse operation binding the contract event 0x45467d61e6f35d1866a79e0828e55247a1cbdbd929532d73c56d73ae64cfe83a.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseOpenForBids(log types.Log) (*OracleRequestContractOpenForBids, error) {
	event := new(OracleRequestContractOpenForBids)
	if err := _OracleRequestContract.contract.UnpackLog(event, "OpenForBids", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractOracleJoinedIterator is returned from FilterOracleJoined and is used to iterate over the raw logs and unpacked data for OracleJoined events raised by the OracleRequestContract contract.
type OracleRequestContractOracleJoinedIterator struct {
	Event *OracleRequestContractOracleJoined // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractOracleJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractOracleJoined)
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
		it.Event = new(OracleRequestContractOracleJoined)
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
func (it *OracleRequestContractOracleJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractOracleJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractOracleJoined represents a OracleJoined event raised by the OracleRequestContract contract.
type OracleRequestContractOracleJoined struct {
	Arg0 common.Address
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOracleJoined is a free log retrieval operation binding the contract event 0x2d62108771d82d92634e2d038b6c7f09290edfee5b1443e25299f04a0cfbca72.
//
// Solidity: event OracleJoined(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterOracleJoined(opts *bind.FilterOpts) (*OracleRequestContractOracleJoinedIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "OracleJoined")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractOracleJoinedIterator{contract: _OracleRequestContract.contract, event: "OracleJoined", logs: logs, sub: sub}, nil
}

// WatchOracleJoined is a free log subscription operation binding the contract event 0x2d62108771d82d92634e2d038b6c7f09290edfee5b1443e25299f04a0cfbca72.
//
// Solidity: event OracleJoined(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchOracleJoined(opts *bind.WatchOpts, sink chan<- *OracleRequestContractOracleJoined) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "OracleJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractOracleJoined)
				if err := _OracleRequestContract.contract.UnpackLog(event, "OracleJoined", log); err != nil {
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

// ParseOracleJoined is a log parse operation binding the contract event 0x2d62108771d82d92634e2d038b6c7f09290edfee5b1443e25299f04a0cfbca72.
//
// Solidity: event OracleJoined(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseOracleJoined(log types.Log) (*OracleRequestContractOracleJoined, error) {
	event := new(OracleRequestContractOracleJoined)
	if err := _OracleRequestContract.contract.UnpackLog(event, "OracleJoined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractOracleLeftIterator is returned from FilterOracleLeft and is used to iterate over the raw logs and unpacked data for OracleLeft events raised by the OracleRequestContract contract.
type OracleRequestContractOracleLeftIterator struct {
	Event *OracleRequestContractOracleLeft // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractOracleLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractOracleLeft)
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
		it.Event = new(OracleRequestContractOracleLeft)
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
func (it *OracleRequestContractOracleLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractOracleLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractOracleLeft represents a OracleLeft event raised by the OracleRequestContract contract.
type OracleRequestContractOracleLeft struct {
	Arg0 common.Address
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOracleLeft is a free log retrieval operation binding the contract event 0x14951d787e2aaf7813982720871de23d576e7d935c0587e12a7fc2c4c960165a.
//
// Solidity: event OracleLeft(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterOracleLeft(opts *bind.FilterOpts) (*OracleRequestContractOracleLeftIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "OracleLeft")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractOracleLeftIterator{contract: _OracleRequestContract.contract, event: "OracleLeft", logs: logs, sub: sub}, nil
}

// WatchOracleLeft is a free log subscription operation binding the contract event 0x14951d787e2aaf7813982720871de23d576e7d935c0587e12a7fc2c4c960165a.
//
// Solidity: event OracleLeft(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchOracleLeft(opts *bind.WatchOpts, sink chan<- *OracleRequestContractOracleLeft) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "OracleLeft")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractOracleLeft)
				if err := _OracleRequestContract.contract.UnpackLog(event, "OracleLeft", log); err != nil {
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

// ParseOracleLeft is a log parse operation binding the contract event 0x14951d787e2aaf7813982720871de23d576e7d935c0587e12a7fc2c4c960165a.
//
// Solidity: event OracleLeft(address arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseOracleLeft(log types.Log) (*OracleRequestContractOracleLeft, error) {
	event := new(OracleRequestContractOracleLeft)
	if err := _OracleRequestContract.contract.UnpackLog(event, "OracleLeft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractOraclePaidIterator is returned from FilterOraclePaid and is used to iterate over the raw logs and unpacked data for OraclePaid events raised by the OracleRequestContract contract.
type OracleRequestContractOraclePaidIterator struct {
	Event *OracleRequestContractOraclePaid // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractOraclePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractOraclePaid)
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
		it.Event = new(OracleRequestContractOraclePaid)
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
func (it *OracleRequestContractOraclePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractOraclePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractOraclePaid represents a OraclePaid event raised by the OracleRequestContract contract.
type OracleRequestContractOraclePaid struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOraclePaid is a free log retrieval operation binding the contract event 0x88fa4dc6dc84fb056628ddf02c361252f313bcdd951c9d28e3b628bffac8ab55.
//
// Solidity: event OraclePaid(address arg0, uint256 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterOraclePaid(opts *bind.FilterOpts) (*OracleRequestContractOraclePaidIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "OraclePaid")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractOraclePaidIterator{contract: _OracleRequestContract.contract, event: "OraclePaid", logs: logs, sub: sub}, nil
}

// WatchOraclePaid is a free log subscription operation binding the contract event 0x88fa4dc6dc84fb056628ddf02c361252f313bcdd951c9d28e3b628bffac8ab55.
//
// Solidity: event OraclePaid(address arg0, uint256 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchOraclePaid(opts *bind.WatchOpts, sink chan<- *OracleRequestContractOraclePaid) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "OraclePaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractOraclePaid)
				if err := _OracleRequestContract.contract.UnpackLog(event, "OraclePaid", log); err != nil {
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

// ParseOraclePaid is a log parse operation binding the contract event 0x88fa4dc6dc84fb056628ddf02c361252f313bcdd951c9d28e3b628bffac8ab55.
//
// Solidity: event OraclePaid(address arg0, uint256 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseOraclePaid(log types.Log) (*OracleRequestContractOraclePaid, error) {
	event := new(OracleRequestContractOraclePaid)
	if err := _OracleRequestContract.contract.UnpackLog(event, "OraclePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractPrintAddressIterator is returned from FilterPrintAddress and is used to iterate over the raw logs and unpacked data for PrintAddress events raised by the OracleRequestContract contract.
type OracleRequestContractPrintAddressIterator struct {
	Event *OracleRequestContractPrintAddress // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractPrintAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractPrintAddress)
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
		it.Event = new(OracleRequestContractPrintAddress)
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
func (it *OracleRequestContractPrintAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractPrintAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractPrintAddress represents a PrintAddress event raised by the OracleRequestContract contract.
type OracleRequestContractPrintAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintAddress is a free log retrieval operation binding the contract event 0xa16adbb4218174483d6638aa2a7cfae613fb09f920505beb947e63144629a875.
//
// Solidity: event PrintAddress(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterPrintAddress(opts *bind.FilterOpts) (*OracleRequestContractPrintAddressIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "PrintAddress")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractPrintAddressIterator{contract: _OracleRequestContract.contract, event: "PrintAddress", logs: logs, sub: sub}, nil
}

// WatchPrintAddress is a free log subscription operation binding the contract event 0xa16adbb4218174483d6638aa2a7cfae613fb09f920505beb947e63144629a875.
//
// Solidity: event PrintAddress(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchPrintAddress(opts *bind.WatchOpts, sink chan<- *OracleRequestContractPrintAddress) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "PrintAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractPrintAddress)
				if err := _OracleRequestContract.contract.UnpackLog(event, "PrintAddress", log); err != nil {
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

// ParsePrintAddress is a log parse operation binding the contract event 0xa16adbb4218174483d6638aa2a7cfae613fb09f920505beb947e63144629a875.
//
// Solidity: event PrintAddress(address arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) ParsePrintAddress(log types.Log) (*OracleRequestContractPrintAddress, error) {
	event := new(OracleRequestContractPrintAddress)
	if err := _OracleRequestContract.contract.UnpackLog(event, "PrintAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractPrintAddressesIterator is returned from FilterPrintAddresses and is used to iterate over the raw logs and unpacked data for PrintAddresses events raised by the OracleRequestContract contract.
type OracleRequestContractPrintAddressesIterator struct {
	Event *OracleRequestContractPrintAddresses // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractPrintAddressesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractPrintAddresses)
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
		it.Event = new(OracleRequestContractPrintAddresses)
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
func (it *OracleRequestContractPrintAddressesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractPrintAddressesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractPrintAddresses represents a PrintAddresses event raised by the OracleRequestContract contract.
type OracleRequestContractPrintAddresses struct {
	Arg0 []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintAddresses is a free log retrieval operation binding the contract event 0x96d75b5fb706f946feb29fd7675ecabcd8ad4262d3b8ead504407c62136aefa2.
//
// Solidity: event PrintAddresses(address[] arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterPrintAddresses(opts *bind.FilterOpts) (*OracleRequestContractPrintAddressesIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "PrintAddresses")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractPrintAddressesIterator{contract: _OracleRequestContract.contract, event: "PrintAddresses", logs: logs, sub: sub}, nil
}

// WatchPrintAddresses is a free log subscription operation binding the contract event 0x96d75b5fb706f946feb29fd7675ecabcd8ad4262d3b8ead504407c62136aefa2.
//
// Solidity: event PrintAddresses(address[] arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchPrintAddresses(opts *bind.WatchOpts, sink chan<- *OracleRequestContractPrintAddresses) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "PrintAddresses")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractPrintAddresses)
				if err := _OracleRequestContract.contract.UnpackLog(event, "PrintAddresses", log); err != nil {
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

// ParsePrintAddresses is a log parse operation binding the contract event 0x96d75b5fb706f946feb29fd7675ecabcd8ad4262d3b8ead504407c62136aefa2.
//
// Solidity: event PrintAddresses(address[] arg0)
func (_OracleRequestContract *OracleRequestContractFilterer) ParsePrintAddresses(log types.Log) (*OracleRequestContractPrintAddresses, error) {
	event := new(OracleRequestContractPrintAddresses)
	if err := _OracleRequestContract.contract.UnpackLog(event, "PrintAddresses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractReleaseRequestDetailsIterator is returned from FilterReleaseRequestDetails and is used to iterate over the raw logs and unpacked data for ReleaseRequestDetails events raised by the OracleRequestContract contract.
type OracleRequestContractReleaseRequestDetailsIterator struct {
	Event *OracleRequestContractReleaseRequestDetails // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractReleaseRequestDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractReleaseRequestDetails)
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
		it.Event = new(OracleRequestContractReleaseRequestDetails)
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
func (it *OracleRequestContractReleaseRequestDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractReleaseRequestDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractReleaseRequestDetails represents a ReleaseRequestDetails event raised by the OracleRequestContract contract.
type OracleRequestContractReleaseRequestDetails struct {
	Arg0 *big.Int
	Arg1 []byte
	Arg2 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReleaseRequestDetails is a free log retrieval operation binding the contract event 0x2510e6950be86d6001a0b1229da4c21fa8e26b6bcf345d0b348fe42894fceb03.
//
// Solidity: event ReleaseRequestDetails(uint256 arg0, bytes arg1, bytes arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterReleaseRequestDetails(opts *bind.FilterOpts) (*OracleRequestContractReleaseRequestDetailsIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "ReleaseRequestDetails")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractReleaseRequestDetailsIterator{contract: _OracleRequestContract.contract, event: "ReleaseRequestDetails", logs: logs, sub: sub}, nil
}

// WatchReleaseRequestDetails is a free log subscription operation binding the contract event 0x2510e6950be86d6001a0b1229da4c21fa8e26b6bcf345d0b348fe42894fceb03.
//
// Solidity: event ReleaseRequestDetails(uint256 arg0, bytes arg1, bytes arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchReleaseRequestDetails(opts *bind.WatchOpts, sink chan<- *OracleRequestContractReleaseRequestDetails) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "ReleaseRequestDetails")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractReleaseRequestDetails)
				if err := _OracleRequestContract.contract.UnpackLog(event, "ReleaseRequestDetails", log); err != nil {
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

// ParseReleaseRequestDetails is a log parse operation binding the contract event 0x2510e6950be86d6001a0b1229da4c21fa8e26b6bcf345d0b348fe42894fceb03.
//
// Solidity: event ReleaseRequestDetails(uint256 arg0, bytes arg1, bytes arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseReleaseRequestDetails(log types.Log) (*OracleRequestContractReleaseRequestDetails, error) {
	event := new(OracleRequestContractReleaseRequestDetails)
	if err := _OracleRequestContract.contract.UnpackLog(event, "ReleaseRequestDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractStatusChangeIterator is returned from FilterStatusChange and is used to iterate over the raw logs and unpacked data for StatusChange events raised by the OracleRequestContract contract.
type OracleRequestContractStatusChangeIterator struct {
	Event *OracleRequestContractStatusChange // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractStatusChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractStatusChange)
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
		it.Event = new(OracleRequestContractStatusChange)
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
func (it *OracleRequestContractStatusChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractStatusChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractStatusChange represents a StatusChange event raised by the OracleRequestContract contract.
type OracleRequestContractStatusChange struct {
	Arg0 *big.Int
	Arg1 uint8
	Arg2 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStatusChange is a free log retrieval operation binding the contract event 0xbd75c0c0fcfb2a06c72850a74ec03c188bd53958646995bde4e3980c2ccde2b9.
//
// Solidity: event StatusChange(uint256 arg0, uint8 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterStatusChange(opts *bind.FilterOpts) (*OracleRequestContractStatusChangeIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "StatusChange")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractStatusChangeIterator{contract: _OracleRequestContract.contract, event: "StatusChange", logs: logs, sub: sub}, nil
}

// WatchStatusChange is a free log subscription operation binding the contract event 0xbd75c0c0fcfb2a06c72850a74ec03c188bd53958646995bde4e3980c2ccde2b9.
//
// Solidity: event StatusChange(uint256 arg0, uint8 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchStatusChange(opts *bind.WatchOpts, sink chan<- *OracleRequestContractStatusChange) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "StatusChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractStatusChange)
				if err := _OracleRequestContract.contract.UnpackLog(event, "StatusChange", log); err != nil {
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

// ParseStatusChange is a log parse operation binding the contract event 0xbd75c0c0fcfb2a06c72850a74ec03c188bd53958646995bde4e3980c2ccde2b9.
//
// Solidity: event StatusChange(uint256 arg0, uint8 arg1, string arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseStatusChange(log types.Log) (*OracleRequestContractStatusChange, error) {
	event := new(OracleRequestContractStatusChange)
	if err := _OracleRequestContract.contract.UnpackLog(event, "StatusChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleRequestContractUpdateTimestampIterator is returned from FilterUpdateTimestamp and is used to iterate over the raw logs and unpacked data for UpdateTimestamp events raised by the OracleRequestContract contract.
type OracleRequestContractUpdateTimestampIterator struct {
	Event *OracleRequestContractUpdateTimestamp // Event containing the contract specifics and raw log

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
func (it *OracleRequestContractUpdateTimestampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleRequestContractUpdateTimestamp)
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
		it.Event = new(OracleRequestContractUpdateTimestamp)
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
func (it *OracleRequestContractUpdateTimestampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleRequestContractUpdateTimestampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleRequestContractUpdateTimestamp represents a UpdateTimestamp event raised by the OracleRequestContract contract.
type OracleRequestContractUpdateTimestamp struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateTimestamp is a free log retrieval operation binding the contract event 0x6a54c2949cfe447c07615c41ee2012ce51a3d9b8acb122e210098f51b2d4798d.
//
// Solidity: event UpdateTimestamp(uint256 arg0, uint256 arg1, uint256 arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterUpdateTimestamp(opts *bind.FilterOpts) (*OracleRequestContractUpdateTimestampIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "UpdateTimestamp")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractUpdateTimestampIterator{contract: _OracleRequestContract.contract, event: "UpdateTimestamp", logs: logs, sub: sub}, nil
}

// WatchUpdateTimestamp is a free log subscription operation binding the contract event 0x6a54c2949cfe447c07615c41ee2012ce51a3d9b8acb122e210098f51b2d4798d.
//
// Solidity: event UpdateTimestamp(uint256 arg0, uint256 arg1, uint256 arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) WatchUpdateTimestamp(opts *bind.WatchOpts, sink chan<- *OracleRequestContractUpdateTimestamp) (event.Subscription, error) {

	logs, sub, err := _OracleRequestContract.contract.WatchLogs(opts, "UpdateTimestamp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleRequestContractUpdateTimestamp)
				if err := _OracleRequestContract.contract.UnpackLog(event, "UpdateTimestamp", log); err != nil {
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

// ParseUpdateTimestamp is a log parse operation binding the contract event 0x6a54c2949cfe447c07615c41ee2012ce51a3d9b8acb122e210098f51b2d4798d.
//
// Solidity: event UpdateTimestamp(uint256 arg0, uint256 arg1, uint256 arg2)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseUpdateTimestamp(log types.Log) (*OracleRequestContractUpdateTimestamp, error) {
	event := new(OracleRequestContractUpdateTimestamp)
	if err := _OracleRequestContract.contract.UnpackLog(event, "UpdateTimestamp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
