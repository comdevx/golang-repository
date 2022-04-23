package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project/handler"
	"project/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var list = []service.UserResponse{}

func init() {
	id1, _ := primitive.ObjectIDFromHex("626145c5161badf80e0e676f")
	id2, _ := primitive.ObjectIDFromHex("626145cc0edb96a4c9198962")
	id3, _ := primitive.ObjectIDFromHex("626145d3abbe3ccaafd5e2bd")
	list = []service.UserResponse{
		{UserID: id1, Username: "test1", Password: "test1", Verified: true, Suspended: false},
		{UserID: id2, Username: "test2", Password: "test2", Verified: true, Suspended: false},
		{UserID: id3, Username: "test3", Password: "test3", Verified: true, Suspended: false},
	}
}

func TestGetUserAll(t *testing.T) {

	t.Run("error server", func(t *testing.T) {

		//Arrange
		var response []service.UserResponse
		expected := service.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}

		userService := &service.UserServiceMock{}
		userService.On("GetUsers").Return(response, service.ErrServerError())

		userHandler := handler.NewUserHandler(userService)

		app := gin.New()
		app.GET("/users", userHandler.GetUsers)

		//Act
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

	t.Run("success", func(t *testing.T) {

		//Arrange
		expected := list

		userService := &service.UserServiceMock{}
		userService.On("GetUsers").Return(expected, nil)

		userHandler := handler.NewUserHandler(userService)

		app := gin.New()
		app.GET("/users", userHandler.GetUsers)

		//Act
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
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

	t.Run("error server", func(t *testing.T) {

		//Arrange
		id := "626145c5161badf80e0e676f"
		expected := service.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&list[0], service.ErrServerError())

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

	t.Run("error no id", func(t *testing.T) {

		//Arrange
		id := "626145c5161badf80e0e676f"
		expected := handler.ErrorResponse{Code: 404, Message: "user not found"}

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&service.UserResponse{}, errors.New("user not found"))

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
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
		id := "626145c5161badf80e0e676f"
		expected := list[0]

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&expected, nil)

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
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
