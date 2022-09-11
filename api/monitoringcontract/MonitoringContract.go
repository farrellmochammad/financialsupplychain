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

// MonitoringContractMonitoring is an auto generated low-level Go binding around an user-defined struct.
type MonitoringContractMonitoring struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}

// MonitoringContractSpawning is an auto generated low-level Go binding around an user-defined struct.
type MonitoringContractSpawning struct {
	Timestamp string
	Weight    *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_monitoringid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pondname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_area\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_district\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fishtype\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_humidity\",\"type\":\"uint256\"}],\"name\":\"addMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"addSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMonitorings\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structMonitoringContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpawnings\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structMonitoringContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"monitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawnings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e0b38038062000e0b833981016040819052620000349162000171565b600880546001600160a01b03191633179055600062000054878262000307565b50600162000063868262000307565b50600262000072858262000307565b50600362000081848262000307565b50600462000090838262000307565b5060056200009f828262000307565b50505050505050620003d3565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000d457600080fd5b81516001600160401b0380821115620000f157620000f1620000ac565b604051601f8301601f19908116603f011681019082821181831017156200011c576200011c620000ac565b816040528381526020925086838588010111156200013957600080fd5b600091505b838210156200015d57858201830151818301840152908201906200013e565b600093810190920192909252949350505050565b60008060008060008060c087890312156200018b57600080fd5b86516001600160401b0380821115620001a357600080fd5b620001b18a838b01620000c2565b97506020890151915080821115620001c857600080fd5b620001d68a838b01620000c2565b96506040890151915080821115620001ed57600080fd5b620001fb8a838b01620000c2565b955060608901519150808211156200021257600080fd5b620002208a838b01620000c2565b945060808901519150808211156200023757600080fd5b620002458a838b01620000c2565b935060a08901519150808211156200025c57600080fd5b506200026b89828a01620000c2565b9150509295509295509295565b600181811c908216806200028d57607f821691505b602082108103620002ae57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200030257600081815260208120601f850160051c81016020861015620002dd5750805b601f850160051c820191505b81811015620002fe57828155600101620002e9565b5050505b505050565b81516001600160401b03811115620003235762000323620000ac565b6200033b8162000334845462000278565b84620002b4565b602080601f8311600181146200037357600084156200035a5750858301515b600019600386901b1c1916600185901b178555620002fe565b600085815260208120601f198616915b82811015620003a45788860151825594840194600190910190840162000383565b5085821015620003c35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610a2880620003e36000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063e939f82c1161005b578063e939f82c146100d6578063ee8e1d51146100eb578063f400d7ca1461010e578063f851a4401461012157600080fd5b806345086489146100825780637092f1fb146100a0578063de97090c146100c1575b600080fd5b61008a61014c565b6040516100979190610610565b60405180910390f35b6100b36100ae36600461069b565b61025b565b6040516100979291906106b4565b6100d46100cf366004610779565b610317565b005b6100de610382565b60405161009791906107be565b6100fe6100f936600461069b565b610474565b6040516100979493929190610826565b6100d461011c366004610855565b61053c565b600854610134906001600160a01b031681565b6040516001600160a01b039091168152602001610097565b60606006805480602002602001604051908101604052809291908181526020016000905b8282101561025257838290600052602060002090600402016040518060800160405290816000820180546101a3906108a9565b80601f01602080910402602001604051908101604052809291908181526020018280546101cf906108a9565b801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152505081526020019060010190610170565b50505050905090565b6007818154811061026b57600080fd5b906000526020600020906002020160009150905080600001805461028e906108a9565b80601f01602080910402602001604051908101604052809291908181526020018280546102ba906108a9565b80156103075780601f106102dc57610100808354040283529160200191610307565b820191906000526020600020905b8154815290600101906020018083116102ea57829003601f168201915b5050505050908060010154905082565b604080518082019091528281526020810182905260078054600181018255600091909152815160029091027fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688019081906103719082610932565b506020820151816001015550505050565b60606007805480602002602001604051908101604052809291908181526020016000905b8282101561025257838290600052602060002090600202016040518060400160405290816000820180546103d9906108a9565b80601f0160208091040260200160405190810160405280929190818152602001828054610405906108a9565b80156104525780601f1061042757610100808354040283529160200191610452565b820191906000526020600020905b81548152906001019060200180831161043557829003601f168201915b50505050508152602001600182015481525050815260200190600101906103a6565b6006818154811061048457600080fd5b90600052602060002090600402016000915090508060000180546104a7906108a9565b80601f01602080910402602001604051908101604052809291908181526020018280546104d3906108a9565b80156105205780601f106104f557610100808354040283529160200191610520565b820191906000526020600020905b81548152906001019060200180831161050357829003601f168201915b5050505050908060010154908060020154908060030154905084565b60408051608081018252858152602081018590529081018390526060810182905260068054600181018255600091909152815160049091027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f019081906105a39082610932565b50602082015181600101556040820151816002015560608201518160030155505050505050565b6000815180845260005b818110156105f0576020818501810151868301820152016105d4565b506000602082860101526020601f19601f83011685010191505092915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561068d57603f1989840301855281516080815181865261065d828701826105ca565b838b0151878c0152898401518a880152606093840151939096019290925250509386019390860190600101610637565b509098975050505050505050565b6000602082840312156106ad57600080fd5b5035919050565b6040815260006106c760408301856105ca565b90508260208301529392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126106fd57600080fd5b813567ffffffffffffffff80821115610718576107186106d6565b604051601f8301601f19908116603f01168101908282118183101715610740576107406106d6565b8160405283815286602085880101111561075957600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561078c57600080fd5b823567ffffffffffffffff8111156107a357600080fd5b6107af858286016106ec565b95602094909401359450505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561068d57888303603f1901855281518051878552610809888601826105ca565b9189015194890194909452948701949250908601906001016107e5565b60808152600061083960808301876105ca565b6020830195909552506040810192909252606090910152919050565b6000806000806080858703121561086b57600080fd5b843567ffffffffffffffff81111561088257600080fd5b61088e878288016106ec565b97602087013597506040870135966060013595509350505050565b600181811c908216806108bd57607f821691505b6020821081036108dd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561092d57600081815260208120601f850160051c8101602086101561090a5750805b601f850160051c820191505b8181101561092957828155600101610916565b5050505b505050565b815167ffffffffffffffff81111561094c5761094c6106d6565b6109608161095a84546108a9565b846108e3565b602080601f831160018114610995576000841561097d5750858301515b600019600386901b1c1916600185901b178555610929565b600085815260208120601f198616915b828110156109c4578886015182559484019460019091019084016109a5565b50858210156109e25787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea264697066735822122037a8568f294ef29a6f8bdc3cc348beceb5d1422fc5c9b3135c20bc7c6161c3f764736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _monitoringid string, _pondid string, _pondname string, _area string, _district string, _fishtype string) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _monitoringid, _pondid, _pondname, _area, _district, _fishtype)
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

// GetMonitorings is a free data retrieval call binding the contract method 0x45086489.
//
// Solidity: function getMonitorings() view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCaller) GetMonitorings(opts *bind.CallOpts) ([]MonitoringContractMonitoring, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getMonitorings")

	if err != nil {
		return *new([]MonitoringContractMonitoring), err
	}

	out0 := *abi.ConvertType(out[0], new([]MonitoringContractMonitoring)).(*[]MonitoringContractMonitoring)

	return out0, err

}

// GetMonitorings is a free data retrieval call binding the contract method 0x45086489.
//
// Solidity: function getMonitorings() view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetMonitorings() ([]MonitoringContractMonitoring, error) {
	return _Api.Contract.GetMonitorings(&_Api.CallOpts)
}

// GetMonitorings is a free data retrieval call binding the contract method 0x45086489.
//
// Solidity: function getMonitorings() view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetMonitorings() ([]MonitoringContractMonitoring, error) {
	return _Api.Contract.GetMonitorings(&_Api.CallOpts)
}

// GetSpawnings is a free data retrieval call binding the contract method 0xe939f82c.
//
// Solidity: function getSpawnings() view returns((string,uint256)[])
func (_Api *ApiCaller) GetSpawnings(opts *bind.CallOpts) ([]MonitoringContractSpawning, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSpawnings")

	if err != nil {
		return *new([]MonitoringContractSpawning), err
	}

	out0 := *abi.ConvertType(out[0], new([]MonitoringContractSpawning)).(*[]MonitoringContractSpawning)

	return out0, err

}

// GetSpawnings is a free data retrieval call binding the contract method 0xe939f82c.
//
// Solidity: function getSpawnings() view returns((string,uint256)[])
func (_Api *ApiSession) GetSpawnings() ([]MonitoringContractSpawning, error) {
	return _Api.Contract.GetSpawnings(&_Api.CallOpts)
}

// GetSpawnings is a free data retrieval call binding the contract method 0xe939f82c.
//
// Solidity: function getSpawnings() view returns((string,uint256)[])
func (_Api *ApiCallerSession) GetSpawnings() ([]MonitoringContractSpawning, error) {
	return _Api.Contract.GetSpawnings(&_Api.CallOpts)
}

// Monitorings is a free data retrieval call binding the contract method 0xee8e1d51.
//
// Solidity: function monitorings(uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiCaller) Monitorings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "monitorings", arg0)

	outstruct := new(struct {
		Timestamp   string
		Weight      *big.Int
		Temperature *big.Int
		Humidity    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Temperature = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Humidity = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Monitorings is a free data retrieval call binding the contract method 0xee8e1d51.
//
// Solidity: function monitorings(uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiSession) Monitorings(arg0 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	return _Api.Contract.Monitorings(&_Api.CallOpts, arg0)
}

// Monitorings is a free data retrieval call binding the contract method 0xee8e1d51.
//
// Solidity: function monitorings(uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiCallerSession) Monitorings(arg0 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	return _Api.Contract.Monitorings(&_Api.CallOpts, arg0)
}

// Spawnings is a free data retrieval call binding the contract method 0x7092f1fb.
//
// Solidity: function spawnings(uint256 ) view returns(string timestamp, uint256 weight)
func (_Api *ApiCaller) Spawnings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp string
	Weight    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawnings", arg0)

	outstruct := new(struct {
		Timestamp string
		Weight    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Spawnings is a free data retrieval call binding the contract method 0x7092f1fb.
//
// Solidity: function spawnings(uint256 ) view returns(string timestamp, uint256 weight)
func (_Api *ApiSession) Spawnings(arg0 *big.Int) (struct {
	Timestamp string
	Weight    *big.Int
}, error) {
	return _Api.Contract.Spawnings(&_Api.CallOpts, arg0)
}

// Spawnings is a free data retrieval call binding the contract method 0x7092f1fb.
//
// Solidity: function spawnings(uint256 ) view returns(string timestamp, uint256 weight)
func (_Api *ApiCallerSession) Spawnings(arg0 *big.Int) (struct {
	Timestamp string
	Weight    *big.Int
}, error) {
	return _Api.Contract.Spawnings(&_Api.CallOpts, arg0)
}

// AddMonitoring is a paid mutator transaction binding the contract method 0xf400d7ca.
//
// Solidity: function addMonitoring(string _timestamp, uint256 _weight, uint256 _temperature, uint256 _humidity) returns()
func (_Api *ApiTransactor) AddMonitoring(opts *bind.TransactOpts, _timestamp string, _weight *big.Int, _temperature *big.Int, _humidity *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addMonitoring", _timestamp, _weight, _temperature, _humidity)
}

// AddMonitoring is a paid mutator transaction binding the contract method 0xf400d7ca.
//
// Solidity: function addMonitoring(string _timestamp, uint256 _weight, uint256 _temperature, uint256 _humidity) returns()
func (_Api *ApiSession) AddMonitoring(_timestamp string, _weight *big.Int, _temperature *big.Int, _humidity *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddMonitoring(&_Api.TransactOpts, _timestamp, _weight, _temperature, _humidity)
}

// AddMonitoring is a paid mutator transaction binding the contract method 0xf400d7ca.
//
// Solidity: function addMonitoring(string _timestamp, uint256 _weight, uint256 _temperature, uint256 _humidity) returns()
func (_Api *ApiTransactorSession) AddMonitoring(_timestamp string, _weight *big.Int, _temperature *big.Int, _humidity *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddMonitoring(&_Api.TransactOpts, _timestamp, _weight, _temperature, _humidity)
}

// AddSpawning is a paid mutator transaction binding the contract method 0xde97090c.
//
// Solidity: function addSpawning(string _timestamp, uint256 _weight) returns()
func (_Api *ApiTransactor) AddSpawning(opts *bind.TransactOpts, _timestamp string, _weight *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addSpawning", _timestamp, _weight)
}

// AddSpawning is a paid mutator transaction binding the contract method 0xde97090c.
//
// Solidity: function addSpawning(string _timestamp, uint256 _weight) returns()
func (_Api *ApiSession) AddSpawning(_timestamp string, _weight *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddSpawning(&_Api.TransactOpts, _timestamp, _weight)
}

// AddSpawning is a paid mutator transaction binding the contract method 0xde97090c.
//
// Solidity: function addSpawning(string _timestamp, uint256 _weight) returns()
func (_Api *ApiTransactorSession) AddSpawning(_timestamp string, _weight *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddSpawning(&_Api.TransactOpts, _timestamp, _weight)
}
