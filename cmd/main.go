package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"financingsupplychain/api"
	__approverContract "financingsupplychain/api/approvercontract"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
)

func main() {
	// address of etherum env
	client, err := ethclient.Dial("http://172.22.112.1:7545")
	if err != nil {
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	auth := getAccountAuth(client, "b5eec637df269767164fe51874f85ac452813098efab8443ac02b939329dad5b")

	//deploying smart contract
	deployedContractAddress, _, _, err := api.DeployApi(auth, client) //api is redirected from api directory from our contract go file
	if err != nil {
		panic(err)
	}

	// create auth and transaction package for deploying smart contract
	approverContractAuth := getAccountAuth(client, "e2c9d72246d20ae6934f37f4b14b8ef60dddda1e784ddb46c4fca4f594d03525")

	//deploying smart contract
	deployedApproverContract, _, _, err := __approverContract.DeployApi(approverContractAuth, client, "31d6cfe0d16ae931b73c59d7e0c089c0", "032f75b3ca02a393196a818328bd32e8") //api is redirected from api directory from our contract go file
	if err != nil {
		panic(err)
	}

	fmt.Println("Deployed contract address : ", deployedContractAddress.Hex())   // print deployed contract address
	fmt.Println("Deployed contract approver : ", deployedApproverContract.Hex()) // print deployed contract address

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// create connection object to connect through are binary go file and deployed contract with help of address
	conn, err := api.NewApi(common.HexToAddress(deployedContractAddress.Hex()), client)
	if err != nil {
		panic(err)
	}

	connApprover, errApprover := __approverContract.NewApi(common.HexToAddress(deployedApproverContract.Hex()), client)
	if errApprover != nil {
		panic(errApprover)
	}

	e.GET("/status", func(c echo.Context) error {
		// usecase
		reply, err := connApprover.GetStatus(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
		if err != nil {
			return err
		}
		fmt.Println("/status")
		return c.JSON(http.StatusOK, reply)
	})

	e.PUT("/status/inreview", func(c echo.Context) error {
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		auth := getAccountAuth(client, v["accountPrivateKey"].(string))

		reply, err := connApprover.SetStatus(auth, 1)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.PUT("/status/reject", func(c echo.Context) error {
		//gets address of account by which amount to be deposite
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		//creating auth object for above account
		auth := getAccountAuth(client, v["accountPrivateKey"].(string))

		reply, err := connApprover.SetStatus(auth, 2)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.PUT("/status/approved", func(c echo.Context) error {
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		auth := getAccountAuth(client, v["accountPrivateKey"].(string))

		reply, err := connApprover.SetStatus(auth, 3)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/transactions", func(c echo.Context) error {
		// usecase
		reply, err := connApprover.GetApprovalTransactions(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
		if err != nil {
			return err
		}
		fmt.Println("/transactions")
		return c.JSON(http.StatusOK, reply)
	})

	e.POST("/transaction", func(c echo.Context) error {
		//gets address of account by which amount to be deposite
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		//creating auth object for above account
		auth := getAccountAuth(client, v["accountPrivateKey"].(string))

		reply, err := connApprover.AddApprovalTransaction(auth, v["from"].(string), v["to"].(string), uint8(v["status"].(float64)), v["filepath"].(string))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	e.GET("/balance", func(c echo.Context) error {
		// usecase
		reply, err := conn.Balance(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
		if err != nil {
			return err
		}
		fmt.Println("/balance")
		return c.JSON(http.StatusOK, reply)
	})

	e.POST("/deposite/:amount", func(c echo.Context) error {
		amount := c.Param("amount")
		amt, _ := strconv.Atoi(amount)

		fmt.Println("/deposite : ", amount, amt)

		//gets address of account by which amount to be deposite
		var v map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			panic(err)
		}

		fmt.Println("/v : ", v)

		//creating auth object for above account
		auth := getAccountAuth(client, v["accountPrivateKey"].(string))

		reply, err := conn.Deposite(auth, big.NewInt(int64(amt)))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	// accountAddressHex := hex.EncodeToString([]byte(accountAddress))
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
