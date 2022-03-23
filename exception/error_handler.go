package exception

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func NewErrorResponse(data interface{}, errors interface{}) (errorResponse ErrorResponse) {
	errorResponse.Data = data
	errorResponse.Error = errors
	return errorResponse
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	errorMessage := ErrorMessage{}
	errr := json.Unmarshal([]byte(err.(string)), &errorMessage)
	httpStatusCode := 0
	var errorM interface{}
	if errr != nil {
		httpStatusCode = 500
		errorM = "error unmarshal json"
	} else {
		httpStatusCode = errorMessage.Code
		errorM = errorMessage.Error
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	errorResponse := NewErrorResponse(nil, errorM)
	erre := json.NewEncoder(w).Encode(errorResponse)
	if erre != nil {
		panic(erre)
	}
}
