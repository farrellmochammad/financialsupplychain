package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

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

	rows, err := db.Query("SELECT * FROM experience")
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

func GetExperience(nik string) __models.Experience {
	db, errDB := sql.Open("sqlite3", "./experience_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM experience")
	if err != nil {
		panic(err)
	}

	var scanner __models.Experience
	var experience __models.Experience

	for rows.Next() {
		err = rows.Scan(&scanner.Nik, &scanner.Name, &scanner.Phone, &scanner.Dob, &scanner.Address, &scanner.StartFarming, &scanner.FishType, &scanner.NumberOfPonds, &scanner.Notes, &scanner.CurrentStatus)

		if err != nil {
			panic(err)
		}

		if nik == scanner.Nik {
			experience = scanner
			defer rows.Close()

			defer db.Close()

			return experience
		}

	}

	defer rows.Close()

	defer db.Close()

	return experience
}
