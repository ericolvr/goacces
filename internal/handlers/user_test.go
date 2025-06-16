package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"newaccess/internal/dto"
	"newaccess/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct{}

func (m *mockUserService) PinExists(_ context.Context, pin string) (*dto.QueryPinReponse, error) {
	if pin == "123456" {
		return &dto.QueryPinReponse{
			Name:     "Teste",
			Profile:  "admin",
			Document: "1234567",
		}, nil
	}
	return nil, repository.ErrNotFound
}

func (m *mockUserService) Create(_ context.Context, user *dto.UserRequest) (int, error) {
	return 42, nil
}
func (m *mockUserService) List(_ context.Context) ([]dto.UserResponse, error) {
	return []dto.UserResponse{{
		ID:         1,
		Name:       "Teste",
		Profile:    "admin",
		Document:   "1234567",
		CardNumber: 654321,
		Status:     true,
		WorkStart:  "08:00",
		WorkEnd:    "17:00",
	}}, nil
}

func (m *mockUserService) FindByID(_ context.Context, id int) (*dto.UserResponse, error) {
	if id == 1 {
		return &dto.UserResponse{
			ID:         1,
			Name:       "Teste",
			Profile:    "admin",
			Document:   "1234567",
			CardNumber: 654321,
			Status:     true,
			WorkStart:  "08:00",
			WorkEnd:    "17:00",
		}, nil
	}
	return nil, nil
}
func (m *mockUserService) Update(_ context.Context, user *dto.UserUpdateRequest) error { return nil }
func (m *mockUserService) Delete(_ context.Context, id int) error                      { return nil }

func TestUserHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.POST("/api/v1/users", h.Create)

	reqBody := dto.UserRequest{
		Name:       "Teste",
		Profile:    "admin",
		Document:   "1234567",
		Pin:        123456,
		Coercion:   1234567,
		CardNumber: 654321,
		Status:     true,
		WorkStart:  "08:00",
		WorkEnd:    "17:00",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.JSONEq(t, `{"id":42}`, resp.Body.String())
}

func TestUserHandler_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.GET("/api/v1/users", h.List)

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Teste")
}

func TestUserHandler_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.GET("/api/v1/users/:id", h.FindByID)

	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Teste")
}

func TestUserHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.PUT("/api/v1/users/:id", h.Update)

	reqBody := dto.UserUpdateRequest{
		ID:         1,
		Name:       "Teste",
		Profile:    "admin",
		Document:   "1234567",
		Pin:        123456,
		Coercion:   1234567,
		CardNumber: 654321,
		Status:     true,
		WorkStart:  "08:00",
		WorkEnd:    "17:00",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"id":1}`, resp.Body.String())
}

func TestUserHandler_PinExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.GET("/api/v1/users/check_pin", h.PinExists)

	// Caso 1: Pin encontrado
	req, _ := http.NewRequest("GET", "/api/v1/users/check_pin?pin=123456", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Teste")

	// Caso 2: Pin não encontrado
	req2, _ := http.NewRequest("GET", "/api/v1/users/check_pin?pin=000000", nil)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)
	assert.Equal(t, http.StatusNotFound, resp2.Code)
	assert.Contains(t, resp2.Body.String(), "not found")

	// Caso 3: Parâmetro ausente
	req3, _ := http.NewRequest("GET", "/api/v1/users/check_pin", nil)
	resp3 := httptest.NewRecorder()
	r.ServeHTTP(resp3, req3)
	assert.Equal(t, http.StatusBadRequest, resp3.Code)
	assert.Contains(t, resp3.Body.String(), "Missing pin parameter")
}

func TestUserHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewUserHandler(&mockUserService{})
	r := gin.Default()
	r.DELETE("/api/v1/users/:id", h.Delete)

	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "success")
}
