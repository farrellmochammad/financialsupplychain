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

const creditKey = "d90636a5a7961f380e91a0f001d0bab43a252b809e2491075c6c7a3140554e94"
const transactionKey = "06adc980a667275e656f1419ecb455624c5d50a0b8feef5feda5e93d4f3b6651"
const creditKeyValue = "767c9BdAD2604F48996Ef9ADFe090DF1942D3e9A"
const transactionKeyValue = "C1F0480Fda79336eD8C4F08b8beCd36A5F9Ed4bB"

var err error

func DeployTransaction() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Init start connection that clean data")
			singleInstance = &single{}
			client, err = ethclient.Dial("http://127.0.0.1:8545")
			if err != nil {
				panic(err)
			}

			// create auth and transaction package for deploying smart contract
			creditContractAuth = getAccountAuth(client, creditKey)

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
			transactionContractAuth = getAccountAuth(client, transactionKey)

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
			client, err = ethclient.Dial("http://127.0.0.1:8545")
			if err != nil {
				panic(err)
			}

			hexaddress := transactionKeyValue
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
			client, err = ethclient.Dial("http://127.0.0.1:8545")
			if err != nil {
				panic(err)
			}

			// create connection object to connect through are binary go file and deployed contract with help of address
			hexaddress := creditKeyValue
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

	return getAccountAuth(client, creditKey)
}

func GetTransactionContractAuth() *bind.TransactOpts {

	return getAccountAuth(client, transactionKey)
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
