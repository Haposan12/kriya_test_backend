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

func (delivery HTTPDelivery) GetUser(w http.ResponseWriter, r *http.Request)  {
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

