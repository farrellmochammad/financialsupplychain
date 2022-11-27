package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InsertSpawningHistory(spawning *__models.Spawning) {
	db, errDB := sql.Open("sqlite3", "./spawning_history.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO spawning_history(Nik, Date, Amount, FishType) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(spawning.Nik, spawning.Date, spawning.Amount, spawning.FishType)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetSpawningHistories() []__models.Spawning {
	db, errDB := sql.Open("sqlite3", "./spawning_history.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM spawning_history")
	if err != nil {
		panic(err)
	}

	var spawningScan __models.Spawning
	var spawnings []__models.Spawning

	for rows.Next() {
		err = rows.Scan(&spawningScan.Nik, &spawningScan.Date, &spawningScan.Amount, &spawningScan.FishType)

		if err != nil {
			panic(err)
		}

		spawnings = append(spawnings, spawningScan)
	}

	defer rows.Close()

	defer db.Close()

	return spawnings
}

func GetSpawningHistory(nik string) []__models.Spawning {
	db, errDB := sql.Open("sqlite3", "./spawning_history.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM spawning_history")
	if err != nil {
		panic(err)
	}

	var spawningScan __models.Spawning
	var spawningResult []__models.Spawning

	for rows.Next() {
		err = rows.Scan(&spawningScan.Nik, &spawningScan.Date, &spawningScan.Amount, &spawningScan.FishType)

		if err != nil {
			panic(err)
		}

		if nik == spawningScan.Nik {
			spawningResult = append(spawningResult, spawningScan)
		}

	}

	defer rows.Close()

	defer db.Close()

	return spawningResult
}

func GetSpawningHistoryAmounts(nik string) []int {
	db, errDB := sql.Open("sqlite3", "./spawning_history.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM spawning_history")
	if err != nil {
		panic(err)
	}

	var spawningScan __models.Spawning
	var spawnings []int

	for rows.Next() {
		err = rows.Scan(&spawningScan.Nik, &spawningScan.Date, &spawningScan.Amount, &spawningScan.FishType)

		if err != nil {
			panic(err)
		}

		if nik == spawningScan.Nik {
			spawnings = append(spawnings, spawningScan.Amount)
		}
	}

	defer rows.Close()

	defer db.Close()

	return spawnings
}
