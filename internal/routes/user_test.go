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
	"newaccess/internal/repository"

	"github.com/gin-gonic/gin"
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

	t.Run("GET /api/v1/users/check_pin", func(t *testing.T) {
		// Caso sucesso
		req, _ := http.NewRequest("GET", "/api/v1/users/check_pin?pin=123456", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
		if !contains(resp.Body.String(), "Teste") {
			t.Errorf("expected body to contain 'Teste', got %s", resp.Body.String())
		}

		// Caso não encontrado
		req2, _ := http.NewRequest("GET", "/api/v1/users/check_pin?pin=000000", nil)
		resp2 := httptest.NewRecorder()
		r.ServeHTTP(resp2, req2)
		if resp2.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp2.Code)
		}

		// Caso parâmetro ausente
		req3, _ := http.NewRequest("GET", "/api/v1/users/check_pin", nil)
		resp3 := httptest.NewRecorder()
		r.ServeHTTP(resp3, req3)
		if resp3.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp3.Code)
		}
	})
}

// helper para contains
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
