// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// CreditContractCredit is an auto generated low-level Go binding around an user-defined struct.
type CreditContractCredit struct {
	CreditAmount *big.Int
	Status       bool
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditTransactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"creditTransactionsInserted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditContract.Credit\",\"name\":\"_credit\",\"type\":\"tuple\"}],\"name\":\"setCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061053e806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633ffc2c931461005157806341eb644f1461007e578063f7522cba146100bc578063fc421ae0146100dd575b600080fd5b61006461005f366004610360565b6100f2565b604080519283529015156020830152015b60405180910390f35b6100ac61008c3660046103a5565b805160208183018101805160018252928201919093012091525460ff1681565b6040519015158152602001610075565b6100cf6100ca3660046103a5565b610140565b604051908152602001610075565b6100f06100eb3660046103e2565b610238565b005b8151602081840181018051600082529282019185019190912091905280548290811061011d57600080fd5b60009182526020909120600290910201805460019091015490925060ff16905082565b6000600182604051610152919061047b565b9081526040519081900360200190205460ff16156101a65760405162461bcd60e51b815260206004820152600d60248201526c1b9a5ac81b9bdd08199bdd5b99609a1b604482015260640160405180910390fd5b6000805b6000846040516101ba919061047b565b90815260405190819003602001902054811015610231576000846040516101e1919061047b565b90815260200160405180910390208181548110610200576102006104aa565b9060005260206000209060020201600001548261021d91906104d6565b915080610229816104ef565b9150506101aa565b5092915050565b600082604051610248919061047b565b90815260405160209181900382018120805460018082018355600092835291849020855160029092020190815592840151928101805460ff191693151593909317909255819061029990859061047b565b908152604051908190036020019020805491151560ff199092169190911790555050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126102e457600080fd5b813567ffffffffffffffff808211156102ff576102ff6102bd565b604051601f8301601f19908116603f01168101908282118183101715610327576103276102bd565b8160405283815286602085880101111561034057600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561037357600080fd5b823567ffffffffffffffff81111561038a57600080fd5b610396858286016102d3565b95602094909401359450505050565b6000602082840312156103b757600080fd5b813567ffffffffffffffff8111156103ce57600080fd5b6103da848285016102d3565b949350505050565b60008082840360608112156103f657600080fd5b833567ffffffffffffffff8082111561040e57600080fd5b61041a878388016102d3565b94506040601f198401121561042e57600080fd5b6040519250604083019150828210818311171561044d5761044d6102bd565b50604090815260208501358252840135801515811461046b57600080fd5b6020820152919491935090915050565b6000825160005b8181101561049c5760208186018101518583015201610482565b506000920191825250919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b808201808211156104e9576104e96104c0565b92915050565b600060018201610501576105016104c0565b506001019056fea2646970667358221220fc9cc3b9824086d4a6f0a6634b982e19c9d456cefd9769b171c39c2a3204d3f864736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// CreditTransactions is a free data retrieval call binding the contract method 0x3ffc2c93.
//
// Solidity: function creditTransactions(string , uint256 ) view returns(uint256 creditAmount, bool status)
func (_Api *ApiCaller) CreditTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	CreditAmount *big.Int
	Status       bool
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "creditTransactions", arg0, arg1)

	outstruct := new(struct {
		CreditAmount *big.Int
		Status       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreditAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CreditTransactions is a free data retrieval call binding the contract method 0x3ffc2c93.
//
// Solidity: function creditTransactions(string , uint256 ) view returns(uint256 creditAmount, bool status)
func (_Api *ApiSession) CreditTransactions(arg0 string, arg1 *big.Int) (struct {
	CreditAmount *big.Int
	Status       bool
}, error) {
	return _Api.Contract.CreditTransactions(&_Api.CallOpts, arg0, arg1)
}

// CreditTransactions is a free data retrieval call binding the contract method 0x3ffc2c93.
//
// Solidity: function creditTransactions(string , uint256 ) view returns(uint256 creditAmount, bool status)
func (_Api *ApiCallerSession) CreditTransactions(arg0 string, arg1 *big.Int) (struct {
	CreditAmount *big.Int
	Status       bool
}, error) {
	return _Api.Contract.CreditTransactions(&_Api.CallOpts, arg0, arg1)
}

// CreditTransactionsInserted is a free data retrieval call binding the contract method 0x41eb644f.
//
// Solidity: function creditTransactionsInserted(string ) view returns(bool)
func (_Api *ApiCaller) CreditTransactionsInserted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "creditTransactionsInserted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditTransactionsInserted is a free data retrieval call binding the contract method 0x41eb644f.
//
// Solidity: function creditTransactionsInserted(string ) view returns(bool)
func (_Api *ApiSession) CreditTransactionsInserted(arg0 string) (bool, error) {
	return _Api.Contract.CreditTransactionsInserted(&_Api.CallOpts, arg0)
}

// CreditTransactionsInserted is a free data retrieval call binding the contract method 0x41eb644f.
//
// Solidity: function creditTransactionsInserted(string ) view returns(bool)
func (_Api *ApiCallerSession) CreditTransactionsInserted(arg0 string) (bool, error) {
	return _Api.Contract.CreditTransactionsInserted(&_Api.CallOpts, arg0)
}

// GetCredit is a free data retrieval call binding the contract method 0xf7522cba.
//
// Solidity: function getCredit(string _nik) view returns(uint256)
func (_Api *ApiCaller) GetCredit(opts *bind.CallOpts, _nik string) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCredit", _nik)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCredit is a free data retrieval call binding the contract method 0xf7522cba.
//
// Solidity: function getCredit(string _nik) view returns(uint256)
func (_Api *ApiSession) GetCredit(_nik string) (*big.Int, error) {
	return _Api.Contract.GetCredit(&_Api.CallOpts, _nik)
}

// GetCredit is a free data retrieval call binding the contract method 0xf7522cba.
//
// Solidity: function getCredit(string _nik) view returns(uint256)
func (_Api *ApiCallerSession) GetCredit(_nik string) (*big.Int, error) {
	return _Api.Contract.GetCredit(&_Api.CallOpts, _nik)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiTransactor) SetCredit(opts *bind.TransactOpts, _nik string, _credit CreditContractCredit) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setCredit", _nik, _credit)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiSession) SetCredit(_nik string, _credit CreditContractCredit) (*types.Transaction, error) {
	return _Api.Contract.SetCredit(&_Api.TransactOpts, _nik, _credit)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiTransactorSession) SetCredit(_nik string, _credit CreditContractCredit) (*types.Transaction, error) {
	return _Api.Contract.SetCredit(&_Api.TransactOpts, _nik, _credit)
}
