package models

type Experience struct {
	Username        string `json:"username" form:"username" query:"username"`
	Nik             string `json:"nik" form:"nik" query:"nik"`
	Date            string `json:"date" form:"date" query:"date"`
	DurationOfYears int    `json:"duration_of_years" form:"duration_of_years" query:"duration_of_years"`
}
