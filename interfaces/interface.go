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
var singleInstanceTransaction *single
var singleInstanceCreditScore *single
var client *ethclient.Client
var deployedCreditContract common.Address
var deployedTransactionContract common.Address
var connCredit *__creditscoreContract.Api
var connTransaction *__transactionsContract.Api
var creditContractAuth *bind.TransactOpts
var transactionContractAuth *bind.TransactOpts
var err error

func DeployTransaction() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Init start connection that clean data")
			singleInstance = &single{}
			client, err = ethclient.Dial("http://172.27.224.1:7545")
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
			fmt.Println("Value Credit Score : ", deployedCreditContract.Hex())

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

			fmt.Println("Value Transaction : ", deployedTransactionContract.Hex())

			// create connection object to connect through are binary go file and deployed contract with help of address
			connTransaction, err = __transactionsContract.NewApi(common.HexToAddress(deployedTransactionContract.Hex()), client)
			if err != nil {
				panic(err)
			}

		}
	}
	return singleInstance
}

func InitConnectionTransaction() *single {
	if singleInstanceTransaction == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceTransaction == nil {
			client, err = ethclient.Dial("http://172.27.224.1:7545")
			if err != nil {
				panic(err)
			}

			hexaddress := "0x0fd5f1b0c91fA9C313AE8B5cE78B7B27A5b87f12"
			// create connection object to connect through are binary go file and deployed contract with help of address
			connTransaction, err = __transactionsContract.NewApi(common.HexToAddress(hexaddress), client)
			if err != nil {
				panic(err)
			}
		}
	}
	return singleInstanceTransaction
}

func InitConnectionCreditScore() *single {
	if singleInstanceCreditScore == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceCreditScore == nil {
			client, err = ethclient.Dial("http://172.27.224.1:7545")
			if err != nil {
				panic(err)
			}

			// create connection object to connect through are binary go file and deployed contract with help of address
			hexaddress := "0x7B9Ed6088909a1e4eCdB79F2caE3E3Bd8eC1dfeF"
			connCredit, err = __creditscoreContract.NewApi(common.HexToAddress(hexaddress), client)
			if err != nil {
				panic(err)
			}
		}
	}
	return singleInstanceCreditScore
}

func CreditScoreContractInterface() *__creditscoreContract.Api {
	InitConnectionCreditScore()

	return connCredit
}

func TransactionsContractInterface() *__transactionsContract.Api {
	InitConnectionTransaction()

	return connTransaction
}

func GetCreditContractAuth() *bind.TransactOpts {

	return getAccountAuth(client, "e7085bed4fc19d2729d839f259b06c02a203f686e2c8044aaa77104f20ecb58b")
}

func GetTransactionContractAuth() *bind.TransactOpts {

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
