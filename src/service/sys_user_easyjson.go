package service

type UpdateUserInfoRequest struct {
	UserId int64  `json:"user_id"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type UpdateUserPasswordRequest struct {
	//UserId      int64  `json:"user_id"`
	UserName string `json:"user_name"`
	//Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
