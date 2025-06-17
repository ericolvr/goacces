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

type mockDeviceService struct{}

func (m *mockDeviceService) Create(_ context.Context, device *dto.DeviceRequest) (int, error) {
	return 101, nil
}
func (m *mockDeviceService) List(_ context.Context) ([]dto.DeviceResponse, error) {
	return []dto.DeviceResponse{{ID: 1, Name: "Device Teste", ServerIP: "192.168.0.1", IP: "10.0.0.1", Port: 8080, Uniorg: "Org1", Timezone: "America/Sao_Paulo"}}, nil
}
func (m *mockDeviceService) FindByID(_ context.Context, id int) (*dto.DeviceResponse, error) {
	if id == 1 {
		return &dto.DeviceResponse{ID: 1, Name: "Device Teste", ServerIP: "192.168.0.1", IP: "10.0.0.1", Port: 8080, Uniorg: "Org1", Timezone: "America/Sao_Paulo"}, nil
	}
	return nil, nil
}
func (m *mockDeviceService) Update(_ context.Context, device *dto.DeviceUpdateRequest) error { return nil }
func (m *mockDeviceService) Delete(_ context.Context, id int) error                      { return nil }

func TestDeviceRoutes_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := handlers.NewDeviceHandler(&mockDeviceService{})
	// Definindo as rotas de device
	r.POST("/api/v1/devices", h.Create)
	r.GET("/api/v1/devices", h.List)
	r.GET("/api/v1/devices/:id", h.FindByID)
	r.PUT("/api/v1/devices/:id", h.Update)
	r.DELETE("/api/v1/devices/:id", h.Delete)

	t.Run("POST /api/v1/devices", func(t *testing.T) {
		payload := dto.DeviceRequest{
			Name:     "Device Teste",
			ServerIP: "192.168.0.1",
			IP:       "10.0.0.1",
			Port:     8080,
			Uniorg:   "Org1",
			Timezone: "America/Sao_Paulo",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/api/v1/devices", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.Code)
		}
	})

	t.Run("GET /api/v1/devices", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/devices", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("GET /api/v1/devices/1", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/devices/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("PUT /api/v1/devices/1", func(t *testing.T) {
		payload := dto.DeviceUpdateRequest{
			ID:       1,
			Name:     "Device Atualizado",
			ServerIP: "192.168.0.1",
			IP:       "10.0.0.1",
			Port:     8080,
			Uniorg:   "Org1",
			Timezone: "America/Sao_Paulo",
		}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest("PUT", "/api/v1/devices/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})

	t.Run("DELETE /api/v1/devices/1", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/devices/1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
		if resp.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.Code)
		}
	})
}
