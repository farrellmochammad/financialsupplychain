package experience

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/montanaflynn/stats"

	"financingsupplychain/models"
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

	experience.SubmitBy = fmt.Sprintf("%v", c.Get("username"))
	if !__repository.NumOfPondsValidation(experience) {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Status": "Maaf jumlah kolam tidak sesuai dengan persyaratan",
		})
	}

	credits := __repository.GetCredit(experience.Nik)
	if !__repository.CreditsHistoryValidation(credits) {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Status": "Maaf data ditolak karena terdapat kredit macet",
		})
	}

	spawnings := stats.LoadRawData(__repository.GetSpawningHistoryAmounts(experience.Nik))

	average, _ := stats.Mean(spawnings)
	__repository.AddSpawningAverage(int(average))

	//need to improve
	creditscore := __repository.GetCurrentAverage()
	__repository.InsertExperienceBlockChain(fmt.Sprintf("%v", c.Get("username")), experience.Nik, experience.NumberOfPonds, int(average), creditscore)

	__repository.InsertExperience(experience)

	funder := &models.Funder{
		FundId:             __repository.NewSHA1Hash(),
		Nik:                experience.Nik,
		SubmittedBy:        fmt.Sprintf("%v", c.Get("username")),
		SubmittedTimestamp: time.Now().String(),
		FishType:           experience.FishType,
		NumberOfPonds:      experience.NumberOfPonds,
		AmountOfFund:       0,
	}

	fmt.Println("Funder : ", funder)
	__repository.InsertFunder(funder)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Experience Success",
	})
}

func UpdateExperience(c echo.Context) error {

	experience := new(__model.Experience)
	if err := c.Bind(&experience); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	__repository.UpdateExperience(experience)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Update Experience Success",
	})
}

func GetExperiences(c echo.Context) error {

	experiences := __repository.GetExperiences()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperiencesByUsername(c echo.Context) error {

	experiences := __repository.GetExperiencesByUsername(fmt.Sprintf("%v", c.Get("username")))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func UploadFile(c echo.Context) error {

	uploadfileurl := &__model.UploadFile{}
	if err := c.Bind(&uploadfileurl); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
			"err":     err,
		})
	}

	__repository.UploadFile(uploadfileurl.FileUrl, uploadfileurl.Nik)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})

}

func GetExperienceByFundId(c echo.Context) error {

	fundid := c.Param("fundid")
	username := fmt.Sprintf("%v", c.Get("username"))

	if c.Get("role") == "analyst" {
		experiences := __repository.GetExperienceBlockchain(fundid)

		if len(experiences) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "Experience not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": experiences,
		})
	}

	if c.Get("role") == "sales" {
		experiences := __repository.GetExperienceBlockchainByUsername(fundid, username)

		if len(experiences) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "Experience not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": experiences,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status": "Error getting experience",
	})
}

func GetExperienceByNik(c echo.Context) error {

	fundid := c.Param("nik")
	username := fmt.Sprintf("%v", c.Get("username"))

	if c.Get("role") == "analyst" {
		experiences := __repository.GetExperienceBlockchain(fundid)

		if len(experiences) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "Experience not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": experiences,
		})
	}

	if c.Get("role") == "sales" {
		experiences := __repository.GetExperienceBlockchainByUsername(fundid, username)

		if len(experiences) == 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "Experience not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": experiences,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status": "Error getting experience",
	})
}
