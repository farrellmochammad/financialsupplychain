package main

import (
	__interface "financingsupplychain/interfaces"

	__authUsecase "financingsupplychain/usecases/auth"
	__CreditUsecase "financingsupplychain/usecases/credit"
	__ExperienceUsecase "financingsupplychain/usecases/experience"
	__MonitoringUsecase "financingsupplychain/usecases/monitoring"
	__PondUsecase "financingsupplychain/usecases/pond"
	__SpawningUsecase "financingsupplychain/usecases/spawning"
	__SpawningHistoryUsecase "financingsupplychain/usecases/spawning_history"

	__middleware "financingsupplychain/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	__interface.StartConnection()
	// address of etherum env
	// client, err := ethclient.Dial("http://172.17.224.1:7545")
	// if err != nil {
	// 	panic(err)
	// }

	// // create auth and transaction package for deploying smart contract
	// approverContractAuth := getAccountAuth(client, "707f9140265ae2467d77f908aa501195d1f816def00100c5dceaae4133097840")

	// //deploying smart contract
	// deployedApproverContract, _, _, err := __approverContract.DeployApi(approverContractAuth, client, "31d6cfe0d16ae931b73c59d7e0c089c0", "032f75b3ca02a393196a818328bd32e8") //api is redirected from api directory from our contract go file
	// if err != nil {
	// 	panic(err)
	// }

	// // create auth and transaction package for deploying smart contract
	// // funderApproverContractAuth := getAccountAuth(client, "e994a0ff4b95cf88b6a611f32ba6c64f8ef4d23ab5c19f3f15c484a42b17be85")

	// // //deploying smart contract
	// // // deployedFunderApproverContract, _, _, err := __funderApproverContract.DeployApi(funderApproverContractAuth, client, "31d6cfe0d16ae931b73c59d7e0c089c0", "032f75b3ca02a393196a818328bd32e8") //api is redirected from api directory from our contract go file
	// // // if err != nil {
	// // // 	panic(err)
	// // // }

	// // // create auth and transaction package for deploying smart contract
	// // monitoringContractAuth := getAccountAuth(client, "e994a0ff4b95cf88b6a611f32ba6c64f8ef4d23ab5c19f3f15c484a42b17be85")

	// // //deploying smart contract
	// // deployedMonitoringContract, _, _, err := __monitoringContract.DeployApi(monitoringContractAuth, client, "31d6cfe0d16ae931b73c59d7e0c089c0", "032f75b3ca02a393196a818328bd32e8", "Agus pod", "Jawa Barat", "Cianjur", "Mujair") //api is redirected from api directory from our contract go file
	// // if err != nil {
	// // 	panic(err)
	// // }

	// // create auth and transaction package for deploying smart contract
	// agreementContractAuth := getAccountAuth(client, "16e11d70f301a1fa7594e9b5e9c8496f41b34b3bad0762720ea360bcdb979b89")

	// //deploying smart contract
	// deployedAgreementContract, _, _, err := __agreementContract.DeployApi(agreementContractAuth, client, "31d6cfe0d16ae931b73c59d7e0c089c0") //api is redirected from api directory from our contract go file
	// if err != nil {
	// 	panic(err)
	// }

	// // create auth and transaction package for deploying smart contract
	// creditContractAuth := getAccountAuth(client, "e7085bed4fc19d2729d839f259b06c02a203f686e2c8044aaa77104f20ecb58b")

	// //deploying smart contract
	// delpoyedCreditContract, _, _, err := __creditscoreContract.DeployApi(creditContractAuth, client) //api is redirected from api directory from our contract go file
	// if err != nil {
	// 	panic(err)
	// }

	// // create auth and transaction package for deploying smart contract
	// transactionContractAuth := getAccountAuth(client, "71723c973dfa7e0e73443705fe3c5cde3e53ee1e889c437ad4a06ab9b78031dc")

	// //deploying smart contract
	// deployedTransactionContract, _, _, err := __transactionsContract.DeployApi(transactionContractAuth, client) //api is redirected from api directory from our contract go file
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Deployed contract approver : ", deployedApproverContract.Hex())              // print deployed contract address
	// fmt.Println("Deployed contract funder approver : ", deployedFunderApproverContract.Hex()) // print deployed contract address

	e := echo.New()

	e.Use(middleware.CORS())

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// create connection object to connect through are binary go file and deployed contract with help of address
	// connApprover, errApprover := __approverContract.NewApi(common.HexToAddress(deployedApproverContract.Hex()), client)
	// if errApprover != nil {
	// 	panic(errApprover)
	// }

	// // // create connection object to connect through are binary go file and deployed contract with help of address
	// // connFunderApprover, errFunderApprover := __funderApproverContract.NewApi(common.HexToAddress(deployedFunderApproverContract.Hex()), client)
	// // if errFunderApprover != nil {
	// // 	panic(errApprover)
	// // }

	// // // create connection object to connect through are binary go file and deployed contract with help of address
	// // connMonitoring, errMonitoring := __monitoringContract.NewApi(common.HexToAddress(deployedMonitoringContract.Hex()), client)
	// // if errMonitoring != nil {
	// // 	panic(errMonitoring)
	// // }

	// // create connection object to connect through are binary go file and deployed contract with help of address
	// connAgreement, errAgreement := __agreementContract.NewApi(common.HexToAddress(deployedAgreementContract.Hex()), client)
	// if errAgreement != nil {
	// 	panic(errAgreement)
	// }

	// // create connection object to connect through are binary go file and deployed contract with help of address
	// connCredit, errCredit := __creditscoreContract.NewApi(common.HexToAddress(delpoyedCreditContract.Hex()), client)
	// if errCredit != nil {
	// 	panic(errCredit)
	// }

	// // create connection object to connect through are binary go file and deployed contract with help of address
	// connTransaction, errTransaction := __transactionsContract.NewApi(common.HexToAddress(deployedTransactionContract.Hex()), client)
	// if errTransaction != nil {
	// 	panic(errCredit)
	// }

	// e.GET("/funders", func(c echo.Context) error {
	// 	// usecase
	// 	reply, err := connFunderApprover.GetFunderApprovalTransactions(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("/funders")
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// type PayloadModel struct {
	// 	PrivateKey  string
	// 	Status      int
	// 	Filepath    string
	// 	timestamp   string
	// 	ModalGiven  int
	// 	CreditScore int
	// }

	// e.POST("/funder", func(c echo.Context) error {
	// 	//gets address of account by which amount to be deposite

	// 	var v PayloadModel
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	//creating auth object for above account
	// 	auth := getAccountAuth(client, v.PrivateKey)

	// 	reply, err := connFunderApprover.AddFunderApprovalTransaction(auth, uint8(v.Status), v.Filepath, v.timestamp, big.NewInt(int64(v.ModalGiven)), big.NewInt(int64(v.CreditScore)))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// e.GET("/monitorings", func(c echo.Context) error {
	// 	// usecase
	// 	reply, err := connMonitoring.GetMonitorings(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("/monitorings")
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// type PayloadMonitoringModel struct {
	// 	PrivateKey  string
	// 	Weight      int
	// 	Timestamp   string
	// 	Temperature int
	// 	Humidity    int
	// }

	// e.POST("/monitoring", func(c echo.Context) error {
	// 	//gets address of account by which amount to be deposite

	// 	var v PayloadMonitoringModel
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	//creating auth object for above account
	// 	auth := getAccountAuth(client, v.PrivateKey)

	// 	reply, err := connMonitoring.AddMonitoring(auth, v.Timestamp, big.NewInt(int64(v.Weight)), big.NewInt(int64(v.Temperature)), big.NewInt(int64(v.Humidity)))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// type PayloadSpwaningModel struct {
	// 	PrivateKey string
	// 	Weight     int
	// 	Timestamp  string
	// }

	e.POST("/login", __middleware.GenerateJWT(__authUsecase.Login))
	e.POST("/register", __authUsecase.Register)
	e.GET("/Testjwt", __middleware.ValidateJWT(__authUsecase.OnlyTest))
	e.POST("/spawning", __middleware.ValidateJWT(__SpawningUsecase.InsertSpawning))
	e.GET("/spawnings", __middleware.ValidateJWT(__SpawningUsecase.GetSpawnings))
	e.GET("/spawning/:spawningid", __middleware.ValidateJWT(__SpawningUsecase.GetSpawning))

	e.POST("/experience", __middleware.ValidateJWT(__ExperienceUsecase.InsertExperience))
	e.GET("/experiences", __middleware.ValidateJWT(__ExperienceUsecase.GetExperiences))
	e.GET("/experience/:nik", __middleware.ValidateJWT(__ExperienceUsecase.GetExperience))
	e.PUT("/uploadfile", __middleware.ValidateJWT(__ExperienceUsecase.UploadFile))

	e.POST("/monitoring", __middleware.ValidateJWT(__MonitoringUsecase.InsertMonitoring))
	e.GET("/monitorings", __middleware.ValidateJWT(__MonitoringUsecase.GetMonitorings))
	e.GET("/monitoring/:fundid", __middleware.ValidateJWT(__MonitoringUsecase.GetMonitoring))

	e.POST("/pond", __middleware.ValidateJWT(__PondUsecase.InsertPond))
	e.GET("/ponds", __middleware.ValidateJWT(__PondUsecase.GetPonds))
	e.GET("/pond_id/:pondid", __middleware.ValidateJWT(__PondUsecase.GetPond))
	e.GET("/pond_fundid/:fundid", __middleware.ValidateJWT(__PondUsecase.GetPondByNik))

	e.POST("/credit", __middleware.ValidateJWT(__CreditUsecase.InsertCredit))
	e.GET("/credits", __middleware.ValidateJWT(__CreditUsecase.GetCredits))
	e.GET("/credit/:creditid", __middleware.ValidateJWT(__CreditUsecase.GetCredit))

	e.POST("/spawning_history", __middleware.ValidateJWT(__SpawningHistoryUsecase.InsertSpawningHistory))
	e.GET("/spawning_histories", __middleware.ValidateJWT(__SpawningHistoryUsecase.GetSpawningsHistories))
	e.GET("/spawning_history/:nik", __middleware.ValidateJWT(__SpawningHistoryUsecase.GetSpawningHistory))

	// e.POST("/monitoring/pond/:pondid", __middleware.ValidateJWT(func(c echo.Context) error {
	// 	var v map[string]interface{}
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	weight, weightStat := new(big.Int).SetString(v["weight"].(string), 0)
	// 	if !weightStat {
	// 		panic(err)
	// 	}

	// 	temperature, temperatureStat := new(big.Int).SetString(v["temperature"].(string), 0)
	// 	if !temperatureStat {
	// 		panic(err)
	// 	}

	// 	humidity, humidityStat := new(big.Int).SetString(v["humidity"].(string), 0)
	// 	if !humidityStat {
	// 		panic(err)
	// 	}

	// 	auth := getAccountAuth(client, "71723c973dfa7e0e73443705fe3c5cde3e53ee1e889c437ad4a06ab9b78031dc")

	// 	fmt.Println("Weight " + weight.String())
	// 	fmt.Println("Temperature #1" + temperature.String())
	// 	fmt.Println("Humidity #1" + humidity.String())

	// 	reply, err := connTransaction.SetMonitoring(auth, c.Param(":pondid"), __transactionsContract.TransactionContractMonitoring{
	// 		Timestamp:   time.Now().String(),
	// 		Weight:      weight,
	// 		Temperature: temperature,
	// 		Humidity:    humidity,
	// 	})

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return err
	// 	}

	// 	fmt.Println("Testing #1")

	// 	fmt.Println("/agreements")
	// 	return c.JSON(http.StatusOK, reply)
	// }),
	// )

	// e.GET("/monitoring/pond/:pondid", __middleware.ValidateJWT(func(c echo.Context) error {

	// 	reply, err := connTransaction.GetMonitoring(&bind.CallOpts{}, c.Param(":pondid"))

	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"data": reply,
	// 	})
	// }),
	// )

	// e.POST("/creditcontract", func(c echo.Context) error {
	// 	var v map[string]interface{}
	// 	err := json.NewDecoder(c.Request().Body).Decode(&v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	auth := getAccountAuth(client, v["accountPrivateKey"].(string))

	// 	var creditAmount, ok = new(big.Int).SetString(v["creditAmount"].(string), 0)
	// 	if !ok {
	// 		panic(err)
	// 	}

	// 	// usecase
	// 	reply, err := connCredit.SetCredit(auth, v["nik"].(string), __creditscoreContract.CreditScoreContractCredit{
	// 		CreditAmount: creditAmount,
	// 		Status:       v["status"].(bool),
	// 	}) // conn call the balance function of deployed smart contract
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("/agreements")
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// e.GET("/creditcontract/credit/:nik", func(c echo.Context) error {
	// 	nik := c.Param("nik")

	// 	// usecase
	// 	reply, err := connCredit.GetCredit(&bind.CallOpts{}, nik) // conn call the balance function of deployed smart contract
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("/agreements")
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// e.GET("/creditcontract/:nik", func(c echo.Context) error {

	// 	// usecase
	// 	reply, err := connCredit.GetCredit(&bind.CallOpts{}, c.Param(":nik")) // conn call the balance function of deployed smart contract
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("/spawnings")
	// 	return c.JSON(http.StatusOK, reply)
	// })

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:2021"))
}
