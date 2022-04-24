package handler_test

import (
	"bytes"
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
		expected := handler.ErrorResponse{Code: 400, Message: "User not found"}

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&service.UserResponse{}, errors.New("User not found"))

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode) {
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

func TestCreateUser(t *testing.T) {

	t.Run("error server", func(t *testing.T) {

		//Arrange
		body := service.NewUserRequest{
			Username: "test",
			Password: "passtest",
		}
		response := service.UserResponse{
			UserID:    primitive.ObjectID{000000000000000000000011},
			Username:  "test",
			Password:  "passtest",
			Verified:  false,
			Suspended: false,
		}
		expected := service.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}

		userService := &service.UserServiceMock{}
		userService.On("NewUser", body).Return(&response, service.ErrServerError())

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.POST("/users", userHandler.NewUser)

		//Act
		toJson, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(toJson)
		req := httptest.NewRequest(http.MethodPost, "/users", buffer)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

	t.Run("error username min 4", func(t *testing.T) {

		//Arrange
		body := service.NewUserRequest{
			Username: "tes",
			Password: "123456",
		}
		response := service.UserResponse{
			UserID:    primitive.ObjectID{000000000000000000000011},
			Username:  "test",
			Password:  "passtest",
			Verified:  false,
			Suspended: false,
		}
		expected := service.AppError{
			Code:    http.StatusUnprocessableEntity,
			Message: "Username at least 4",
		}

		userService := &service.UserServiceMock{}
		userService.On("NewUser", body).Return(&response,
			expected)

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.POST("/users", userHandler.NewUser)

		//Act
		toJson, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(toJson)
		req := httptest.NewRequest(http.MethodPost, "/users", buffer)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

	t.Run("error password min 6", func(t *testing.T) {

		//Arrange
		body := service.NewUserRequest{
			Username: "test",
			Password: "12345",
		}
		response := service.UserResponse{
			UserID:    primitive.ObjectID{000000000000000000000011},
			Username:  "test",
			Password:  "passtest",
			Verified:  false,
			Suspended: false,
		}
		expected := service.AppError{
			Code:    http.StatusUnprocessableEntity,
			Message: "Password at least 6",
		}

		userService := &service.UserServiceMock{}
		userService.On("NewUser", body).Return(&response, expected)

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.POST("/users", userHandler.NewUser)

		//Act
		toJson, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(toJson)
		req := httptest.NewRequest(http.MethodPost, "/users", buffer)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusUnprocessableEntity, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})

	t.Run("error EOF", func(t *testing.T) {

		//Arrange
		body := service.NewUserRequest{
			Username: "test",
			Password: "passtest",
		}
		expected := handler.ErrorResponse{Code: 400, Message: "EOF"}

		userService := &service.UserServiceMock{}
		userService.On("NewUser", body).Return(&service.UserResponse{}, errors.New("EOF"))

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.POST("/users", userHandler.NewUser)

		//Act
		toJson, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(toJson)
		req := httptest.NewRequest(http.MethodPost, "/users", buffer)
		req.Header.Set("Content-Type", "application/xml")
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		fmt.Println("msg", res.Body.String(), res.Result().StatusCode)
		if assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode) {
			e, _ := json.Marshal(expected)
			assert.Equal(t, errors.New(string(e)), errors.New(res.Body.String()))
		}
	})

	t.Run("success", func(t *testing.T) {

		//Arrange
		body := service.NewUserRequest{
			Username: "test",
			Password: "passtest",
		}
		expected := service.UserResponse{
			UserID:    primitive.ObjectID{000000000000000000000011},
			Username:  "test",
			Password:  "passtest",
			Verified:  false,
			Suspended: false,
		}

		userService := &service.UserServiceMock{}
		userService.On("NewUser", body).Return(&expected, nil)

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.POST("/users", userHandler.NewUser)

		//Act
		toJson, _ := json.Marshal(body)
		buffer := bytes.NewBuffer(toJson)
		req := httptest.NewRequest(http.MethodPost, "/users", buffer)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		app.ServeHTTP(res, req)
		defer res.Result().Body.Close()

		//Assert
		if assert.Equal(t, http.StatusCreated, res.Result().StatusCode) {
			mock, _ := json.Marshal(expected)
			assert.Equal(t, string(mock), res.Body.String())
		}
	})
}
