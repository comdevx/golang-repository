package handler

import (
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorValidateResponse struct {
	Code   int        `json:"code"`
	Errors []ErrorMsg `json:"errors"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Should be greater than " + fe.Param()
	case "max":
		return "Should be less than " + fe.Param()
	}
	return "Unknown error"
}

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case service.AppError:
		c.AbortWithStatusJSON(e.Code, ErrorResponse{Code: e.Code, Message: e.Message})
	case validator.ValidationErrors:
		list := make([]ErrorMsg, len(e))
		for i, fe := range e {
			list[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorValidateResponse{Code: http.StatusBadRequest, Errors: list})
	case error:
		status := http.StatusBadRequest
		c.AbortWithStatusJSON(status, ErrorResponse{Code: status, Message: e.Error()})
	}
}
