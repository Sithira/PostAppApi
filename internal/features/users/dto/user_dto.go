package dto

type CreateUserRequest struct {
	FullName    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type CreateUserResponse struct {
	UserID    int64  `json:"user_id"`
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}
