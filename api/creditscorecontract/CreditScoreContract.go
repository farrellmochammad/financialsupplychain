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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spawningaverage\",\"type\":\"uint256\"}],\"name\":\"addSpawningAverage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit[]\",\"name\":\"_credits\",\"type\":\"tuple[]\"}],\"name\":\"creditHistoryScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditTransactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit[]\",\"name\":\"_credits\",\"type\":\"tuple[]\"}],\"name\":\"creditValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_yoe\",\"type\":\"uint256\"}],\"name\":\"experienceScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_durationofexperience\",\"type\":\"uint256\"}],\"name\":\"experienceValidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_yoe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nop\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit[]\",\"name\":\"_credits\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"_spawnings\",\"type\":\"tuple[]\"}],\"name\":\"generateCreditScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getPond\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Pond[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"}],\"name\":\"getSpawning\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpawningCurrentAverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nop\",\"type\":\"uint256\"}],\"name\":\"pondScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pondlength\",\"type\":\"uint256\"}],\"name\":\"pondValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pondsInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creditAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structCreditScoreContract.Credit\",\"name\":\"_credit\",\"type\":\"tuple\"}],\"name\":\"setCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fishtype\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Pond\",\"name\":\"_pond\",\"type\":\"tuple\"}],\"name\":\"setPond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning\",\"name\":\"_spawning\",\"type\":\"tuple\"}],\"name\":\"setSpawning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningAverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"_spawnings\",\"type\":\"tuple[]\"}],\"name\":\"spawningScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"internalType\":\"structCreditScoreContract.Spawning[]\",\"name\":\"_spawnings\",\"type\":\"tuple[]\"}],\"name\":\"spawningValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spawningsInformation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpawning\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"spawningDate\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200005e6040518060400160405280601b81526020017f4f776e657220636f6e7472616374206465706c6f7965642062793a0000000000815250336200007660201b62000eac1760201c565b600080546001600160a01b031916331790556200014a565b620000c382826040516024016200008f929190620000e8565b60408051601f198184030181529190526020810180516001600160e01b0390811663319af33360e01b17909152620000c716565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b604081526000835180604084015260005b81811015620001185760208187018101516060868401015201620000f9565b50600060608285018101919091526001600160a01b03949094166020840152601f01601f191690910190910192915050565b611847806200015a6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806398a81c21116100b8578063bfbda90d1161007c578063bfbda90d146102fc578063cbb37a1d1461030f578063e62741601461032f578063e7e577e414610350578063f7522cba14610370578063fc421ae01461038357600080fd5b806398a81c211461028f5780639f532bcc146102a2578063a620da68146102b5578063bace5b30146102d6578063bb280b12146102e957600080fd5b8063314bd9641161010a578063314bd964146101ae5780633ffc2c93146101d157806361a23d9b146101f9578063693a4d371461020e5780638da5cb5b1461025157806391d23d441461027c57600080fd5b8063082c95b6146101475780632131f27c1461016d57806321956c8c1461018057806326e119f51461019357806326e87955146101a6575b600080fd5b61015a610155366004610f16565b610396565b6040519081526020015b60405180910390f35b61015a61017b366004610f16565b6103b7565b61015a61018e366004610f16565b61041e565b61015a6101a1366004610f16565b61049b565b61015a6104fa565b6101c16101bc36600461110f565b61052e565b6040519015158152602001610164565b6101e46101df36600461114b565b610608565b60408051928352901515602083015201610164565b61020c61020736600461118f565b610656565b005b61020c61021c366004610f16565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155565b600054610264906001600160a01b031681565b6040516001600160a01b039091168152602001610164565b6101c161028a3660046112e8565b6106b9565b6101c161029d366004610f16565b610707565b61015a6102b036600461131c565b610722565b6102c86102c336600461114b565b61076f565b6040516101649291906113e2565b61015a6102e436600461110f565b6108cf565b61015a6102f73660046112e8565b610986565b61020c61030a366004611407565b610a75565b61032261031d36600461146a565b610ac6565b604051610164919061149e565b61034261033d36600461114b565b610bdf565b604051610164929190611512565b61036361035e36600461146a565b610c34565b604051610164919061152b565b61015a61037e36600461146a565b610dca565b61020c6103913660046115a2565b610e56565b600481815481106103a657600080fd5b600091825260209091200154905081565b600060028210156103ca57506000919050565b600282101580156103db5750600582105b156103e857506019919050565b600582101580156103f95750600782105b1561040657506026919050565b6007821061041657506035919050565b506000919050565b6000600282101561043157506000919050565b600282101580156104425750600482105b1561044f57506024919050565b600482101580156104605750600682105b1561046d57506035919050565b6006821015801561047e5750600982105b1561048b57506046919050565b600a82106104165750605e919050565b600060028210156104ae57506000919050565b600282101580156104bf5750600582105b156104cc57506011919050565b600582101580156104dd5750600882105b156104ea57506021919050565b600882106104165750602d919050565b600480546000919061050e90600190611606565b8154811061051e5761051e61161f565b9060005260206000200154905090565b600080805b83518110156105795783818151811061054e5761054e61161f565b602002602001015160000151826105659190611635565b91508061057181611648565b915050610533565b5082516105869082611661565b600480549192509061059a90600190611606565b815481106105aa576105aa61161f565b906000526020600020015481106105c45750600192915050565b600480546105d490600190611606565b815481106105e4576105e461161f565b90600052602060002001548110156105ff5750600092915050565b50600192915050565b8151602081840181018051600182529282019185019190912091905280548290811061063357600080fd5b60009182526020909120600290910201805460019091015490925060ff16905082565b6002826040516106669190611683565b90815260405160209181900382019020805460018101825560009182529190208251839260020290910190819061069d9082611728565b50602082015160018201906106b29082611728565b5050505050565b60006001815b8351811015610700578381815181106106da576106da61161f565b6020026020010151602001516106ee579015905b806106f881611648565b9150506106bf565b5092915050565b6000600282101561071a57506000919050565b506001919050565b600061072d826108cf565b61073684610986565b61073f8661041e565b610748886103b7565b6107529190611635565b61075c9190611635565b6107669190611635565b95945050505050565b8151602081840181018051600282529282019185019190912091905280548290811061079a57600080fd5b9060005260206000209060020201600091509150508060000180546107be9061169f565b80601f01602080910402602001604051908101604052809291908181526020018280546107ea9061169f565b80156108375780601f1061080c57610100808354040283529160200191610837565b820191906000526020600020905b81548152906001019060200180831161081a57829003601f168201915b50505050509080600101805461084c9061169f565b80601f01602080910402602001604051908101604052809291908181526020018280546108789061169f565b80156108c55780601f1061089a576101008083540402835291602001916108c5565b820191906000526020600020905b8154815290600101906020018083116108a857829003601f168201915b5050505050905082565b600080805b835181101561091a578381815181106108ef576108ef61161f565b602002602001015160000151826109069190611635565b91508061091281611648565b9150506108d4565b50600a81101561092d5750600f92915050565b6064811015801561093e5750609681105b1561094c5750606492915050565b6096811015801561095d575060c881105b1561096b575060e192915050565b60c8811061097d575061015e92915050565b50600092915050565b600080805b83518110156109f4578381815181106109a6576109a661161f565b6020026020010151602001516109e2578381815181106109c8576109c861161f565b602002602001015160000151826109df9190611635565b91505b806109ec81611648565b91505061098b565b5080600003610a065750600092915050565b60018110158015610a175750600381105b15610a255750600f92915050565b60038110158015610a365750600581105b15610a445750601b92915050565b60058110158015610a555750600881105b15610a635750602b92915050565b600881111561097d5750603c92915050565b600382604051610a859190611683565b908152604051602091819003820190208054600180820183556000928352918390208451600290920201908155918301518392918201906106b29082611728565b6060600382604051610ad89190611683565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610bd4578382906000526020600020906002020160405180604001604052908160008201548152602001600182018054610b439061169f565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6f9061169f565b8015610bbc5780601f10610b9157610100808354040283529160200191610bbc565b820191906000526020600020905b815481529060010190602001808311610b9f57829003601f168201915b50505050508152505081526020019060010190610b06565b505050509050919050565b81516020818401810180516003825292820191850191909120919052805482908110610c0a57600080fd5b90600052602060002090600202016000915091505080600001549080600101805461084c9061169f565b6060600282604051610c469190611683565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b82821015610bd45783829060005260206000209060020201604051806040016040529081600082018054610ca79061169f565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd39061169f565b8015610d205780601f10610cf557610100808354040283529160200191610d20565b820191906000526020600020905b815481529060010190602001808311610d0357829003601f168201915b50505050508152602001600182018054610d399061169f565b80601f0160208091040260200160405190810160405280929190818152602001828054610d659061169f565b8015610db25780601f10610d8757610100808354040283529160200191610db2565b820191906000526020600020905b815481529060010190602001808311610d9557829003601f168201915b50505050508152505081526020019060010190610c74565b600080805b600184604051610ddf9190611683565b9081526040519081900360200190205481101561070057600184604051610e069190611683565b90815260200160405180910390208181548110610e2557610e2561161f565b90600052602060002090600202016000015482610e429190611635565b915080610e4e81611648565b915050610dcf565b600182604051610e669190611683565b90815260405160209181900382019020805460018082018355600092835291839020845160029092020190815592909101519101805460ff191691151591909117905550565b610ef18282604051602401610ec29291906117e7565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052610ef5565b5050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b600060208284031215610f2857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610f6757610f67610f2f565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610f9557610f95610f2f565b604052919050565b60006001600160401b03821115610fb657610fb6610f2f565b5060051b60200190565b600082601f830112610fd157600080fd5b81356001600160401b03811115610fea57610fea610f2f565b610ffd601f8201601f1916602001610f6d565b81815284602083860101111561101257600080fd5b816020850160208301376000918101602001919091529392505050565b60006040828403121561104157600080fd5b611049610f45565b90508135815260208201356001600160401b0381111561106857600080fd5b61107484828501610fc0565b60208301525092915050565b600082601f83011261109157600080fd5b813560206110a66110a183610f9d565b610f6d565b82815260059290921b840181019181810190868411156110c557600080fd5b8286015b848110156111045780356001600160401b038111156110e85760008081fd5b6110f68986838b010161102f565b8452509183019183016110c9565b509695505050505050565b60006020828403121561112157600080fd5b81356001600160401b0381111561113757600080fd5b61114384828501611080565b949350505050565b6000806040838503121561115e57600080fd5b82356001600160401b0381111561117457600080fd5b61118085828601610fc0565b95602094909401359450505050565b600080604083850312156111a257600080fd5b82356001600160401b03808211156111b957600080fd5b6111c586838701610fc0565b935060208501359150808211156111db57600080fd5b90840190604082870312156111ef57600080fd5b6111f7610f45565b82358281111561120657600080fd5b61121288828601610fc0565b82525060208301358281111561122757600080fd5b61123388828601610fc0565b6020830152508093505050509250929050565b60006040828403121561125857600080fd5b611260610f45565b9050813581526020820135801515811461127957600080fd5b602082015292915050565b600082601f83011261129557600080fd5b813560206112a56110a183610f9d565b82815260069290921b840181019181810190868411156112c457600080fd5b8286015b84811015611104576112da8882611246565b8352918301916040016112c8565b6000602082840312156112fa57600080fd5b81356001600160401b0381111561131057600080fd5b61114384828501611284565b6000806000806080858703121561133257600080fd5b843593506020850135925060408501356001600160401b038082111561135757600080fd5b61136388838901611284565b9350606087013591508082111561137957600080fd5b5061138687828801611080565b91505092959194509250565b60005b838110156113ad578181015183820152602001611395565b50506000910152565b600081518084526113ce816020860160208601611392565b601f01601f19169290920160200192915050565b6040815260006113f560408301856113b6565b828103602084015261076681856113b6565b6000806040838503121561141a57600080fd5b82356001600160401b038082111561143157600080fd5b61143d86838701610fc0565b9350602085013591508082111561145357600080fd5b506114608582860161102f565b9150509250929050565b60006020828403121561147c57600080fd5b81356001600160401b0381111561149257600080fd5b61114384828501610fc0565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561150457888303603f190185528151805184528701518784018790526114f1878501826113b6565b95880195935050908601906001016114c5565b509098975050505050505050565b82815260406020820152600061114360408301846113b6565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561150457888303603f1901855281518051878552611576888601826113b6565b91890151858303868b015291905061158e81836113b6565b968901969450505090860190600101611552565b600080606083850312156115b557600080fd5b82356001600160401b038111156115cb57600080fd5b6115d785828601610fc0565b9250506115e78460208501611246565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611619576116196115f0565b92915050565b634e487b7160e01b600052603260045260246000fd5b80820180821115611619576116196115f0565b60006001820161165a5761165a6115f0565b5060010190565b60008261167e57634e487b7160e01b600052601260045260246000fd5b500490565b60008251611695818460208701611392565b9190910192915050565b600181811c908216806116b357607f821691505b6020821081036116d357634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561172357600081815260208120601f850160051c810160208610156117005750805b601f850160051c820191505b8181101561171f5782815560010161170c565b5050505b505050565b81516001600160401b0381111561174157611741610f2f565b6117558161174f845461169f565b846116d9565b602080601f83116001811461178a57600084156117725750858301515b600019600386901b1c1916600185901b17855561171f565b600085815260208120601f198616915b828110156117b95788860151825594840194600190910190840161179a565b50858210156117d75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6040815260006117fa60408301856113b6565b905060018060a01b0383166020830152939250505056fea26469706673582212202bc31b9aef46fd2d87309e7d98c7fe71ff87b86b022b4624bd09d069e13b670d64736f6c63430008110033",
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

// CreditHistoryScore is a free data retrieval call binding the contract method 0xbb280b12.
//
// Solidity: function creditHistoryScore((uint256,bool)[] _credits) pure returns(uint256)
func (_Api *ApiCaller) CreditHistoryScore(opts *bind.CallOpts, _credits []CreditScoreContractCredit) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "creditHistoryScore", _credits)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditHistoryScore is a free data retrieval call binding the contract method 0xbb280b12.
//
// Solidity: function creditHistoryScore((uint256,bool)[] _credits) pure returns(uint256)
func (_Api *ApiSession) CreditHistoryScore(_credits []CreditScoreContractCredit) (*big.Int, error) {
	return _Api.Contract.CreditHistoryScore(&_Api.CallOpts, _credits)
}

// CreditHistoryScore is a free data retrieval call binding the contract method 0xbb280b12.
//
// Solidity: function creditHistoryScore((uint256,bool)[] _credits) pure returns(uint256)
func (_Api *ApiCallerSession) CreditHistoryScore(_credits []CreditScoreContractCredit) (*big.Int, error) {
	return _Api.Contract.CreditHistoryScore(&_Api.CallOpts, _credits)
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

// ExperienceScore is a free data retrieval call binding the contract method 0x2131f27c.
//
// Solidity: function experienceScore(uint256 _yoe) pure returns(uint256)
func (_Api *ApiCaller) ExperienceScore(opts *bind.CallOpts, _yoe *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "experienceScore", _yoe)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExperienceScore is a free data retrieval call binding the contract method 0x2131f27c.
//
// Solidity: function experienceScore(uint256 _yoe) pure returns(uint256)
func (_Api *ApiSession) ExperienceScore(_yoe *big.Int) (*big.Int, error) {
	return _Api.Contract.ExperienceScore(&_Api.CallOpts, _yoe)
}

// ExperienceScore is a free data retrieval call binding the contract method 0x2131f27c.
//
// Solidity: function experienceScore(uint256 _yoe) pure returns(uint256)
func (_Api *ApiCallerSession) ExperienceScore(_yoe *big.Int) (*big.Int, error) {
	return _Api.Contract.ExperienceScore(&_Api.CallOpts, _yoe)
}

// ExperienceValidation is a free data retrieval call binding the contract method 0x26e119f5.
//
// Solidity: function experienceValidation(uint256 _durationofexperience) pure returns(uint256)
func (_Api *ApiCaller) ExperienceValidation(opts *bind.CallOpts, _durationofexperience *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "experienceValidation", _durationofexperience)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExperienceValidation is a free data retrieval call binding the contract method 0x26e119f5.
//
// Solidity: function experienceValidation(uint256 _durationofexperience) pure returns(uint256)
func (_Api *ApiSession) ExperienceValidation(_durationofexperience *big.Int) (*big.Int, error) {
	return _Api.Contract.ExperienceValidation(&_Api.CallOpts, _durationofexperience)
}

// ExperienceValidation is a free data retrieval call binding the contract method 0x26e119f5.
//
// Solidity: function experienceValidation(uint256 _durationofexperience) pure returns(uint256)
func (_Api *ApiCallerSession) ExperienceValidation(_durationofexperience *big.Int) (*big.Int, error) {
	return _Api.Contract.ExperienceValidation(&_Api.CallOpts, _durationofexperience)
}

// GenerateCreditScore is a free data retrieval call binding the contract method 0x9f532bcc.
//
// Solidity: function generateCreditScore(uint256 _yoe, uint256 _nop, (uint256,bool)[] _credits, (uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiCaller) GenerateCreditScore(opts *bind.CallOpts, _yoe *big.Int, _nop *big.Int, _credits []CreditScoreContractCredit, _spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "generateCreditScore", _yoe, _nop, _credits, _spawnings)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateCreditScore is a free data retrieval call binding the contract method 0x9f532bcc.
//
// Solidity: function generateCreditScore(uint256 _yoe, uint256 _nop, (uint256,bool)[] _credits, (uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiSession) GenerateCreditScore(_yoe *big.Int, _nop *big.Int, _credits []CreditScoreContractCredit, _spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	return _Api.Contract.GenerateCreditScore(&_Api.CallOpts, _yoe, _nop, _credits, _spawnings)
}

// GenerateCreditScore is a free data retrieval call binding the contract method 0x9f532bcc.
//
// Solidity: function generateCreditScore(uint256 _yoe, uint256 _nop, (uint256,bool)[] _credits, (uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiCallerSession) GenerateCreditScore(_yoe *big.Int, _nop *big.Int, _credits []CreditScoreContractCredit, _spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	return _Api.Contract.GenerateCreditScore(&_Api.CallOpts, _yoe, _nop, _credits, _spawnings)
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

// PondScore is a free data retrieval call binding the contract method 0x21956c8c.
//
// Solidity: function pondScore(uint256 _nop) pure returns(uint256)
func (_Api *ApiCaller) PondScore(opts *bind.CallOpts, _nop *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "pondScore", _nop)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PondScore is a free data retrieval call binding the contract method 0x21956c8c.
//
// Solidity: function pondScore(uint256 _nop) pure returns(uint256)
func (_Api *ApiSession) PondScore(_nop *big.Int) (*big.Int, error) {
	return _Api.Contract.PondScore(&_Api.CallOpts, _nop)
}

// PondScore is a free data retrieval call binding the contract method 0x21956c8c.
//
// Solidity: function pondScore(uint256 _nop) pure returns(uint256)
func (_Api *ApiCallerSession) PondScore(_nop *big.Int) (*big.Int, error) {
	return _Api.Contract.PondScore(&_Api.CallOpts, _nop)
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

// SpawningScore is a free data retrieval call binding the contract method 0xbace5b30.
//
// Solidity: function spawningScore((uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiCaller) SpawningScore(opts *bind.CallOpts, _spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "spawningScore", _spawnings)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpawningScore is a free data retrieval call binding the contract method 0xbace5b30.
//
// Solidity: function spawningScore((uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiSession) SpawningScore(_spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	return _Api.Contract.SpawningScore(&_Api.CallOpts, _spawnings)
}

// SpawningScore is a free data retrieval call binding the contract method 0xbace5b30.
//
// Solidity: function spawningScore((uint256,string)[] _spawnings) pure returns(uint256)
func (_Api *ApiCallerSession) SpawningScore(_spawnings []CreditScoreContractSpawning) (*big.Int, error) {
	return _Api.Contract.SpawningScore(&_Api.CallOpts, _spawnings)
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
