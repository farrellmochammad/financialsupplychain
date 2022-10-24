package spawning

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertSpawning(c echo.Context) error {

	spawning := new(__model.Spawning)
	if err := c.Bind(&spawning); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InserSpawning(spawning)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Insert Spawning Success",
	})
}
