package auth

import (
	"database/sql"
	"fmt"
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
	var user UserLogin
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     "some role",
	})

	hmacSampleSecret := []byte("123156")
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func Register(c echo.Context) error {

	db, errDB := sql.Open("sqlite3", "./username_db.db")
	if errDB != nil {
		panic(errDB)
	}

	records := `INSERT INTO users(Username, Password, Role) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}

	user := new(UserRegister)
	if err = c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	_, err = query.Exec(user.Username, user.Password, user.Role)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Login Success",
	})
}

func OnlyTest(c echo.Context) error {
	fmt.Println("only test", c.Get("username"))
	fmt.Println("only test", c.Get("role"))
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Login Success",
	})
}

type UserRegister struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Role     string `json:"role" form:"role" query:"role"`
}

type UserLogin struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Role     string `json:"role" form:"role" query:"role"`
}
