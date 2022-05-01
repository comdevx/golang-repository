package handler

import (
	"net/http"
	service "project/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authenService service.AuthenService
}

func NewAuthenHandler(authenService service.AuthenService) AuthHandler {
	return AuthHandler{authenService: authenService}
}

func (h AuthHandler) Login(c *gin.Context) {

	auth := service.AuthenBody{}

	if err := c.ShouldBindJSON(&auth); err != nil {
		handleError(c, err.(validator.ValidationErrors))
		return
	}

	result, err := h.authenService.Login(auth)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h AuthHandler) ChangePassword(c *gin.Context) {

	body := service.PasswordForm{}

	if err := c.ShouldBindJSON(&body); err != nil {
		handleError(c, err.(validator.ValidationErrors))
	}

	result, err := h.authenService.ChangePassword(body)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, result)
}
