package repositories

import (
	"database/sql"
	__transactionContract "financingsupplychain/api/transactioncontract"
	__interface "financingsupplychain/interfaces"
	__models "financingsupplychain/models"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func InsertMonitoringPondBlockChain(fundid string, monitoring *__models.MonitoringPond) {
	fmt.Println("Fund Id : ", fundid)
	_, err := __interface.TransactionsContractInterface().SetMonitoring(__interface.GetTransactionContractAuth(), fundid, __transactionContract.TransactionContractMonitoring{
		Timestamp:   time.Now().String(),
		Weight:      big.NewInt(int64(monitoring.Weight)),
		Temperature: big.NewInt(int64(monitoring.Temperature)),
		Humidity:    big.NewInt(int64(monitoring.Humidity)),
	}) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}
}

func GetMonitoringBlockChain(fundid string) []__transactionContract.TransactionContractMonitoring {
	fmt.Println("Fund Id  2: ", fundid)
	reply, err := __interface.TransactionsContractInterface().GetMonitoring(&bind.CallOpts{}, fundid) // conn call the balance function of deployed smart contract
	// conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply
}

func InsertMonitoring(monitoring *__models.Monitoring) {
	db, errDB := sql.Open("sqlite3", "./monitoring_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO monitoring(FundId, Nik, Name, TotalSpawning, FishType) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(monitoring.FundId, monitoring.Nik, monitoring.Name, monitoring.TotalSpawning, monitoring.FishType)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetMonitorings() []__models.Monitoring {
	db, errDB := sql.Open("sqlite3", "./monitoring_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM monitoring")
	if err != nil {
		panic(err)
	}

	var scanner __models.Monitoring

	var monitorings []__models.Monitoring

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.Name, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		monitorings = append(monitorings, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return monitorings
}

func GetMonitoring(fundid string) []__models.Monitoring {
	db, errDB := sql.Open("sqlite3", "./monitoring_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM monitoring")
	if err != nil {
		panic(err)
	}

	var scanner __models.Monitoring
	var monitorings []__models.Monitoring

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.Name, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		if fundid == scanner.FundId {
			monitorings = append(monitorings, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return monitorings
}
