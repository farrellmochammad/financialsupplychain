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

// CreditScoreContractCredit is an auto generated low-level Go binding around an user-defined struct.
type CreditScoreContractCredit struct {
	CreditAmount *big.Int
	Status       bool
}

// CreditScoreContractPond is an auto generated low-level Go binding around an user-defined struct.
type CreditScoreContractPond struct {
	Id       string
	Fishtype string
}

// CreditScoreContractSpawning is an auto generated low-level Go binding around an user-defined struct.
type CreditScoreContractSpawning struct {
	TotalSpawning *big.Int
	SpawningDate  string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spawningaverage\",\"type\":\"uint256\"}],\"name\":\"addSpawningAverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditTransactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit[]\",\"name\":\"_credits\",\"type\":\"tuple[]\"}],\"name\":\"creditValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getPond\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Pond[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpawningCurrentAverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pondlength\",\"type\":\"uint256\"}],\"name\":\"pondValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondsInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit\",\"name\":\"_credit\",\"type\":\"tuple\"}],\"name\":\"setCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Pond\",\"name\":\"_pond\",\"type\":\"tuple\"}],\"name\":\"setPond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningAverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"_spawnings\",\"type\":\"tuple[]\"}],\"name\":\"spawningValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningsInformation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061005a6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a00000000008152503361007160201b610ac21760201c565b600080546001600160a01b0319163317905561013f565b6100ba82826040516024016100879291906100df565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b179091526100be16565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b8181101561010d57602081870181015160608684010152016100f0565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b6113e5806200014f6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806398a81c2111610097578063e627416011610066578063e62741601461027b578063e7e577e41461029c578063f7522cba146102bc578063fc421ae0146102cf57600080fd5b806398a81c2114610214578063a620da6814610227578063bfbda90d14610248578063cbb37a1d1461025b57600080fd5b806361a23d9b116100d357806361a23d9b1461017e578063693a4d37146101935780638da5cb5b146101d657806391d23d441461020157600080fd5b8063082c95b61461010557806326e879551461012b578063314bd964146101335780633ffc2c9314610156575b600080fd5b610118610113366004610b2c565b6102e2565b6040519081526020015b60405180910390f35b610118610303565b610146610141366004610c96565b610337565b6040519015158152602001610122565b610169610164366004610d4b565b610411565b60408051928352901515602083015201610122565b61019161018c366004610d8f565b61045f565b005b6101916101a1366004610b2c565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155565b6000546101e9906001600160a01b031681565b6040516001600160a01b039091168152602001610122565b61014661020f366004610e84565b6104c2565b610146610222366004610b2c565b610510565b61023a610235366004610d4b565b61052b565b604051610122929190610f6f565b610191610256366004610f9d565b61068b565b61026e610269366004611000565b6106dc565b604051610122919061103c565b61028e610289366004610d4b565b6107f5565b6040516101229291906110b0565b6102af6102aa366004611000565b61084a565b60405161012291906110c9565b6101186102ca366004611000565b6109e0565b6101916102dd366004611140565b610a6c565b600481815481106102f257600080fd5b600091825260209091200154905081565b6004805460009190610317906001906111a4565b81548110610327576103276111bd565b9060005260206000200154905090565b600080805b835181101561038257838181518110610357576103576111bd565b6020026020010151600001518261036e91906111d3565b91508061037a816111e6565b91505061033c565b50825161038f90826111ff565b60048054919250906103a3906001906111a4565b815481106103b3576103b36111bd565b906000526020600020015481106103cd5750600192915050565b600480546103dd906001906111a4565b815481106103ed576103ed6111bd565b90600052602060002001548110156104085750600092915050565b50600192915050565b8151602081840181018051600182529282019185019190912091905280548290811061043c57600080fd5b60009182526020909120600290910201805460019091015490925060ff16905082565b60028260405161046f9190611221565b9081526040516020918190038201902080546001810182556000918252919020825183926002029091019081906104a690826112c6565b50602082015160018201906104bb90826112c6565b5050505050565b60006001815b8351811015610509578381815181106104e3576104e36111bd565b6020026020010151602001516104f7579015905b80610501816111e6565b9150506104c8565b5092915050565b6000600282101561052357506000919050565b506001919050565b8151602081840181018051600282529282019185019190912091905280548290811061055657600080fd5b90600052602060002090600202016000915091505080600001805461057a9061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546105a69061123d565b80156105f35780601f106105c8576101008083540402835291602001916105f3565b820191906000526020600020905b8154815290600101906020018083116105d657829003601f168201915b5050505050908060010180546106089061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546106349061123d565b80156106815780601f1061065657610100808354040283529160200191610681565b820191906000526020600020905b81548152906001019060200180831161066457829003601f168201915b5050505050905082565b60038260405161069b9190611221565b908152604051602091819003820190208054600180820183556000928352918390208451600290920201908155918301518392918201906104bb90826112c6565b60606003826040516106ee9190611221565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ea5783829060005260206000209060020201604051806040016040529081600082015481526020016001820180546107599061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546107859061123d565b80156107d25780601f106107a7576101008083540402835291602001916107d2565b820191906000526020600020905b8154815290600101906020018083116107b557829003601f168201915b5050505050815250508152602001906001019061071c565b505050509050919050565b8151602081840181018051600382529282019185019190912091905280548290811061082057600080fd5b9060005260206000209060020201600091509150508060000154908060010180546106089061123d565b606060028260405161085c9190611221565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ea57838290600052602060002090600202016040518060400160405290816000820180546108bd9061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546108e99061123d565b80156109365780601f1061090b57610100808354040283529160200191610936565b820191906000526020600020905b81548152906001019060200180831161091957829003601f168201915b5050505050815260200160018201805461094f9061123d565b80601f016020809104026020016040519081016040528092919081815260200182805461097b9061123d565b80156109c85780601f1061099d576101008083540402835291602001916109c8565b820191906000526020600020905b8154815290600101906020018083116109ab57829003601f168201915b5050505050815250508152602001906001019061088a565b600080805b6001846040516109f59190611221565b9081526040519081900360200190205481101561050957600184604051610a1c9190611221565b90815260200160405180910390208181548110610a3b57610a3b6111bd565b90600052602060002090600202016000015482610a5891906111d3565b915080610a64816111e6565b9150506109e5565b600182604051610a7c9190611221565b90815260405160209181900382019020805460018082018355600092835291839020845160029092020190815592909101519101805460ff191691151591909117905550565b610b078282604051602401610ad8929190611385565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052610b0b565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b600060208284031215610b3e57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610b7d57610b7d610b45565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610bab57610bab610b45565b604052919050565b60006001600160401b03821115610bcc57610bcc610b45565b5060051b60200190565b600082601f830112610be757600080fd5b81356001600160401b03811115610c0057610c00610b45565b610c13601f8201601f1916602001610b83565b818152846020838601011115610c2857600080fd5b816020850160208301376000918101602001919091529392505050565b600060408284031215610c5757600080fd5b610c5f610b5b565b90508135815260208201356001600160401b03811115610c7e57600080fd5b610c8a84828501610bd6565b60208301525092915050565b60006020808385031215610ca957600080fd5b82356001600160401b0380821115610cc057600080fd5b818501915085601f830112610cd457600080fd5b8135610ce7610ce282610bb3565b610b83565b81815260059190911b83018401908481019088831115610d0657600080fd5b8585015b83811015610d3e57803585811115610d225760008081fd5b610d308b89838a0101610c45565b845250918601918601610d0a565b5098975050505050505050565b60008060408385031215610d5e57600080fd5b82356001600160401b03811115610d7457600080fd5b610d8085828601610bd6565b95602094909401359450505050565b60008060408385031215610da257600080fd5b82356001600160401b0380821115610db957600080fd5b610dc586838701610bd6565b93506020850135915080821115610ddb57600080fd5b9084019060408287031215610def57600080fd5b610df7610b5b565b823582811115610e0657600080fd5b610e1288828601610bd6565b825250602083013582811115610e2757600080fd5b610e3388828601610bd6565b6020830152508093505050509250929050565b600060408284031215610e5857600080fd5b610e60610b5b565b90508135815260208201358015158114610e7957600080fd5b602082015292915050565b60006020808385031215610e9757600080fd5b82356001600160401b03811115610ead57600080fd5b8301601f81018513610ebe57600080fd5b8035610ecc610ce282610bb3565b81815260069190911b82018301908381019087831115610eeb57600080fd5b928401925b82841015610f1457610f028885610e46565b82528482019150604084019350610ef0565b979650505050505050565b60005b83811015610f3a578181015183820152602001610f22565b50506000910152565b60008151808452610f5b816020860160208601610f1f565b601f01601f19169290920160200192915050565b604081526000610f826040830185610f43565b8281036020840152610f948185610f43565b95945050505050565b60008060408385031215610fb057600080fd5b82356001600160401b0380821115610fc757600080fd5b610fd386838701610bd6565b93506020850135915080821115610fe957600080fd5b50610ff685828601610c45565b9150509250929050565b60006020828403121561101257600080fd5b81356001600160401b0381111561102857600080fd5b61103484828501610bd6565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156110a257888303603f1901855281518051845287015187840187905261108f87850182610f43565b9588019593505090860190600101611063565b509098975050505050505050565b8281526040602082015260006110346040830184610f43565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156110a257888303603f190185528151805187855261111488860182610f43565b91890151858303868b015291905061112c8183610f43565b9689019694505050908601906001016110f0565b6000806060838503121561115357600080fd5b82356001600160401b0381111561116957600080fd5b61117585828601610bd6565b9250506111858460208501610e46565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b818103818111156111b7576111b761118e565b92915050565b634e487b7160e01b600052603260045260246000fd5b808201808211156111b7576111b761118e565b6000600182016111f8576111f861118e565b5060010190565b60008261121c57634e487b7160e01b600052601260045260246000fd5b500490565b60008251611233818460208701610f1f565b9190910192915050565b600181811c9082168061125157607f821691505b60208210810361127157634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156112c157600081815260208120601f850160051c8101602086101561129e5750805b601f850160051c820191505b818110156112bd578281556001016112aa565b5050505b505050565b81516001600160401b038111156112df576112df610b45565b6112f3816112ed845461123d565b84611277565b602080601f83116001811461132857600084156113105750858301515b600019600386901b1c1916600185901b1785556112bd565b600085815260208120601f198616915b8281101561135757888601518255948401946001909101908401611338565b50858210156113755787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6040815260006113986040830185610f43565b905060018060a01b0383166020830152939250505056fea26469706673582212201139696d2d75a3baa13d4ac331ad4549006c7a1627e4ab812ee20f7acbb8036864736f6c63430008110033",
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

// CreditValidation is a free data retrieval call binding the contract method 0x91d23d44.
//
// Solidity: function creditValidation((uint256,bool)[] _credits) pure returns(bool)
func (_Api *ApiCaller) CreditValidation(opts *bind.CallOpts, _credits []CreditScoreContractCredit) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "creditValidation", _credits)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditValidation is a free data retrieval call binding the contract method 0x91d23d44.
//
// Solidity: function creditValidation((uint256,bool)[] _credits) pure returns(bool)
func (_Api *ApiSession) CreditValidation(_credits []CreditScoreContractCredit) (bool, error) {
	return _Api.Contract.CreditValidation(&_Api.CallOpts, _credits)
}

// CreditValidation is a free data retrieval call binding the contract method 0x91d23d44.
//
// Solidity: function creditValidation((uint256,bool)[] _credits) pure returns(bool)
func (_Api *ApiCallerSession) CreditValidation(_credits []CreditScoreContractCredit) (bool, error) {
	return _Api.Contract.CreditValidation(&_Api.CallOpts, _credits)
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

// GetPond is a free data retrieval call binding the contract method 0xe7e577e4.
//
// Solidity: function getPond(string _nik) view returns((string,string)[])
func (_Api *ApiCaller) GetPond(opts *bind.CallOpts, _nik string) ([]CreditScoreContractPond, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getPond", _nik)

	if err != nil {
		return *new([]CreditScoreContractPond), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditScoreContractPond)).(*[]CreditScoreContractPond)

	return out0, err

}

// GetPond is a free data retrieval call binding the contract method 0xe7e577e4.
//
// Solidity: function getPond(string _nik) view returns((string,string)[])
func (_Api *ApiSession) GetPond(_nik string) ([]CreditScoreContractPond, error) {
	return _Api.Contract.GetPond(&_Api.CallOpts, _nik)
}

// GetPond is a free data retrieval call binding the contract method 0xe7e577e4.
//
// Solidity: function getPond(string _nik) view returns((string,string)[])
func (_Api *ApiCallerSession) GetPond(_nik string) ([]CreditScoreContractPond, error) {
	return _Api.Contract.GetPond(&_Api.CallOpts, _nik)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((uint256,string)[])
func (_Api *ApiCaller) GetSpawning(opts *bind.CallOpts, _nik string) ([]CreditScoreContractSpawning, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSpawning", _nik)

	if err != nil {
		return *new([]CreditScoreContractSpawning), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditScoreContractSpawning)).(*[]CreditScoreContractSpawning)

	return out0, err

}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((uint256,string)[])
func (_Api *ApiSession) GetSpawning(_nik string) ([]CreditScoreContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _nik)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((uint256,string)[])
func (_Api *ApiCallerSession) GetSpawning(_nik string) ([]CreditScoreContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _nik)
}

// GetSpawningCurrentAverage is a free data retrieval call binding the contract method 0x26e87955.
//
// Solidity: function getSpawningCurrentAverage() view returns(uint256)
func (_Api *ApiCaller) GetSpawningCurrentAverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSpawningCurrentAverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpawningCurrentAverage is a free data retrieval call binding the contract method 0x26e87955.
//
// Solidity: function getSpawningCurrentAverage() view returns(uint256)
func (_Api *ApiSession) GetSpawningCurrentAverage() (*big.Int, error) {
	return _Api.Contract.GetSpawningCurrentAverage(&_Api.CallOpts)
}

// GetSpawningCurrentAverage is a free data retrieval call binding the contract method 0x26e87955.
//
// Solidity: function getSpawningCurrentAverage() view returns(uint256)
func (_Api *ApiCallerSession) GetSpawningCurrentAverage() (*big.Int, error) {
	return _Api.Contract.GetSpawningCurrentAverage(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// PondValidation is a free data retrieval call binding the contract method 0x98a81c21.
//
// Solidity: function pondValidation(uint256 _pondlength) pure returns(bool)
func (_Api *ApiCaller) PondValidation(opts *bind.CallOpts, _pondlength *big.Int) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pondValidation", _pondlength)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PondValidation is a free data retrieval call binding the contract method 0x98a81c21.
//
// Solidity: function pondValidation(uint256 _pondlength) pure returns(bool)
func (_Api *ApiSession) PondValidation(_pondlength *big.Int) (bool, error) {
	return _Api.Contract.PondValidation(&_Api.CallOpts, _pondlength)
}

// PondValidation is a free data retrieval call binding the contract method 0x98a81c21.
//
// Solidity: function pondValidation(uint256 _pondlength) pure returns(bool)
func (_Api *ApiCallerSession) PondValidation(_pondlength *big.Int) (bool, error) {
	return _Api.Contract.PondValidation(&_Api.CallOpts, _pondlength)
}

// PondsInformation is a free data retrieval call binding the contract method 0xa620da68.
//
// Solidity: function pondsInformation(string , uint256 ) view returns(string id, string fishtype)
func (_Api *ApiCaller) PondsInformation(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Id       string
	Fishtype string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pondsInformation", arg0, arg1)

	outstruct := new(struct {
		Id       string
		Fishtype string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Fishtype = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// PondsInformation is a free data retrieval call binding the contract method 0xa620da68.
//
// Solidity: function pondsInformation(string , uint256 ) view returns(string id, string fishtype)
func (_Api *ApiSession) PondsInformation(arg0 string, arg1 *big.Int) (struct {
	Id       string
	Fishtype string
}, error) {
	return _Api.Contract.PondsInformation(&_Api.CallOpts, arg0, arg1)
}

// PondsInformation is a free data retrieval call binding the contract method 0xa620da68.
//
// Solidity: function pondsInformation(string , uint256 ) view returns(string id, string fishtype)
func (_Api *ApiCallerSession) PondsInformation(arg0 string, arg1 *big.Int) (struct {
	Id       string
	Fishtype string
}, error) {
	return _Api.Contract.PondsInformation(&_Api.CallOpts, arg0, arg1)
}

// SpawningAverage is a free data retrieval call binding the contract method 0x082c95b6.
//
// Solidity: function spawningAverage(uint256 ) view returns(uint256)
func (_Api *ApiCaller) SpawningAverage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawningAverage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpawningAverage is a free data retrieval call binding the contract method 0x082c95b6.
//
// Solidity: function spawningAverage(uint256 ) view returns(uint256)
func (_Api *ApiSession) SpawningAverage(arg0 *big.Int) (*big.Int, error) {
	return _Api.Contract.SpawningAverage(&_Api.CallOpts, arg0)
}

// SpawningAverage is a free data retrieval call binding the contract method 0x082c95b6.
//
// Solidity: function spawningAverage(uint256 ) view returns(uint256)
func (_Api *ApiCallerSession) SpawningAverage(arg0 *big.Int) (*big.Int, error) {
	return _Api.Contract.SpawningAverage(&_Api.CallOpts, arg0)
}

// SpawningValidation is a free data retrieval call binding the contract method 0x314bd964.
//
// Solidity: function spawningValidation((uint256,string)[] _spawnings) view returns(bool)
func (_Api *ApiCaller) SpawningValidation(opts *bind.CallOpts, _spawnings []CreditScoreContractSpawning) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawningValidation", _spawnings)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpawningValidation is a free data retrieval call binding the contract method 0x314bd964.
//
// Solidity: function spawningValidation((uint256,string)[] _spawnings) view returns(bool)
func (_Api *ApiSession) SpawningValidation(_spawnings []CreditScoreContractSpawning) (bool, error) {
	return _Api.Contract.SpawningValidation(&_Api.CallOpts, _spawnings)
}

// SpawningValidation is a free data retrieval call binding the contract method 0x314bd964.
//
// Solidity: function spawningValidation((uint256,string)[] _spawnings) view returns(bool)
func (_Api *ApiCallerSession) SpawningValidation(_spawnings []CreditScoreContractSpawning) (bool, error) {
	return _Api.Contract.SpawningValidation(&_Api.CallOpts, _spawnings)
}

// SpawningsInformation is a free data retrieval call binding the contract method 0xe6274160.
//
// Solidity: function spawningsInformation(string , uint256 ) view returns(uint256 totalSpawning, string spawningDate)
func (_Api *ApiCaller) SpawningsInformation(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	TotalSpawning *big.Int
	SpawningDate  string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawningsInformation", arg0, arg1)

	outstruct := new(struct {
		TotalSpawning *big.Int
		SpawningDate  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSpawning = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SpawningDate = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// SpawningsInformation is a free data retrieval call binding the contract method 0xe6274160.
//
// Solidity: function spawningsInformation(string , uint256 ) view returns(uint256 totalSpawning, string spawningDate)
func (_Api *ApiSession) SpawningsInformation(arg0 string, arg1 *big.Int) (struct {
	TotalSpawning *big.Int
	SpawningDate  string
}, error) {
	return _Api.Contract.SpawningsInformation(&_Api.CallOpts, arg0, arg1)
}

// SpawningsInformation is a free data retrieval call binding the contract method 0xe6274160.
//
// Solidity: function spawningsInformation(string , uint256 ) view returns(uint256 totalSpawning, string spawningDate)
func (_Api *ApiCallerSession) SpawningsInformation(arg0 string, arg1 *big.Int) (struct {
	TotalSpawning *big.Int
	SpawningDate  string
}, error) {
	return _Api.Contract.SpawningsInformation(&_Api.CallOpts, arg0, arg1)
}

// AddSpawningAverage is a paid mutator transaction binding the contract method 0x693a4d37.
//
// Solidity: function addSpawningAverage(uint256 _spawningaverage) returns()
func (_Api *ApiTransactor) AddSpawningAverage(opts *bind.TransactOpts, _spawningaverage *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addSpawningAverage", _spawningaverage)
}

// AddSpawningAverage is a paid mutator transaction binding the contract method 0x693a4d37.
//
// Solidity: function addSpawningAverage(uint256 _spawningaverage) returns()
func (_Api *ApiSession) AddSpawningAverage(_spawningaverage *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddSpawningAverage(&_Api.TransactOpts, _spawningaverage)
}

// AddSpawningAverage is a paid mutator transaction binding the contract method 0x693a4d37.
//
// Solidity: function addSpawningAverage(uint256 _spawningaverage) returns()
func (_Api *ApiTransactorSession) AddSpawningAverage(_spawningaverage *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddSpawningAverage(&_Api.TransactOpts, _spawningaverage)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiTransactor) SetCredit(opts *bind.TransactOpts, _nik string, _credit CreditScoreContractCredit) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setCredit", _nik, _credit)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiSession) SetCredit(_nik string, _credit CreditScoreContractCredit) (*types.Transaction, error) {
	return _Api.Contract.SetCredit(&_Api.TransactOpts, _nik, _credit)
}

// SetCredit is a paid mutator transaction binding the contract method 0xfc421ae0.
//
// Solidity: function setCredit(string _nik, (uint256,bool) _credit) returns()
func (_Api *ApiTransactorSession) SetCredit(_nik string, _credit CreditScoreContractCredit) (*types.Transaction, error) {
	return _Api.Contract.SetCredit(&_Api.TransactOpts, _nik, _credit)
}

// SetPond is a paid mutator transaction binding the contract method 0x61a23d9b.
//
// Solidity: function setPond(string _nik, (string,string) _pond) returns()
func (_Api *ApiTransactor) SetPond(opts *bind.TransactOpts, _nik string, _pond CreditScoreContractPond) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPond", _nik, _pond)
}

// SetPond is a paid mutator transaction binding the contract method 0x61a23d9b.
//
// Solidity: function setPond(string _nik, (string,string) _pond) returns()
func (_Api *ApiSession) SetPond(_nik string, _pond CreditScoreContractPond) (*types.Transaction, error) {
	return _Api.Contract.SetPond(&_Api.TransactOpts, _nik, _pond)
}

// SetPond is a paid mutator transaction binding the contract method 0x61a23d9b.
//
// Solidity: function setPond(string _nik, (string,string) _pond) returns()
func (_Api *ApiTransactorSession) SetPond(_nik string, _pond CreditScoreContractPond) (*types.Transaction, error) {
	return _Api.Contract.SetPond(&_Api.TransactOpts, _nik, _pond)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xbfbda90d.
//
// Solidity: function setSpawning(string _nik, (uint256,string) _spawning) returns()
func (_Api *ApiTransactor) SetSpawning(opts *bind.TransactOpts, _nik string, _spawning CreditScoreContractSpawning) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setSpawning", _nik, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xbfbda90d.
//
// Solidity: function setSpawning(string _nik, (uint256,string) _spawning) returns()
func (_Api *ApiSession) SetSpawning(_nik string, _spawning CreditScoreContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _nik, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xbfbda90d.
//
// Solidity: function setSpawning(string _nik, (uint256,string) _spawning) returns()
func (_Api *ApiTransactorSession) SetSpawning(_nik string, _spawning CreditScoreContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _nik, _spawning)
}
