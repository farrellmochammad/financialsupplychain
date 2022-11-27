package spawning

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertSpawningHistory(c echo.Context) error {

	spawning := new(__model.Spawning)
	if err := c.Bind(&spawning); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertSpawningHistory(spawning)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Spawning Success",
	})
}

func GetSpawningsHistories(c echo.Context) error {

	spawnings := __repository.GetSpawningHistories()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": spawnings,
	})
}

func GetSpawningHistory(c echo.Context) error {
	nik := c.Param("nik")

	spawning := __repository.GetSpawningHistory(nik)

	if len(spawning) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "nik not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": spawning,
	})

}
