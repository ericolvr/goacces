package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"
	"newaccess/internal/dto"
	"newaccess/internal/handlers"

	"github.com/gin-gonic/gin"
)

type mockUserService struct{}

func (m *mockUserService) Create(_ context.Context, user *dto.UserRequest) (int, error) {
	return 1, nil
}
func (m *mockUserService) List(_ context.Context) ([]dto.UserResponse, error) {
	return []dto.UserResponse{{ID: 1, Name: "Teste"}}, nil
}
func (m *mockUserService) FindByID(_ context.Context, id int) (*dto.UserResponse, error) {
	if id == 1 {
		return &dto.UserResponse{ID: 1, Name: "Teste"}, nil
	}
	return nil, nil
}
func (m *mockUserService) Update(_ context.Context, user *dto.UserUpdateRequest) error { return nil }
func (m *mockUserService) Delete(_ context.Context, id int) error                      { return nil }

func TestUserRoutes_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := handlers.NewUserHandler(&mockUserService{})
	UserRoutes(r, h)

	t.Run("POST /api/v1/users", func(t *testing.T) {
		payload := map[string]interface{}{"name": "Teste"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.Code)
		}
	})

	t.Run("GET /api/v1/users", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("GET /api/v1/users/1", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("PUT /api/v1/users/1", func(t *testing.T) {
		payload := map[string]interface{}{"name": "Teste"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("DELETE /api/v1/users/1", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})
}
