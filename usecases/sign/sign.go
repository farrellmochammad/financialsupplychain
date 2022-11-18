package sign

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertSigned(c echo.Context) error {
	sign := new(__model.Signed)
	if err := c.Bind(&sign); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertSigned(sign)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert sign Success",
	})
}

func GetSigns(c echo.Context) error {

	signs := __repository.GetSigns()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": signs,
	})
}

func GetSign(c echo.Context) error {

	signid := c.Param("signid")

	signs := __repository.GetSign(signid)

	if len(signs) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "sign data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": signs,
	})

}