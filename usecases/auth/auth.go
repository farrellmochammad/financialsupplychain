package auth

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "myname" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != "myname123" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func Register(c echo.Context) error {

	db, errDB := sql.Open("sqlite3", "./users_db.db")
	if errDB != nil {
		panic(errDB)
	}

	// users_table := `CREATE TABLE users (
	//     id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	//     "FirstName" TEXT,
	//     "LastName" TEXT,
	//     "Dept" TEXT,
	//     "Salary" INT);`
	// query, err := db.Prepare(users_table)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// query.Exec()
	// fmt.Println("Table created successfully!")

	records := `INSERT INTO users(FirstName, LastName, Dept, Salary) VALUES (?, ?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	FirstName := "a"
	LastName := "b"
	Dept := "231"
	Salary := 3123
	_, err = query.Exec(FirstName, LastName, Dept, Salary)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": "dwad",
	})
}
