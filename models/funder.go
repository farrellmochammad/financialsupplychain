package models

type Funder struct {
	FundId             string `json:"fund_id" form:"fund_id" query:"fund_id"`
	Nik                string `json:"nik" form:"nik" query:"nik"`
	SubmittedBy        string `json:"submitted_by" form:"submitted_by" query:"submitted_by"`
	SubmittedTimestamp string `json:"submitted_timestamp" form:"submitted_timestamp" query:"submitted_timestamp"`
	ProposedBy         string `json:"proposed_by" form:"proposed_by" query:"proposed_by"`
	ProposedTimestamp  string `json:"propsoed_timestamp" form:"propsoed_timestamp" query:"propsoed_timestamp"`
	FundedBy           string `json:"funded_by" form:"funded_by" query:"funded_by"`
	FundedTimestamp    string `json:"funded_timestamp" form:"funded_timestamp" query:"funded_timestamp"`
	FileUrl            string `json:"file_url" form:"file_url" query:"file_url"`
	FishType           string `json:"fish_type" form:"fish_type" query:"fish_type"`
	Status             string `json:"status" form:"status" query:"status"`
	NumberOfPonds      int    `json:"number_of_ponds" form:"number_of_ponds" query:"number_of_ponds"`
	AmountOfFund       int    `json:"amount_of_fund" form:"amount_of_fund" query:"amount_of_fund"`
}
