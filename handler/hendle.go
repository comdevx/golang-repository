package handler

import (
	"bank/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, e)
	case error:
		c.JSON(http.StatusInternalServerError, e)
	}
}
