package model

type SignupData struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignupSuccess struct {
	Success bool `json:"success"`
}

type SignInData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInSucess struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}
