package user

type Service interface {
	CreateUserService(requestData CreateUserRequestModel) (err error)
	UpdateUserService(requestData UpdateUserRequestModel) (err error)
	DeleteUserService(uuid string) (err error)
	GetListUserService(offset string) (listUser []GetUserListResponse, err error)
	GetUserDetailData(uuid string) (userDetailData GetUserDetailResponse, err error)
}

type Repository interface {
	InsertUser(data CreateUserModel) (err error)
	UpdateUser(data UpdateUserRequestModel) (err error)
	DeleteUser(uuid string) (err error)
	GetListUser(offset int) (listUser []GetUserListResponse, err error)
	GetUserDetail(uuid string) (userDetail GetUserDetailResponse, err error)
}
