package dto

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type PasswordReset struct {
	Email string `json:"email"`
	Code string `json:"code"`
	Password string `json:"password"`
}