package experience

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/montanaflynn/stats"

	__model "financingsupplychain/models"
	__repository "financingsupplychain/repositories"
)

func InsertExperience(c echo.Context) error {

	experience := new(__model.Experience)
	if err := c.Bind(&experience); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	if !__repository.NumOfPondsValidation(experience) {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"Status": "Maaf jumlah kolam tidak sesuai dengan persyaratan",
		})
	}

	credits := __repository.GetCredit(experience.Nik)
	if !__repository.CreditsHistoryValidation(credits) {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"Status": "Maaf hdata ditolak karena terdapat kredit macet",
		})
	}

	spawnings := stats.LoadRawData(__repository.GetSpawningHistoryAmounts(experience.Nik))

	average, _ := stats.Mean(spawnings)
	__repository.AddSpawningAverage(int(average))
	__repository.GetCurrentAverage()
	__repository.InsertExperience(experience)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"Statys": "Insert Experience Success",
	})
}

func GetExperiences(c echo.Context) error {

	experiences := __repository.GetExperiences()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperience(c echo.Context) error {

	nik := c.Param("nik")

	experiences := __repository.GetExperience(nik)

	if len(experiences) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Experience not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})

}
