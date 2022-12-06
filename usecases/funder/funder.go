package funder

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

func GetFunderBlockChain(c echo.Context) error {

	nik := c.Param("nik")

	funders := __repository.GetFunderBlockChain(nik)

	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funders not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func InsertFunderBlockChain(c echo.Context) error {
	funder := new(__model.Funder)
	if err := c.Bind(&funder); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	fundid := __repository.NewSHA1Hash()
	funder.FundId = fundid
	__repository.InsertFundingBlockchain(fundid, fmt.Sprintf("%v", c.Get("username")), funder.NumberOfPonds, funder.AmountOfFund)
	__repository.InsertFunder(funder)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Funder Success",
	})
}

func InsertFunderSubmission(c echo.Context) error {
	fund_submission := new(__model.FundSubmission)
	if err := c.Bind(&fund_submission); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
		})
	}

	if !__repository.NumOfPondsValidationByInt(fund_submission.NumberOfPonds) {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Status": "Maaf jumlah kolam tidak sesuai dengan persyaratan",
		})
	}

	credits := __repository.GetCredit(fund_submission.Nik)
	if !__repository.CreditsHistoryValidation(credits) {
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"Status": "Maaf data ditolak karena terdapat kredit macet",
		})
	}

	spawnings := stats.LoadRawData(__repository.GetSpawningHistoryAmounts(fund_submission.Nik))

	average, _ := stats.Mean(spawnings)
	__repository.AddSpawningAverage(int(average))

	spawning_histories := __repository.GetSpawningHistory(fund_submission.Nik)

	//Convert start farming to years of experience
	startfarming, errConvert := time.Parse("2006-01-02", fund_submission.StartFarming)

	if errConvert != nil {
		fmt.Println(errConvert)
		return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
			"status": "Cannot convert date to years of experience",
		})
	}

	currenttime := time.Now()

	yoe := int(currenttime.Sub(startfarming).Hours() / 24 / 365)

	creditscore := __repository.GenerateCreditScoreBlockChain(yoe, fund_submission.NumberOfPonds, credits, spawning_histories)

	funder := &models.Funder{
		FundId:             __repository.NewSHA1Hash(),
		Nik:                fund_submission.Nik,
		SubmittedBy:        fmt.Sprintf("%v", c.Get("username")),
		SubmittedTimestamp: time.Now().String(),
		FishType:           fund_submission.FishType,
		NumberOfPonds:      fund_submission.NumberOfPonds,
		AmountOfFund:       0,
	}

	__repository.InsertFunder(funder)

	__repository.InsertExperienceBlockChain(fmt.Sprintf("%v", c.Get("username")), funder.FundId, fund_submission.NumberOfPonds, int(average), creditscore)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "Insert Experience Success",
	})
}

func GetFundersBySales(c echo.Context) error {
	funders := __repository.GetFundersBySales(fmt.Sprintf("%v", c.Get("username")))
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})
}

func GetFundersByAnalyst(c echo.Context) error {
	funders := __repository.GetFundersByAnalyst(fmt.Sprintf("%v", c.Get("username")))
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})
}

func GetFunders(c echo.Context) error {

	funders := __repository.GetFunders()
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func GetFundersByFunder(c echo.Context) error {

	funders := __repository.GetFunderByFunders()
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func GetFundersByApproveFunder(c echo.Context) error {

	funders := __repository.GetFundersByApproveFunder(fmt.Sprintf("%v", c.Get("username")))
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func GetAllStatus(c echo.Context) error {

	fundid := c.Param("fund_id")
	funders := __repository.GetAllStatus(fundid)
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func GetFunderByNik(c echo.Context) error {

	nik := c.Param("nik")

	funders := __repository.GetFunderByNik(nik)
	if len(funders) == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Funder not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": funders,
	})

}

func UploadFileFunder(c echo.Context) error {

	uploadfileurl := &__model.UploadFileFund{}
	if err := c.Bind(&uploadfileurl); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
			"err":     err,
		})
	}

	__repository.UploadFileFundId(uploadfileurl.FileUrl, uploadfileurl.FundId, fmt.Sprintf("%v", c.Get("username")))

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "success",
	})

}

func InsertFunder(c echo.Context) error {

	insertfund := &__model.InsertFund{}
	if err := c.Bind(&insertfund); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
			"err":     err,
		})
	}

	__repository.InsertFund(insertfund.AmountOfFund, insertfund.FundId, fmt.Sprintf("%v", c.Get("username")))
	__repository.InsertApproveFundStatusFundingBlockchain(insertfund.FundId, fmt.Sprintf("%v", c.Get("username")))

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "success",
	})

}

func InsertRejectedFunder(c echo.Context) error {

	insertfund := &__model.InsertFund{}
	if err := c.Bind(&insertfund); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"message": "failed",
			"err":     err,
		})
	}

	__repository.InsertRejectedFundStatusFundingBlockchain(insertfund.FundId, fmt.Sprintf("%v", c.Get("username")))

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "success",
	})

}

func GetTracingHistories(c echo.Context) error {
	fundid := c.Param("fund_id")

	monitorings := __repository.GetMonitoringBlockChain(fundid)

	approvertransactions := __repository.GetAllStatus(fundid)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"monitorings":  monitorings,
			"transactions": approvertransactions,
		},
	})
}
