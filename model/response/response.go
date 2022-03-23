package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jeruktutut2/backend-user/exception"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func ToResponse(data interface{}, errors interface{}) (response Response) {
	response.Data = data
	response.Error = errors
	return response
}

func ResponseHttp(w http.ResponseWriter, httpStatusCode int, data interface{}, errors interface{}) {
	response := ToResponse(data, errors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	err := json.NewEncoder(w).Encode(response)
	exception.PanicIfError(err)
}

func ResponseHttpContext(w http.ResponseWriter, err error) {
	if err == context.Canceled || err == context.DeadlineExceeded {
		response := ToResponse(nil, "timeout or lost connection")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(408)
		err := json.NewEncoder(w).Encode(response)
		exception.PanicIfError(err)
	}
}
