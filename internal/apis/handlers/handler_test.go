package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"goapibackend/internal/domain/dto"
	"goapibackend/mocks"
)

type HealtCheckResponse struct {
	Message string `json:"message"`
}

func TestHealthCheckHandler(t *testing.T) {
	mockService := mocks.MockUserService{}
	handler := Handler{
		UserService: &mockService,
	}

	expectedhealthResponse := HealtCheckResponse{Message: "Service Is Running Fine"}

	// We have to invoke the handler

	// We have to create a dummy gin server
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// Add some routes to it
	r.GET("/", handler.Healthcheck)
	// We have to create a request object too
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	// We have to convert the json body into struct
	var respo HealtCheckResponse
	err := json.Unmarshal(w.Body.Bytes(), &respo)
	assert.Nil(t, err)

	assert.Equal(t, expectedhealthResponse, respo)

}

func TestHandler_GetAllUsers(t *testing.T) {

	mockService := mocks.MockUserService{}
	var users []dto.UserDto
	users = append(users, dto.UserDto{
		FirstName: "Demo",
		LastName:  "Test",
		Email:     "Test",
		Location:  "Test",
		Schedule:  "Test",
		Password:  "Test",
		Degree:    "Test",
	})
	mockService.On("GetAllUsers").Return(users, nil)
	handler := Handler{
		UserService: &mockService,
	}

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// Add some routes to it
	r.GET("/users", handler.GetAllUsers)
	// We have to create a request object too
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestHandler_SignUp(t *testing.T) {
	mockService := mocks.MockUserService{}
	userDto := dto.UserDto{
		FirstName: "Demo",
		LastName:  "Test",
		Email:     "Test",
		Location:  "Test",
		Schedule:  "Test",
		Password:  "Test",
		Degree:    "Test",
	}
	handler := Handler{
		UserService: &mockService,
	}
	mockService.On("Signup", &userDto).Return(uint(1), nil)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	// Add some routes to it
	r.POST("/signup", handler.SignUp)

	body, err := json.Marshal(userDto)
	assert.Nil(t, err)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(body))
	r.ServeHTTP(w, req)
	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)

}
