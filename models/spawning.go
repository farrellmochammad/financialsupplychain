package models

type Spawning struct {
	SpawningId    string `json:"spawning_id" form:"spawning_id" query:"spawning_id"`
	Nik           string `json:"nik" form:"nik" query:"nik"`
	Date          string `json:"date" form:"date" query:"date"`
	Amount        int    `json:"amount" form:"amount" query:"amount"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
	TotalSpawning int    `json:"total_spawning" form:"total_spawning" query:"total_spawning"`
	SpawningDate  string `json:"spawning_date" form:"spawning_date" query:"spawning_date"`
}
