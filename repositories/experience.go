package repositories

import (
	"database/sql"
	__creditContract "financingsupplychain/api/creditscorecontract"
	__transactionContract "financingsupplychain/api/transactioncontract"
	__interface "financingsupplychain/interfaces"
	__models "financingsupplychain/models"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		}
	}
	return singleInstance
}

func NumOfPondsValidation(experience *__models.Experience) bool {
	reply, err := __interface.CreditScoreContractInterface().PondValidation(&bind.CallOpts{}, big.NewInt(int64(experience.NumberOfPonds))) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply

}

func CreditsHistoryValidation(credits []__models.Credit) bool {
	var credits_object []__creditContract.CreditScoreContractCredit

	for i := 0; i < len(credits); i++ {
		credit := __creditContract.CreditScoreContractCredit{
			CreditAmount: big.NewInt(int64(credits[i].CreditAmount)),
			Status:       credits[i].Status,
		}
		credits_object = append(credits_object, credit)
	}

	reply, err := __interface.CreditScoreContractInterface().CreditValidation(&bind.CallOpts{}, credits_object) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply

}

func AddSpawningAverage(average int) {

	_, err := __interface.CreditScoreContractInterface().AddSpawningAverage(__interface.GetCreditContractAuth(), big.NewInt(int64(average))) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

}

func GetCurrentAverage() *big.Int {
	reply, err := __interface.CreditScoreContractInterface().GetSpawningCurrentAverage(&bind.CallOpts{}) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply
}

func InsertExperienceBlockChain(username string, nik string, numofponds int, spawningaverage int, creditscore *big.Int) {
	_, err := __interface.TransactionsContractInterface().SetApproverTransaction(__interface.GetTransactionContractAuth(), nik, __transactionContract.TransactionContractApproverTransaction{
		Submitby:        username,
		Status:          "Approved",
		Timestamp:       time.Now().String(),
		Numofponds:      big.NewInt(int64(numofponds)),
		Spawningaverage: big.NewInt(int64(spawningaverage)),
		Creditscore:     creditscore,
	}) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}
}

func GetExperienceBlockchain(nik string) []__transactionContract.TransactionContractApproverTransaction {
	reply, err := __interface.TransactionsContractInterface().GetApproverTransaction(&bind.CallOpts{}, nik) // conn call the balance function of deployed smart contract
	// conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply
}

func UploadFile(fileurl string, nik string) {
	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	stmt, errStmt := db.Prepare("update experience set UrlFile=? where Nik=?")
	if errStmt != nil {
		panic(errDB)
	}

	_, err := stmt.Exec(fileurl, nik)

	if err != nil {
		log.Fatal(err)
	}
}

func InsertExperience(experience *__models.Experience) {
	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO experience(Nik, Name, Phone, Dob, Address, StartFarming, FishType, NumberOfPonds, Notes, CurrentStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(experience.Nik, experience.Name, experience.Phone, experience.Dob, experience.Address, experience.StartFarming, experience.FishType, experience.NumberOfPonds, experience.Notes, "Draft")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetExperiences() []__models.Experience {
	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT Nik,Name,Phone,Dob,Address,StartFarming,FishType,NumberOfPonds,Notes,CurrentStatus FROM experience")
	if err != nil {
		panic(err)
	}

	var scanner __models.Experience

	var experiences []__models.Experience

	for rows.Next() {
		err = rows.Scan(&scanner.Nik, &scanner.Name, &scanner.Phone, &scanner.Dob, &scanner.Address, &scanner.StartFarming, &scanner.FishType, &scanner.NumberOfPonds, &scanner.Notes, &scanner.CurrentStatus)

		if err != nil {
			panic(err)
		}

		experiences = append(experiences, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return experiences
}

func GetExperience(nik string) []__models.Experience {
	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("select Nik,Name,Phone,Dob,Address,StartFarming,FishType,NumberOfPonds,Notes,CurrentStatus FROM experience")
	if err != nil {
		panic(err)
	}

	var scanner __models.Experience
	var experiences []__models.Experience

	for rows.Next() {
		err = rows.Scan(&scanner.Nik, &scanner.Name, &scanner.Phone, &scanner.Dob, &scanner.Address, &scanner.StartFarming, &scanner.FishType, &scanner.NumberOfPonds, &scanner.Notes, &scanner.CurrentStatus)

		if err != nil {
			panic(err)
		}

		if nik == scanner.Nik {
			experiences = append(experiences, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return experiences
}
