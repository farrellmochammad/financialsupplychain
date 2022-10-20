package models

type UserRegister struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Role     string `json:"role" form:"role" query:"role"`
}

type UserLogin struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Role     string `json:"role" form:"role" query:"role"`
}
