package repositories

import (
	"database/sql"
	__transactionContract "financingsupplychain/api/transactioncontract"
	__interface "financingsupplychain/interfaces"
	__model "financingsupplychain/models"
	"fmt"
	"log"

	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var funderlock = &sync.Mutex{}

type fundersingleton struct {
}

var funderInstance *fundersingleton

func getFunderInstance() *fundersingleton {
	if funderInstance == nil {
		funderlock.Lock()
		defer funderlock.Unlock()
		if funderInstance == nil {
			funderInstance = &fundersingleton{}
		}
	}
	return funderInstance
}

func InsertFundingBlockchain(fundid string, username string, numofponds int, amountoffund int) {
	_, err := __interface.TransactionsContractInterface().SetFunderApproverTransaction(__interface.GetTransactionContractAuth(), fundid, __transactionContract.TransactionContractFunderApproverTransaction{
		Funder:        username,
		Timestamp:     time.Now().String(),
		Numberofponds: big.NewInt(int64(numofponds)),
		Amountoffund:  big.NewInt(int64(amountoffund)),
	}) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}
}

func GetFunderBlockChain(fundid string) []__transactionContract.TransactionContractFunderApproverTransaction {
	reply, err := __interface.TransactionsContractInterface().GetFunderApproverTransaction(&bind.CallOpts{}, fundid) // conn call the balance function of deployed smart contract
	// conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply
}

func InsertFunder(funder *__model.Funder) {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	fmt.Println("Fundid Funder : ", funder.FundId)

	records := `INSERT INTO funder(FundId, Nik, SubmittedBy, SubmittedTimestamp,FundedBy, FundedTimestamp,  FishType, NumberOfPonds, AmountOfFund) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(funder.FundId, funder.Nik, funder.SubmittedBy, funder.SubmittedTimestamp, funder.FundedBy, funder.FundedTimestamp, funder.FishType, funder.NumberOfPonds, funder.AmountOfFund)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetFundersBySales(username string) []__model.Funder {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT FundId,Nik,FishType,NumberOfPonds,AmountOfFund, SubmittedBy FROM funder")
	if err != nil {
		panic(err)
	}

	var scanner __model.Funder

	var funders []__model.Funder

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.FishType, &scanner.NumberOfPonds, &scanner.AmountOfFund, &scanner.SubmittedBy)

		if err != nil {
			panic(err)
		}

		if scanner.SubmittedBy == username {
			funders = append(funders, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return funders
}

func GetFunders() []__model.Funder {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT FundId,Nik,FishType,NumberOfPonds,AmountOfFund FROM funder")
	if err != nil {
		panic(err)
	}

	var scanner __model.Funder

	var funders []__model.Funder

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.FishType, &scanner.NumberOfPonds, &scanner.AmountOfFund)

		if err != nil {
			panic(err)
		}

		funders = append(funders, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return funders
}

func GetFunderByFunders() []__model.Funder {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT FundId,Nik,ProposedBy, SubmittedTimestamp, FishType,NumberOfPonds,AmountOfFund FROM funder")
	if err != nil {
		panic(err)
	}

	var scanner __model.Funder

	var funders []__model.Funder

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.ProposedBy, &scanner.SubmittedTimestamp, &scanner.FishType, &scanner.NumberOfPonds, &scanner.AmountOfFund)

		if err != nil {
			panic(err)
		}

		if len(scanner.ProposedBy) != 0 {
			funders = append(funders, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return funders
}

func GetFunderByNik(nik string) []__model.Funder {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT FundId,Nik,SubmittedBy, SubmittedTimestamp, FishType,NumberOfPonds,AmountOfFund FROM funder")
	if err != nil {
		panic(err)
	}

	var scanner __model.Funder

	var funders []__model.Funder

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.SubmittedBy, &scanner.SubmittedTimestamp, &scanner.FishType, &scanner.NumberOfPonds, &scanner.AmountOfFund)

		if err != nil {
			panic(err)
		}

		if scanner.Nik == nik {
			funders = append(funders, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return funders
}

func UploadFileFundId(fileurl string, fundid string, username string) {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	stmt, errStmt := db.Prepare("update funder set FileUrl=?, ProposedBy=? where FundId=?")
	if errStmt != nil {
		panic(errStmt)
	}

	_, err := stmt.Exec(fileurl, username, fundid)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func InsertFund(amountoffund int, fundid string) {
	db, errDB := sql.Open("sqlite3", "./funder_db.db")
	if errDB != nil {
		panic(errDB)
	}

	stmt, errStmt := db.Prepare("update funder set AmountOfFund=? where FundId=?")
	if errStmt != nil {
		panic(errStmt)
	}

	_, err := stmt.Exec(amountoffund, fundid)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
