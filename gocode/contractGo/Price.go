// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scf

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ScfABI is the input ABI used to generate the binding from.
const ScfABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Scf is an auto generated Go binding around an Ethereum contract.
type Scf struct {
	ScfCaller     // Read-only binding to the contract
	ScfTransactor // Write-only binding to the contract
	ScfFilterer   // Log filterer for contract events
}

// ScfCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScfSession struct {
	Contract     *Scf              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScfCallerSession struct {
	Contract *ScfCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ScfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScfTransactorSession struct {
	Contract     *ScfTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScfRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScfRaw struct {
	Contract *Scf // Generic contract binding to access the raw methods on
}

// ScfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScfCallerRaw struct {
	Contract *ScfCaller // Generic read-only contract binding to access the raw methods on
}

// ScfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScfTransactorRaw struct {
	Contract *ScfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScf creates a new instance of Scf, bound to a specific deployed contract.
func NewScf(address common.Address, backend bind.ContractBackend) (*Scf, error) {
	contract, err := bindScf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Scf{ScfCaller: ScfCaller{contract: contract}, ScfTransactor: ScfTransactor{contract: contract}, ScfFilterer: ScfFilterer{contract: contract}}, nil
}

// NewScfCaller creates a new read-only instance of Scf, bound to a specific deployed contract.
func NewScfCaller(address common.Address, caller bind.ContractCaller) (*ScfCaller, error) {
	contract, err := bindScf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScfCaller{contract: contract}, nil
}

// NewScfTransactor creates a new write-only instance of Scf, bound to a specific deployed contract.
func NewScfTransactor(address common.Address, transactor bind.ContractTransactor) (*ScfTransactor, error) {
	contract, err := bindScf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScfTransactor{contract: contract}, nil
}

// NewScfFilterer creates a new log filterer instance of Scf, bound to a specific deployed contract.
func NewScfFilterer(address common.Address, filterer bind.ContractFilterer) (*ScfFilterer, error) {
	contract, err := bindScf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScfFilterer{contract: contract}, nil
}

// bindScf binds a generic wrapper to an already deployed contract.
func bindScf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scf *ScfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scf.Contract.ScfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scf *ScfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scf.Contract.ScfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scf *ScfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scf.Contract.ScfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Scf *ScfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Scf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Scf *ScfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Scf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Scf *ScfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Scf.Contract.contract.Transact(opts, method, params...)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Scf *ScfCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Scf.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Scf *ScfSession) GetLatestPrice() (*big.Int, error) {
	return _Scf.Contract.GetLatestPrice(&_Scf.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Scf *ScfCallerSession) GetLatestPrice() (*big.Int, error) {
	return _Scf.Contract.GetLatestPrice(&_Scf.CallOpts)
}
