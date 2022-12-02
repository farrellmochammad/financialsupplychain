package models

type Experience struct {
	Nik           string `json:"nik" form:"nik" query:"nik"`
	Name          string `json:"name" form:"name" query:"name"`
	Phone         string `json:"phone" form:"phone" query:"phone"`
	SubmitBy      string `json:"submit_by" form:"submit_by" query:"submit_by"`
	ProposeBy     string `json:"proposed_by" form:"proposed_by" query:"proposed_by"`
	ApprovedBy    string `json:"approved_by" form:"approved_by" query:"approved_by"`
	Dob           string `json:"dob" form:"dob" query:"dob"`
	Address       string `json:"address" form:"address" query:"address"`
	StartFarming  string `json:"start_farming" form:"start_farming" query:"start_farming"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
	NumberOfPonds int    `json:"number_of_ponds" form:"number_of_ponds" query:"number_of_ponds"`
	Notes         string `json:"notes" form:"notes" query:"notes"`
	CurrentStatus string `json:"current_status" form:"current_status" query:"current_status"`
	UrlFile       string `json:"url_file" form:"url_file" query:"url_file"`
}

type UploadFile struct {
	Nik     string `json:"nik" form:"nik" query:"nik"`
	FileUrl string `json:"file_url" form:"file_url" query:"file_url"`
}

type UploadFileFund struct {
	FundId  string `json:"fund_id" form:"fund_id" query:"fund_id"`
	FileUrl string `json:"file_url" form:"file_url" query:"file_url"`
}

type InsertFund struct {
	FundId       string `json:"fund_id" form:"fund_id" query:"fund_id"`
	AmountOfFund int    `json:"amount_of_fund" form:"amount_of_fund" query:"amount_of_fund"`
}

type FundSubmission struct {
	Nik           string `json:"nik" form:"nik" query:"nik"`
	StartFarming  string `json:"start_farming" form:"start_farming" query:"start_farming"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
	NumberOfPonds int    `json:"number_of_ponds" form:"number_of_ponds" query:"number_of_ponds"`
}
