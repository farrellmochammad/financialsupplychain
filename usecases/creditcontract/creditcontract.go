package credit

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__creditscorecontract "financingsupplychain/api/creditscorecontract"
	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertCredit(c echo.Context, creditcontract *__creditscorecontract.Api) error {

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
