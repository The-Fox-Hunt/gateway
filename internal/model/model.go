package model

type SignupData struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignupSuccess struct {
	Success bool `json:"success"`
}
