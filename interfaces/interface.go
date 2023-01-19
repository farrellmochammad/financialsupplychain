package interfaces

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	__creditscoreContract "financingsupplychain/api/creditscorecontract"
	__transactionsContract "financingsupplychain/api/transactioncontract"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
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
			GanacheIp := readGanacheIp()
			client, err = ethclient.Dial(GanacheIp)
			if err != nil {
				panic(err)
			}

			CreditKey := readCreditPrivateKey()
			// create auth and transaction package for deploying smart contract
			creditContractAuth = getAccountAuth(client, CreditKey)

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

			TransactionKey := readTransactionPrivateKey()
			// create auth and transaction package for deploying smart contract
			transactionContractAuth = getAccountAuth(client, TransactionKey)

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

			data := Environment{
				GanacheHost:           GanacheIp,
				CreditPrivateKey:      CreditKey,
				TransactionPrivateKey: TransactionKey,
				CreditKey:             deployedCreditContract.Hex(),
				TransactionKey:        deployedTransactionContract.Hex(),
			}

			writeToJSON(data)

		}
	}
	return singleInstance
}

func InitConnectionTransaction() *single {
	if singleInstanceTransaction == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstanceTransaction == nil {
			client, err = ethclient.Dial(readGanacheIp())
			if err != nil {
				panic(err)
			}

			hexaddress := readTransactionKey()
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
			client, err = ethclient.Dial(readGanacheIp())
			if err != nil {
				panic(err)
			}

			// create connection object to connect through are binary go file and deployed contract with help of address
			hexaddress := readCreditKey()
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

	return getAccountAuth(client, readCreditPrivateKey())
}

func GetTransactionContractAuth() *bind.TransactOpts {

	return getAccountAuth(client, readTransactionPrivateKey())
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

func writeToJSON(env Environment) {
	file, _ := json.MarshalIndent(env, "", " ")

	_ = ioutil.WriteFile("env.json", file, 0644)
}

func readCreditKey() string {
	// Open our jsonFile
	jsonFile, err := os.Open("env.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result["CreditKey"])

}

func readGanacheIp() string {
	// Open our jsonFile
	jsonFile, err := os.Open("env.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result["GanacheHost"])

}

func readTransactionKey() string {
	// Open our jsonFile
	jsonFile, err := os.Open("env.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result["TransactionKey"])
}

func readCreditPrivateKey() string {
	// Open our jsonFile
	jsonFile, err := os.Open("env.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result["CreditPrivateKey"])
}

func readTransactionPrivateKey() string {
	// Open our jsonFile
	jsonFile, err := os.Open("env.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result["TransactionPrivateKey"])
}

type Environment struct {
	GanacheHost           string
	CreditPrivateKey      string
	TransactionPrivateKey string
	CreditKey             string
	TransactionKey        string
}
