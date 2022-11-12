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
	Fundid        string
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b6200124e1760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611d53806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063c91a2bd21161008c578063d2892d1011610066578063d2892d10146101ec578063e137df8314610211578063f734b1b314610231578063fbaa96cf1461024457600080fd5b8063c91a2bd2146101a6578063cbb37a1d146101b9578063cdc693b1146101d957600080fd5b80632455f66d146100d457806379248d94146101015780637c5acbe2146101245780638da5cb5b14610146578063a264efc614610171578063aaa5004e14610191575b600080fd5b6100e76100e23660046113ed565b610264565b6040516100f8959493929190611482565b60405180910390f35b61011461010f3660046113ed565b61045e565b6040516100f894939291906114d0565b6101376101323660046113ed565b610542565b6040516100f8939291906114ff565b600054610159906001600160a01b031681565b6040516001600160a01b0390911681526020016100f8565b61018461017f366004611535565b6106a8565b6040516100f89190611572565b6101a461019f3660046115fd565b6107d5565b005b6101a46101b43660046116bf565b610859565b6101cc6101c7366004611535565b6108fc565b6040516100f891906117af565b6101a46101e7366004611832565b610a9c565b6101ff6101fa3660046113ed565b610b1f565b6040516100f8969594939291906118e4565b61022461021f366004611535565b610d1f565b6040516100f8919061193c565b6101a461023f3660046119e3565b610f5b565b610257610252366004611535565b611008565b6040516100f89190611add565b8151602081840181018051600282529282019185019190912091905280548290811061028f57600080fd5b9060005260206000209060050201600091509150508060000180546102b390611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546102df90611b8e565b801561032c5780601f106103015761010080835404028352916020019161032c565b820191906000526020600020905b81548152906001019060200180831161030f57829003601f168201915b50505050509080600101805461034190611b8e565b80601f016020809104026020016040519081016040528092919081815260200182805461036d90611b8e565b80156103ba5780601f1061038f576101008083540402835291602001916103ba565b820191906000526020600020905b81548152906001019060200180831161039d57829003601f168201915b5050505050908060020180546103cf90611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546103fb90611b8e565b80156104485780601f1061041d57610100808354040283529160200191610448565b820191906000526020600020905b81548152906001019060200180831161042b57829003601f168201915b5050505050908060030154908060040154905085565b8151602081840181018051600382529282019185019190912091905280548290811061048957600080fd5b9060005260206000209060040201600091509150508060000180546104ad90611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546104d990611b8e565b80156105265780601f106104fb57610100808354040283529160200191610526565b820191906000526020600020905b81548152906001019060200180831161050957829003601f168201915b5050505050908060010154908060020154908060030154905084565b8151602081840181018051600482529282019185019190912091905280548290811061056d57600080fd5b90600052602060002090600302016000915091505080600001805461059190611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546105bd90611b8e565b801561060a5780601f106105df5761010080835404028352916020019161060a565b820191906000526020600020905b8154815290600101906020018083116105ed57829003601f168201915b50505050509080600101805461061f90611b8e565b80601f016020809104026020016040519081016040528092919081815260200182805461064b90611b8e565b80156106985780601f1061066d57610100808354040283529160200191610698565b820191906000526020600020905b81548152906001019060200180831161067b57829003601f168201915b5050505050908060020154905083565b60606003826040516106ba9190611bc8565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ca578382906000526020600020906004020160405180608001604052908160008201805461071b90611b8e565b80601f016020809104026020016040519081016040528092919081815260200182805461074790611b8e565b80156107945780601f1061076957610100808354040283529160200191610794565b820191906000526020600020905b81548152906001019060200180831161077757829003601f168201915b505050505081526020016001820154815260200160028201548152602001600382015481525050815260200190600101906106e8565b505050509050919050565b6000546001600160a01b031633146107ec57600080fd5b6004826040516107fc9190611bc8565b9081526040516020918190038201902080546001810182556000918252919020825183926003029091019081906108339082611c33565b50602082015160018201906108489082611c33565b506040820151816002015550505050565b6000546001600160a01b0316331461087057600080fd5b6002826040516108809190611bc8565b9081526040516020918190038201902080546001810182556000918252919020825183926005029091019081906108b79082611c33565b50602082015160018201906108cc9082611c33565b50604082015160028201906108e19082611c33565b50606082015181600301556080820151816004015550505050565b606060048260405161090e9190611bc8565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ca578382906000526020600020906003020160405180606001604052908160008201805461096f90611b8e565b80601f016020809104026020016040519081016040528092919081815260200182805461099b90611b8e565b80156109e85780601f106109bd576101008083540402835291602001916109e8565b820191906000526020600020905b8154815290600101906020018083116109cb57829003601f168201915b50505050508152602001600182018054610a0190611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2d90611b8e565b8015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b505050505081526020016002820154815250508152602001906001019061093c565b6000546001600160a01b03163314610ab357600080fd5b600382604051610ac39190611bc8565b908152604051602091819003820190208054600181018255600091825291902082518392600402909101908190610afa9082611c33565b5060208201518160010155604082015181600201556060820151816003015550505050565b81516020818401810180516001825292820191850191909120919052805482908110610b4a57600080fd5b906000526020600020906006020160009150915050806000018054610b6e90611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9a90611b8e565b8015610be75780601f10610bbc57610100808354040283529160200191610be7565b820191906000526020600020905b815481529060010190602001808311610bca57829003601f168201915b505050505090806001018054610bfc90611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2890611b8e565b8015610c755780601f10610c4a57610100808354040283529160200191610c75565b820191906000526020600020905b815481529060010190602001808311610c5857829003601f168201915b505050505090806002018054610c8a90611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb690611b8e565b8015610d035780601f10610cd857610100808354040283529160200191610d03565b820191906000526020600020905b815481529060010190602001808311610ce657829003601f168201915b5050505050908060030154908060040154908060050154905086565b6060600282604051610d319190611bc8565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ca57838290600052602060002090600502016040518060a0016040529081600082018054610d9290611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610dbe90611b8e565b8015610e0b5780601f10610de057610100808354040283529160200191610e0b565b820191906000526020600020905b815481529060010190602001808311610dee57829003601f168201915b50505050508152602001600182018054610e2490611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5090611b8e565b8015610e9d5780601f10610e7257610100808354040283529160200191610e9d565b820191906000526020600020905b815481529060010190602001808311610e8057829003601f168201915b50505050508152602001600282018054610eb690611b8e565b80601f0160208091040260200160405190810160405280929190818152602001828054610ee290611b8e565b8015610f2f5780601f10610f0457610100808354040283529160200191610f2f565b820191906000526020600020905b815481529060010190602001808311610f1257829003601f168201915b505050505081526020016003820154815260200160048201548152505081526020019060010190610d5f565b6000546001600160a01b03163314610f7257600080fd5b600182604051610f829190611bc8565b908152604051602091819003820190208054600181018255600091825291902082518392600602909101908190610fb99082611c33565b5060208201516001820190610fce9082611c33565b5060408201516002820190610fe39082611c33565b50606082015181600301556080820151816004015560a0820151816005015550505050565b606060018260405161101a9190611bc8565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156107ca57838290600052602060002090600602016040518060c001604052908160008201805461107b90611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546110a790611b8e565b80156110f45780601f106110c9576101008083540402835291602001916110f4565b820191906000526020600020905b8154815290600101906020018083116110d757829003601f168201915b5050505050815260200160018201805461110d90611b8e565b80601f016020809104026020016040519081016040528092919081815260200182805461113990611b8e565b80156111865780601f1061115b57610100808354040283529160200191611186565b820191906000526020600020905b81548152906001019060200180831161116957829003601f168201915b5050505050815260200160028201805461119f90611b8e565b80601f01602080910402602001604051908101604052809291908181526020018280546111cb90611b8e565b80156112185780601f106111ed57610100808354040283529160200191611218565b820191906000526020600020905b8154815290600101906020018083116111fb57829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505081526020019060010190611048565b6112938282604051602401611264929190611cf3565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611297565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112f1576112f16112b8565b60405290565b60405160a0810167ffffffffffffffff811182821017156112f1576112f16112b8565b6040516080810167ffffffffffffffff811182821017156112f1576112f16112b8565b60405160c0810167ffffffffffffffff811182821017156112f1576112f16112b8565b600082601f83011261137157600080fd5b813567ffffffffffffffff8082111561138c5761138c6112b8565b604051601f8301601f19908116603f011681019082821181831017156113b4576113b46112b8565b816040528381528660208588010111156113cd57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561140057600080fd5b823567ffffffffffffffff81111561141757600080fd5b61142385828601611360565b95602094909401359450505050565b60005b8381101561144d578181015183820152602001611435565b50506000910152565b6000815180845261146e816020860160208601611432565b601f01601f19169290920160200192915050565b60a08152600061149560a0830188611456565b82810360208401526114a78188611456565b905082810360408401526114bb8187611456565b60608401959095525050608001529392505050565b6080815260006114e36080830187611456565b6020830195909552506040810192909252606090910152919050565b6060815260006115126060830186611456565b82810360208401526115248186611456565b915050826040830152949350505050565b60006020828403121561154757600080fd5b813567ffffffffffffffff81111561155e57600080fd5b61156a84828501611360565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115ef57603f198984030185528151608081518186526115bf82870182611456565b838b0151878c0152898401518a880152606093840151939096019290925250509386019390860190600101611599565b509098975050505050505050565b6000806040838503121561161057600080fd5b823567ffffffffffffffff8082111561162857600080fd5b61163486838701611360565b9350602085013591508082111561164a57600080fd5b908401906060828703121561165e57600080fd5b6116666112ce565b82358281111561167557600080fd5b61168188828601611360565b82525060208301358281111561169657600080fd5b6116a288828601611360565b602083015250604083013560408201528093505050509250929050565b600080604083850312156116d257600080fd5b823567ffffffffffffffff808211156116ea57600080fd5b6116f686838701611360565b9350602085013591508082111561170c57600080fd5b9084019060a0828703121561172057600080fd5b6117286112f7565b82358281111561173757600080fd5b61174388828601611360565b82525060208301358281111561175857600080fd5b61176488828601611360565b60208301525060408301358281111561177c57600080fd5b61178888828601611360565b60408301525060608301356060820152608083013560808201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115ef57603f198984030185528151606081518186526117fc82870182611456565b915050888201518582038a8701526118148282611456565b928901519589019590955250948701949250908601906001016117d6565b6000806040838503121561184557600080fd5b823567ffffffffffffffff8082111561185d57600080fd5b61186986838701611360565b9350602085013591508082111561187f57600080fd5b908401906080828703121561189357600080fd5b61189b61131a565b8235828111156118aa57600080fd5b6118b688828601611360565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b60c0815260006118f760c0830189611456565b82810360208401526119098189611456565b9050828103604084015261191d8188611456565b60608401969096525050608081019290925260a0909101529392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115ef57603f19898403018552815160a0815181865261198982870182611456565b915050888201518582038a8701526119a18282611456565b91505087820151858203898701526119b98282611456565b60608481015190880152608093840151939096019290925250509386019390860190600101611963565b600080604083850312156119f657600080fd5b823567ffffffffffffffff80821115611a0e57600080fd5b611a1a86838701611360565b93506020850135915080821115611a3057600080fd5b9084019060c08287031215611a4457600080fd5b611a4c61133d565b823582811115611a5b57600080fd5b611a6788828601611360565b825250602083013582811115611a7c57600080fd5b611a8888828601611360565b602083015250604083013582811115611aa057600080fd5b611aac88828601611360565b604083015250606083013560608201526080830135608082015260a083013560a08201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115ef57603f19898403018552815160c08151818652611b2a82870182611456565b915050888201518582038a870152611b428282611456565b9150508782015185820389870152611b5a8282611456565b606084810151908801526080808501519088015260a093840151939096019290925250509386019390860190600101611b04565b600181811c90821680611ba257607f821691505b602082108103611bc257634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611bda818460208701611432565b9190910192915050565b601f821115611c2e57600081815260208120601f850160051c81016020861015611c0b5750805b601f850160051c820191505b81811015611c2a57828155600101611c17565b5050505b505050565b815167ffffffffffffffff811115611c4d57611c4d6112b8565b611c6181611c5b8454611b8e565b84611be4565b602080601f831160018114611c965760008415611c7e5750858301515b600019600386901b1c1916600185901b178555611c2a565b600085815260208120601f198616915b82811015611cc557888601518255948401946001909101908401611ca6565b5085821015611ce35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611d066040830185611456565b905060018060a01b0383166020830152939250505056fea264697066735822122088d1b0abfd83afa7095d4ef543d2065f31fe44ab8e7175f8fc85ad791201e32e64736f6c63430008110033",
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
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string fundid, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCaller) FunderApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Fundid        string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderApproverTransactions", arg0, arg1)

	outstruct := new(struct {
		Fundid        string
		Funder        string
		Timestamp     string
		Numberofponds *big.Int
		Amountoffund  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fundid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Funder = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Numberofponds = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amountoffund = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string fundid, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Fundid        string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string fundid, string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCallerSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Fundid        string
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
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
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256)[])
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
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256)[])
func (_Api *ApiSession) GetFunderApproverTransaction(_nik string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _nik)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _nik) view returns((string,string,string,uint256,uint256)[])
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

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc91a2bd2.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactor) SetFunderApproverTransaction(opts *bind.TransactOpts, _nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setFunderApproverTransaction", _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc91a2bd2.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiSession) SetFunderApproverTransaction(_nik string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _nik, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0xc91a2bd2.
//
// Solidity: function setFunderApproverTransaction(string _nik, (string,string,string,uint256,uint256) _funderapprovertransaction) returns()
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
