package http

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// RestErr Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// RestError Rest error struct
type RestError struct {
	Code             int         `json:"code,omitempty"`
	ErrorCode        string      `json:"error"`
	ErrorDescription string      `json:"error_description"`
	ErrCauses        interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrorCode, e.ErrorDescription, e.ErrCauses)
}

func NewBadRequest(code string, causes interface{}) RestError {
	message, ok := causes.(string)
	if ok {
		return RestError{
			Code:             http.StatusBadRequest,
			ErrorCode:        errors.New(code).Error(),
			ErrorDescription: message,
		}
	}
	return RestError{
		ErrorCode:        errors.New(causes.(string)).Error(),
		ErrorDescription: "",
	}
}

func NewInternalServerError(causes interface{}) RestError {
	return RestError{
		Code:             http.StatusBadRequest,
		ErrorCode:        errors.New("Internal Server Error").Error(),
		ErrorDescription: causes.(string),
	}
}
