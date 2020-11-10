package user

type Service interface {
	GetListUserService(offset string) (listUser []GetUserListResponse, err error)
	GetUserDetailData(uuid string) (userDetailData GetUserDetailResponse, err error)
}

type Repository interface {
	GetListUser(offset int) (listUser []GetUserListResponse, err error)
	GetUserDetail(uuid string) (userDetail GetUserDetailResponse, err error)
}
