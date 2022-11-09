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

// TransactionContractApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractApproverTransaction struct {
	Reviewer  string
	Status    uint8
	Filepath  string
	Timestamp string
}

// TransactionContractFunderApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractFunderApproverTransaction struct {
	Reviewer   string
	Status     uint8
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}

// TransactionContractMonitoring is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractMonitoring struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}

// TransactionContractSpawning is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractSpawning struct {
	Pondid    string
	Timestamp string
	Weight    *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getApproverValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getFunderApproverValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"enumTransactionContract.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b620014971760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b612001806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063a264efc611610097578063d2892d1011610066578063d2892d1014610244578063e137df8314610267578063fbaa96cf14610287578063fd100e17146102a757600080fd5b8063a264efc6146101de578063aaa5004e146101fe578063cbb37a1d14610211578063cdc693b11461023157600080fd5b806366ce64cb116100d357806366ce64cb1461015b57806379248d941461016e5780637c5acbe2146101915780638da5cb5b146101b357600080fd5b80631793993c146100fa5780632455f66d146101225780634b870c7714610146575b600080fd5b61010d610108366004611613565b6102ba565b60405190151581526020015b60405180910390f35b610135610130366004611650565b610362565b60405161011995949392919061171d565b61015961015436600461178a565b610565565b005b61010d610169366004611613565b610620565b61018161017c366004611650565b6106c1565b6040516101199493929190611877565b6101a461019f366004611650565b6107a5565b604051610119939291906118a6565b6000546101c6906001600160a01b031681565b6040516001600160a01b039091168152602001610119565b6101f16101ec366004611613565b61090b565b60405161011991906118dc565b61015961020c366004611967565b610a38565b61022461021f366004611613565b610ace565b6040516101199190611a29565b61015961023f366004611aac565b610c6e565b610257610252366004611650565b610d03565b6040516101199493929190611b5e565b61027a610275366004611613565b610f00565b6040516101199190611baf565b61029a610295366004611613565b611165565b6040516101199190611c61565b6101596102b5366004611d08565b6113c0565b600080805b6002846040516102cf9190611dff565b9081526040519081900360200190205481101561035b5760036002856040516102f89190611dff565b9081526020016040518091039020828154811061031757610317611e1b565b600091825260209091206001600590920201015460ff16600381111561033f5761033f6116e5565b0361034957600191505b8061035381611e31565b9150506102bf565b5092915050565b8151602081840181018051600282529282019185019190912091905280548290811061038d57600080fd5b9060005260206000209060050201600091509150508060000180546103b190611e58565b80601f01602080910402602001604051908101604052809291908181526020018280546103dd90611e58565b801561042a5780601f106103ff5761010080835404028352916020019161042a565b820191906000526020600020905b81548152906001019060200180831161040d57829003601f168201915b5050506001840154600285018054949560ff90921694919350915061044e90611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461047a90611e58565b80156104c75780601f1061049c576101008083540402835291602001916104c7565b820191906000526020600020905b8154815290600101906020018083116104aa57829003601f168201915b5050505050908060030180546104dc90611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461050890611e58565b80156105555780601f1061052a57610100808354040283529160200191610555565b820191906000526020600020905b81548152906001019060200180831161053857829003601f168201915b5050505050908060040154905085565b6000546001600160a01b0316331461057c57600080fd5b60018260405161058c9190611dff565b9081526040516020918190038201902080546001810182556000918252919020825183926004029091019081906105c39082611ee1565b50602082015160018083018054909160ff19909116908360038111156105eb576105eb6116e5565b0217905550604082015160028201906106049082611ee1565b50606082015160038201906106199082611ee1565b5050505050565b600080805b6001846040516106359190611dff565b9081526040519081900360200190205481101561035b57600360018560405161065e9190611dff565b9081526020016040518091039020828154811061067d5761067d611e1b565b600091825260209091206001600490920201015460ff1660038111156106a5576106a56116e5565b036106af57600191505b806106b981611e31565b915050610625565b815160208184018101805160038252928201918501919091209190528054829081106106ec57600080fd5b90600052602060002090600402016000915091505080600001805461071090611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461073c90611e58565b80156107895780601f1061075e57610100808354040283529160200191610789565b820191906000526020600020905b81548152906001019060200180831161076c57829003601f168201915b5050505050908060010154908060020154908060030154905084565b815160208184018101805160048252928201918501919091209190528054829081106107d057600080fd5b9060005260206000209060030201600091509150508060000180546107f490611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461082090611e58565b801561086d5780601f106108425761010080835404028352916020019161086d565b820191906000526020600020905b81548152906001019060200180831161085057829003601f168201915b50505050509080600101805461088290611e58565b80601f01602080910402602001604051908101604052809291908181526020018280546108ae90611e58565b80156108fb5780601f106108d0576101008083540402835291602001916108fb565b820191906000526020600020905b8154815290600101906020018083116108de57829003601f168201915b5050505050908060020154905083565b606060038260405161091d9190611dff565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a2d578382906000526020600020906004020160405180608001604052908160008201805461097e90611e58565b80601f01602080910402602001604051908101604052809291908181526020018280546109aa90611e58565b80156109f75780601f106109cc576101008083540402835291602001916109f7565b820191906000526020600020905b8154815290600101906020018083116109da57829003601f168201915b5050505050815260200160018201548152602001600282015481526020016003820154815250508152602001906001019061094b565b505050509050919050565b6000546001600160a01b03163314610a4f57600080fd5b610a58826102ba565b610a6157600080fd5b600482604051610a719190611dff565b908152604051602091819003820190208054600181018255600091825291902082518392600302909101908190610aa89082611ee1565b5060208201516001820190610abd9082611ee1565b506040820151816002015550505050565b6060600482604051610ae09190611dff565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a2d5783829060005260206000209060030201604051806060016040529081600082018054610b4190611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6d90611e58565b8015610bba5780601f10610b8f57610100808354040283529160200191610bba565b820191906000526020600020905b815481529060010190602001808311610b9d57829003601f168201915b50505050508152602001600182018054610bd390611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610bff90611e58565b8015610c4c5780601f10610c2157610100808354040283529160200191610c4c565b820191906000526020600020905b815481529060010190602001808311610c2f57829003601f168201915b5050505050815260200160028201548152505081526020019060010190610b0e565b6000546001600160a01b03163314610c8557600080fd5b610c8e826102ba565b610c9757600080fd5b600382604051610ca79190611dff565b908152604051602091819003820190208054600181018255600091825291902082518392600402909101908190610cde9082611ee1565b5060208201518160010155604082015181600201556060820151816003015550505050565b81516020818401810180516001825292820191850191909120919052805482908110610d2e57600080fd5b906000526020600020906004020160009150915050806000018054610d5290611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610d7e90611e58565b8015610dcb5780601f10610da057610100808354040283529160200191610dcb565b820191906000526020600020905b815481529060010190602001808311610dae57829003601f168201915b5050506001840154600285018054949560ff909216949193509150610def90611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1b90611e58565b8015610e685780601f10610e3d57610100808354040283529160200191610e68565b820191906000526020600020905b815481529060010190602001808311610e4b57829003601f168201915b505050505090806003018054610e7d90611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610ea990611e58565b8015610ef65780601f10610ecb57610100808354040283529160200191610ef6565b820191906000526020600020905b815481529060010190602001808311610ed957829003601f168201915b5050505050905084565b6060600282604051610f129190611dff565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a2d57838290600052602060002090600502016040518060a0016040529081600082018054610f7390611e58565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9f90611e58565b8015610fec5780601f10610fc157610100808354040283529160200191610fec565b820191906000526020600020905b815481529060010190602001808311610fcf57829003601f168201915b5050509183525050600182015460209091019060ff166003811115611013576110136116e5565b6003811115611024576110246116e5565b815260200160028201805461103890611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461106490611e58565b80156110b15780601f10611086576101008083540402835291602001916110b1565b820191906000526020600020905b81548152906001019060200180831161109457829003601f168201915b505050505081526020016003820180546110ca90611e58565b80601f01602080910402602001604051908101604052809291908181526020018280546110f690611e58565b80156111435780601f1061111857610100808354040283529160200191611143565b820191906000526020600020905b81548152906001019060200180831161112657829003601f168201915b5050505050815260200160048201548152505081526020019060010190610f40565b60606001826040516111779190611dff565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a2d57838290600052602060002090600402016040518060800160405290816000820180546111d890611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461120490611e58565b80156112515780601f1061122657610100808354040283529160200191611251565b820191906000526020600020905b81548152906001019060200180831161123457829003601f168201915b5050509183525050600182015460209091019060ff166003811115611278576112786116e5565b6003811115611289576112896116e5565b815260200160028201805461129d90611e58565b80601f01602080910402602001604051908101604052809291908181526020018280546112c990611e58565b80156113165780601f106112eb57610100808354040283529160200191611316565b820191906000526020600020905b8154815290600101906020018083116112f957829003601f168201915b5050505050815260200160038201805461132f90611e58565b80601f016020809104026020016040519081016040528092919081815260200182805461135b90611e58565b80156113a85780601f1061137d576101008083540402835291602001916113a8565b820191906000526020600020905b81548152906001019060200180831161138b57829003601f168201915b505050505081525050815260200190600101906111a5565b6000546001600160a01b031633146113d757600080fd5b6113e082610620565b6113e957600080fd5b6002826040516113f99190611dff565b9081526040516020918190038201902080546001810182556000918252919020825183926005029091019081906114309082611ee1565b50602082015160018083018054909160ff1990911690836003811115611458576114586116e5565b0217905550604082015160028201906114719082611ee1565b50606082015160038201906114869082611ee1565b506080820151816004015550505050565b6114dc82826040516024016114ad929190611fa1565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b1790526114e0565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561153a5761153a611501565b60405290565b6040516060810167ffffffffffffffff8111828210171561153a5761153a611501565b60405160a0810167ffffffffffffffff8111828210171561153a5761153a611501565b600082601f83011261159757600080fd5b813567ffffffffffffffff808211156115b2576115b2611501565b604051601f8301601f19908116603f011681019082821181831017156115da576115da611501565b816040528381528660208588010111156115f357600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561162557600080fd5b813567ffffffffffffffff81111561163c57600080fd5b61164884828501611586565b949350505050565b6000806040838503121561166357600080fd5b823567ffffffffffffffff81111561167a57600080fd5b61168685828601611586565b95602094909401359450505050565b60005b838110156116b0578181015183820152602001611698565b50506000910152565b600081518084526116d1816020860160208601611695565b601f01601f19169290920160200192915050565b634e487b7160e01b600052602160045260246000fd5b6004811061171957634e487b7160e01b600052602160045260246000fd5b9052565b60a08152600061173060a08301886116b9565b61173d60208401886116fb565b828103604084015261174f81876116b9565b9050828103606084015261176381866116b9565b9150508260808301529695505050505050565b80356004811061178557600080fd5b919050565b6000806040838503121561179d57600080fd5b823567ffffffffffffffff808211156117b557600080fd5b6117c186838701611586565b935060208501359150808211156117d757600080fd5b90840190608082870312156117eb57600080fd5b6117f3611517565b82358281111561180257600080fd5b61180e88828601611586565b82525061181d60208401611776565b602082015260408301358281111561183457600080fd5b61184088828601611586565b60408301525060608301358281111561185857600080fd5b61186488828601611586565b6060830152508093505050509250929050565b60808152600061188a60808301876116b9565b6020830195909552506040810192909252606090910152919050565b6060815260006118b960608301866116b9565b82810360208401526118cb81866116b9565b915050826040830152949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561195957603f19898403018552815160808151818652611929828701826116b9565b838b0151878c0152898401518a880152606093840151939096019290925250509386019390860190600101611903565b509098975050505050505050565b6000806040838503121561197a57600080fd5b823567ffffffffffffffff8082111561199257600080fd5b61199e86838701611586565b935060208501359150808211156119b457600080fd5b90840190606082870312156119c857600080fd5b6119d0611540565b8235828111156119df57600080fd5b6119eb88828601611586565b825250602083013582811115611a0057600080fd5b611a0c88828601611586565b602083015250604083013560408201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561195957603f19898403018552815160608151818652611a76828701826116b9565b915050888201518582038a870152611a8e82826116b9565b92890151958901959095525094870194925090860190600101611a50565b60008060408385031215611abf57600080fd5b823567ffffffffffffffff80821115611ad757600080fd5b611ae386838701611586565b93506020850135915080821115611af957600080fd5b9084019060808287031215611b0d57600080fd5b611b15611517565b823582811115611b2457600080fd5b611b3088828601611586565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b608081526000611b7160808301876116b9565b611b7e60208401876116fb565b8281036040840152611b9081866116b9565b90508281036060840152611ba481856116b9565b979650505050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561195957603f19898403018552815160a08151818652611bfc828701826116b9565b91505088820151611c0f8a8701826116fb565b508782015185820389870152611c2582826116b9565b91505060608083015186830382880152611c3f83826116b9565b6080948501519790940196909652505094870194925090860190600101611bd6565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561195957603f19898403018552815160808151818652611cae828701826116b9565b91505088820151611cc18a8701826116fb565b508782015185820389870152611cd782826116b9565b91505060608083015192508582038187015250611cf481836116b9565b968901969450505090860190600101611c88565b60008060408385031215611d1b57600080fd5b823567ffffffffffffffff80821115611d3357600080fd5b611d3f86838701611586565b93506020850135915080821115611d5557600080fd5b9084019060a08287031215611d6957600080fd5b611d71611563565b823582811115611d8057600080fd5b611d8c88828601611586565b825250611d9b60208401611776565b6020820152604083013582811115611db257600080fd5b611dbe88828601611586565b604083015250606083013582811115611dd657600080fd5b611de288828601611586565b606083015250608083013560808201528093505050509250929050565b60008251611e11818460208701611695565b9190910192915050565b634e487b7160e01b600052603260045260246000fd5b600060018201611e5157634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c90821680611e6c57607f821691505b602082108103611e8c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115611edc57600081815260208120601f850160051c81016020861015611eb95750805b601f850160051c820191505b81811015611ed857828155600101611ec5565b5050505b505050565b815167ffffffffffffffff811115611efb57611efb611501565b611f0f81611f098454611e58565b84611e92565b602080601f831160018114611f445760008415611f2c5750858301515b600019600386901b1c1916600185901b178555611ed8565b600085815260208120601f198616915b82811015611f7357888601518255948401946001909101908401611f54565b5085821015611f915787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611fb460408301856116b9565b905060018060a01b0383166020830152939250505056fea2646970667358221220861905b326f8e3bac0a9b6c2c21cd46930bc1aeac982cade333c2d019ed82ae164736f6c63430008110033",
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

// ApproverTransactions is a free data retrieval call binding the contract method 0xd2892d10.
//
// Solidity: function approverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp)
func (_Api *ApiCaller) ApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Reviewer  string
	Status    uint8
	Filepath  string
	Timestamp string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "approverTransactions", arg0, arg1)

	outstruct := new(struct {
		Reviewer  string
		Status    uint8
		Filepath  string
		Timestamp string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reviewer = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Filepath = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// ApproverTransactions is a free data retrieval call binding the contract method 0xd2892d10.
//
// Solidity: function approverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp)
func (_Api *ApiSession) ApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer  string
	Status    uint8
	Filepath  string
	Timestamp string
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// ApproverTransactions is a free data retrieval call binding the contract method 0xd2892d10.
//
// Solidity: function approverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp)
func (_Api *ApiCallerSession) ApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer  string
	Status    uint8
	Filepath  string
	Timestamp string
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiCaller) FunderApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     uint8
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderApproverTransactions", arg0, arg1)

	outstruct := new(struct {
		Reviewer   string
		Status     uint8
		Filepath   string
		Timestamp  string
		Modalgiven *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reviewer = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Filepath = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Modalgiven = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     uint8
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, uint8 status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiCallerSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     uint8
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _nik) view returns((string,uint8,string,string)[])
func (_Api *ApiCaller) GetApproverTransaction(opts *bind.CallOpts, _nik string) ([]TransactionContractApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApproverTransaction", _nik)

	if err != nil {
		return *new([]TransactionContractApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractApproverTransaction)).(*[]TransactionContractApproverTransaction)

	return out0, err

}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _nik) view returns((string,uint8,string,string)[])
func (_Api *ApiSession) GetApproverTransaction(_nik string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _nik)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _nik) view returns((string,uint8,string,string)[])
func (_Api *ApiCallerSession) GetApproverTransaction(_nik string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _nik)
}

// GetApproverValidation is a free data retrieval call binding the contract method 0x66ce64cb.
//
// Solidity: function getApproverValidation(string _nik) view returns(bool)
func (_Api *ApiCaller) GetApproverValidation(opts *bind.CallOpts, _nik string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApproverValidation", _nik)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetApproverValidation is a free data retrieval call binding the contract method 0x66ce64cb.
//
// Solidity: function getApproverValidation(string _nik) view returns(bool)
func (_Api *ApiSession) GetApproverValidation(_nik string) (bool, error) {
	return _Api.Contract.GetApproverValidation(&_Api.CallOpts, _nik)
}

// GetApproverValidation is a free data retrieval call binding the contract method 0x66ce64cb.
//
// Solidity: function getApproverValidation(string _nik) view returns(bool)
func (_Api *ApiCallerSession) GetApproverValidation(_nik string) (bool, error) {
	return _Api.Contract.GetApproverValidation(&_Api.CallOpts, _nik)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,uint8,string,string,uint256)[])
func (_Api *ApiCaller) GetFunderApproverTransaction(opts *bind.CallOpts, _nik string) ([]TransactionContractFunderApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getFunderApproverTransaction", _nik)

	if err != nil {
		return *new([]TransactionContractFunderApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractFunderApproverTransaction)).(*[]TransactionContractFunderApproverTransaction)

	return out0, err

}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,uint8,string,string,uint256)[])
func (_Api *ApiSession) GetFunderApproverTransaction(_nik string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _nik)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,uint8,string,string,uint256)[])
func (_Api *ApiCallerSession) GetFunderApproverTransaction(_nik string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _nik)
}

// GetFunderApproverValidation is a free data retrieval call binding the contract method 0x1793993c.
//
// Solidity: function getFunderApproverValidation(string _nik) view returns(bool)
func (_Api *ApiCaller) GetFunderApproverValidation(opts *bind.CallOpts, _nik string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getFunderApproverValidation", _nik)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFunderApproverValidation is a free data retrieval call binding the contract method 0x1793993c.
//
// Solidity: function getFunderApproverValidation(string _nik) view returns(bool)
func (_Api *ApiSession) GetFunderApproverValidation(_nik string) (bool, error) {
	return _Api.Contract.GetFunderApproverValidation(&_Api.CallOpts, _nik)
}

// GetFunderApproverValidation is a free data retrieval call binding the contract method 0x1793993c.
//
// Solidity: function getFunderApproverValidation(string _nik) view returns(bool)
func (_Api *ApiCallerSession) GetFunderApproverValidation(_nik string) (bool, error) {
	return _Api.Contract.GetFunderApproverValidation(&_Api.CallOpts, _nik)
}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _pondid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCaller) GetMonitoring(opts *bind.CallOpts, _pondid string) ([]TransactionContractMonitoring, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getMonitoring", _pondid)

	if err != nil {
		return *new([]TransactionContractMonitoring), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractMonitoring)).(*[]TransactionContractMonitoring)

	return out0, err

}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _pondid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetMonitoring(_pondid string) ([]TransactionContractMonitoring, error) {
	return _Api.Contract.GetMonitoring(&_Api.CallOpts, _pondid)
}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _pondid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetMonitoring(_pondid string) ([]TransactionContractMonitoring, error) {
	return _Api.Contract.GetMonitoring(&_Api.CallOpts, _pondid)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((string,string,uint256)[])
func (_Api *ApiCaller) GetSpawning(opts *bind.CallOpts, _nik string) ([]TransactionContractSpawning, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSpawning", _nik)

	if err != nil {
		return *new([]TransactionContractSpawning), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractSpawning)).(*[]TransactionContractSpawning)

	return out0, err

}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((string,string,uint256)[])
func (_Api *ApiSession) GetSpawning(_nik string) ([]TransactionContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _nik)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _nik) view returns((string,string,uint256)[])
func (_Api *ApiCallerSession) GetSpawning(_nik string) ([]TransactionContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _nik)
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

// PondMonitorings is a free data retrieval call binding the contract method 0x79248d94.
//
// Solidity: function pondMonitorings(string , uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiCaller) PondMonitorings(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pondMonitorings", arg0, arg1)

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

// PondMonitorings is a free data retrieval call binding the contract method 0x79248d94.
//
// Solidity: function pondMonitorings(string , uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiSession) PondMonitorings(arg0 string, arg1 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	return _Api.Contract.PondMonitorings(&_Api.CallOpts, arg0, arg1)
}

// PondMonitorings is a free data retrieval call binding the contract method 0x79248d94.
//
// Solidity: function pondMonitorings(string , uint256 ) view returns(string timestamp, uint256 weight, uint256 temperature, uint256 humidity)
func (_Api *ApiCallerSession) PondMonitorings(arg0 string, arg1 *big.Int) (struct {
	Timestamp   string
	Weight      *big.Int
	Temperature *big.Int
	Humidity    *big.Int
}, error) {
	return _Api.Contract.PondMonitorings(&_Api.CallOpts, arg0, arg1)
}

// SpawningHistory is a free data retrieval call binding the contract method 0x7c5acbe2.
//
// Solidity: function spawningHistory(string , uint256 ) view returns(string pondid, string timestamp, uint256 weight)
func (_Api *ApiCaller) SpawningHistory(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Pondid    string
	Timestamp string
	Weight    *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawningHistory", arg0, arg1)

	outstruct := new(struct {
		Pondid    string
		Timestamp string
		Weight    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pondid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Weight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SpawningHistory is a free data retrieval call binding the contract method 0x7c5acbe2.
//
// Solidity: function spawningHistory(string , uint256 ) view returns(string pondid, string timestamp, uint256 weight)
func (_Api *ApiSession) SpawningHistory(arg0 string, arg1 *big.Int) (struct {
	Pondid    string
	Timestamp string
	Weight    *big.Int
}, error) {
	return _Api.Contract.SpawningHistory(&_Api.CallOpts, arg0, arg1)
}

// SpawningHistory is a free data retrieval call binding the contract method 0x7c5acbe2.
//
// Solidity: function spawningHistory(string , uint256 ) view returns(string pondid, string timestamp, uint256 weight)
func (_Api *ApiCallerSession) SpawningHistory(arg0 string, arg1 *big.Int) (struct {
	Pondid    string
	Timestamp string
	Weight    *big.Int
}, error) {
	return _Api.Contract.SpawningHistory(&_Api.CallOpts, arg0, arg1)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0x4b870c77.
//
// Solidity: function setApproverTransaction(string _nik, (string,uint8,string,string) _approvertransaction) returns()
func (_Api *ApiTransactor) SetApproverTransaction(opts *bind.TransactOpts, _nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApproverTransaction", _nik, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0x4b870c77.
//
// Solidity: function setApproverTransaction(string _nik, (string,uint8,string,string) _approvertransaction) returns()
func (_Api *ApiSession) SetApproverTransaction(_nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _nik, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0x4b870c77.
//
// Solidity: function setApproverTransaction(string _nik, (string,uint8,string,string) _approvertransaction) returns()
func (_Api *ApiTransactorSession) SetApproverTransaction(_nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _nik, _approvertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xfd100e17.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,uint8,string,string,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactor) SetFunderApproverTransaction(opts *bind.TransactOpts, _nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setFunderApproverTransaction", _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xfd100e17.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,uint8,string,string,uint256) _funderapprovertransaction) returns()
func (_Api *ApiSession) SetFunderApproverTransaction(_nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xfd100e17.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,uint8,string,string,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactorSession) SetFunderApproverTransaction(_nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _nik, _funderapprovertransaction)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _pondid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiTransactor) SetMonitoring(opts *bind.TransactOpts, _pondid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMonitoring", _pondid, _monitoring)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _pondid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiSession) SetMonitoring(_pondid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.Contract.SetMonitoring(&_Api.TransactOpts, _pondid, _monitoring)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _pondid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiTransactorSession) SetMonitoring(_pondid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.Contract.SetMonitoring(&_Api.TransactOpts, _pondid, _monitoring)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _nik, (string,string,uint256) _spawning) returns()
func (_Api *ApiTransactor) SetSpawning(opts *bind.TransactOpts, _nik string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setSpawning", _nik, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _nik, (string,string,uint256) _spawning) returns()
func (_Api *ApiSession) SetSpawning(_nik string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _nik, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _nik, (string,string,uint256) _spawning) returns()
func (_Api *ApiTransactorSession) SetSpawning(_nik string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _nik, _spawning)
}
