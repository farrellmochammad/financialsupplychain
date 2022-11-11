package models

type Credit struct {
	CreditId     string `json:"credit_id" form:"credit_id" query:"credit_id"`
	Username     string `json:"username" form:"username" query:"username"`
	Nik          string `json:"nik" form:"nik" query:"nik"`
	CreditAmount int    `json:"credit_amount" form:"credit_amount" query:"credit_amount"`
	Status       bool   `json:"status" form:"status" query:"status"`
}
