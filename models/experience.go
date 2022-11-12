package models

type Experience struct {
	Nik           string `json:"nik" form:"nik" query:"nik"`
	Name          string `json:"name" form:"name" query:"name"`
	Phone         string `json:"phone" form:"phone" query:"phone"`
	Dob           string `json:"dob" form:"dob" query:"dob"`
	Address       string `json:"address" form:"address" query:"address"`
	StartFarming  string `json:"start_farming" form:"start_farming" query:"start_farming"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
	NumberOfPonds int    `json:"number_of_ponds" form:"number_of_ponds" query:"number_of_ponds"`
	Notes         string `json:"notes" form:"notes" query:"notes"`
	CurrentStatus string `json:"current_status" form:"current_status" query:"current_status"`
}

type UploadFile struct {
	Nik     string `json:"nik" form:"nik" query:"nik"`
	FileUrl string `json:"file_url" form:"file_url" query:"file_url"`
}
