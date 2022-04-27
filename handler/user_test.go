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
)

var list = service.UserListResponse{}

func init() {
	list = service.UserListResponse{
		List: []service.UserResponse{
			{ID: 1, Username: "test1", Password: "test1", Verified: true, Suspended: false},
			{ID: 2, Username: "test2", Password: "test2", Verified: true, Suspended: false},
			{ID: 3, Username: "test3", Password: "test3", Verified: true, Suspended: false},
		},
		Total: 3,
	}
}

func TestGetUserAll(t *testing.T) {

	t.Run("error server", func(t *testing.T) {

		//Arrange
		page := 1
		limit := 10
		var response service.UserListResponse
		expected := service.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}

		userService := &service.UserServiceMock{}
		userService.On("GetUsers", page, limit).Return(response, service.ErrServerError())

		userHandler := handler.NewUserHandler(userService)

		app := gin.New()
		app.GET("/users", userHandler.GetUsers)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users?page=%v&limit=%v", page, limit), nil)
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
		page := 1
		limit := 10
		expected := list

		userService := &service.UserServiceMock{}
		userService.On("GetUsers", page, limit).Return(expected, nil)

		userHandler := handler.NewUserHandler(userService)

		app := gin.New()
		app.GET("/users", userHandler.GetUsers)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users?page=%v&limit=%v", page, limit), nil)
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
		id := 1
		expected := service.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected error",
		}

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&list.List[0], service.ErrServerError())

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%v", id), nil)
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
		id := 1
		expected := handler.ErrorResponse{Code: 400, Message: "User not found"}

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&service.UserResponse{}, errors.New("User not found"))

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%v", id), nil)
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
		id := 1
		expected := list.List[0]

		userService := &service.UserServiceMock{}
		userService.On("GetUser", id).Return(&expected, nil)

		userHandler := handler.NewUserHandler(userService)
		app := gin.New()
		app.GET("/users/:user_id", userHandler.GetUser)

		//Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%v", id), nil)
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
			ID:        1,
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
			ID:        1,
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
			ID:        1,
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
			ID:        1,
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
