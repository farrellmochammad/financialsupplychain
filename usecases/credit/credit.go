package credit

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertCredit(c echo.Context) error {

	credit := new(__model.Credit)
	if err := c.Bind(&credit); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertCredit(credit)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Credit Success",
	})
}

func GetCredits(c echo.Context) error {

	credits := __repository.GetCredits()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": credits,
	})
}

func GetCredit(c echo.Context) error {

	credit_id := c.Param("creditid")

	credit := __repository.GetCredit(credit_id)

	if len(credit) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Credit not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": credit,
	})

}
