package handler

import (
	"net/http"
	service "project/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetUsers(c *gin.Context) {

	users, err := h.userService.GetUsers()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) GetUser(c *gin.Context) {

	id := c.Param("user_id")

	users, err := h.userService.GetUser(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// func (h userHandler) NewUser(c *gin.Context) {

// 	user := service.NewUserRequest{}

// 	if err := c.Bind(&user); err != nil {
// 		handleError(c, err)
// 		return
// 	}

// 	result, err := h.userService.NewUser(user)
// 	if err != nil {
// 		handleError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, result)
// }
