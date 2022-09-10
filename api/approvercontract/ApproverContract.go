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

// ApproverContractApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type ApproverContractApproverTransaction struct {
	From     string
	To       string
	Status   uint8
	Filepath string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_to\",\"type\":\"string\"},{\"internalType\":\"enumApproverContract.Status\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_filepath\",\"type\":\"string\"}],\"name\":\"addApprovalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvertransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"enumApproverContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovalTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"enumApproverContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"}],\"internalType\":\"structApproverContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumApproverContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"setApproveTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumApproverContract.Status\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumApproverContract.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000f1c38038062000f1c83398101604081905262000034916200015f565b600580546001600160a01b03191633179055600062000054838262000258565b50600162000063828262000258565b506200006e62000076565b505062000324565b6005546001600160a01b031633146200008e57600080fd5b6003805460ff19169055565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c257600080fd5b81516001600160401b0380821115620000df57620000df6200009a565b604051601f8301601f19908116603f011681019082821181831017156200010a576200010a6200009a565b816040528381526020925086838588010111156200012757600080fd5b600091505b838210156200014b57858201830151818301840152908201906200012c565b600093810190920192909252949350505050565b600080604083850312156200017357600080fd5b82516001600160401b03808211156200018b57600080fd5b6200019986838701620000b0565b93506020850151915080821115620001b057600080fd5b50620001bf85828601620000b0565b9150509250929050565b600181811c90821680620001de57607f821691505b602082108103620001ff57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200025357600081815260208120601f850160051c810160208610156200022e5750805b601f850160051c820191505b818110156200024f578281556001016200023a565b5050505b505050565b81516001600160401b038111156200027457620002746200009a565b6200028c81620002858454620001c9565b8462000205565b602080601f831160018114620002c45760008415620002ab5750858301515b600019600386901b1c1916600185901b1785556200024f565b600085815260208120601f198616915b82811015620002f557888601518255948401946001909101908401620002d4565b5085821015620003145787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610be880620003346000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806383c338791161006657806383c33879146101035780638c3c4c62146101165780639ed928511461011e578063c2e3805114610131578063f851a4401461015457600080fd5b8063200d2ed2146100a35780632e49d78b146100c65780632e4b7d34146100db5780634e69d560146100e3578063641f22b8146100ee575b600080fd5b6003546100b09060ff1681565b6040516100bd9190610741565b60405180910390f35b6100d96100d4366004610769565b61017f565b005b6100d96101bc565b60035460ff166100b0565b6100f66101e2565b6040516100bd91906107d1565b6100d9610111366004610929565b61042a565b6100d96104de565b6100d961012c3660046109c2565b610501565b61014461013f3660046109ff565b610528565b6040516100bd9493929190610a18565b600554610167906001600160a01b031681565b6040516001600160a01b0390911681526020016100bd565b6005546001600160a01b0316331461019657600080fd5b6003805482919060ff1916600183838111156101b4576101b4610709565b021790555050565b6005546001600160a01b031633146101d357600080fd5b6003805460ff19166002179055565b60606004805480602002602001604051908101604052809291908181526020016000905b82821015610421578382906000526020600020906004020160405180608001604052908160008201805461023990610a69565b80601f016020809104026020016040519081016040528092919081815260200182805461026590610a69565b80156102b25780601f10610287576101008083540402835291602001916102b2565b820191906000526020600020905b81548152906001019060200180831161029557829003601f168201915b505050505081526020016001820180546102cb90610a69565b80601f01602080910402602001604051908101604052809291908181526020018280546102f790610a69565b80156103445780601f1061031957610100808354040283529160200191610344565b820191906000526020600020905b81548152906001019060200180831161032757829003601f168201915b5050509183525050600282015460209091019060ff16600381111561036b5761036b610709565b600381111561037c5761037c610709565b815260200160038201805461039090610a69565b80601f01602080910402602001604051908101604052809291908181526020018280546103bc90610a69565b80156104095780601f106103de57610100808354040283529160200191610409565b820191906000526020600020905b8154815290600101906020018083116103ec57829003601f168201915b50505050508152505081526020019060010190610206565b50505050905090565b6004604051806080016040528086815260200185815260200184600381111561045557610455610709565b8152602090810184905282546001810184556000938452922081519192600402019081906104839082610af2565b50602082015160018201906104989082610af2565b50604082015160028201805460ff191660018360038111156104bc576104bc610709565b0217905550606082015160038201906104d59082610af2565b50505050505050565b6005546001600160a01b031633146104f557600080fd5b6003805460ff19169055565b6005546001600160a01b0316331461051857600080fd5b60026105248282610af2565b5050565b6004818154811061053857600080fd5b906000526020600020906004020160009150905080600001805461055b90610a69565b80601f016020809104026020016040519081016040528092919081815260200182805461058790610a69565b80156105d45780601f106105a9576101008083540402835291602001916105d4565b820191906000526020600020905b8154815290600101906020018083116105b757829003601f168201915b5050505050908060010180546105e990610a69565b80601f016020809104026020016040519081016040528092919081815260200182805461061590610a69565b80156106625780601f1061063757610100808354040283529160200191610662565b820191906000526020600020905b81548152906001019060200180831161064557829003601f168201915b5050506002840154600385018054949560ff90921694919350915061068690610a69565b80601f01602080910402602001604051908101604052809291908181526020018280546106b290610a69565b80156106ff5780601f106106d4576101008083540402835291602001916106ff565b820191906000526020600020905b8154815290600101906020018083116106e257829003601f168201915b5050505050905084565b634e487b7160e01b600052602160045260246000fd5b6004811061073d57634e487b7160e01b600052602160045260246000fd5b9052565b6020810161074f828461071f565b92915050565b80356004811061076457600080fd5b919050565b60006020828403121561077b57600080fd5b61078482610755565b9392505050565b6000815180845260005b818110156107b157602081850181015186830182015201610795565b506000602082860101526020601f19601f83011685010191505092915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561087857603f1989840301855281516080815181865261081e8287018261078b565b915050888201518582038a870152610836828261078b565b915050878201516108498987018261071f565b5060608083015192508582038187015250610864818361078b565b9689019694505050908601906001016107f8565b509098975050505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126108ad57600080fd5b813567ffffffffffffffff808211156108c8576108c8610886565b604051601f8301601f19908116603f011681019082821181831017156108f0576108f0610886565b8160405283815286602085880101111561090957600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561093f57600080fd5b843567ffffffffffffffff8082111561095757600080fd5b6109638883890161089c565b9550602087013591508082111561097957600080fd5b6109858883890161089c565b945061099360408801610755565b935060608701359150808211156109a957600080fd5b506109b68782880161089c565b91505092959194509250565b6000602082840312156109d457600080fd5b813567ffffffffffffffff8111156109eb57600080fd5b6109f78482850161089c565b949350505050565b600060208284031215610a1157600080fd5b5035919050565b608081526000610a2b608083018761078b565b8281036020840152610a3d818761078b565b9050610a4c604084018661071f565b8281036060840152610a5e818561078b565b979650505050505050565b600181811c90821680610a7d57607f821691505b602082108103610a9d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610aed57600081815260208120601f850160051c81016020861015610aca5750805b601f850160051c820191505b81811015610ae957828155600101610ad6565b5050505b505050565b815167ffffffffffffffff811115610b0c57610b0c610886565b610b2081610b1a8454610a69565b84610aa3565b602080601f831160018114610b555760008415610b3d5750858301515b600019600386901b1c1916600185901b178555610ae9565b600085815260208120601f198616915b82811015610b8457888601518255948401946001909101908401610b65565b5085821015610ba25787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212200fa9e2104297a63d45723e64d6f88aff1872f58388b842ea9d28a174c876b61164736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _pondid string, _nik string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _pondid, _nik)
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

// Approvertransactions is a free data retrieval call binding the contract method 0xc2e38051.
//
// Solidity: function approvertransactions(uint256 ) view returns(string from, string to, uint8 status, string filepath)
func (_Api *ApiCaller) Approvertransactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From     string
	To       string
	Status   uint8
	Filepath string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "approvertransactions", arg0)

	outstruct := new(struct {
		From     string
		To       string
		Status   uint8
		Filepath string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.To = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Filepath = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Approvertransactions is a free data retrieval call binding the contract method 0xc2e38051.
//
// Solidity: function approvertransactions(uint256 ) view returns(string from, string to, uint8 status, string filepath)
func (_Api *ApiSession) Approvertransactions(arg0 *big.Int) (struct {
	From     string
	To       string
	Status   uint8
	Filepath string
}, error) {
	return _Api.Contract.Approvertransactions(&_Api.CallOpts, arg0)
}

// Approvertransactions is a free data retrieval call binding the contract method 0xc2e38051.
//
// Solidity: function approvertransactions(uint256 ) view returns(string from, string to, uint8 status, string filepath)
func (_Api *ApiCallerSession) Approvertransactions(arg0 *big.Int) (struct {
	From     string
	To       string
	Status   uint8
	Filepath string
}, error) {
	return _Api.Contract.Approvertransactions(&_Api.CallOpts, arg0)
}

// GetApprovalTransactions is a free data retrieval call binding the contract method 0x641f22b8.
//
// Solidity: function getApprovalTransactions() view returns((string,string,uint8,string)[])
func (_Api *ApiCaller) GetApprovalTransactions(opts *bind.CallOpts) ([]ApproverContractApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApprovalTransactions")

	if err != nil {
		return *new([]ApproverContractApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]ApproverContractApproverTransaction)).(*[]ApproverContractApproverTransaction)

	return out0, err

}

// GetApprovalTransactions is a free data retrieval call binding the contract method 0x641f22b8.
//
// Solidity: function getApprovalTransactions() view returns((string,string,uint8,string)[])
func (_Api *ApiSession) GetApprovalTransactions() ([]ApproverContractApproverTransaction, error) {
	return _Api.Contract.GetApprovalTransactions(&_Api.CallOpts)
}

// GetApprovalTransactions is a free data retrieval call binding the contract method 0x641f22b8.
//
// Solidity: function getApprovalTransactions() view returns((string,string,uint8,string)[])
func (_Api *ApiCallerSession) GetApprovalTransactions() ([]ApproverContractApproverTransaction, error) {
	return _Api.Contract.GetApprovalTransactions(&_Api.CallOpts)
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

// AddApprovalTransaction is a paid mutator transaction binding the contract method 0x83c33879.
//
// Solidity: function addApprovalTransaction(string _from, string _to, uint8 _status, string _filepath) returns()
func (_Api *ApiTransactor) AddApprovalTransaction(opts *bind.TransactOpts, _from string, _to string, _status uint8, _filepath string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addApprovalTransaction", _from, _to, _status, _filepath)
}

// AddApprovalTransaction is a paid mutator transaction binding the contract method 0x83c33879.
//
// Solidity: function addApprovalTransaction(string _from, string _to, uint8 _status, string _filepath) returns()
func (_Api *ApiSession) AddApprovalTransaction(_from string, _to string, _status uint8, _filepath string) (*types.Transaction, error) {
	return _Api.Contract.AddApprovalTransaction(&_Api.TransactOpts, _from, _to, _status, _filepath)
}

// AddApprovalTransaction is a paid mutator transaction binding the contract method 0x83c33879.
//
// Solidity: function addApprovalTransaction(string _from, string _to, uint8 _status, string _filepath) returns()
func (_Api *ApiTransactorSession) AddApprovalTransaction(_from string, _to string, _status uint8, _filepath string) (*types.Transaction, error) {
	return _Api.Contract.AddApprovalTransaction(&_Api.TransactOpts, _from, _to, _status, _filepath)
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
