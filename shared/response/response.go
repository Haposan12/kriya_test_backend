package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	responseTemplate := ResponseModel{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)

	resp, err := json.Marshal(responseTemplate)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)
}

func SendSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	HTTPResponse(w, data, message, http.StatusOK)
}

func SendErrorResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	HTTPResponse(w, data, message, statusCode)
}
