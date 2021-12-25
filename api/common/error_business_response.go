package common

import (
	"fmt"
	"go-hexagonal-auth/business"
	"net/http"
)

type errorBusinessResponseCode string

const (
	errResponseServer errorBusinessResponseCode = "response_server_error"
	errInternalServerError errorBusinessResponseCode = "internal_server_error"
	errHasBeenModified     errorBusinessResponseCode = "data_has_been modified"
	errNotFound            errorBusinessResponseCode = "data_not_found"
	errInvalidSpec         errorBusinessResponseCode = "invalid_spec"
)

//BusinessResponse default payload response
type BusinessResponse struct {
	Code    errorBusinessResponseCode `json:"code"`
	Message string                    `json:"message"`
	Data    interface{}               `json:"data"`
}

//NewErrorBusinessResponse Response return choosen http status like 400 bad request 422 unprocessable entity, ETC, based on responseCode
func NewErrorBusinessResponse(err error) (int, BusinessResponse) {
	return errorMapping(err)
}

//errorMapping error for missing header key with given value
func errorMapping(err error) (int, BusinessResponse) {
	fmt.Println(err)
	switch err {
	default:
		return newInternalServerErrorResponse(err.Error())
	case business.ErrNotFound:
		return newNotFoundResponse()
	case business.ErrInvalidSpec:
		return newValidationResponse(err.Error())
	case business.ErrHasBeenModified:
		return newHasBeedModifiedResponse()
	}
}

//newInternalServerErrorResponse default internal server error response
func newInternalServerErrorResponse(err string) (int, BusinessResponse) {
	return http.StatusInternalServerError, BusinessResponse{
		errResponseServer,
		"error "+err,
		map[string]interface{}{},
	}
}

//newHasBeedModifiedResponse failed to validate request payload
func newHasBeedModifiedResponse() (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errHasBeenModified,
		"Data has been modified",
		map[string]interface{}{},
	}
}

//newNotFoundResponse default not found error response
func newNotFoundResponse() (int, BusinessResponse) {
	return http.StatusNotFound, BusinessResponse{
		errNotFound,
		"Data Not found",
		map[string]interface{}{},
	}
}

//newValidationResponse failed to validate request payload
func newValidationResponse(message string) (int, BusinessResponse) {
	return http.StatusBadRequest, BusinessResponse{
		errInvalidSpec,
		"Validation failed " + message,
		map[string]interface{}{},
	}
}
