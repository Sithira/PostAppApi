package dto

type RegisterRequest struct {
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	PasswordRetyped string `json:"password_retyped"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}
