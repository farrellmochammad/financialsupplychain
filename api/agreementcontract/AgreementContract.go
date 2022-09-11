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

// AgreementContractAgreementHistory is an auto generated low-level Go binding around an user-defined struct.
type AgreementContractAgreementHistory struct {
	Timestamp        string
	Approverfunderid string
	Monitoringid     string
	Approverid       string
	Signed           string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_agreementid\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_approverfunderid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_monitoringid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_approverid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signed\",\"type\":\"string\"}],\"name\":\"addAgreement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"agreementhistories\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approverfunderid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"monitoringid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approverid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signed\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAgreements\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approverfunderid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"monitoringid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approverid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signed\",\"type\":\"string\"}],\"internalType\":\"structAgreementContract.AgreementHistory[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000f0f38038062000f0f833981016040819052620000349162000072565b600180546001600160a01b031916331790556000620000548282620001d6565b5050620002a2565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200008657600080fd5b82516001600160401b03808211156200009e57600080fd5b818501915085601f830112620000b357600080fd5b815181811115620000c857620000c86200005c565b604051601f8201601f19908116603f01168101908382118183101715620000f357620000f36200005c565b8160405282815288868487010111156200010c57600080fd5b600093505b8284101562000130578484018601518185018701529285019262000111565b600086848301015280965050505050505092915050565b600181811c908216806200015c57607f821691505b6020821081036200017d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001d157600081815260208120601f850160051c81016020861015620001ac5750805b601f850160051c820191505b81811015620001cd57828155600101620001b8565b5050505b505050565b81516001600160401b03811115620001f257620001f26200005c565b6200020a8162000203845462000147565b8462000183565b602080601f831160018114620002425760008415620002295750858301515b600019600386901b1c1916600185901b178555620001cd565b600085815260208120601f198616915b82811015620002735788860151825594840194600190910190840162000252565b5085821015620002925787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610c5d80620002b26000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633dd1721e1461005157806394c0d35314610066578063cf7e924214610093578063f851a440146100a8575b600080fd5b61006461005f366004610869565b6100d3565b005b61007961007436600461093b565b61019f565b60405161008a95949392919061099a565b60405180910390f35b61009b61048d565b60405161008a9190610a07565b6001546100bb906001600160a01b031681565b6040516001600160a01b03909116815260200161008a565b6040805160a08101825286815260208101869052908101849052606081018390526080810182905260028054600181018255600091909152815160059091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace019081906101419082610b67565b50602082015160018201906101569082610b67565b506040820151600282019061016b9082610b67565b50606082015160038201906101809082610b67565b50608082015160048201906101959082610b67565b5050505050505050565b600281815481106101af57600080fd5b90600052602060002090600502016000915090508060000180546101d290610ade565b80601f01602080910402602001604051908101604052809291908181526020018280546101fe90610ade565b801561024b5780601f106102205761010080835404028352916020019161024b565b820191906000526020600020905b81548152906001019060200180831161022e57829003601f168201915b50505050509080600101805461026090610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461028c90610ade565b80156102d95780601f106102ae576101008083540402835291602001916102d9565b820191906000526020600020905b8154815290600101906020018083116102bc57829003601f168201915b5050505050908060020180546102ee90610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461031a90610ade565b80156103675780601f1061033c57610100808354040283529160200191610367565b820191906000526020600020905b81548152906001019060200180831161034a57829003601f168201915b50505050509080600301805461037c90610ade565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890610ade565b80156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b50505050509080600401805461040a90610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461043690610ade565b80156104835780601f1061045857610100808354040283529160200191610483565b820191906000526020600020905b81548152906001019060200180831161046657829003601f168201915b5050505050905085565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156107bd57838290600052602060002090600502016040518060a00160405290816000820180546104e490610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461051090610ade565b801561055d5780601f106105325761010080835404028352916020019161055d565b820191906000526020600020905b81548152906001019060200180831161054057829003601f168201915b5050505050815260200160018201805461057690610ade565b80601f01602080910402602001604051908101604052809291908181526020018280546105a290610ade565b80156105ef5780601f106105c4576101008083540402835291602001916105ef565b820191906000526020600020905b8154815290600101906020018083116105d257829003601f168201915b5050505050815260200160028201805461060890610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461063490610ade565b80156106815780601f1061065657610100808354040283529160200191610681565b820191906000526020600020905b81548152906001019060200180831161066457829003601f168201915b5050505050815260200160038201805461069a90610ade565b80601f01602080910402602001604051908101604052809291908181526020018280546106c690610ade565b80156107135780601f106106e857610100808354040283529160200191610713565b820191906000526020600020905b8154815290600101906020018083116106f657829003601f168201915b5050505050815260200160048201805461072c90610ade565b80601f016020809104026020016040519081016040528092919081815260200182805461075890610ade565b80156107a55780601f1061077a576101008083540402835291602001916107a5565b820191906000526020600020905b81548152906001019060200180831161078857829003601f168201915b505050505081525050815260200190600101906104b1565b50505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126107ed57600080fd5b813567ffffffffffffffff80821115610808576108086107c6565b604051601f8301601f19908116603f01168101908282118183101715610830576108306107c6565b8160405283815286602085880101111561084957600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a0868803121561088157600080fd5b853567ffffffffffffffff8082111561089957600080fd5b6108a589838a016107dc565b965060208801359150808211156108bb57600080fd5b6108c789838a016107dc565b955060408801359150808211156108dd57600080fd5b6108e989838a016107dc565b945060608801359150808211156108ff57600080fd5b61090b89838a016107dc565b9350608088013591508082111561092157600080fd5b5061092e888289016107dc565b9150509295509295909350565b60006020828403121561094d57600080fd5b5035919050565b6000815180845260005b8181101561097a5760208185018101518683018201520161095e565b506000602082860101526020601f19601f83011685010191505092915050565b60a0815260006109ad60a0830188610954565b82810360208401526109bf8188610954565b905082810360408401526109d38187610954565b905082810360608401526109e78186610954565b905082810360808401526109fb8185610954565b98975050505050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015610ad057603f19898403018552815160a08151818652610a5482870182610954565b915050888201518582038a870152610a6c8282610954565b9150508782015185820389870152610a848282610954565b91505060608083015186830382880152610a9e8382610954565b9250505060808083015192508582038187015250610abc8183610954565b968901969450505090860190600101610a2e565b509098975050505050505050565b600181811c90821680610af257607f821691505b602082108103610b1257634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b6257600081815260208120601f850160051c81016020861015610b3f5750805b601f850160051c820191505b81811015610b5e57828155600101610b4b565b5050505b505050565b815167ffffffffffffffff811115610b8157610b816107c6565b610b9581610b8f8454610ade565b84610b18565b602080601f831160018114610bca5760008415610bb25750858301515b600019600386901b1c1916600185901b178555610b5e565b600085815260208120601f198616915b82811015610bf957888601518255948401946001909101908401610bda565b5085821015610c175787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220ba29c17d44d09f78d4159b2434eec521605d1a9a62ee55c0f50e8f2afb01fdd864736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _agreementid string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _agreementid)
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCallerSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Agreementhistories is a free data retrieval call binding the contract method 0x94c0d353.
//
// Solidity: function agreementhistories(uint256 ) view returns(string timestamp, string approverfunderid, string monitoringid, string approverid, string signed)
func (_Api *ApiCaller) Agreementhistories(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp        string
	Approverfunderid string
	Monitoringid     string
	Approverid       string
	Signed           string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "agreementhistories", arg0)

	outstruct := new(struct {
		Timestamp        string
		Approverfunderid string
		Monitoringid     string
		Approverid       string
		Signed           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Approverfunderid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Monitoringid = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Approverid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Signed = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Agreementhistories is a free data retrieval call binding the contract method 0x94c0d353.
//
// Solidity: function agreementhistories(uint256 ) view returns(string timestamp, string approverfunderid, string monitoringid, string approverid, string signed)
func (_Api *ApiSession) Agreementhistories(arg0 *big.Int) (struct {
	Timestamp        string
	Approverfunderid string
	Monitoringid     string
	Approverid       string
	Signed           string
}, error) {
	return _Api.Contract.Agreementhistories(&_Api.CallOpts, arg0)
}

// Agreementhistories is a free data retrieval call binding the contract method 0x94c0d353.
//
// Solidity: function agreementhistories(uint256 ) view returns(string timestamp, string approverfunderid, string monitoringid, string approverid, string signed)
func (_Api *ApiCallerSession) Agreementhistories(arg0 *big.Int) (struct {
	Timestamp        string
	Approverfunderid string
	Monitoringid     string
	Approverid       string
	Signed           string
}, error) {
	return _Api.Contract.Agreementhistories(&_Api.CallOpts, arg0)
}

// GetAgreements is a free data retrieval call binding the contract method 0xcf7e9242.
//
// Solidity: function getAgreements() view returns((string,string,string,string,string)[])
func (_Api *ApiCaller) GetAgreements(opts *bind.CallOpts) ([]AgreementContractAgreementHistory, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAgreements")

	if err != nil {
		return *new([]AgreementContractAgreementHistory), err
	}

	out0 := *abi.ConvertType(out[0], new([]AgreementContractAgreementHistory)).(*[]AgreementContractAgreementHistory)

	return out0, err

}

// GetAgreements is a free data retrieval call binding the contract method 0xcf7e9242.
//
// Solidity: function getAgreements() view returns((string,string,string,string,string)[])
func (_Api *ApiSession) GetAgreements() ([]AgreementContractAgreementHistory, error) {
	return _Api.Contract.GetAgreements(&_Api.CallOpts)
}

// GetAgreements is a free data retrieval call binding the contract method 0xcf7e9242.
//
// Solidity: function getAgreements() view returns((string,string,string,string,string)[])
func (_Api *ApiCallerSession) GetAgreements() ([]AgreementContractAgreementHistory, error) {
	return _Api.Contract.GetAgreements(&_Api.CallOpts)
}

// AddAgreement is a paid mutator transaction binding the contract method 0x3dd1721e.
//
// Solidity: function addAgreement(string _timestamp, string _approverfunderid, string _monitoringid, string _approverid, string _signed) returns()
func (_Api *ApiTransactor) AddAgreement(opts *bind.TransactOpts, _timestamp string, _approverfunderid string, _monitoringid string, _approverid string, _signed string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addAgreement", _timestamp, _approverfunderid, _monitoringid, _approverid, _signed)
}

// AddAgreement is a paid mutator transaction binding the contract method 0x3dd1721e.
//
// Solidity: function addAgreement(string _timestamp, string _approverfunderid, string _monitoringid, string _approverid, string _signed) returns()
func (_Api *ApiSession) AddAgreement(_timestamp string, _approverfunderid string, _monitoringid string, _approverid string, _signed string) (*types.Transaction, error) {
	return _Api.Contract.AddAgreement(&_Api.TransactOpts, _timestamp, _approverfunderid, _monitoringid, _approverid, _signed)
}

// AddAgreement is a paid mutator transaction binding the contract method 0x3dd1721e.
//
// Solidity: function addAgreement(string _timestamp, string _approverfunderid, string _monitoringid, string _approverid, string _signed) returns()
func (_Api *ApiTransactorSession) AddAgreement(_timestamp string, _approverfunderid string, _monitoringid string, _approverid string, _signed string) (*types.Transaction, error) {
	return _Api.Contract.AddAgreement(&_Api.TransactOpts, _timestamp, _approverfunderid, _monitoringid, _approverid, _signed)
}
