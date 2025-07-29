package model

type SignUpRequest struct {
	Username        string `json:"username" binding:"required,min=5,max=60"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type SignInResponse struct {
	User         *UserResponse `json:"profile,omitempty"`
	AccessToken  string        `json:"access_token,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
}

type RefreshRequest struct {
	RefreshToken string `json:"token"`
}

// type StoreSession struct {
// 	Username  string
// 	UserID    string
// 	UserAgent string
// 	Token     string
// }