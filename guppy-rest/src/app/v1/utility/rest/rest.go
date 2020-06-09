package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse types
type APIResponse struct {
	Code          int               `json:"code"`
	Name          string            `json:"name,omitempty"`
	Error         string            `json:"error,omitempty"`
	Meta          interface{}       `json:"meta,omitempty"`
	Results       interface{}       `json:"data,omitempty"`
	Message       string            `json:"message"`
	MessageDetail map[string]string `json:"message_detail,omitempty"`
}

// SuccessResponse (200)
func SuccessResponse(context *gin.Context, payload interface{}, meta interface{}, message string) {
	response := APIResponse{
		Code:    http.StatusOK,
		Meta:    meta,
		Results: payload,
		Message: message,
	}
	context.JSON(http.StatusOK, response)
}

// InvalidParameterResponse (400)
func InvalidParameterResponse(context *gin.Context, err error) {
	response := APIResponse{
		Code:    http.StatusOK,
		Name:    "VALIDATION_FAILURE",
		Error:   err.Error(),
		Message: "The given data was invalid.",
	}
	context.JSON(http.StatusOK, response)
}

// UnauthorizedResponse (401)
func UnauthorizedResponse(context *gin.Context, err error) {
	response := APIResponse{
		Code:    http.StatusUnauthorized,
		Name:    "AUTHENTICATION_FAILURE",
		Error:   err.Error(),
		Message: "Authentication failed due to invalid authentication credentials or a missing Authorization header.",
	}
	context.JSON(http.StatusOK, response)
}

// ForbiddenResponse (403)
func ForbiddenResponse(context *gin.Context, err error) {
	response := APIResponse{
		Code:    http.StatusForbidden,
		Name:    "FORBIDDEN",
		Error:   err.Error(),
		Message: "You can not access this resource. Please contact the Admin.",
	}
	context.JSON(http.StatusOK, response)
}

// NotFoundResponse (404)
func NotFoundResponse(context *gin.Context, err error) {
	response := APIResponse{
		Code:    http.StatusNotFound,
		Name:    "ROUTE_FAILURE",
		Error:   err.Error(),
		Message: "Method or Route not found",
	}
	context.JSON(http.StatusOK, response)
}

// UnprocessableEntityResponse (422)
func UnprocessableEntityResponse(context *gin.Context, err error, messageDetail map[string]string) {
	response := APIResponse{
		Code:          http.StatusUnprocessableEntity,
		Name:          "UNPROCESSABLE_ENTITY",
		Error:         err.Error(),
		Message:       "The given data is unprocessable entity.",
		MessageDetail: messageDetail,
	}
	context.JSON(http.StatusOK, response)
}

// ErrorResponse (500)
func ErrorResponse(context *gin.Context, err error) {
	response := APIResponse{
		Code:    http.StatusInternalServerError,
		Name:    "INTERNAL_SERVICE_ERROR",
		Error:   err.Error(),
		Message: "Something went wrong",
	}
	context.JSON(http.StatusOK, response)
}
