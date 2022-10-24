package models

type Spawning struct {
	SpawningId string `json:"spawning_id" form:"spawning_id" query:"spawning_id"`
	Nik        string `json:"nik" form:"nik" query:"nik"`
	Date       string `json:"date" form:"date" query:"date"`
	Amount     int    `json:"amount" form:"amount" query:"amount"`
	FishType   string `json:"fish_type" form:"fish_type" query:"fish_type"`
}
