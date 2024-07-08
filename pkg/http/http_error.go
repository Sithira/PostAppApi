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
	Code             int         `json:"code"`
	ErrorCode        string      `json:"status,omitempty"`
	ErrorDescription string      `json:"error,omitempty"`
	ErrCauses        interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrorCode, e.ErrorDescription, e.ErrCauses)
}

func NewBadRequest(causes interface{}) RestError {
	return RestError{
		Code:             http.StatusBadRequest,
		ErrorCode:        errors.New("Bad Request").Error(),
		ErrorDescription: "",
	}
}
