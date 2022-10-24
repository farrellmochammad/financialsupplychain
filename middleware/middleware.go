package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		hmacSampleSecret := []byte("123156")
		tokenString, err := token.SignedString(hmacSampleSecret)

		fmt.Println(tokenString, err)
		return next(c)
	}
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer: "
		if !strings.Contains(token, b) {
			return c.JSON(401, map[string]interface{}{"status": "failed", "message": "login failed, token not valid"})
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			return c.JSON(401, map[string]interface{}{"status": "failed", "message": "Authroization not valid"})
		}

		// Validate token
		hmacSampleSecret := "123156"
		valid, err := validateToken(t[1], hmacSampleSecret)
		if err != nil {
			return c.JSON(401, map[string]interface{}{"status": "failed", "message": "Token expired, please login back"})
		}

		if valid.Claims.(jwt.MapClaims)["role"] != "admin" {
			return c.JSON(401, map[string]interface{}{"status": "failed", "message": "User doesn't have permission"})
		}

		fmt.Println("Hollaa")

		c.Set("username", valid.Claims.(jwt.MapClaims)["username"])
		c.Set("role", valid.Claims.(jwt.MapClaims)["role"])

		return next(c)
	}
}

func validateToken(t string, k string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})

	return token, err
}
