package models

type Pond struct {
	PondId        string `json:"pond_id" form:"pond_id" query:"pond_id"`
	FundId        string `json:"fund_id" form:"fund_id" query:"fund_id"`
	TotalSpawning int    `json:"total_spawning" form:"total_spawning" query:"total_spawning"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
}
