package handler

import (
	process "bank/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	custSer process.UserProcess
}

func NewUserHandler(custSer process.UserProcess) userHandler {
	return userHandler{custSer: custSer}
}

func (h userHandler) GetUsers(c *gin.Context) {

	users, err := h.custSer.GetUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) GetUser(c *gin.Context) {

	id := c.Param("user_id")
	users, err := h.custSer.GetUser(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) NewUser(c *gin.Context) {

	user := process.NewUserRequest{}

	if err := c.Bind(&user); err != nil {
		handleError(c, err)
		return
	}

	result, err := h.custSer.NewUser(user)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
