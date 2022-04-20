package controller

import (
	"net/http"
	errs "project/helper/errs"
	process "project/process"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userCon process.UserProcess
}

func NewUserController(userCon process.UserProcess) userController {
	return userController{userCon: userCon}
}

func (h userController) GetUsers(c *gin.Context) {

	users, err := h.userCon.GetUsers()
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userController) GetUser(c *gin.Context) {

	id := c.Param("user_id")
	users, err := h.userCon.GetUser(id)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userController) NewUser(c *gin.Context) {

	user := process.NewUserRequest{}

	if err := c.Bind(&user); err != nil {
		errs.HandleError(c, err)
		return
	}

	result, err := h.userCon.NewUser(user)
	if err != nil {
		errs.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
