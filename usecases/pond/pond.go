package monitoring

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertPond(c echo.Context) error {
	pond := new(__model.Pond)
	fmt.Println(pond)
	if err := c.Bind(&pond); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertPond(pond)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Insert Pond Success",
	})
}

func GetPonds(c echo.Context) error {

	ponds := __repository.GetPonds()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": ponds,
	})
}

func GetPond(c echo.Context) error {

	pondid := c.Param("pondid")

	ponds := __repository.GetPond(pondid)

	if len(ponds) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "pond data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": ponds,
	})

}

func GetPondByNik(c echo.Context) error {

	fundid := c.Param("fundid")

	ponds := __repository.GetPondByFundId(fundid)

	if len(ponds) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "pond data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": ponds,
	})

}
