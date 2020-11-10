package user

type UserModel struct {
	UserName string
	Email    string
	IsActive *bool
}

type GetUserListResponse struct {
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Status   IsActiveModel `json:"status"`
}

type IsActiveModel struct {
	IsActive *bool `json:"is_active"`
}

type GetUserDetailRequestData struct {
	UUID string
}

type GetUserDetailResponse struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	RoleName string `json:"role_name"`
}
