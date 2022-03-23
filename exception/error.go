package exception

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Code  int
	Error interface{}
}

func NewErrorMessage(code int, errors interface{}) (errorMessage ErrorMessage) {
	errorMessage.Code = code
	errorMessage.Error = errors
	return errorMessage
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorAndRollback(err error, tx *sql.Tx) {
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			panic(err)
		}
		panic(err)
	}
}

func LogFatallnIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func PanicIfErrorValidator(err error) {
	if err != nil {
		errors := []string{}
		er := ""
		for _, err := range err.(validator.ValidationErrors) {

			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
			er = err.Field() + " is " + err.ActualTag()
			errors = append(errors, er)
		}
		errMessage := NewErrorMessage(400, errors)
		errMessageJson, err := json.Marshal(errMessage)
		PanicIfError(err)
		panic(string(errMessageJson))
	}
}
