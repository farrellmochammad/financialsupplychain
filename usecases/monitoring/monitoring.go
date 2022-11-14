package monitoring

import (
	"net/http"

	"github.com/labstack/echo/v4"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertMonitoring(c echo.Context) error {
	monitoring := new(__model.Monitoring)
	if err := c.Bind(&monitoring); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertMonitoring(monitoring)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Monitoring Success",
	})
}

func GetMonitorings(c echo.Context) error {

	monitorings := __repository.GetMonitorings()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": monitorings,
	})
}

func GetMonitoring(c echo.Context) error {

	fundid := c.Param("fundid")

	monitorings := __repository.GetMonitoring(fundid)

	if len(monitorings) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "monitoring data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": monitorings,
	})

}

func InsertMonitoringBlockChain(c echo.Context) error {
	monitoring := new(__model.MonitoringPond)
	if err := c.Bind(&monitoring); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.InsertMonitoringPondBlockChain(monitoring.FundId, monitoring)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Monitoring Success",
	})
}

func GetMonitoringBlockChain(c echo.Context) error {

	fundid := c.Param("fundid")

	monitorings := __repository.GetMonitoringBlockChain(fundid)

	if len(monitorings) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Monitoring not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": monitorings,
	})

}
