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

type SignInSuccess struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type ChangePasswordData struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordSuccess struct {
	Success bool `json:"success"`
}

type Key string

const (
	Username Key = "username"
)
