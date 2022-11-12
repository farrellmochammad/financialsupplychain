package models

type Funder struct {
	FundId        string `json:"fund_id" form:"fund_id" query:"fund_id"`
	Nik           string `json:"nik" form:"nik" query:"nik"`
	FishType      string `json:"fish_type" form:"fish_type" query:"fish_type"`
	NumberOfPonds int    `json:"number_of_ponds" form:"number_of_ponds" query:"number_of_ponds"`
	AmountOfFund  int    `json:"amount_of_fund" form:"amount_of_fund" query:"amount_of_fund"`
}
