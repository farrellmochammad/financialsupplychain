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

	spawning_histories := __repository.GetSpawningHistory(experience.Nik)

	//Convert start farming to years of experience
	startfarming, errConvert := time.Parse("2006-01-02", experience.StartFarming)

	if errConvert != nil {
		fmt.Println(errConvert)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"status": "Cannot convert date to years of experience",
		})
	}

	currenttime := time.Now()

	yoe := int(currenttime.Sub(startfarming).Hours() / 24 / 365)

	creditscore := __repository.GenerateCreditScoreBlockChain(yoe, experience.NumberOfPonds, credits, spawning_histories)

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

	__repository.InsertFunder(funder)

	__repository.InsertExperienceBlockChain(fmt.Sprintf("%v", c.Get("username")), funder.FundId, experience.NumberOfPonds, int(average), creditscore)
	__repository.InsertPendingFundStatusFundingBlockchain(funder.FundId, fmt.Sprintf("%v", c.Get("username")))

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

func GetExperienceSales(c echo.Context) error {

	experiences := __repository.GetExperiencesByUsername(fmt.Sprintf("%v", c.Get("username")))

	if len(experiences) == 0 {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"status": "Tidak ada data experience",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperienceAnalyst(c echo.Context) error {

	experiences := __repository.GetExperiencesByAnalystUsername(fmt.Sprintf("%v", c.Get("username")))

	if len(experiences) == 0 {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"status": "Tidak ada data experience",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperienceFunder(c echo.Context) error {

	experiences := __repository.GetExperiencesByFunderUsername(fmt.Sprintf("%v", c.Get("username")))

	if len(experiences) == 0 {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"status": "Tidak ada data experience",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperiencesSalesBlockchain(c echo.Context) error {

	fundid := c.Param("fund_id")

	experiences := __repository.GetExperienceBlockchainByUsername(fundid, fmt.Sprintf("%v", c.Get("username")))

	if len(experiences) == 0 {
		return c.JSON(http.StatusAccepted, map[string]interface{}{
			"status": "Tidak ada data experience",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperiences(c echo.Context) error {

	experiences := __repository.GetExperiences()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experiences,
	})
}

func GetExperiencesByUsername(c echo.Context) error {

	fundid := c.Param("fundid")

	experiences := __repository.GetExperienceBlockchainByUsername(fundid, fmt.Sprintf("%v", c.Get("username")))

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

	fundid := c.Param("fund_id")
	username := fmt.Sprintf("%v", c.Get("username"))

	if c.Get("role") == "analyst" || c.Get("role") == "funder" {
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
