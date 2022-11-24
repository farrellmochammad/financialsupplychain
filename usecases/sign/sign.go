package sign

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"financingsupplychain/models"
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

func GetSignsBySales(c echo.Context) error {

	funded := __repository.GetFundersBySales(fmt.Sprintf("%v", c.Get("username")))

	var fundids []string
	for i := 0; i < len(funded); i++ {
		fundids = append(fundids, funded[i].FundId)
	}

	signs := __repository.GetSignsByFundId(fundids)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": signs,
	})
}

func GetSign(c echo.Context) error {

	fundid := c.Param("fundid")

	sign := __repository.GetSign(fundid)

	if (models.Signed{} == sign) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "sign data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": sign,
	})

}
