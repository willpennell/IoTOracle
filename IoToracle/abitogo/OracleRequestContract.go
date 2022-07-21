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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"OpenForBids\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OracleJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"OracleLeft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"ReleaseRequestDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleRequesterContract.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"StatusChange\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"cancelRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackFID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_IoTID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_dataType\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_requiredResult\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_numberOfOracles\",\"type\":\"uint32\"}],\"name\":\"createRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_finalResult\",\"type\":\"bytes\"}],\"name\":\"deliverResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getDataType\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getNumberOfOracles\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_orcAddr\",\"type\":\"address\"}],\"name\":\"getOracleForRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"getPHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinAsOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaveOracleNetwork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackFID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"IoTID\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"dataType\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"pHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"numberOfOracles\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"oracleCounter\",\"type\":\"uint32\"},{\"internalType\":\"enumOracleRequesterContract.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetPHash is a free data retrieval call binding the contract method 0x829c22fd.
//
// Solidity: function getPHash(uint256 _requestID) view returns(bytes32)
func (_OracleRequestContract *OracleRequestContractCaller) GetPHash(opts *bind.CallOpts, _requestID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OracleRequestContract.contract.Call(opts, &out, "getPHash", _requestID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPHash is a free data retrieval call binding the contract method 0x829c22fd.
//
// Solidity: function getPHash(uint256 _requestID) view returns(bytes32)
func (_OracleRequestContract *OracleRequestContractSession) GetPHash(_requestID *big.Int) ([32]byte, error) {
	return _OracleRequestContract.Contract.GetPHash(&_OracleRequestContract.CallOpts, _requestID)
}

// GetPHash is a free data retrieval call binding the contract method 0x829c22fd.
//
// Solidity: function getPHash(uint256 _requestID) view returns(bytes32)
func (_OracleRequestContract *OracleRequestContractCallerSession) GetPHash(_requestID *big.Int) ([32]byte, error) {
	return _OracleRequestContract.Contract.GetPHash(&_OracleRequestContract.CallOpts, _requestID)
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

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, bytes32 pHash, uint32 numberOfOracles, uint32 oracleCounter, uint8 status)
func (_OracleRequestContract *OracleRequestContractCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	PHash           [32]byte
	NumberOfOracles uint32
	OracleCounter   uint32
	Status          uint8
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
		PHash           [32]byte
		NumberOfOracles uint32
		OracleCounter   uint32
		Status          uint8
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
	outstruct.PHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.NumberOfOracles = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.OracleCounter = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, bytes32 pHash, uint32 numberOfOracles, uint32 oracleCounter, uint8 status)
func (_OracleRequestContract *OracleRequestContractSession) Requests(arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	PHash           [32]byte
	NumberOfOracles uint32
	OracleCounter   uint32
	Status          uint8
}, error) {
	return _OracleRequestContract.Contract.Requests(&_OracleRequestContract.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestID, address requester, address callbackAddress, bytes callbackFID, bytes IoTID, bytes dataType, bytes32 pHash, uint32 numberOfOracles, uint32 oracleCounter, uint8 status)
func (_OracleRequestContract *OracleRequestContractCallerSession) Requests(arg0 *big.Int) (struct {
	RequestID       *big.Int
	Requester       common.Address
	CallbackAddress common.Address
	CallbackFID     []byte
	IoTID           []byte
	DataType        []byte
	PHash           [32]byte
	NumberOfOracles uint32
	OracleCounter   uint32
	Status          uint8
}, error) {
	return _OracleRequestContract.Contract.Requests(&_OracleRequestContract.CallOpts, arg0)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactor) CancelRequest(opts *bind.TransactOpts, _requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "cancelRequest", _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequest(&_OracleRequestContract.TransactOpts, _requestID)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 _requestID) returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) CancelRequest(_requestID *big.Int) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CancelRequest(&_OracleRequestContract.TransactOpts, _requestID)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8979a3c2.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, bytes _requiredResult, uint32 _numberOfOracles) returns(uint256)
func (_OracleRequestContract *OracleRequestContractTransactor) CreateRequest(opts *bind.TransactOpts, _callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _requiredResult []byte, _numberOfOracles uint32) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "createRequest", _callbackAddress, _callbackFID, _IoTID, _dataType, _requiredResult, _numberOfOracles)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8979a3c2.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, bytes _requiredResult, uint32 _numberOfOracles) returns(uint256)
func (_OracleRequestContract *OracleRequestContractSession) CreateRequest(_callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _requiredResult []byte, _numberOfOracles uint32) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CreateRequest(&_OracleRequestContract.TransactOpts, _callbackAddress, _callbackFID, _IoTID, _dataType, _requiredResult, _numberOfOracles)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8979a3c2.
//
// Solidity: function createRequest(address _callbackAddress, bytes _callbackFID, bytes _IoTID, bytes _dataType, bytes _requiredResult, uint32 _numberOfOracles) returns(uint256)
func (_OracleRequestContract *OracleRequestContractTransactorSession) CreateRequest(_callbackAddress common.Address, _callbackFID []byte, _IoTID []byte, _dataType []byte, _requiredResult []byte, _numberOfOracles uint32) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.CreateRequest(&_OracleRequestContract.TransactOpts, _callbackAddress, _callbackFID, _IoTID, _dataType, _requiredResult, _numberOfOracles)
}

// DeliverResponse is a paid mutator transaction binding the contract method 0xea317d5d.
//
// Solidity: function deliverResponse(uint256 _requestID, bytes _finalResult) returns()
func (_OracleRequestContract *OracleRequestContractTransactor) DeliverResponse(opts *bind.TransactOpts, _requestID *big.Int, _finalResult []byte) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "deliverResponse", _requestID, _finalResult)
}

// DeliverResponse is a paid mutator transaction binding the contract method 0xea317d5d.
//
// Solidity: function deliverResponse(uint256 _requestID, bytes _finalResult) returns()
func (_OracleRequestContract *OracleRequestContractSession) DeliverResponse(_requestID *big.Int, _finalResult []byte) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult)
}

// DeliverResponse is a paid mutator transaction binding the contract method 0xea317d5d.
//
// Solidity: function deliverResponse(uint256 _requestID, bytes _finalResult) returns()
func (_OracleRequestContract *OracleRequestContractTransactorSession) DeliverResponse(_requestID *big.Int, _finalResult []byte) (*types.Transaction, error) {
	return _OracleRequestContract.Contract.DeliverResponse(&_OracleRequestContract.TransactOpts, _requestID, _finalResult)
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() returns(bool)
func (_OracleRequestContract *OracleRequestContractTransactor) JoinAsOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleRequestContract.contract.Transact(opts, "joinAsOracle")
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() returns(bool)
func (_OracleRequestContract *OracleRequestContractSession) JoinAsOracle() (*types.Transaction, error) {
	return _OracleRequestContract.Contract.JoinAsOracle(&_OracleRequestContract.TransactOpts)
}

// JoinAsOracle is a paid mutator transaction binding the contract method 0x950e1da9.
//
// Solidity: function joinAsOracle() returns(bool)
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
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOpenForBids is a free log retrieval operation binding the contract event 0x9fa3818a0bcd92e6bdf76b585ed9feea4ad7d3467dc3770c2af818d43f9581a5.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterOpenForBids(opts *bind.FilterOpts) (*OracleRequestContractOpenForBidsIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "OpenForBids")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractOpenForBidsIterator{contract: _OracleRequestContract.contract, event: "OpenForBids", logs: logs, sub: sub}, nil
}

// WatchOpenForBids is a free log subscription operation binding the contract event 0x9fa3818a0bcd92e6bdf76b585ed9feea4ad7d3467dc3770c2af818d43f9581a5.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1)
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

// ParseOpenForBids is a log parse operation binding the contract event 0x9fa3818a0bcd92e6bdf76b585ed9feea4ad7d3467dc3770c2af818d43f9581a5.
//
// Solidity: event OpenForBids(uint256 arg0, bytes arg1)
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
	Arg0 uint8
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStatusChange is a free log retrieval operation binding the contract event 0x346e0fa9e112a7376778dcb0fe0fdfe51db577b3d0b1592a40266db0b192f822.
//
// Solidity: event StatusChange(uint8 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) FilterStatusChange(opts *bind.FilterOpts) (*OracleRequestContractStatusChangeIterator, error) {

	logs, sub, err := _OracleRequestContract.contract.FilterLogs(opts, "StatusChange")
	if err != nil {
		return nil, err
	}
	return &OracleRequestContractStatusChangeIterator{contract: _OracleRequestContract.contract, event: "StatusChange", logs: logs, sub: sub}, nil
}

// WatchStatusChange is a free log subscription operation binding the contract event 0x346e0fa9e112a7376778dcb0fe0fdfe51db577b3d0b1592a40266db0b192f822.
//
// Solidity: event StatusChange(uint8 arg0, string arg1)
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

// ParseStatusChange is a log parse operation binding the contract event 0x346e0fa9e112a7376778dcb0fe0fdfe51db577b3d0b1592a40266db0b192f822.
//
// Solidity: event StatusChange(uint8 arg0, string arg1)
func (_OracleRequestContract *OracleRequestContractFilterer) ParseStatusChange(log types.Log) (*OracleRequestContractStatusChange, error) {
	event := new(OracleRequestContractStatusChange)
	if err := _OracleRequestContract.contract.UnpackLog(event, "StatusChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
