package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case AppError:
		c.JSON(e.Code, e)
	case error:
		c.JSON(http.StatusInternalServerError, e)
	}
}
