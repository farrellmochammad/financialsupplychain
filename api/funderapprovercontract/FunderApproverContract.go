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

// FunderApproverContractFunderApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type FunderApproverContractFunderApproverTransaction struct {
	Status      uint8
	Filepath    string
	Timestamp   string
	Modalgiven  *big.Int
	Creditscore *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_funderapproverid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_agentapproverid\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_modalgiven\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creditscore\",\"type\":\"uint256\"}],\"name\":\"addFunderApprovalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderapprovertransactions\",\"outputs\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFunderApprovalTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structFunderApproverContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"setApproveTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumFunderApproverContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000dfa38038062000dfa83398101604081905262000034916200015f565b600580546001600160a01b03191633179055600062000054838262000258565b50600162000063828262000258565b506200006e62000076565b505062000324565b6005546001600160a01b031633146200008e57600080fd5b6003805460ff19169055565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c257600080fd5b81516001600160401b0380821115620000df57620000df6200009a565b604051601f8301601f19908116603f011681019082821181831017156200010a576200010a6200009a565b816040528381526020925086838588010111156200012757600080fd5b600091505b838210156200014b57858201830151818301840152908201906200012c565b600093810190920192909252949350505050565b600080604083850312156200017357600080fd5b82516001600160401b03808211156200018b57600080fd5b6200019986838701620000b0565b93506020850151915080821115620001b057600080fd5b50620001bf85828601620000b0565b9150509250929050565b600181811c90821680620001de57607f821691505b602082108103620001ff57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200025357600081815260208120601f850160051c810160208610156200022e5750805b601f850160051c820191505b818110156200024f578281556001016200023a565b5050505b505050565b81516001600160401b038111156200027457620002746200009a565b6200028c81620002858454620001c9565b8462000205565b602080601f831160018114620002c45760008415620002ab5750858301515b600019600386901b1c1916600185901b1785556200024f565b600085815260208120601f198616915b82811015620002f557888601518255948401946001909101908401620002d4565b5085821015620003145787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610ac680620003346000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637b9b9e4e116100665780637b9b9e4e146101035780638c3c4c62146101275780639ed928511461012f578063e11fb98814610142578063f851a4401461015557600080fd5b8063200d2ed2146100a35780632e49d78b146100c65780632e4b7d34146100db5780632f12f5d2146100e35780634e69d560146100f8575b600080fd5b6003546100b09060ff1681565b6040516100bd9190610645565b60405180910390f35b6100d96100d436600461066d565b610180565b005b6100d96101bd565b6100eb6101e3565b6040516100bd91906106d5565b60035460ff166100b0565b610116610111366004610781565b6103a7565b6040516100bd95949392919061079a565b6100d96104fd565b6100d961013d366004610884565b610520565b6100d96101503660046108c1565b610547565b600554610168906001600160a01b031681565b6040516001600160a01b0390911681526020016100bd565b6005546001600160a01b0316331461019757600080fd5b6003805482919060ff1916600183838111156101b5576101b561060d565b021790555050565b6005546001600160a01b031633146101d457600080fd5b6003805460ff19166002179055565b60606004805480602002602001604051908101604052809291908181526020016000905b8282101561039e576000848152602090206040805160a08101909152600584029091018054829060ff1660038111156102425761024261060d565b60038111156102535761025361060d565b815260200160018201805461026790610947565b80601f016020809104026020016040519081016040528092919081815260200182805461029390610947565b80156102e05780601f106102b5576101008083540402835291602001916102e0565b820191906000526020600020905b8154815290600101906020018083116102c357829003601f168201915b505050505081526020016002820180546102f990610947565b80601f016020809104026020016040519081016040528092919081815260200182805461032590610947565b80156103725780601f1061034757610100808354040283529160200191610372565b820191906000526020600020905b81548152906001019060200180831161035557829003601f168201915b505050505081526020016003820154815260200160048201548152505081526020019060010190610207565b50505050905090565b600481815481106103b757600080fd5b60009182526020909120600590910201805460018201805460ff9092169350906103e090610947565b80601f016020809104026020016040519081016040528092919081815260200182805461040c90610947565b80156104595780601f1061042e57610100808354040283529160200191610459565b820191906000526020600020905b81548152906001019060200180831161043c57829003601f168201915b50505050509080600201805461046e90610947565b80601f016020809104026020016040519081016040528092919081815260200182805461049a90610947565b80156104e75780601f106104bc576101008083540402835291602001916104e7565b820191906000526020600020905b8154815290600101906020018083116104ca57829003601f168201915b5050505050908060030154908060040154905085565b6005546001600160a01b0316331461051457600080fd5b6003805460ff19169055565b6005546001600160a01b0316331461053757600080fd5b600261054382826109d0565b5050565b60046040518060a001604052808760038111156105665761056661060d565b815260208082018890526040820187905260608201869052608090910184905282546001818101855560009485529190932082516005909402018054929390929091839160ff1916908360038111156105c1576105c161060d565b0217905550602082015160018201906105da90826109d0565b50604082015160028201906105ef90826109d0565b50606082015181600301556080820151816004015550505050505050565b634e487b7160e01b600052602160045260246000fd5b6004811061064157634e487b7160e01b600052602160045260246000fd5b9052565b602081016106538284610623565b92915050565b80356004811061066857600080fd5b919050565b60006020828403121561067f57600080fd5b61068882610659565b9392505050565b6000815180845260005b818110156106b557602081850181015186830182015201610699565b506000602082860101526020601f19601f83011685010191505092915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561077357603f19898403018552815160a061071c858351610623565b88820151818a8701526107318287018261068f565b9150508782015185820389870152610749828261068f565b606084810151908801526080938401519390960192909252505093860193908601906001016106fc565b509098975050505050505050565b60006020828403121561079357600080fd5b5035919050565b6107a48187610623565b60a0602082015260006107ba60a083018761068f565b82810360408401526107cc818761068f565b60608401959095525050608001529392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261080857600080fd5b813567ffffffffffffffff80821115610823576108236107e1565b604051601f8301601f19908116603f0116810190828211818310171561084b5761084b6107e1565b8160405283815286602085880101111561086457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561089657600080fd5b813567ffffffffffffffff8111156108ad57600080fd5b6108b9848285016107f7565b949350505050565b600080600080600060a086880312156108d957600080fd5b6108e286610659565b9450602086013567ffffffffffffffff808211156108ff57600080fd5b61090b89838a016107f7565b9550604088013591508082111561092157600080fd5b5061092e888289016107f7565b9598949750949560608101359550608001359392505050565b600181811c9082168061095b57607f821691505b60208210810361097b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156109cb57600081815260208120601f850160051c810160208610156109a85750805b601f850160051c820191505b818110156109c7578281556001016109b4565b5050505b505050565b815167ffffffffffffffff8111156109ea576109ea6107e1565b6109fe816109f88454610947565b84610981565b602080601f831160018114610a335760008415610a1b5750858301515b600019600386901b1c1916600185901b1785556109c7565b600085815260208120601f198616915b82811015610a6257888601518255948401946001909101908401610a43565b5085821015610a805787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220684ec667c310aefde68a68168568ce0a3df94ea3826c3ad1a227214391f499a164736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _funderapproverid string, _agentapproverid string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _funderapproverid, _agentapproverid)
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

// Funderapprovertransactions is a free data retrieval call binding the contract method 0x7b9b9e4e.
//
// Solidity: function funderapprovertransactions(uint256 ) view returns(uint8 status, string filepath, string timestamp, uint256 modalgiven, uint256 creditscore)
func (_Api *ApiCaller) Funderapprovertransactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status      uint8
	Filepath    string
	Timestamp   string
	Modalgiven  *big.Int
	Creditscore *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderapprovertransactions", arg0)

	outstruct := new(struct {
		Status      uint8
		Filepath    string
		Timestamp   string
		Modalgiven  *big.Int
		Creditscore *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Filepath = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Modalgiven = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Creditscore = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Funderapprovertransactions is a free data retrieval call binding the contract method 0x7b9b9e4e.
//
// Solidity: function funderapprovertransactions(uint256 ) view returns(uint8 status, string filepath, string timestamp, uint256 modalgiven, uint256 creditscore)
func (_Api *ApiSession) Funderapprovertransactions(arg0 *big.Int) (struct {
	Status      uint8
	Filepath    string
	Timestamp   string
	Modalgiven  *big.Int
	Creditscore *big.Int
}, error) {
	return _Api.Contract.Funderapprovertransactions(&_Api.CallOpts, arg0)
}

// Funderapprovertransactions is a free data retrieval call binding the contract method 0x7b9b9e4e.
//
// Solidity: function funderapprovertransactions(uint256 ) view returns(uint8 status, string filepath, string timestamp, uint256 modalgiven, uint256 creditscore)
func (_Api *ApiCallerSession) Funderapprovertransactions(arg0 *big.Int) (struct {
	Status      uint8
	Filepath    string
	Timestamp   string
	Modalgiven  *big.Int
	Creditscore *big.Int
}, error) {
	return _Api.Contract.Funderapprovertransactions(&_Api.CallOpts, arg0)
}

// GetFunderApprovalTransactions is a free data retrieval call binding the contract method 0x2f12f5d2.
//
// Solidity: function getFunderApprovalTransactions() view returns((uint8,string,string,uint256,uint256)[])
func (_Api *ApiCaller) GetFunderApprovalTransactions(opts *bind.CallOpts) ([]FunderApproverContractFunderApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getFunderApprovalTransactions")

	if err != nil {
		return *new([]FunderApproverContractFunderApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]FunderApproverContractFunderApproverTransaction)).(*[]FunderApproverContractFunderApproverTransaction)

	return out0, err

}

// GetFunderApprovalTransactions is a free data retrieval call binding the contract method 0x2f12f5d2.
//
// Solidity: function getFunderApprovalTransactions() view returns((uint8,string,string,uint256,uint256)[])
func (_Api *ApiSession) GetFunderApprovalTransactions() ([]FunderApproverContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApprovalTransactions(&_Api.CallOpts)
}

// GetFunderApprovalTransactions is a free data retrieval call binding the contract method 0x2f12f5d2.
//
// Solidity: function getFunderApprovalTransactions() view returns((uint8,string,string,uint256,uint256)[])
func (_Api *ApiCallerSession) GetFunderApprovalTransactions() ([]FunderApproverContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApprovalTransactions(&_Api.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Api *ApiCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Api *ApiSession) GetStatus() (uint8, error) {
	return _Api.Contract.GetStatus(&_Api.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Api *ApiCallerSession) GetStatus() (uint8, error) {
	return _Api.Contract.GetStatus(&_Api.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Api *ApiCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Api *ApiSession) Status() (uint8, error) {
	return _Api.Contract.Status(&_Api.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Api *ApiCallerSession) Status() (uint8, error) {
	return _Api.Contract.Status(&_Api.CallOpts)
}

// AddFunderApprovalTransaction is a paid mutator transaction binding the contract method 0xe11fb988.
//
// Solidity: function addFunderApprovalTransaction(uint8 _status, string _filepath, string _timestamp, uint256 _modalgiven, uint256 _creditscore) returns()
func (_Api *ApiTransactor) AddFunderApprovalTransaction(opts *bind.TransactOpts, _status uint8, _filepath string, _timestamp string, _modalgiven *big.Int, _creditscore *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addFunderApprovalTransaction", _status, _filepath, _timestamp, _modalgiven, _creditscore)
}

// AddFunderApprovalTransaction is a paid mutator transaction binding the contract method 0xe11fb988.
//
// Solidity: function addFunderApprovalTransaction(uint8 _status, string _filepath, string _timestamp, uint256 _modalgiven, uint256 _creditscore) returns()
func (_Api *ApiSession) AddFunderApprovalTransaction(_status uint8, _filepath string, _timestamp string, _modalgiven *big.Int, _creditscore *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddFunderApprovalTransaction(&_Api.TransactOpts, _status, _filepath, _timestamp, _modalgiven, _creditscore)
}

// AddFunderApprovalTransaction is a paid mutator transaction binding the contract method 0xe11fb988.
//
// Solidity: function addFunderApprovalTransaction(uint8 _status, string _filepath, string _timestamp, uint256 _modalgiven, uint256 _creditscore) returns()
func (_Api *ApiTransactorSession) AddFunderApprovalTransaction(_status uint8, _filepath string, _timestamp string, _modalgiven *big.Int, _creditscore *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddFunderApprovalTransaction(&_Api.TransactOpts, _status, _filepath, _timestamp, _modalgiven, _creditscore)
}

// RejectStatus is a paid mutator transaction binding the contract method 0x2e4b7d34.
//
// Solidity: function rejectStatus() returns()
func (_Api *ApiTransactor) RejectStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "rejectStatus")
}

// RejectStatus is a paid mutator transaction binding the contract method 0x2e4b7d34.
//
// Solidity: function rejectStatus() returns()
func (_Api *ApiSession) RejectStatus() (*types.Transaction, error) {
	return _Api.Contract.RejectStatus(&_Api.TransactOpts)
}

// RejectStatus is a paid mutator transaction binding the contract method 0x2e4b7d34.
//
// Solidity: function rejectStatus() returns()
func (_Api *ApiTransactorSession) RejectStatus() (*types.Transaction, error) {
	return _Api.Contract.RejectStatus(&_Api.TransactOpts)
}

// ResetStatus is a paid mutator transaction binding the contract method 0x8c3c4c62.
//
// Solidity: function resetStatus() returns()
func (_Api *ApiTransactor) ResetStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "resetStatus")
}

// ResetStatus is a paid mutator transaction binding the contract method 0x8c3c4c62.
//
// Solidity: function resetStatus() returns()
func (_Api *ApiSession) ResetStatus() (*types.Transaction, error) {
	return _Api.Contract.ResetStatus(&_Api.TransactOpts)
}

// ResetStatus is a paid mutator transaction binding the contract method 0x8c3c4c62.
//
// Solidity: function resetStatus() returns()
func (_Api *ApiTransactorSession) ResetStatus() (*types.Transaction, error) {
	return _Api.Contract.ResetStatus(&_Api.TransactOpts)
}

// SetApproveTimestamp is a paid mutator transaction binding the contract method 0x9ed92851.
//
// Solidity: function setApproveTimestamp(string _timestamp) returns()
func (_Api *ApiTransactor) SetApproveTimestamp(opts *bind.TransactOpts, _timestamp string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApproveTimestamp", _timestamp)
}

// SetApproveTimestamp is a paid mutator transaction binding the contract method 0x9ed92851.
//
// Solidity: function setApproveTimestamp(string _timestamp) returns()
func (_Api *ApiSession) SetApproveTimestamp(_timestamp string) (*types.Transaction, error) {
	return _Api.Contract.SetApproveTimestamp(&_Api.TransactOpts, _timestamp)
}

// SetApproveTimestamp is a paid mutator transaction binding the contract method 0x9ed92851.
//
// Solidity: function setApproveTimestamp(string _timestamp) returns()
func (_Api *ApiTransactorSession) SetApproveTimestamp(_timestamp string) (*types.Transaction, error) {
	return _Api.Contract.SetApproveTimestamp(&_Api.TransactOpts, _timestamp)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_Api *ApiTransactor) SetStatus(opts *bind.TransactOpts, _status uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setStatus", _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_Api *ApiSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _Api.Contract.SetStatus(&_Api.TransactOpts, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x2e49d78b.
//
// Solidity: function setStatus(uint8 _status) returns()
func (_Api *ApiTransactorSession) SetStatus(_status uint8) (*types.Transaction, error) {
	return _Api.Contract.SetStatus(&_Api.TransactOpts, _status)
}
