package repositories

import (
	"database/sql"
	__models "financingsupplychain/models"
	"log"
)

func InsertUsers(user *__models.UserRegister) {
	db, errDB := sql.Open("sqlite3", "./username_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO users(Username, Password, Role) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec(user.Username, user.Password, user.Role)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func GetUsers(username string) __models.UserLogin {
	db, errDB := sql.Open("sqlite3", "./username_db.db")
	if errDB != nil {
		panic(errDB)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var userScan __models.UserLogin

	for rows.Next() {
		err = rows.Scan(&userScan.Username, &userScan.Password, &userScan.Role)

		if err != nil {
			panic(err)
		}

		if userScan.Username == username {
			defer rows.Close()

			defer db.Close()

			return userScan
		}

	}

	defer rows.Close()

	defer db.Close()

	return userScan
}
