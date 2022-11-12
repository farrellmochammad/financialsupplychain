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
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}

// TransactionContractFunderApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractFunderApproverTransaction struct {
	Reviewer   string
	Status     string
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"reviewer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"filepath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"modalgiven\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b620013691760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611eac806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063c3292eda1161008c578063d2892d1011610066578063d2892d10146101ec578063e137df8314610211578063f734b1b314610231578063fbaa96cf1461024457600080fd5b8063c3292eda146101a6578063cbb37a1d146101b9578063cdc693b1146101d957600080fd5b80632455f66d146100d457806379248d94146101015780637c5acbe2146101245780638da5cb5b14610146578063a264efc614610171578063aaa5004e14610191575b600080fd5b6100e76100e2366004611508565b610264565b6040516100f895949392919061159d565b60405180910390f35b61011461010f366004611508565b6104e6565b6040516100f894939291906115fd565b610137610132366004611508565b6105ca565b6040516100f89392919061162c565b600054610159906001600160a01b031681565b6040516001600160a01b0390911681526020016100f8565b61018461017f366004611662565b610730565b6040516100f8919061169f565b6101a461019f36600461172a565b61085d565b005b6101a46101b43660046117ec565b6108e1565b6101cc6101c7366004611662565b61098f565b6040516100f891906118f6565b6101a46101e7366004611979565b610b2f565b6101ff6101fa366004611508565b610bb2565b6040516100f896959493929190611a2b565b61022461021f366004611662565b610db2565b6040516100f89190611a83565b6101a461023f366004611b3c565b611076565b610257610252366004611662565b611123565b6040516100f89190611c36565b8151602081840181018051600282529282019185019190912091905280548290811061028f57600080fd5b9060005260206000209060050201600091509150508060000180546102b390611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546102df90611ce7565b801561032c5780601f106103015761010080835404028352916020019161032c565b820191906000526020600020905b81548152906001019060200180831161030f57829003601f168201915b50505050509080600101805461034190611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461036d90611ce7565b80156103ba5780601f1061038f576101008083540402835291602001916103ba565b820191906000526020600020905b81548152906001019060200180831161039d57829003601f168201915b5050505050908060020180546103cf90611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546103fb90611ce7565b80156104485780601f1061041d57610100808354040283529160200191610448565b820191906000526020600020905b81548152906001019060200180831161042b57829003601f168201915b50505050509080600301805461045d90611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461048990611ce7565b80156104d65780601f106104ab576101008083540402835291602001916104d6565b820191906000526020600020905b8154815290600101906020018083116104b957829003601f168201915b5050505050908060040154905085565b8151602081840181018051600382529282019185019190912091905280548290811061051157600080fd5b90600052602060002090600402016000915091505080600001805461053590611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461056190611ce7565b80156105ae5780601f10610583576101008083540402835291602001916105ae565b820191906000526020600020905b81548152906001019060200180831161059157829003601f168201915b5050505050908060010154908060020154908060030154905084565b815160208184018101805160048252928201918501919091209190528054829081106105f557600080fd5b90600052602060002090600302016000915091505080600001805461061990611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461064590611ce7565b80156106925780601f1061066757610100808354040283529160200191610692565b820191906000526020600020905b81548152906001019060200180831161067557829003601f168201915b5050505050908060010180546106a790611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546106d390611ce7565b80156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b5050505050908060020154905083565b60606003826040516107429190611d21565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561085257838290600052602060002090600402016040518060800160405290816000820180546107a390611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546107cf90611ce7565b801561081c5780601f106107f15761010080835404028352916020019161081c565b820191906000526020600020905b8154815290600101906020018083116107ff57829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152505081526020019060010190610770565b505050509050919050565b6000546001600160a01b0316331461087457600080fd5b6004826040516108849190611d21565b9081526040516020918190038201902080546001810182556000918252919020825183926003029091019081906108bb9082611d8c565b50602082015160018201906108d09082611d8c565b506040820151816002015550505050565b6000546001600160a01b031633146108f857600080fd5b6002826040516109089190611d21565b90815260405160209181900382019020805460018101825560009182529190208251839260050290910190819061093f9082611d8c565b50602082015160018201906109549082611d8c565b50604082015160028201906109699082611d8c565b506060820151600382019061097e9082611d8c565b506080820151816004015550505050565b60606004826040516109a19190611d21565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156108525783829060005260206000209060030201604051806060016040529081600082018054610a0290611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2e90611ce7565b8015610a7b5780601f10610a5057610100808354040283529160200191610a7b565b820191906000526020600020905b815481529060010190602001808311610a5e57829003601f168201915b50505050508152602001600182018054610a9490611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac090611ce7565b8015610b0d5780601f10610ae257610100808354040283529160200191610b0d565b820191906000526020600020905b815481529060010190602001808311610af057829003601f168201915b50505050508152602001600282015481525050815260200190600101906109cf565b6000546001600160a01b03163314610b4657600080fd5b600382604051610b569190611d21565b908152604051602091819003820190208054600181018255600091825291902082518392600402909101908190610b8d9082611d8c565b5060208201518160010155604082015181600201556060820151816003015550505050565b81516020818401810180516001825292820191850191909120919052805482908110610bdd57600080fd5b906000526020600020906006020160009150915050806000018054610c0190611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2d90611ce7565b8015610c7a5780601f10610c4f57610100808354040283529160200191610c7a565b820191906000526020600020905b815481529060010190602001808311610c5d57829003601f168201915b505050505090806001018054610c8f90611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbb90611ce7565b8015610d085780601f10610cdd57610100808354040283529160200191610d08565b820191906000526020600020905b815481529060010190602001808311610ceb57829003601f168201915b505050505090806002018054610d1d90611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610d4990611ce7565b8015610d965780601f10610d6b57610100808354040283529160200191610d96565b820191906000526020600020905b815481529060010190602001808311610d7957829003601f168201915b5050505050908060030154908060040154908060050154905086565b6060600282604051610dc49190611d21565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561085257838290600052602060002090600502016040518060a0016040529081600082018054610e2590611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5190611ce7565b8015610e9e5780601f10610e7357610100808354040283529160200191610e9e565b820191906000526020600020905b815481529060010190602001808311610e8157829003601f168201915b50505050508152602001600182018054610eb790611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ee390611ce7565b8015610f305780601f10610f0557610100808354040283529160200191610f30565b820191906000526020600020905b815481529060010190602001808311610f1357829003601f168201915b50505050508152602001600282018054610f4990611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7590611ce7565b8015610fc25780601f10610f9757610100808354040283529160200191610fc2565b820191906000526020600020905b815481529060010190602001808311610fa557829003601f168201915b50505050508152602001600382018054610fdb90611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461100790611ce7565b80156110545780601f1061102957610100808354040283529160200191611054565b820191906000526020600020905b81548152906001019060200180831161103757829003601f168201915b5050505050815260200160048201548152505081526020019060010190610df2565b6000546001600160a01b0316331461108d57600080fd5b60018260405161109d9190611d21565b9081526040516020918190038201902080546001810182556000918252919020825183926006029091019081906110d49082611d8c565b50602082015160018201906110e99082611d8c565b50604082015160028201906110fe9082611d8c565b50606082015181600301556080820151816004015560a0820151816005015550505050565b60606001826040516111359190611d21565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561085257838290600052602060002090600602016040518060c001604052908160008201805461119690611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546111c290611ce7565b801561120f5780601f106111e45761010080835404028352916020019161120f565b820191906000526020600020905b8154815290600101906020018083116111f257829003601f168201915b5050505050815260200160018201805461122890611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461125490611ce7565b80156112a15780601f10611276576101008083540402835291602001916112a1565b820191906000526020600020905b81548152906001019060200180831161128457829003601f168201915b505050505081526020016002820180546112ba90611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546112e690611ce7565b80156113335780601f1061130857610100808354040283529160200191611333565b820191906000526020600020905b81548152906001019060200180831161131657829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505081526020019060010190611163565b6113ae828260405160240161137f929190611e4c565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b1790526113b2565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561140c5761140c6113d3565b60405290565b60405160a0810167ffffffffffffffff8111828210171561140c5761140c6113d3565b6040516080810167ffffffffffffffff8111828210171561140c5761140c6113d3565b60405160c0810167ffffffffffffffff8111828210171561140c5761140c6113d3565b600082601f83011261148c57600080fd5b813567ffffffffffffffff808211156114a7576114a76113d3565b604051601f8301601f19908116603f011681019082821181831017156114cf576114cf6113d3565b816040528381528660208588010111156114e857600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561151b57600080fd5b823567ffffffffffffffff81111561153257600080fd5b61153e8582860161147b565b95602094909401359450505050565b60005b83811015611568578181015183820152602001611550565b50506000910152565b6000815180845261158981602086016020860161154d565b601f01601f19169290920160200192915050565b60a0815260006115b060a0830188611571565b82810360208401526115c28188611571565b905082810360408401526115d68187611571565b905082810360608401526115ea8186611571565b9150508260808301529695505050505050565b6080815260006116106080830187611571565b6020830195909552506040810192909252606090910152919050565b60608152600061163f6060830186611571565b82810360208401526116518186611571565b915050826040830152949350505050565b60006020828403121561167457600080fd5b813567ffffffffffffffff81111561168b57600080fd5b6116978482850161147b565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561171c57603f198984030185528151608081518186526116ec82870182611571565b838b0151878c0152898401518a8801526060938401519390960192909252505093860193908601906001016116c6565b509098975050505050505050565b6000806040838503121561173d57600080fd5b823567ffffffffffffffff8082111561175557600080fd5b6117618683870161147b565b9350602085013591508082111561177757600080fd5b908401906060828703121561178b57600080fd5b6117936113e9565b8235828111156117a257600080fd5b6117ae8882860161147b565b8252506020830135828111156117c357600080fd5b6117cf8882860161147b565b602083015250604083013560408201528093505050509250929050565b600080604083850312156117ff57600080fd5b823567ffffffffffffffff8082111561181757600080fd5b6118238683870161147b565b9350602085013591508082111561183957600080fd5b9084019060a0828703121561184d57600080fd5b611855611412565b82358281111561186457600080fd5b6118708882860161147b565b82525060208301358281111561188557600080fd5b6118918882860161147b565b6020830152506040830135828111156118a957600080fd5b6118b58882860161147b565b6040830152506060830135828111156118cd57600080fd5b6118d98882860161147b565b606083015250608083013560808201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561171c57603f1989840301855281516060815181865261194382870182611571565b915050888201518582038a87015261195b8282611571565b9289015195890195909552509487019492509086019060010161191d565b6000806040838503121561198c57600080fd5b823567ffffffffffffffff808211156119a457600080fd5b6119b08683870161147b565b935060208501359150808211156119c657600080fd5b90840190608082870312156119da57600080fd5b6119e2611435565b8235828111156119f157600080fd5b6119fd8882860161147b565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b60c081526000611a3e60c0830189611571565b8281036020840152611a508189611571565b90508281036040840152611a648188611571565b60608401969096525050608081019290925260a0909101529392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561171c57603f19898403018552815160a08151818652611ad082870182611571565b915050888201518582038a870152611ae88282611571565b9150508782015185820389870152611b008282611571565b91505060608083015186830382880152611b1a8382611571565b6080948501519790940196909652505094870194925090860190600101611aaa565b60008060408385031215611b4f57600080fd5b823567ffffffffffffffff80821115611b6757600080fd5b611b738683870161147b565b93506020850135915080821115611b8957600080fd5b9084019060c08287031215611b9d57600080fd5b611ba5611458565b823582811115611bb457600080fd5b611bc08882860161147b565b825250602083013582811115611bd557600080fd5b611be18882860161147b565b602083015250604083013582811115611bf957600080fd5b611c058882860161147b565b604083015250606083013560608201526080830135608082015260a083013560a08201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561171c57603f19898403018552815160c08151818652611c8382870182611571565b915050888201518582038a870152611c9b8282611571565b9150508782015185820389870152611cb38282611571565b606084810151908801526080808501519088015260a093840151939096019290925250509386019390860190600101611c5d565b600181811c90821680611cfb57607f821691505b602082108103611d1b57634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611d3381846020870161154d565b9190910192915050565b601f821115611d8757600081815260208120601f850160051c81016020861015611d645750805b601f850160051c820191505b81811015611d8357828155600101611d70565b5050505b505050565b815167ffffffffffffffff811115611da657611da66113d3565b611dba81611db48454611ce7565b84611d3d565b602080601f831160018114611def5760008415611dd75750858301515b600019600386901b1c1916600185901b178555611d83565b600085815260208120601f198616915b82811015611e1e57888601518255948401946001909101908401611dff565b5085821015611e3c5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611e5f6040830185611571565b905060018060a01b0383166020830152939250505056fea264697066735822122043173fd8c9749fc1b5ad459b52adad294ad2da2c0f83850ed375dfc14c19842464736f6c63430008110033",
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
// Solidity: function approverTransactions(string , uint256 ) view returns(string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiCaller) ApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "approverTransactions", arg0, arg1)

	outstruct := new(struct {
		Submitby        string
		Status          string
		Timestamp       string
		Numofponds      *big.Int
		Spawningaverage *big.Int
		Creditscore     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Submitby = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Numofponds = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Spawningaverage = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Creditscore = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApproverTransactions is a free data retrieval call binding the contract method 0xd2892d10.
//
// Solidity: function approverTransactions(string , uint256 ) view returns(string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiSession) ApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// ApproverTransactions is a free data retrieval call binding the contract method 0xd2892d10.
//
// Solidity: function approverTransactions(string , uint256 ) view returns(string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiCallerSession) ApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, string status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiCaller) FunderApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     string
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderApproverTransactions", arg0, arg1)

	outstruct := new(struct {
		Reviewer   string
		Status     string
		Filepath   string
		Timestamp  string
		Modalgiven *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reviewer = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Filepath = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Modalgiven = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, string status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     string
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string reviewer, string status, string filepath, string timestamp, uint256 modalgiven)
func (_Api *ApiCallerSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Reviewer   string
	Status     string
	Filepath   string
	Timestamp  string
	Modalgiven *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256,uint256)[])
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
// Solidity: function getApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetApproverTransaction(_nik string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _nik)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetApproverTransaction(_nik string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _nik)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,string,uint256)[])
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
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,string,uint256)[])
func (_Api *ApiSession) GetFunderApproverTransaction(_nik string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _nik)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,string,uint256)[])
func (_Api *ApiCallerSession) GetFunderApproverTransaction(_nik string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _nik)
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

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _nik, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactor) SetApproverTransaction(opts *bind.TransactOpts, _nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApproverTransaction", _nik, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _nik, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiSession) SetApproverTransaction(_nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _nik, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _nik, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactorSession) SetApproverTransaction(_nik string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _nik, _approvertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc3292eda.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,string,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactor) SetFunderApproverTransaction(opts *bind.TransactOpts, _nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setFunderApproverTransaction", _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc3292eda.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,string,uint256) _funderapprovertransaction) returns()
func (_Api *ApiSession) SetFunderApproverTransaction(_nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc3292eda.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,string,uint256) _funderapprovertransaction) returns()
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
