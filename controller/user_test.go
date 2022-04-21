package controller_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project/controller"
	"project/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetUserAll(t *testing.T) {

	id1, _ := primitive.ObjectIDFromHex("626145c5161badf80e0e676f")
	id2, _ := primitive.ObjectIDFromHex("626145cc0edb96a4c9198962")
	id3, _ := primitive.ObjectIDFromHex("626145d3abbe3ccaafd5e2bd")
	list := []service.UserResponse{
		{UserID: id1, Username: "test1", Password: "test1", Verified: true, Suspended: false},
		{UserID: id2, Username: "test2", Password: "test2", Verified: true, Suspended: false},
		{UserID: id3, Username: "test3", Password: "test3", Verified: true, Suspended: false},
	}

	t.Run("success", func(t *testing.T) {

		//Arrange
		expected := list

		userService := service.NewUserServiceMock()
		userService.On("GetUsers").Return(expected, nil)

		userController := controller.NewUserController(userService)

		app := gin.New()
		app.GET("/users", userController.GetUsers)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)

		//Act
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

}

func TestGetUser(t *testing.T) {

	id1, _ := primitive.ObjectIDFromHex("626145c5161badf80e0e676f")
	id2, _ := primitive.ObjectIDFromHex("626145cc0edb96a4c9198962")
	id3, _ := primitive.ObjectIDFromHex("626145d3abbe3ccaafd5e2bd")
	list := []service.UserResponse{
		{UserID: id1, Username: "test1", Password: "test1", Verified: true, Suspended: false},
		{UserID: id2, Username: "test2", Password: "test2", Verified: true, Suspended: false},
		{UserID: id3, Username: "test3", Password: "test3", Verified: true, Suspended: false},
	}

	t.Run("error no id", func(t *testing.T) {

		//Arrange
		id := "626145c5161badf80e0e676f"
		expected := controller.ErrorResponse{Code: 404, Message: "user not found"}

		userService := service.NewUserServiceMock()
		userService.On("GetUser", id).Return(&service.UserResponse{}, errors.New("user not found"))

		userController := controller.NewUserController(userService)

		app := gin.New()
		app.GET("/users/:user_id", userController.GetUser)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)

		//Act
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusNotFound, res.Result().StatusCode) {
			e, _ := json.Marshal(expected)
			assert.Equal(t, errors.New(string(e)), errors.New(res.Body.String()))
		}
	})

	t.Run("success", func(t *testing.T) {

		//Arrange
		id := "626145c5161badf80e0e676f1"
		expected := service.UserResponse{UserID: id1, Username: "test1", Password: "test1", Verified: true, Suspended: false}

		userService := service.NewUserServiceMock()
		userService.On("GetUser", id).Return(&list[0], nil)

		userController := controller.NewUserController(userService)

		app := gin.New()
		app.GET("/users/:user_id", userController.GetUser)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)

		//Act
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusOK, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})
}
