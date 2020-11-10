package user

import "strconv"

type ServiceImpl struct {
	userRepository RepositoryPostgres
}

func NewService(userRepository RepositoryPostgres) *ServiceImpl {
	return &ServiceImpl{userRepository: userRepository}
}

func (svc ServiceImpl) GetListUserService(offset string) (listUser []GetUserListResponse, err error) {
	// convert offset into int
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return
	}

	if offset == "" {
		offsetInt = 0
	}
	listUser, err = svc.userRepository.GetListUser(offsetInt)

	return
}

func (svc ServiceImpl) GetUserDetailData(uuid string) (userDetailData GetUserDetailResponse, err error)  {
	userDetailData, err = svc.userRepository.GetUserDetail(uuid)

	return
}
