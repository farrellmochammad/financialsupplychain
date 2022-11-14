package auth

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var user __model.UserLogin
	err := c.Bind(&user)

	userQuery := __repository.GetUsers(user.Username)

	if user.Username != userQuery.Username {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Username not exist",
		})
	}

	if user.Password != userQuery.Password {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Wrong password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userQuery.Username,
		"role":     userQuery.Role,
	})

	hmacSampleSecret := []byte("123156")
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token":      tokenString,
		"permission": userQuery.Role,
		"username":   userQuery.Username,
	})
}

func Register(c echo.Context) error {

	user := new(__model.UserRegister)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertUsers(user)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Register Success",
	})
}

func OnlyTest(c echo.Context) error {
	fmt.Println("only test", c.Get("username"))
	fmt.Println("only test", c.Get("role"))
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Login Success",
	})
}
