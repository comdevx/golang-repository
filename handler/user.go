package handler

import (
	"errors"
	"net/http"
	service "project/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetUsers(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	pageToInt, _ := strconv.Atoi(page)
	limitToInt, _ := strconv.Atoi(limit)

	if pageToInt < 1 {
		handleError(c, errors.New("Page value less than 1"))
		return
	}

	if limitToInt < 1 || limitToInt > 100 {
		handleError(c, errors.New("Limit values less than 1 or greater than 100"))
		return
	}

	users, err := h.userService.GetUsers(pageToInt, limitToInt)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) GetUser(c *gin.Context) {

	id := c.Param("user_id")
	toInt, _ := strconv.Atoi(id)
	users, err := h.userService.GetUser(toInt)
	if err != nil {
		handleError(c, err.(validator.ValidationErrors))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) NewUser(c *gin.Context) {

	user := service.NewUserRequest{}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err.(validator.ValidationErrors))
		return
	}

	result, err := h.userService.NewUser(user)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
