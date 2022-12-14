package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InsertPond(pond *__models.Pond) {
	db, errDB := sql.Open("sqlite3", "./pond_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO pond(PondId, FundId, TotalSpawning, FishType) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(pond.PondId, pond.FundId, pond.TotalSpawning, pond.FishType)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetPonds() []__models.Pond {
	db, errDB := sql.Open("sqlite3", "./pond_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM pond")
	if err != nil {
		panic(err)
	}

	var scanner __models.Pond

	var ponds []__models.Pond

	for rows.Next() {
		err = rows.Scan(&scanner.PondId, &scanner.FundId, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		ponds = append(ponds, scanner)

	}

	defer rows.Close()

	defer db.Close()

	return ponds
}

func GetPond(pondid string) []__models.Pond {
	db, errDB := sql.Open("sqlite3", "./pond_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM pond")
	if err != nil {
		panic(err)
	}

	var scanner __models.Pond
	var ponds []__models.Pond

	for rows.Next() {
		err = rows.Scan(&scanner.PondId, &scanner.FundId, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		if pondid == scanner.PondId {
			ponds = append(ponds, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return ponds
}

func GetPondByFundId(fundid string) []__models.Pond {
	db, errDB := sql.Open("sqlite3", "./pond_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM pond")
	if err != nil {
		panic(err)
	}

	var scanner __models.Pond
	var ponds []__models.Pond

	for rows.Next() {
		err = rows.Scan(&scanner.PondId, &scanner.FundId, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		if fundid == scanner.FundId {
			ponds = append(ponds, scanner)
		}

	}

	defer rows.Close()

	defer db.Close()

	return ponds
}
