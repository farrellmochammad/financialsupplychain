package funder

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

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

	__repository.UploadFileFundId(uploadfileurl.FileUrl, uploadfileurl.FundId)

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

	__repository.InsertFund(insertfund.AmountOfFund, insertfund.FundId)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "success",
	})

}
