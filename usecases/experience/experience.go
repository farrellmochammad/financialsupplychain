package experience

import (
	"net/http"

	"github.com/labstack/echo/v4"

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

	experience := __repository.GetExperience(nik)

	if (__model.Experience{}) == experience {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "Experience not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": experience,
	})

}
