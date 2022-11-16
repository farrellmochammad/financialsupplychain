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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funderApproverTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getFunderApproverTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"}],\"name\":\"getMonitoring\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondMonitorings\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"submitby\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spawningaverage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditscore\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.ApproverTransaction\",\"name\":\"_approvertransaction\",\"type\":\"tuple\"}],\"name\":\"setApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"fundid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funder\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numberofponds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountoffund\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.FunderApproverTransaction\",\"name\":\"_funderapprovertransaction\",\"type\":\"tuple\"}],\"name\":\"setFunderApproverTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_pondid\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"temperature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"humidity\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Monitoring\",\"name\":\"_monitoring\",\"type\":\"tuple\"}],\"name\":\"setMonitoring\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structTransactionContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningHistory\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"pondid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b620013471760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611e95806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063aaa5004e1161008c578063c91a2bd211610066578063c91a2bd2146101f4578063cbb37a1d14610207578063cdc693b114610227578063e137df831461023a57600080fd5b8063aaa5004e146101b7578063bd882eb2146101cc578063be3ca5a7146101e157600080fd5b80632455f66d146100d457806358c6600d1461010157806379248d94146101275780637c5acbe21461014a5780638da5cb5b1461016c578063a264efc614610197575b600080fd5b6100e76100e23660046114e6565b61025a565b6040516100f895949392919061157b565b60405180910390f35b61011461010f3660046115c9565b610454565b6040516100f897969594939291906115e2565b61013a6101353660046114e6565b6106c6565b6040516100f8949392919061164f565b61015d6101583660046114e6565b6107aa565b6040516100f89392919061167e565b60005461017f906001600160a01b031681565b6040516001600160a01b0390911681526020016100f8565b6101aa6101a53660046116b4565b610910565b6040516100f891906116f1565b6101ca6101c536600461177c565b610a3d565b005b6101d4610ac1565b6040516100f8919061183e565b6101ca6101ef36600461190a565b610d86565b6101ca610202366004611a04565b610e45565b61021a6102153660046116b4565b610ee8565b6040516100f89190611af4565b6101ca610235366004611b77565b611088565b61024d6102483660046116b4565b61110b565b6040516100f89190611c29565b8151602081840181018051600282529282019185019190912091905280548290811061028557600080fd5b9060005260206000209060050201600091509150508060000180546102a990611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546102d590611cd0565b80156103225780601f106102f757610100808354040283529160200191610322565b820191906000526020600020905b81548152906001019060200180831161030557829003601f168201915b50505050509080600101805461033790611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461036390611cd0565b80156103b05780601f10610385576101008083540402835291602001916103b0565b820191906000526020600020905b81548152906001019060200180831161039357829003601f168201915b5050505050908060020180546103c590611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546103f190611cd0565b801561043e5780601f106104135761010080835404028352916020019161043e565b820191906000526020600020905b81548152906001019060200180831161042157829003601f168201915b5050505050908060030154908060040154905085565b6001818154811061046457600080fd5b906000526020600020906007020160009150905080600001805461048790611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546104b390611cd0565b80156105005780601f106104d557610100808354040283529160200191610500565b820191906000526020600020905b8154815290600101906020018083116104e357829003601f168201915b50505050509080600101805461051590611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461054190611cd0565b801561058e5780601f106105635761010080835404028352916020019161058e565b820191906000526020600020905b81548152906001019060200180831161057157829003601f168201915b5050505050908060020180546105a390611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546105cf90611cd0565b801561061c5780601f106105f15761010080835404028352916020019161061c565b820191906000526020600020905b8154815290600101906020018083116105ff57829003601f168201915b50505050509080600301805461063190611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461065d90611cd0565b80156106aa5780601f1061067f576101008083540402835291602001916106aa565b820191906000526020600020905b81548152906001019060200180831161068d57829003601f168201915b5050505050908060040154908060050154908060060154905087565b815160208184018101805160038252928201918501919091209190528054829081106106f157600080fd5b90600052602060002090600402016000915091505080600001805461071590611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461074190611cd0565b801561078e5780601f106107635761010080835404028352916020019161078e565b820191906000526020600020905b81548152906001019060200180831161077157829003601f168201915b5050505050908060010154908060020154908060030154905084565b815160208184018101805160048252928201918501919091209190528054829081106107d557600080fd5b9060005260206000209060030201600091509150508060000180546107f990611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461082590611cd0565b80156108725780601f1061084757610100808354040283529160200191610872565b820191906000526020600020905b81548152906001019060200180831161085557829003601f168201915b50505050509080600101805461088790611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546108b390611cd0565b80156109005780601f106108d557610100808354040283529160200191610900565b820191906000526020600020905b8154815290600101906020018083116108e357829003601f168201915b5050505050908060020154905083565b60606003826040516109229190611d0a565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a32578382906000526020600020906004020160405180608001604052908160008201805461098390611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546109af90611cd0565b80156109fc5780601f106109d1576101008083540402835291602001916109fc565b820191906000526020600020905b8154815290600101906020018083116109df57829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152505081526020019060010190610950565b505050509050919050565b6000546001600160a01b03163314610a5457600080fd5b600482604051610a649190611d0a565b908152604051602091819003820190208054600181018255600091825291902082518392600302909101908190610a9b9082611d75565b5060208201516001820190610ab09082611d75565b506040820151816002015550505050565b60606001805480602002602001604051908101604052809291908181526020016000905b82821015610d7d57838290600052602060002090600702016040518060e0016040529081600082018054610b1890611cd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4490611cd0565b8015610b915780601f10610b6657610100808354040283529160200191610b91565b820191906000526020600020905b815481529060010190602001808311610b7457829003601f168201915b50505050508152602001600182018054610baa90611cd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd690611cd0565b8015610c235780601f10610bf857610100808354040283529160200191610c23565b820191906000526020600020905b815481529060010190602001808311610c0657829003601f168201915b50505050508152602001600282018054610c3c90611cd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610c6890611cd0565b8015610cb55780601f10610c8a57610100808354040283529160200191610cb5565b820191906000526020600020905b815481529060010190602001808311610c9857829003601f168201915b50505050508152602001600382018054610cce90611cd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfa90611cd0565b8015610d475780601f10610d1c57610100808354040283529160200191610d47565b820191906000526020600020905b815481529060010190602001808311610d2a57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152505081526020019060010190610ae5565b50505050905090565b6000546001600160a01b03163314610d9d57600080fd5b600180548082018255600091909152815182916007027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601908190610de29082611d75565b5060208201516001820190610df79082611d75565b5060408201516002820190610e0c9082611d75565b5060608201516003820190610e219082611d75565b506080820151816004015560a0820151816005015560c08201518160060155505050565b6000546001600160a01b03163314610e5c57600080fd5b600282604051610e6c9190611d0a565b908152604051602091819003820190208054600181018255600091825291902082518392600502909101908190610ea39082611d75565b5060208201516001820190610eb89082611d75565b5060408201516002820190610ecd9082611d75565b50606082015181600301556080820151816004015550505050565b6060600482604051610efa9190611d0a565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a325783829060005260206000209060030201604051806060016040529081600082018054610f5b90611cd0565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8790611cd0565b8015610fd45780601f10610fa957610100808354040283529160200191610fd4565b820191906000526020600020905b815481529060010190602001808311610fb757829003601f168201915b50505050508152602001600182018054610fed90611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461101990611cd0565b80156110665780601f1061103b57610100808354040283529160200191611066565b820191906000526020600020905b81548152906001019060200180831161104957829003601f168201915b5050505050815260200160028201548152505081526020019060010190610f28565b6000546001600160a01b0316331461109f57600080fd5b6003826040516110af9190611d0a565b9081526040516020918190038201902080546001810182556000918252919020825183926004029091019081906110e69082611d75565b5060208201518160010155604082015181600201556060820151816003015550505050565b606060028260405161111d9190611d0a565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610a3257838290600052602060002090600502016040518060a001604052908160008201805461117e90611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546111aa90611cd0565b80156111f75780601f106111cc576101008083540402835291602001916111f7565b820191906000526020600020905b8154815290600101906020018083116111da57829003601f168201915b5050505050815260200160018201805461121090611cd0565b80601f016020809104026020016040519081016040528092919081815260200182805461123c90611cd0565b80156112895780601f1061125e57610100808354040283529160200191611289565b820191906000526020600020905b81548152906001019060200180831161126c57829003601f168201915b505050505081526020016002820180546112a290611cd0565b80601f01602080910402602001604051908101604052809291908181526020018280546112ce90611cd0565b801561131b5780601f106112f05761010080835404028352916020019161131b565b820191906000526020600020905b8154815290600101906020018083116112fe57829003601f168201915b50505050508152602001600382015481526020016004820154815250508152602001906001019061114b565b61138c828260405160240161135d929190611e35565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611390565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156113ea576113ea6113b1565b60405290565b60405160e0810167ffffffffffffffff811182821017156113ea576113ea6113b1565b60405160a0810167ffffffffffffffff811182821017156113ea576113ea6113b1565b6040516080810167ffffffffffffffff811182821017156113ea576113ea6113b1565b600082601f83011261146a57600080fd5b813567ffffffffffffffff80821115611485576114856113b1565b604051601f8301601f19908116603f011681019082821181831017156114ad576114ad6113b1565b816040528381528660208588010111156114c657600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156114f957600080fd5b823567ffffffffffffffff81111561151057600080fd5b61151c85828601611459565b95602094909401359450505050565b60005b8381101561154657818101518382015260200161152e565b50506000910152565b6000815180845261156781602086016020860161152b565b601f01601f19169290920160200192915050565b60a08152600061158e60a083018861154f565b82810360208401526115a0818861154f565b905082810360408401526115b4818761154f565b60608401959095525050608001529392505050565b6000602082840312156115db57600080fd5b5035919050565b60e0815260006115f560e083018a61154f565b8281036020840152611607818a61154f565b9050828103604084015261161b818961154f565b9050828103606084015261162f818861154f565b6080840196909652505060a081019290925260c090910152949350505050565b608081526000611662608083018761154f565b6020830195909552506040810192909252606090910152919050565b606081526000611691606083018661154f565b82810360208401526116a3818661154f565b915050826040830152949350505050565b6000602082840312156116c657600080fd5b813567ffffffffffffffff8111156116dd57600080fd5b6116e984828501611459565b949350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561176e57603f1989840301855281516080815181865261173e8287018261154f565b838b0151878c0152898401518a880152606093840151939096019290925250509386019390860190600101611718565b509098975050505050505050565b6000806040838503121561178f57600080fd5b823567ffffffffffffffff808211156117a757600080fd5b6117b386838701611459565b935060208501359150808211156117c957600080fd5b90840190606082870312156117dd57600080fd5b6117e56113c7565b8235828111156117f457600080fd5b61180088828601611459565b82525060208301358281111561181557600080fd5b61182188828601611459565b602083015250604083013560408201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561176e57603f19898403018552815160e0815181865261188b8287018261154f565b915050888201518582038a8701526118a3828261154f565b91505087820151858203898701526118bb828261154f565b915050606080830151868303828801526118d5838261154f565b6080858101519089015260a0808601519089015260c09485015194909701939093525050509386019390860190600101611865565b60006020828403121561191c57600080fd5b813567ffffffffffffffff8082111561193457600080fd5b9083019060e0828603121561194857600080fd5b6119506113f0565b82358281111561195f57600080fd5b61196b87828601611459565b82525060208301358281111561198057600080fd5b61198c87828601611459565b6020830152506040830135828111156119a457600080fd5b6119b087828601611459565b6040830152506060830135828111156119c857600080fd5b6119d487828601611459565b6060830152506080830135608082015260a083013560a082015260c083013560c082015280935050505092915050565b60008060408385031215611a1757600080fd5b823567ffffffffffffffff80821115611a2f57600080fd5b611a3b86838701611459565b93506020850135915080821115611a5157600080fd5b9084019060a08287031215611a6557600080fd5b611a6d611413565b823582811115611a7c57600080fd5b611a8888828601611459565b825250602083013582811115611a9d57600080fd5b611aa988828601611459565b602083015250604083013582811115611ac157600080fd5b611acd88828601611459565b60408301525060608301356060820152608083013560808201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561176e57603f19898403018552815160608151818652611b418287018261154f565b915050888201518582038a870152611b59828261154f565b92890151958901959095525094870194925090860190600101611b1b565b60008060408385031215611b8a57600080fd5b823567ffffffffffffffff80821115611ba257600080fd5b611bae86838701611459565b93506020850135915080821115611bc457600080fd5b9084019060808287031215611bd857600080fd5b611be0611436565b823582811115611bef57600080fd5b611bfb88828601611459565b8252506020830135602082015260408301356040820152606083013560608201528093505050509250929050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561176e57603f19898403018552815160a08151818652611c768287018261154f565b915050888201518582038a870152611c8e828261154f565b9150508782015185820389870152611ca6828261154f565b60608481015190880152608093840151939096019290925250509386019390860190600101611c50565b600181811c90821680611ce457607f821691505b602082108103611d0457634e487b7160e01b600052602260045260246000fd5b50919050565b60008251611d1c81846020870161152b565b9190910192915050565b601f821115611d7057600081815260208120601f850160051c81016020861015611d4d5750805b601f850160051c820191505b81811015611d6c57828155600101611d59565b5050505b505050565b815167ffffffffffffffff811115611d8f57611d8f6113b1565b611da381611d9d8454611cd0565b84611d26565b602080601f831160018114611dd85760008415611dc05750858301515b600019600386901b1c1916600185901b178555611d6c565b600085815260208120601f198616915b82811015611e0757888601518255948401946001909101908401611de8565b5085821015611e255787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b604081526000611e48604083018561154f565b905060018060a01b0383166020830152939250505056fea26469706673582212208bdfa2ff6e8b5fa1cee0139111d1584c32847ab0a940d367483d99b6330e397264736f6c63430008110033",
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
