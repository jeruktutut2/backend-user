package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jeruktutut2/backend-user/exception"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	fmt.Println("r.Body:", r.Body)
	err := json.NewDecoder(r.Body).Decode(result)
	exception.PanicIfError(err)
}
