package handler

import (
	"net/http"
	"project/service"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case service.AppError:
		c.AbortWithStatusJSON(e.Code, ErrorResponse{Code: e.Code, Message: e.Message})
	case error:
		status := http.StatusBadRequest
		c.AbortWithStatusJSON(status, ErrorResponse{Code: status, Message: e.Error()})
	}
}
