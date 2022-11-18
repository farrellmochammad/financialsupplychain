package main

import (
	// __interface "financingsupplychain/interfaces"

	__authUsecase "financingsupplychain/usecases/auth"
	__CreditUsecase "financingsupplychain/usecases/credit"
	__ExperienceUsecase "financingsupplychain/usecases/experience"
	__FunderUsecase "financingsupplychain/usecases/funder"
	__MonitoringUsecase "financingsupplychain/usecases/monitoring"
	__PondUsecase "financingsupplychain/usecases/pond"
	__SignUsecase "financingsupplychain/usecases/sign"
	__SpawningUsecase "financingsupplychain/usecases/spawning"
	__SpawningHistoryUsecase "financingsupplychain/usecases/spawning_history"

	__middleware "financingsupplychain/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// __interface.StartConnection()

	e := echo.New()

	e.Use(middleware.CORS())

	e.POST("/login", __middleware.GenerateJWT(__authUsecase.Login))
	e.POST("/register", __authUsecase.Register)

	e.POST("/spawning", __middleware.ValidateJWT(__SpawningUsecase.InsertSpawning))
	e.GET("/spawnings", __middleware.ValidateJWT(__SpawningUsecase.GetSpawnings))
	e.GET("/spawning/:spawningid", __middleware.ValidateJWT(__SpawningUsecase.GetSpawning))

	e.POST("/experience", __middleware.ValidateJWT(__ExperienceUsecase.InsertExperience))
	e.PUT("/experience", __middleware.ValidateJWT(__ExperienceUsecase.UpdateExperience))
	e.GET("/experiences", __middleware.ValidateJWT(__ExperienceUsecase.GetExperiences))
	e.GET("/experience/:nik", __middleware.ValidateJWT(__ExperienceUsecase.GetExperience))
	e.PUT("/uploadfile", __middleware.ValidateJWT(__ExperienceUsecase.UploadFile))

	e.GET("/funders", __middleware.ValidateJWT(__FunderUsecase.GetFunders))
	e.GET("/funder/:nik", __middleware.ValidateJWT(__FunderUsecase.GetFunderBlockChain))
	e.POST("/funder", __middleware.ValidateJWT(__FunderUsecase.InsertFunderBlockChain))

	e.POST("/monitoring", __middleware.ValidateJWT(__MonitoringUsecase.InsertMonitoring))
	e.GET("/monitorings", __middleware.ValidateJWT(__MonitoringUsecase.GetMonitorings))
	e.GET("/monitoring/:nik", __middleware.ValidateJWT(__MonitoringUsecase.GetMonitoring))

	e.POST("/monitoring_pond", __middleware.ValidateJWT(__MonitoringUsecase.InsertMonitoringBlockChain))
	e.GET("/monitoring_pond/:fundid", __middleware.ValidateJWT(__MonitoringUsecase.GetMonitoringBlockChain))

	e.POST("/pond", __middleware.ValidateJWT(__PondUsecase.InsertPond))
	e.GET("/ponds", __middleware.ValidateJWT(__PondUsecase.GetPonds))
	e.GET("/pond_id/:pondid", __middleware.ValidateJWT(__PondUsecase.GetPond))
	e.GET("/pond_fundid/:fundid", __middleware.ValidateJWT(__PondUsecase.GetPondByNik))

	e.POST("/sign", __middleware.ValidateJWT(__SignUsecase.InsertSigned))
	e.GET("/signs", __middleware.ValidateJWT(__SignUsecase.GetSigns))
	e.GET("/sign/:signid", __middleware.ValidateJWT(__SignUsecase.GetSign))

	e.POST("/credit", __middleware.ValidateJWT(__CreditUsecase.InsertCredit))
	e.GET("/credits", __middleware.ValidateJWT(__CreditUsecase.GetCredits))
	e.GET("/credit/:creditid", __middleware.ValidateJWT(__CreditUsecase.GetCredit))

	e.POST("/spawning_history", __middleware.ValidateJWT(__SpawningHistoryUsecase.InsertSpawningHistory))
	e.GET("/spawning_histories", __middleware.ValidateJWT(__SpawningHistoryUsecase.GetSpawningsHistories))
	e.GET("/spawning_history/:nik", __middleware.ValidateJWT(__SpawningHistoryUsecase.GetSpawningHistory))

	e.Logger.Fatal(e.Start("127.0.0.1:2021"))
}
