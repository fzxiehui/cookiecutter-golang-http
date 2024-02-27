package request

type RegisterRequest struct {
	Username string `json:"Username" binding:"required" example:"alan"`
	Password string `json:"Password" binding:"required" example:"123456"`
	Email    string `json:"EMail" binding:"required,email" example:"1234@gmail.com"`
}

type LoginRequest struct {
	Username string `json:"UserName" binding:"required" example:"alan"`
	Password string `json:"Password" binding:"required" example:"123456"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"NickName" example:"alan"`
	Email    string `json:"EMail" binding:"required,email" example:"1234@gmail.com"`
	Avatar   string `json:"Avatar" example:"xxxx"`
}
