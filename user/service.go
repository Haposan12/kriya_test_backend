package user

import (
	"encoding/json"
	"kriya_test_backend/utils"
	"strconv"
)

type ServiceImpl struct {
	userRepository RepositoryPostgres
}

func NewService(userRepository RepositoryPostgres) *ServiceImpl {
	return &ServiceImpl{userRepository: userRepository}
}

func (svc ServiceImpl) CreateUserService(requestData CreateUserRequestModel) (err error) {
	// Generate uuid using utils function
	var data CreateUserModel
	var isActiveModel IsActiveModel

	statusActiveUser := true
	// set user status true
	isActiveModel.IsActive = &statusActiveUser

	requestData.Status = isActiveModel

	// marshalled request data
	marshalled, err := json.Marshal(requestData)
	if err != nil {
		return
	}

	// generate uuid
	uuid := utils.GenerateNewUUID()

	data.UUID = uuid
	data.Data = string(marshalled)
	data.RoleID = UserRoleID

	// insert user into db
	err = svc.userRepository.InsertUser(data)

	return
}

func (svc ServiceImpl) UpdateUserService(requestData UpdateUserRequestModel) (err error) {
	err = svc.userRepository.UpdateUser(requestData)

	return
}

func (svc ServiceImpl) DeleteUserService(uuid string) (err error)  {
	err = svc.userRepository.DeleteUser(uuid)

	return
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

func (svc ServiceImpl) GetUserDetailData(uuid string) (userDetailData GetUserDetailResponse, err error) {
	userDetailData, err = svc.userRepository.GetUserDetail(uuid)

	return
}
