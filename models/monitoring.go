package models

type Monitoring struct {
	FundId        string `json:"fund_id" form:"fund_id" query:"fund_id"`
	Nik           string `json:"nik" form:"nik" query:"nik"`
	Name          string `json:"name" form:"name" query:"name"`
	TotalSpawning int    `json:"total_spawning" form:"total_spawning" query:"total_spawning"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
}

type MonitoringPond struct {
	FundId      string `json:"fund_id" form:"fund_id" query:"fund_id"`
	Weight      int    `json:"weight" form:"weight" query:"weight"`
	Temperature int    `json:"temperature" form:"temperature" query:"temperature"`
	Humidity    int    `json:"humidity" form:"humidity" query:"humidity"`
}
