package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InserSpawning(spawning *__models.Spawning) {
	db, errDB := sql.Open("sqlite3", "./spawning_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO spawning(SpawningId, Nik, Date, Amount, FishType) VALUES (?, ?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(spawning.SpawningId, spawning.Nik, spawning.Date, spawning.Amount, spawning.FishType)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetSpawnings(spawningid string) []__models.Spawning {
	db, errDB := sql.Open("sqlite3", "./spawning_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM spawning")
	if err != nil {
		panic(err)
	}

	var spawningScan __models.Spawning
	var spawnings []__models.Spawning

	for rows.Next() {
		err = rows.Scan(&spawningScan.SpawningId, &spawningScan.Nik, &spawningScan.Date, &spawningScan.Amount, &spawningScan.FishType)

		if err != nil {
			panic(err)
		}

		spawnings = append(spawnings, spawningScan)
	}

	defer rows.Close()

	defer db.Close()

	return spawnings
}
