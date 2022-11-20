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
	Nik             string
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}

// TransactionContractFunderApproverTransaction is an auto generated low-level Go binding around an user-defined struct.
type TransactionContractFunderApproverTransaction struct {
	Fundid        string
	Nik           string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b620014381760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611fb6806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063aaa5004e1161008c578063be3ca5a711610066578063be3ca5a7146101f7578063cbb37a1d1461020a578063cdc298b21461022a578063cdc693b11461023d57600080fd5b8063aaa5004e146101b8578063ad948224146101cd578063bd882eb2146101e257600080fd5b806358c6600d146100d45780636f847e631461010357806379248d94146101285780637c5acbe21461014b5780638da5cb5b1461016d578063a264efc614610198575b600080fd5b6100e76100e23660046114a2565b610250565b6040516100fa979695949392919061150b565b60405180910390f35b6101166101113660046114a2565b6104c2565b6040516100fa96959493929190611578565b61013b610136366004611710565b61072e565b6040516100fa9493929190611755565b61015e610159366004611710565b610812565b6040516100fa93929190611784565b600054610180906001600160a01b031681565b6040516001600160a01b0390911681526020016100fa565b6101ab6101a63660046117ba565b610978565b6040516100fa91906117f7565b6101cb6101c6366004611882565b610aa5565b005b6101d5610b29565b6040516100fa9190611944565b6101ea610de4565b6040516100fa9190611a06565b6101cb610205366004611ad2565b6110a0565b61021d6102183660046117ba565b61115f565b6040516100fa9190611bcc565b6101cb610238366004611c4f565b6112ff565b6101cb61024b366004611d3f565b6113b5565b6001818154811061026057600080fd5b906000526020600020906007020160009150905080600001805461028390611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546102af90611df1565b80156102fc5780601f106102d1576101008083540402835291602001916102fc565b820191906000526020600020905b8154815290600101906020018083116102df57829003601f168201915b50505050509080600101805461031190611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461033d90611df1565b801561038a5780601f1061035f5761010080835404028352916020019161038a565b820191906000526020600020905b81548152906001019060200180831161036d57829003601f168201915b50505050509080600201805461039f90611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546103cb90611df1565b80156104185780601f106103ed57610100808354040283529160200191610418565b820191906000526020600020905b8154815290600101906020018083116103fb57829003601f168201915b50505050509080600301805461042d90611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461045990611df1565b80156104a65780601f1061047b576101008083540402835291602001916104a6565b820191906000526020600020905b81548152906001019060200180831161048957829003601f168201915b5050505050908060040154908060050154908060060154905087565b600281815481106104d257600080fd5b90600052602060002090600602016000915090508060000180546104f590611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461052190611df1565b801561056e5780601f106105435761010080835404028352916020019161056e565b820191906000526020600020905b81548152906001019060200180831161055157829003601f168201915b50505050509080600101805461058390611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546105af90611df1565b80156105fc5780601f106105d1576101008083540402835291602001916105fc565b820191906000526020600020905b8154815290600101906020018083116105df57829003601f168201915b50505050509080600201805461061190611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461063d90611df1565b801561068a5780601f1061065f5761010080835404028352916020019161068a565b820191906000526020600020905b81548152906001019060200180831161066d57829003601f168201915b50505050509080600301805461069f90611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546106cb90611df1565b80156107185780601f106106ed57610100808354040283529160200191610718565b820191906000526020600020905b8154815290600101906020018083116106fb57829003601f168201915b5050505050908060040154908060050154905086565b8151602081840181018051600382529282019185019190912091905280548290811061075957600080fd5b90600052602060002090600402016000915091505080600001805461077d90611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546107a990611df1565b80156107f65780601f106107cb576101008083540402835291602001916107f6565b820191906000526020600020905b8154815290600101906020018083116107d957829003601f168201915b5050505050908060010154908060020154908060030154905084565b8151602081840181018051600482529282019185019190912091905280548290811061083d57600080fd5b90600052602060002090600302016000915091505080600001805461086190611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461088d90611df1565b80156108da5780601f106108af576101008083540402835291602001916108da565b820191906000526020600020905b8154815290600101906020018083116108bd57829003601f168201915b5050505050908060010180546108ef90611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461091b90611df1565b80156109685780601f1061093d57610100808354040283529160200191610968565b820191906000526020600020905b81548152906001019060200180831161094b57829003601f168201915b5050505050908060020154905083565b606060038260405161098a9190611e2b565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a9a57838290600052602060002090600402016040518060800160405290816000820180546109eb90611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1790611df1565b8015610a645780601f10610a3957610100808354040283529160200191610a64565b820191906000526020600020905b815481529060010190602001808311610a4757829003601f168201915b505050505081526020016001820154815260200160028201548152602001600382015481525050815260200190600101906109b8565b505050509050919050565b6000546001600160a01b03163314610abc57600080fd5b600482604051610acc9190611e2b565b908152604051602091819003820190208054600181018255600091825291902082518392600302909101908190610b039082611e96565b5060208201516001820190610b189082611e96565b506040820151816002015550505050565b60606002805480602002602001604051908101604052809291908181526020016000905b82821015610ddb57838290600052602060002090600602016040518060c0016040529081600082018054610b8090611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610bac90611df1565b8015610bf95780601f10610bce57610100808354040283529160200191610bf9565b820191906000526020600020905b815481529060010190602001808311610bdc57829003601f168201915b50505050508152602001600182018054610c1290611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3e90611df1565b8015610c8b5780601f10610c6057610100808354040283529160200191610c8b565b820191906000526020600020905b815481529060010190602001808311610c6e57829003601f168201915b50505050508152602001600282018054610ca490611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd090611df1565b8015610d1d5780601f10610cf257610100808354040283529160200191610d1d565b820191906000526020600020905b815481529060010190602001808311610d0057829003601f168201915b50505050508152602001600382018054610d3690611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6290611df1565b8015610daf5780601f10610d8457610100808354040283529160200191610daf565b820191906000526020600020905b815481529060010190602001808311610d9257829003601f168201915b505050505081526020016004820154815260200160058201548152505081526020019060010190610b4d565b50505050905090565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610ddb57838290600052602060002090600702016040518060e0016040529081600082018054610e3b90611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6790611df1565b8015610eb45780601f10610e8957610100808354040283529160200191610eb4565b820191906000526020600020905b815481529060010190602001808311610e9757829003601f168201915b50505050508152602001600182018054610ecd90611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610ef990611df1565b8015610f465780601f10610f1b57610100808354040283529160200191610f46565b820191906000526020600020905b815481529060010190602001808311610f2957829003601f168201915b50505050508152602001600282018054610f5f90611df1565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8b90611df1565b8015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b50505050508152602001600382018054610ff190611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461101d90611df1565b801561106a5780601f1061103f5761010080835404028352916020019161106a565b820191906000526020600020905b81548152906001019060200180831161104d57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152505081526020019060010190610e08565b6000546001600160a01b031633146110b757600080fd5b600180548082018255600091909152815182916007027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6019081906110fc9082611e96565b50602082015160018201906111119082611e96565b50604082015160028201906111269082611e96565b506060820151600382019061113b9082611e96565b506080820151816004015560a0820151816005015560c08201518160060155505050565b60606004826040516111719190611e2b565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a9a57838290600052602060002090600302016040518060600160405290816000820180546111d290611df1565b80601f01602080910402602001604051908101604052809291908181526020018280546111fe90611df1565b801561124b5780601f106112205761010080835404028352916020019161124b565b820191906000526020600020905b81548152906001019060200180831161122e57829003601f168201915b5050505050815260200160018201805461126490611df1565b80601f016020809104026020016040519081016040528092919081815260200182805461129090611df1565b80156112dd5780601f106112b2576101008083540402835291602001916112dd565b820191906000526020600020905b8154815290600101906020018083116112c057829003601f168201915b505050505081526020016002820154815250508152602001906001019061119f565b6000546001600160a01b0316331461131657600080fd5b60028054600181018255600091909152815182916006027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0190819061135c9082611e96565b50602082015160018201906113719082611e96565b50604082015160028201906113869082611e96565b506060820151600382019061139b9082611e96565b506080820151816004015560a08201518160050155505050565b6000546001600160a01b031633146113cc57600080fd5b6003826040516113dc9190611e2b565b9081526040516020918190038201902080546001810182556000918252919020825183926004029091019081906114139082611e96565b5060208201518160010155604082015181600201556060820151816003015550505050565b61147d828260405160240161144e929190611f56565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611481565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6000602082840312156114b457600080fd5b5035919050565b60005b838110156114d65781810151838201526020016114be565b50506000910152565b600081518084526114f78160208601602086016114bb565b601f01601f19169290920160200192915050565b60e08152600061151e60e083018a6114df565b8281036020840152611530818a6114df565b9050828103604084015261154481896114df565b9050828103606084015261155881886114df565b6080840196909652505060a081019290925260c090910152949350505050565b60c08152600061158b60c08301896114df565b828103602084015261159d81896114df565b905082810360408401526115b181886114df565b905082810360608401526115c581876114df565b6080840195909552505060a00152949350505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715611614576116146115db565b60405290565b60405160e0810167ffffffffffffffff81118282101715611614576116146115db565b60405160c0810167ffffffffffffffff81118282101715611614576116146115db565b6040516080810167ffffffffffffffff81118282101715611614576116146115db565b600082601f83011261169457600080fd5b813567ffffffffffffffff808211156116af576116af6115db565b604051601f8301601f19908116603f011681019082821181831017156116d7576116d76115db565b816040528381528660208588010111156116f057600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561172357600080fd5b823567ffffffffffffffff81111561173a57600080fd5b61174685828601611683565b95602094909401359450505050565b60808152600061176860808301876114df565b6020830195909552506040810192909252606090910152919050565b60608152600061179760608301866114df565b82810360208401526117a981866114df565b915050826040830152949350505050565b6000602082840312156117cc57600080fd5b813567ffffffffffffffff8111156117e357600080fd5b6117ef84828501611683565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561187457603f19898403018552815160808151818652611844828701826114df565b838b0151878c0152898401518a88015260609384015193909601929092525050938601939086019060010161181e565b509098975050505050505050565b6000806040838503121561189557600080fd5b823567ffffffffffffffff808211156118ad57600080fd5b6118b986838701611683565b935060208501359150808211156118cf57600080fd5b90840190606082870312156118e357600080fd5b6118eb6115f1565b8235828111156118fa57600080fd5b61190688828601611683565b82525060208301358281111561191b57600080fd5b61192788828601611683565b602083015250604083013560408201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561187457603f19898403018552815160c08151818652611991828701826114df565b915050888201518582038a8701526119a982826114df565b91505087820151858203898701526119c182826114df565b915050606080830151868303828801526119db83826114df565b6080858101519089015260a0948501519490970193909352505050938601939086019060010161196b565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561187457603f19898403018552815160e08151818652611a53828701826114df565b915050888201518582038a870152611a6b82826114df565b9150508782015185820389870152611a8382826114df565b91505060608083015186830382880152611a9d83826114df565b6080858101519089015260a0808601519089015260c09485015194909701939093525050509386019390860190600101611a2d565b600060208284031215611ae457600080fd5b813567ffffffffffffffff80821115611afc57600080fd5b9083019060e08286031215611b1057600080fd5b611b1861161a565b823582811115611b2757600080fd5b611b3387828601611683565b825250602083013582811115611b4857600080fd5b611b5487828601611683565b602083015250604083013582811115611b6c57600080fd5b611b7887828601611683565b604083015250606083013582811115611b9057600080fd5b611b9c87828601611683565b6060830152506080830135608082015260a083013560a082015260c083013560c082015280935050505092915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561187457603f19898403018552815160608151818652611c19828701826114df565b915050888201518582038a870152611c3182826114df565b92890151958901959095525094870194925090860190600101611bf3565b600060208284031215611c6157600080fd5b813567ffffffffffffffff80821115611c7957600080fd5b9083019060c08286031215611c8d57600080fd5b611c9561163d565b823582811115611ca457600080fd5b611cb087828601611683565b825250602083013582811115611cc557600080fd5b611cd187828601611683565b602083015250604083013582811115611ce957600080fd5b611cf587828601611683565b604083015250606083013582811115611d0d57600080fd5b611d1987828601611683565b6060830152506080830135608082015260a083013560a082015280935050505092915050565b60008060408385031215611d5257600080fd5b823567ffffffffffffffff80821115611d6a57600080fd5b611d7686838701611683565b93506020850135915080821115611d8c57600080fd5b9084019060808287031215611da057600080fd5b611da8611660565b823582811115611db757600080fd5b611dc388828601611683565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b600181811c90821680611e0557607f821691505b602082108103611e2557634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611e3d8184602087016114bb565b9190910192915050565b601f821115611e9157600081815260208120601f850160051c81016020861015611e6e5750805b601f850160051c820191505b81811015611e8d57828155600101611e7a565b5050505b505050565b815167ffffffffffffffff811115611eb057611eb06115db565b611ec481611ebe8454611df1565b84611e47565b602080601f831160018114611ef95760008415611ee15750858301515b600019600386901b1c1916600185901b178555611e8d565b600085815260208120601f198616915b82811015611f2857888601518255948401946001909101908401611f09565b5085821015611f465787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611f6960408301856114df565b905060018060a01b0383166020830152939250505056fea26469706673582212208a79314dc8b041772d7ce5641c096072c073837739e301a285c8aee8625d9cd864736f6c63430008110033",
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

// ApproverTransactions is a free data retrieval call binding the contract method 0x58c6600d.
//
// Solidity: function approverTransactions(uint256 ) view returns(string nik, string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiCaller) ApproverTransactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Nik             string
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "approverTransactions", arg0)

	outstruct := new(struct {
		Nik             string
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

	outstruct.Nik = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Submitby = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Numofponds = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Spawningaverage = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Creditscore = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApproverTransactions is a free data retrieval call binding the contract method 0x58c6600d.
//
// Solidity: function approverTransactions(uint256 ) view returns(string nik, string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiSession) ApproverTransactions(arg0 *big.Int) (struct {
	Nik             string
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0)
}

// ApproverTransactions is a free data retrieval call binding the contract method 0x58c6600d.
//
// Solidity: function approverTransactions(uint256 ) view returns(string nik, string submitby, string status, string timestamp, uint256 numofponds, uint256 spawningaverage, uint256 creditscore)
func (_Api *ApiCallerSession) ApproverTransactions(arg0 *big.Int) (struct {
	Nik             string
	Submitby        string
	Status          string
	Timestamp       string
	Numofponds      *big.Int
	Spawningaverage *big.Int
	Creditscore     *big.Int
}, error) {
	return _Api.Contract.ApproverTransactions(&_Api.CallOpts, arg0)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x6f847e63.
//
// Solidity: function funderApproverTransactions(uint256 ) view returns(string fundid, string nik, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCaller) FunderApproverTransactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Fundid        string
	Nik           string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderApproverTransactions", arg0)

	outstruct := new(struct {
		Fundid        string
		Nik           string
		Funder        string
		Timestamp     string
		Numberofponds *big.Int
		Amountoffund  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fundid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nik = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Funder = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Numberofponds = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Amountoffund = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x6f847e63.
//
// Solidity: function funderApproverTransactions(uint256 ) view returns(string fundid, string nik, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiSession) FunderApproverTransactions(arg0 *big.Int) (struct {
	Fundid        string
	Nik           string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x6f847e63.
//
// Solidity: function funderApproverTransactions(uint256 ) view returns(string fundid, string nik, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCallerSession) FunderApproverTransactions(arg0 *big.Int) (struct {
	Fundid        string
	Nik           string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xbd882eb2.
//
// Solidity: function getApproverTransaction() view returns((string,string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiCaller) GetApproverTransaction(opts *bind.CallOpts) ([]TransactionContractApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApproverTransaction")

	if err != nil {
		return *new([]TransactionContractApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractApproverTransaction)).(*[]TransactionContractApproverTransaction)

	return out0, err

}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xbd882eb2.
//
// Solidity: function getApproverTransaction() view returns((string,string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetApproverTransaction() ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xbd882eb2.
//
// Solidity: function getApproverTransaction() view returns((string,string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetApproverTransaction() ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xad948224.
//
// Solidity: function getFunderApproverTransaction() view returns((string,string,string,string,uint256,uint256)[])
func (_Api *ApiCaller) GetFunderApproverTransaction(opts *bind.CallOpts) ([]TransactionContractFunderApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getFunderApproverTransaction")

	if err != nil {
		return *new([]TransactionContractFunderApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractFunderApproverTransaction)).(*[]TransactionContractFunderApproverTransaction)

	return out0, err

}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xad948224.
//
// Solidity: function getFunderApproverTransaction() view returns((string,string,string,string,uint256,uint256)[])
func (_Api *ApiSession) GetFunderApproverTransaction() ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xad948224.
//
// Solidity: function getFunderApproverTransaction() view returns((string,string,string,string,uint256,uint256)[])
func (_Api *ApiCallerSession) GetFunderApproverTransaction() ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts)
}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _fundid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCaller) GetMonitoring(opts *bind.CallOpts, _fundid string) ([]TransactionContractMonitoring, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getMonitoring", _fundid)

	if err != nil {
		return *new([]TransactionContractMonitoring), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractMonitoring)).(*[]TransactionContractMonitoring)

	return out0, err

}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _fundid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetMonitoring(_fundid string) ([]TransactionContractMonitoring, error) {
	return _Api.Contract.GetMonitoring(&_Api.CallOpts, _fundid)
}

// GetMonitoring is a free data retrieval call binding the contract method 0xa264efc6.
//
// Solidity: function getMonitoring(string _fundid) view returns((string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetMonitoring(_fundid string) ([]TransactionContractMonitoring, error) {
	return _Api.Contract.GetMonitoring(&_Api.CallOpts, _fundid)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _fundid) view returns((string,string,uint256)[])
func (_Api *ApiCaller) GetSpawning(opts *bind.CallOpts, _fundid string) ([]TransactionContractSpawning, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSpawning", _fundid)

	if err != nil {
		return *new([]TransactionContractSpawning), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractSpawning)).(*[]TransactionContractSpawning)

	return out0, err

}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _fundid) view returns((string,string,uint256)[])
func (_Api *ApiSession) GetSpawning(_fundid string) ([]TransactionContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _fundid)
}

// GetSpawning is a free data retrieval call binding the contract method 0xcbb37a1d.
//
// Solidity: function getSpawning(string _fundid) view returns((string,string,uint256)[])
func (_Api *ApiCallerSession) GetSpawning(_fundid string) ([]TransactionContractSpawning, error) {
	return _Api.Contract.GetSpawning(&_Api.CallOpts, _fundid)
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

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xbe3ca5a7.
//
// Solidity: function setApproverTransaction((string,string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactor) SetApproverTransaction(opts *bind.TransactOpts, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApproverTransaction", _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xbe3ca5a7.
//
// Solidity: function setApproverTransaction((string,string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiSession) SetApproverTransaction(_approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xbe3ca5a7.
//
// Solidity: function setApproverTransaction((string,string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactorSession) SetApproverTransaction(_approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _approvertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xcdc298b2.
//
// Solidity: function setFunderApproverTransaction((string,string,string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactor) SetFunderApproverTransaction(opts *bind.TransactOpts, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setFunderApproverTransaction", _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xcdc298b2.
//
// Solidity: function setFunderApproverTransaction((string,string,string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiSession) SetFunderApproverTransaction(_funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xcdc298b2.
//
// Solidity: function setFunderApproverTransaction((string,string,string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactorSession) SetFunderApproverTransaction(_funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _funderapprovertransaction)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _fundid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiTransactor) SetMonitoring(opts *bind.TransactOpts, _fundid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setMonitoring", _fundid, _monitoring)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _fundid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiSession) SetMonitoring(_fundid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.Contract.SetMonitoring(&_Api.TransactOpts, _fundid, _monitoring)
}

// SetMonitoring is a paid mutator transaction binding the contract method 0xcdc693b1.
//
// Solidity: function setMonitoring(string _fundid, (string,uint256,uint256,uint256) _monitoring) returns()
func (_Api *ApiTransactorSession) SetMonitoring(_fundid string, _monitoring TransactionContractMonitoring) (*types.Transaction, error) {
	return _Api.Contract.SetMonitoring(&_Api.TransactOpts, _fundid, _monitoring)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _fundid, (string,string,uint256) _spawning) returns()
func (_Api *ApiTransactor) SetSpawning(opts *bind.TransactOpts, _fundid string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setSpawning", _fundid, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _fundid, (string,string,uint256) _spawning) returns()
func (_Api *ApiSession) SetSpawning(_fundid string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _fundid, _spawning)
}

// SetSpawning is a paid mutator transaction binding the contract method 0xaaa5004e.
//
// Solidity: function setSpawning(string _fundid, (string,string,uint256) _spawning) returns()
func (_Api *ApiTransactorSession) SetSpawning(_fundid string, _spawning TransactionContractSpawning) (*types.Transaction, error) {
	return _Api.Contract.SetSpawning(&_Api.TransactOpts, _fundid, _spawning)
}
