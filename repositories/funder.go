package repositories

import (
	"database/sql"
	__transactionContract "financingsupplychain/api/transactioncontract"
	__interface "financingsupplychain/interfaces"
	__model "financingsupplychain/models"
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

func InsertFundingBlockchain(nik string, fundid string, username string, numofponds int, amountoffund int) {
	_, err := __interface.TransactionsContractInterface().SetFunderApproverTransaction(__interface.GetTransactionContractAuth(), nik, __transactionContract.TransactionContractFunderApproverTransaction{
		Fundid:        fundid,
		Funder:        username,
		Timestamp:     time.Now().String(),
		Numberofponds: big.NewInt(int64(numofponds)),
		Amountoffund:  big.NewInt(int64(amountoffund)),
	}) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}
}

func GetFunderBlockChain(nik string) []__transactionContract.TransactionContractFunderApproverTransaction {
	reply, err := __interface.TransactionsContractInterface().GetFunderApproverTransaction(&bind.CallOpts{}, nik) // conn call the balance function of deployed smart contract
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

	records := `INSERT INTO funder(FundId, Nik, FishType, NumberOfPonds, AmountOfFund) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(funder.FundId, funder.Nik, funder.FishType, funder.NumberOfPonds, funder.AmountOfFund)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
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
