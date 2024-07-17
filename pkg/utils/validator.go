package utils

import (
	httperror "RestApiBackend/pkg/http"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidateStruct(stub interface{}) (*int, error) {
	validate := validator.New()
	if err := validate.Struct(stub); err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		errorList := make(map[string]string)
		for _, fieldErr := range errs {
			errorList[fieldErr.Field()] = fieldErr.Tag()
		}
		return nil, httperror.NewRestError(http.StatusBadRequest, "ERR_000", errorList)
	}
	return nil, nil
}
