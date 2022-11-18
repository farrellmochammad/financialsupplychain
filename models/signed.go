package models

type Signed struct {
	SignId  string `json:"sign_id" form:"sign_id" query:"sign_id"`
	FundId  string `json:"fund_id" form:"fund_id" query:"fund_id"`
	SignUrl string `json:"sign_url" form:"sign_url" query:"sign_url"`
}
