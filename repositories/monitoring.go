package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

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

func GetMonitoring(nik string) __models.Monitoring {
	db, errDB := sql.Open("sqlite3", "./monitoring_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM monitoring")
	if err != nil {
		panic(err)
	}

	var scanner __models.Monitoring
	var monitoring __models.Monitoring

	for rows.Next() {
		err = rows.Scan(&scanner.FundId, &scanner.Nik, &scanner.Name, &scanner.TotalSpawning, &scanner.FishType)

		if err != nil {
			panic(err)
		}

		if nik == scanner.Nik {
			monitoring = scanner

			defer rows.Close()

			defer db.Close()

			return monitoring
		}

	}

	defer rows.Close()

	defer db.Close()

	return monitoring
}
