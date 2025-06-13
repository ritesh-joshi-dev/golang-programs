package model

type RegistrationData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `josn:"password"`
	MobileNo string `json:"mobile_no"`
}
