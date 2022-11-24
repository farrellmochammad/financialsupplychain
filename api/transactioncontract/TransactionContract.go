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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fundid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b620011af1760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611ca2806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063bed34bba1161008c578063d2892d1011610066578063d2892d1014610229578063e137df831461024e578063f734b1b31461026e578063fbaa96cf1461028157600080fd5b8063bed34bba146101d3578063cbb37a1d146101f6578063cdc693b11461021657600080fd5b80638da5cb5b116100c85780638da5cb5b1461016057806398e6405f1461018b578063a264efc6146101a0578063aaa5004e146101c057600080fd5b80632455f66d146100ef57806379248d941461011b5780637c5acbe21461013e575b600080fd5b6101026100fd36600461132b565b6102a1565b60405161011294939291906113c0565b60405180910390f35b61012e61012936600461132b565b61040d565b60405161011294939291906113f9565b61015161014c36600461132b565b6104f1565b60405161011293929190611428565b600054610173906001600160a01b031681565b6040516001600160a01b039091168152602001610112565b61019e61019936600461145e565b610657565b005b6101b36101ae36600461152a565b6106e5565b6040516101129190611567565b61019e6101ce3660046115f2565b610812565b6101e66101e13660046116b4565b610896565b6040519015158152602001610112565b61020961020436600461152a565b6108ef565b6040516101129190611718565b61019e61022436600461179b565b610a8f565b61023c61023736600461132b565b610b12565b6040516101129695949392919061184d565b61026161025c36600461152a565b610d12565b60405161011291906118a5565b61019e61027c366004611932565b610ebc565b61029461028f36600461152a565b610f69565b6040516101129190611a2c565b815160208184018101805160028252928201918501919091209190528054829081106102cc57600080fd5b9060005260206000209060040201600091509150508060000180546102f090611add565b80601f016020809104026020016040519081016040528092919081815260200182805461031c90611add565b80156103695780601f1061033e57610100808354040283529160200191610369565b820191906000526020600020905b81548152906001019060200180831161034c57829003601f168201915b50505050509080600101805461037e90611add565b80601f01602080910402602001604051908101604052809291908181526020018280546103aa90611add565b80156103f75780601f106103cc576101008083540402835291602001916103f7565b820191906000526020600020905b8154815290600101906020018083116103da57829003601f168201915b5050505050908060020154908060030154905084565b8151602081840181018051600382529282019185019190912091905280548290811061043857600080fd5b90600052602060002090600402016000915091505080600001805461045c90611add565b80601f016020809104026020016040519081016040528092919081815260200182805461048890611add565b80156104d55780601f106104aa576101008083540402835291602001916104d5565b820191906000526020600020905b8154815290600101906020018083116104b857829003601f168201915b5050505050908060010154908060020154908060030154905084565b8151602081840181018051600482529282019185019190912091905280548290811061051c57600080fd5b90600052602060002090600302016000915091505080600001805461054090611add565b80601f016020809104026020016040519081016040528092919081815260200182805461056c90611add565b80156105b95780601f1061058e576101008083540402835291602001916105b9565b820191906000526020600020905b81548152906001019060200180831161059c57829003601f168201915b5050505050908060010180546105ce90611add565b80601f01602080910402602001604051908101604052809291908181526020018280546105fa90611add565b80156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b5050505050908060020154905083565b6000546001600160a01b0316331461066e57600080fd5b60028260405161067e9190611b17565b9081526040516020918190038201902080546001810182556000918252919020825183926004029091019081906106b59082611b82565b50602082015160018201906106ca9082611b82565b50604082015181600201556060820151816003015550505050565b60606003826040516106f79190611b17565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610807578382906000526020600020906004020160405180608001604052908160008201805461075890611add565b80601f016020809104026020016040519081016040528092919081815260200182805461078490611add565b80156107d15780601f106107a6576101008083540402835291602001916107d1565b820191906000526020600020905b8154815290600101906020018083116107b457829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152505081526020019060010190610725565b505050509050919050565b6000546001600160a01b0316331461082957600080fd5b6004826040516108399190611b17565b9081526040516020918190038201902080546001810182556000918252919020825183926003029091019081906108709082611b82565b50602082015160018201906108859082611b82565b506040820151816002015550505050565b6000816040516020016108a99190611b17565b60405160208183030381529060405280519060200120836040516020016108d09190611b17565b6040516020818303038152906040528051906020012014905092915050565b60606004826040516109019190611b17565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610807578382906000526020600020906003020160405180606001604052908160008201805461096290611add565b80601f016020809104026020016040519081016040528092919081815260200182805461098e90611add565b80156109db5780601f106109b0576101008083540402835291602001916109db565b820191906000526020600020905b8154815290600101906020018083116109be57829003601f168201915b505050505081526020016001820180546109f490611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2090611add565b8015610a6d5780601f10610a4257610100808354040283529160200191610a6d565b820191906000526020600020905b815481529060010190602001808311610a5057829003601f168201915b505050505081526020016002820154815250508152602001906001019061092f565b6000546001600160a01b03163314610aa657600080fd5b600382604051610ab69190611b17565b908152604051602091819003820190208054600181018255600091825291902082518392600402909101908190610aed9082611b82565b5060208201518160010155604082015181600201556060820151816003015550505050565b81516020818401810180516001825292820191850191909120919052805482908110610b3d57600080fd5b906000526020600020906006020160009150915050806000018054610b6190611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8d90611add565b8015610bda5780601f10610baf57610100808354040283529160200191610bda565b820191906000526020600020905b815481529060010190602001808311610bbd57829003601f168201915b505050505090806001018054610bef90611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1b90611add565b8015610c685780601f10610c3d57610100808354040283529160200191610c68565b820191906000526020600020905b815481529060010190602001808311610c4b57829003601f168201915b505050505090806002018054610c7d90611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca990611add565b8015610cf65780601f10610ccb57610100808354040283529160200191610cf6565b820191906000526020600020905b815481529060010190602001808311610cd957829003601f168201915b5050505050908060030154908060040154908060050154905086565b6060600282604051610d249190611b17565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156108075783829060005260206000209060040201604051806080016040529081600082018054610d8590611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610db190611add565b8015610dfe5780601f10610dd357610100808354040283529160200191610dfe565b820191906000526020600020905b815481529060010190602001808311610de157829003601f168201915b50505050508152602001600182018054610e1790611add565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4390611add565b8015610e905780601f10610e6557610100808354040283529160200191610e90565b820191906000526020600020905b815481529060010190602001808311610e7357829003601f168201915b505050505081526020016002820154815260200160038201548152505081526020019060010190610d52565b6000546001600160a01b03163314610ed357600080fd5b600182604051610ee39190611b17565b908152604051602091819003820190208054600181018255600091825291902082518392600602909101908190610f1a9082611b82565b5060208201516001820190610f2f9082611b82565b5060408201516002820190610f449082611b82565b50606082015181600301556080820151816004015560a0820151816005015550505050565b6060600182604051610f7b9190611b17565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561080757838290600052602060002090600602016040518060c0016040529081600082018054610fdc90611add565b80601f016020809104026020016040519081016040528092919081815260200182805461100890611add565b80156110555780601f1061102a57610100808354040283529160200191611055565b820191906000526020600020905b81548152906001019060200180831161103857829003601f168201915b5050505050815260200160018201805461106e90611add565b80601f016020809104026020016040519081016040528092919081815260200182805461109a90611add565b80156110e75780601f106110bc576101008083540402835291602001916110e7565b820191906000526020600020905b8154815290600101906020018083116110ca57829003601f168201915b5050505050815260200160028201805461110090611add565b80601f016020809104026020016040519081016040528092919081815260200182805461112c90611add565b80156111795780601f1061114e57610100808354040283529160200191611179565b820191906000526020600020905b81548152906001019060200180831161115c57829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152505081526020019060010190610fa9565b6111f482826040516024016111c5929190611c42565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b1790526111f8565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561125257611252611219565b60405290565b6040516060810167ffffffffffffffff8111828210171561125257611252611219565b60405160c0810167ffffffffffffffff8111828210171561125257611252611219565b600082601f8301126112af57600080fd5b813567ffffffffffffffff808211156112ca576112ca611219565b604051601f8301601f19908116603f011681019082821181831017156112f2576112f2611219565b8160405283815286602085880101111561130b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561133e57600080fd5b823567ffffffffffffffff81111561135557600080fd5b6113618582860161129e565b95602094909401359450505050565b60005b8381101561138b578181015183820152602001611373565b50506000910152565b600081518084526113ac816020860160208601611370565b601f01601f19169290920160200192915050565b6080815260006113d36080830187611394565b82810360208401526113e58187611394565b604084019590955250506060015292915050565b60808152600061140c6080830187611394565b6020830195909552506040810192909252606090910152919050565b60608152600061143b6060830186611394565b828103602084015261144d8186611394565b915050826040830152949350505050565b6000806040838503121561147157600080fd5b823567ffffffffffffffff8082111561148957600080fd5b6114958683870161129e565b935060208501359150808211156114ab57600080fd5b90840190608082870312156114bf57600080fd5b6114c761122f565b8235828111156114d657600080fd5b6114e28882860161129e565b8252506020830135828111156114f757600080fd5b6115038882860161129e565b60208301525060408301356040820152606083013560608201528093505050509250929050565b60006020828403121561153c57600080fd5b813567ffffffffffffffff81111561155357600080fd5b61155f8482850161129e565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115e457603f198984030185528151608081518186526115b482870182611394565b838b0151878c0152898401518a88015260609384015193909601929092525050938601939086019060010161158e565b509098975050505050505050565b6000806040838503121561160557600080fd5b823567ffffffffffffffff8082111561161d57600080fd5b6116298683870161129e565b9350602085013591508082111561163f57600080fd5b908401906060828703121561165357600080fd5b61165b611258565b82358281111561166a57600080fd5b6116768882860161129e565b82525060208301358281111561168b57600080fd5b6116978882860161129e565b602083015250604083013560408201528093505050509250929050565b600080604083850312156116c757600080fd5b823567ffffffffffffffff808211156116df57600080fd5b6116eb8683870161129e565b9350602085013591508082111561170157600080fd5b5061170e8582860161129e565b9150509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115e457603f1989840301855281516060815181865261176582870182611394565b915050888201518582038a87015261177d8282611394565b9289015195890195909552509487019492509086019060010161173f565b600080604083850312156117ae57600080fd5b823567ffffffffffffffff808211156117c657600080fd5b6117d28683870161129e565b935060208501359150808211156117e857600080fd5b90840190608082870312156117fc57600080fd5b61180461122f565b82358281111561181357600080fd5b61181f8882860161129e565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b60c08152600061186060c0830189611394565b82810360208401526118728189611394565b905082810360408401526118868188611394565b60608401969096525050608081019290925260a0909101529392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115e457603f198984030185528151608081518186526118f282870182611394565b915050888201518582038a87015261190a8282611394565b838a0151878b01526060938401519390960192909252505093860193908601906001016118cc565b6000806040838503121561194557600080fd5b823567ffffffffffffffff8082111561195d57600080fd5b6119698683870161129e565b9350602085013591508082111561197f57600080fd5b9084019060c0828703121561199357600080fd5b61199b61127b565b8235828111156119aa57600080fd5b6119b68882860161129e565b8252506020830135828111156119cb57600080fd5b6119d78882860161129e565b6020830152506040830135828111156119ef57600080fd5b6119fb8882860161129e565b604083015250606083013560608201526080830135608082015260a083013560a08201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156115e457603f19898403018552815160c08151818652611a7982870182611394565b915050888201518582038a870152611a918282611394565b9150508782015185820389870152611aa98282611394565b606084810151908801526080808501519088015260a093840151939096019290925250509386019390860190600101611a53565b600181811c90821680611af157607f821691505b602082108103611b1157634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611b29818460208701611370565b9190910192915050565b601f821115611b7d57600081815260208120601f850160051c81016020861015611b5a5750805b601f850160051c820191505b81811015611b7957828155600101611b66565b5050505b505050565b815167ffffffffffffffff811115611b9c57611b9c611219565b611bb081611baa8454611add565b84611b33565b602080601f831160018114611be55760008415611bcd5750858301515b600019600386901b1c1916600185901b178555611b79565b600085815260208120601f198616915b82811015611c1457888601518255948401946001909101908401611bf5565b5085821015611c325787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611c556040830185611394565b905060018060a01b0383166020830152939250505056fea26469706673582212203e80631a38c054c48ed2bd9e517c6cfca9734a4468cbf15832ebf0a4d364e6d364736f6c63430008110033",
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

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "compareStrings", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiSession) CompareStrings(a string, b string) (bool, error) {
	return _Api.Contract.CompareStrings(&_Api.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Api *ApiCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _Api.Contract.CompareStrings(&_Api.CallOpts, a, b)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCaller) FunderApproverTransactions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "funderApproverTransactions", arg0, arg1)

	outstruct := new(struct {
		Funder        string
		Timestamp     string
		Numberofponds *big.Int
		Amountoffund  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Funder = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Numberofponds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amountoffund = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// FunderApproverTransactions is a free data retrieval call binding the contract method 0x2455f66d.
//
// Solidity: function funderApproverTransactions(string , uint256 ) view returns(string funder, string timestamp, uint256 numberofponds, uint256 amountoffund)
func (_Api *ApiCallerSession) FunderApproverTransactions(arg0 string, arg1 *big.Int) (struct {
	Funder        string
	Timestamp     string
	Numberofponds *big.Int
	Amountoffund  *big.Int
}, error) {
	return _Api.Contract.FunderApproverTransactions(&_Api.CallOpts, arg0, arg1)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _fundid) view returns((string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiCaller) GetApproverTransaction(opts *bind.CallOpts, _fundid string) ([]TransactionContractApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getApproverTransaction", _fundid)

	if err != nil {
		return *new([]TransactionContractApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractApproverTransaction)).(*[]TransactionContractApproverTransaction)

	return out0, err

}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _fundid) view returns((string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiSession) GetApproverTransaction(_fundid string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _fundid)
}

// GetApproverTransaction is a free data retrieval call binding the contract method 0xfbaa96cf.
//
// Solidity: function getApproverTransaction(string _fundid) view returns((string,string,string,uint256,uint256,uint256)[])
func (_Api *ApiCallerSession) GetApproverTransaction(_fundid string) ([]TransactionContractApproverTransaction, error) {
	return _Api.Contract.GetApproverTransaction(&_Api.CallOpts, _fundid)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _fundid) view returns((string,string,uint256,uint256)[])
func (_Api *ApiCaller) GetFunderApproverTransaction(opts *bind.CallOpts, _fundid string) ([]TransactionContractFunderApproverTransaction, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getFunderApproverTransaction", _fundid)

	if err != nil {
		return *new([]TransactionContractFunderApproverTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]TransactionContractFunderApproverTransaction)).(*[]TransactionContractFunderApproverTransaction)

	return out0, err

}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _fundid) view returns((string,string,uint256,uint256)[])
func (_Api *ApiSession) GetFunderApproverTransaction(_fundid string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _fundid)
}

// GetFunderApproverTransaction is a free data retrieval call binding the contract method 0xe137df83.
//
// Solidity: function getFunderApproverTransaction(string _fundid) view returns((string,string,uint256,uint256)[])
func (_Api *ApiCallerSession) GetFunderApproverTransaction(_fundid string) ([]TransactionContractFunderApproverTransaction, error) {
	return _Api.Contract.GetFunderApproverTransaction(&_Api.CallOpts, _fundid)
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

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _fundid, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactor) SetApproverTransaction(opts *bind.TransactOpts, _fundid string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setApproverTransaction", _fundid, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _fundid, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiSession) SetApproverTransaction(_fundid string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _fundid, _approvertransaction)
}

// SetApproverTransaction is a paid mutator transaction binding the contract method 0xf734b1b3.
//
// Solidity: function setApproverTransaction(string _fundid, (string,string,string,uint256,uint256,uint256) _approvertransaction) returns()
func (_Api *ApiTransactorSession) SetApproverTransaction(_fundid string, _approvertransaction TransactionContractApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetApproverTransaction(&_Api.TransactOpts, _fundid, _approvertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0x98e6405f.
//
// Solidity: function setFunderApproverTransaction(string _fundid, (string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactor) SetFunderApproverTransaction(opts *bind.TransactOpts, _fundid string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setFunderApproverTransaction", _fundid, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0x98e6405f.
//
// Solidity: function setFunderApproverTransaction(string _fundid, (string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiSession) SetFunderApproverTransaction(_fundid string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _fundid, _funderapprovertransaction)
}

// SetFunderApproverTransaction is a paid mutator transaction binding the contract method 0x98e6405f.
//
// Solidity: function setFunderApproverTransaction(string _fundid, (string,string,uint256,uint256) _funderapprovertransaction) returns()
func (_Api *ApiTransactorSession) SetFunderApproverTransaction(_fundid string, _funderapprovertransaction TransactionContractFunderApproverTransaction) (*types.Transaction, error) {
	return _Api.Contract.SetFunderApproverTransaction(&_Api.TransactOpts, _fundid, _funderapprovertransaction)
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
