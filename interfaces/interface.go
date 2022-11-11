package interfaces

import (
	"context"
	"crypto/ecdsa"
	__creditscoreContract "financingsupplychain/api/creditscorecontract"
	__transactionsContract "financingsupplychain/api/transactioncontract"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single
var client *ethclient.Client
var deployedCreditContract common.Address
var deployedTransactionContract common.Address
var connCredit *__creditscoreContract.Api
var connTransaction *__transactionsContract.Api
var creditContractAuth *bind.TransactOpts
var transactionContractAuth *bind.TransactOpts
var err error

func StartConnection() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
			client, err = ethclient.Dial("http://172.17.224.1:7545")
			if err != nil {
				panic(err)
			}

			// create auth and transaction package for deploying smart contract
			creditContractAuth = getAccountAuth(client, "e7085bed4fc19d2729d839f259b06c02a203f686e2c8044aaa77104f20ecb58b")

			//deploying smart contract
			deployedCreditContract, _, _, err = __creditscoreContract.DeployApi(creditContractAuth, client) //api is redirected from api directory from our contract go file
			if err != nil {
				panic(err)
			}

			// create connection object to connect through are binary go file and deployed contract with help of address
			connCredit, err = __creditscoreContract.NewApi(common.HexToAddress(deployedCreditContract.Hex()), client)
			if err != nil {
				panic(err)
			}

			// create auth and transaction package for deploying smart contract
			transactionContractAuth = getAccountAuth(client, "71723c973dfa7e0e73443705fe3c5cde3e53ee1e889c437ad4a06ab9b78031dc")

			//deploying smart contract
			deployedTransactionContract, _, _, err = __transactionsContract.DeployApi(transactionContractAuth, client) //api is redirected from api directory from our contract go file
			if err != nil {
				panic(err)
			}

			// create connection object to connect through are binary go file and deployed contract with help of address
			connTransaction, err = __transactionsContract.NewApi(common.HexToAddress(deployedTransactionContract.Hex()), client)
			if err != nil {
				panic(err)
			}

		}
	}
	return singleInstance
}

func CreditScoreContractInterface() *__creditscoreContract.Api {
	StartConnection()

	return connCredit
}

func TransactionsContractInterface() *__transactionsContract.Api {
	StartConnection()

	return connTransaction
}

func GetCreditContractAuth() *bind.TransactOpts {

	return getAccountAuth(client, "e7085bed4fc19d2729d839f259b06c02a203f686e2c8044aaa77104f20ecb58b")
}

func GetTransactionContractAuth() *bind.TransactOpts {
	StartConnection()

	return getAccountAuth(client, "71723c973dfa7e0e73443705fe3c5cde3e53ee1e889c437ad4a06ab9b78031dc")
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
