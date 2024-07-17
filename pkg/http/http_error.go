package http_error

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrBadRequest         = "Bad request"
	ErrEmailAlreadyExists = "User with given email already exists"
	ErrNoSuchUser         = "User not found"
	ErrWrongCredentials   = "Wrong Credentials"
	ErrNotFound           = "Not Found"
	ErrUnauthorized       = "Unauthorized"
	ErrForbidden          = "Forbidden"
	ErrBadQueryParams     = "Invalid query params"
)

var (
	BadRequest            = errors.New("Bad request")
	JsonFormatFailed      = errors.New("Request Format Failed")
	WrongCredentials      = errors.New("Wrong Credentials")
	NotFound              = errors.New("Not Found")
	Unauthorized          = errors.New("Unauthorized")
	Forbidden             = errors.New("Forbidden")
	PermissionDenied      = errors.New("Permission Denied")
	ExpiredCSRFError      = errors.New("Expired CSRF token")
	WrongCSRFToken        = errors.New("Wrong CSRF token")
	CSRFNotPresented      = errors.New("CSRF not presented")
	NotRequiredFields     = errors.New("No such required fields")
	BadQueryParams        = errors.New("Invalid query params")
	InternalServerError   = errors.New("Internal Server Error")
	RequestTimeoutError   = errors.New("Request Timeout")
	ExistsEmailError      = errors.New("User with given email already exists")
	InvalidJWTToken       = errors.New("Invalid JWT token")
	InvalidJWTClaims      = errors.New("Invalid JWT claims")
	NotAllowedImageHeader = errors.New("Not allowed image header")
	NoCookie              = errors.New("not found cookie header")
)

// RestErr Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// RestError Rest error struct
type RestError struct {
	ErrStatus        int         `json:"status,omitempty"`
	ErrorCode        string      `json:"error_code,omitempty"`
	ErrorDescription interface{} `json:"error_description"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrorCode, e.ErrorDescription)
}

// Status Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrorDescription
}

// NewRestError New Rest Error
func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus:        status,
		ErrorCode:        err,
		ErrorDescription: causes,
	}
}

// NewRestErrorWithMessage New Rest Error With Message
func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus:        status,
		ErrorCode:        err,
		ErrorDescription: causes,
	}
}

// NewRestErrorFromBytes New Rest Error From Bytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// NewBadRequestError New Bad Request Error
func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus:        http.StatusBadRequest,
		ErrorCode:        BadRequest.Error(),
		ErrorDescription: causes,
	}
}

// NewNotFoundError New Not Found Error
func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus:        http.StatusNotFound,
		ErrorCode:        NotFound.Error(),
		ErrorDescription: causes,
	}
}

// NewUnauthorizedError New Unauthorized Error
func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrStatus:        http.StatusUnauthorized,
		ErrorCode:        Unauthorized.Error(),
		ErrorDescription: causes,
	}
}

// NewForbiddenError New Forbidden Error
func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus:        http.StatusForbidden,
		ErrorCode:        Forbidden.Error(),
		ErrorDescription: causes,
	}
}

// NewInternalServerError New Internal Server Error
func NewInternalServerError(causes interface{}) RestErr {
	result := RestError{
		ErrStatus:        http.StatusInternalServerError,
		ErrorCode:        InternalServerError.Error(),
		ErrorDescription: causes,
	}
	return result
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error(), err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error(), err)
	case strings.Contains(err.Error(), "SQLSTATE"):
		return parseSqlErrors(err)
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	case strings.Contains(err.Error(), "UUID"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewRestError(http.StatusUnauthorized, Unauthorized.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func parseSqlErrors(err error) RestErr {
	if strings.Contains(err.Error(), "23505") {
		return NewRestError(http.StatusBadRequest, ExistsEmailError.Error(), err)
	}

	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
}

func parseValidatorError(err error) RestErr {
	if strings.Contains(err.Error(), "Password") {
		return NewRestError(http.StatusBadRequest, "Invalid password, min length 6", err)
	}

	if strings.Contains(err.Error(), "Email") {
		return NewRestError(http.StatusBadRequest, "Invalid email", err)
	}

	return NewRestError(http.StatusBadRequest, BadRequest.Error(), err)
}

// ErrorResponse Error response
func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
