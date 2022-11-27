package repositories

import (
	"crypto/sha1"
	"database/sql"
	__creditContract "financingsupplychain/api/creditscorecontract"
	__interface "financingsupplychain/interfaces"
	__models "financingsupplychain/models"
	"fmt"
	"log"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GenerateCreditScoreBlockChain(yoe int, nop int, credits []__models.Credit, spawnings []__models.Spawning) *big.Int {
	var credits_object []__creditContract.CreditScoreContractCredit

	for i := 0; i < len(credits); i++ {
		credit := __creditContract.CreditScoreContractCredit{
			CreditAmount: big.NewInt(int64(credits[i].CreditAmount)),
			Status:       credits[i].Status,
		}
		credits_object = append(credits_object, credit)
	}

	var spawning_object []__creditContract.CreditScoreContractSpawning

	for i := 0; i < len(spawnings); i++ {
		spawning := __creditContract.CreditScoreContractSpawning{
			TotalSpawning: big.NewInt(int64(spawnings[i].TotalSpawning)),
			SpawningDate:  spawnings[i].SpawningDate,
		}
		spawning_object = append(spawning_object, spawning)
	}

	reply, err := __interface.CreditScoreContractInterface().GenerateCreditScore(&bind.CallOpts{}, big.NewInt(int64(yoe)), big.NewInt(int64(nop)), credits_object, spawning_object) // conn call the balance function of deployed smart contract
	if err != nil {
		panic(err)
	}

	return reply
}

func InsertCredit(credit *__models.Credit) {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO credit(CreditId, Username, Nik, CreditAmount, Status) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(NewSHA1Hash(), credit.Username, credit.Nik, credit.CreditAmount, credit.Status)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetCredits() []__models.Credit {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM credit")
	if err != nil {
		panic(err)
	}

	var scanner __models.Credit

	var credits []__models.Credit

	for rows.Next() {
		err = rows.Scan(&scanner.CreditId, &scanner.Username, &scanner.Nik, &scanner.CreditAmount, &scanner.Status)

		if err != nil {
			panic(err)
		}

		credits = append(credits, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return credits
}

func GetCredit(nik string) []__models.Credit {
	db, errDB := sql.Open("sqlite3", "./credit_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM credit")
	if err != nil {
		panic(err)
	}

	var scanner __models.Credit
	var credit []__models.Credit

	for rows.Next() {
		err = rows.Scan(&scanner.CreditId, &scanner.Username, &scanner.Nik, &scanner.CreditAmount, &scanner.Status)

		if err != nil {
			panic(err)
		}

		if nik == scanner.Nik {
			credit = append(credit, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return credit
}

// NewSHA1Hash generates a new SHA1 hash based on
// a random number of characters.
func NewSHA1Hash(n ...int) string {
	noRandomCharacters := 32

	if len(n) > 0 {
		noRandomCharacters = n[0]
	}

	randString := RandomString(noRandomCharacters)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
