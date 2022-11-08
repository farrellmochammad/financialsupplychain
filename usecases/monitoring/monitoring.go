package monitoring

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertMonitoring(c echo.Context) error {
	fmt.Println("Insert monitoring")
	monitoring := new(__model.Monitoring)
	if err := c.Bind(&monitoring); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertMonitoring(monitoring)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Insert Monitoring Success",
	})
}

func GetMonitorings(c echo.Context) error {

	monitorings := __repository.GetMonitorings()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": monitorings,
	})
}

func GetMonitoring(c echo.Context) error {

	nik := c.Param("fundid")

	monitoring := __repository.GetMonitoring(nik)

	if (__model.Monitoring{}) == monitoring {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "monitoring data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": monitoring,
	})

}
