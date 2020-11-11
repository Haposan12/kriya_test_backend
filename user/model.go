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
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

type CreateUserRequestModel struct {
	Email string `json:"email"`
	Status IsActiveModel `json:"status"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UpdateUserRequestModel struct {
	UUID string `json:"uuid"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserModel struct {
	UUID string
	Data string
	RoleID string
}

type DeleteUserRequestModel struct {
	UUID string `json:"uuid"`
}
