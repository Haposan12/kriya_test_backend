package user

import (
	"encoding/json"
	"kriya_test_backend/shared/response"
	"net/http"
)

type HTTPDelivery struct {
	userService Service
}

func NewDelivery(userService Service) *HTTPDelivery {
	return &HTTPDelivery{userService: userService}
}

func (delivery HTTPDelivery) PostUser(w http.ResponseWriter, r *http.Request) {
	var requestData CreateUserRequestModel

	var err error

	// decode body request to struct request data
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = delivery.userService.CreateUserService(requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, "User Inserted Successfully", http.StatusText(http.StatusOK))
}

func (delivery HTTPDelivery) UpdateUser(w http.ResponseWriter, r *http.Request)  {
	var err error
	var requestData UpdateUserRequestModel

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = delivery.userService.UpdateUserService(requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, "User Updated Successfully", http.StatusText(http.StatusOK))
}

func (delivery HTTPDelivery) DeleteUser(w http.ResponseWriter, r *http.Request)  {
	var requestData DeleteUserRequestModel
	var err error

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = delivery.userService.DeleteUserService(requestData.UUID)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, "User Deleted Successfully", http.StatusText(http.StatusOK))
}

func (delivery HTTPDelivery) GetListUser(w http.ResponseWriter, r *http.Request) {
	// Get offset from url param
	offset := r.URL.Query().Get("offset")

	listUser, err := delivery.userService.GetListUserService(offset)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, listUser, http.StatusText(http.StatusOK))
}

func (delivery HTTPDelivery) GetUser(w http.ResponseWriter, r *http.Request) {
	var requestData GetUserDetailRequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// get user data detail from user service
	data, err := delivery.userService.GetUserDetailData(requestData.UUID)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, data, http.StatusText(http.StatusOK))
}
