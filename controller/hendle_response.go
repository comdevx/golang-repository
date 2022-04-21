package controller

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
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{Code: http.StatusNotFound, Message: e.Error()})
	}
}
