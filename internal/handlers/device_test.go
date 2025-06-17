package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"

	"newaccess/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockDeviceService struct{}

func (m *mockDeviceService) Create(_ context.Context, device *dto.DeviceRequest) (int, error) {
	return 99, nil
}
func (m *mockDeviceService) List(_ context.Context) ([]dto.DeviceResponse, error) {
	return []dto.DeviceResponse{{
		ID:       1,
		Name:     "Device Teste",
		ServerIP: "192.168.0.1",
		IP:       "10.0.0.1",
		Port:     8080,
		Uniorg:   "Org1",
		Timezone: "America/Sao_Paulo",
	}}, nil
}
func (m *mockDeviceService) FindByID(_ context.Context, id int) (*dto.DeviceResponse, error) {
	if id == 1 {
		return &dto.DeviceResponse{
			ID:       1,
			Name:     "Device Teste",
			ServerIP: "192.168.0.1",
			IP:       "10.0.0.1",
			Port:     8080,
			Uniorg:   "Org1",
			Timezone: "America/Sao_Paulo",
		}, nil
	}
	return nil, nil
}
func (m *mockDeviceService) Update(_ context.Context, device *dto.DeviceUpdateRequest) error {
	return nil
}
func (m *mockDeviceService) Delete(_ context.Context, id int) error {
	return nil
}

func TestDeviceHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewDeviceHandler(&mockDeviceService{})
	r.POST("/devices", h.Create)

	payload := dto.DeviceRequest{
		Name:     "Device Teste",
		ServerIP: "192.168.0.1",
		IP:       "10.0.0.1",
		Port:     8080,
		Uniorg:   "Org1",
		Timezone: "America/Sao_Paulo",
	}
	body, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/devices", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, float64(99), resp["id"])
}

func TestDeviceHandler_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewDeviceHandler(&mockDeviceService{})
	r.GET("/devices", h.List)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/devices", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []dto.DeviceResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, "Device Teste", resp[0].Name)
}

func TestDeviceHandler_FindByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewDeviceHandler(&mockDeviceService{})
	r.GET("/devices/:id", h.FindByID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/devices/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp dto.DeviceResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Device Teste", resp.Name)
}

func TestDeviceHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewDeviceHandler(&mockDeviceService{})
	r.PUT("/devices/:id", h.Update)
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
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/devices/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, float64(1), resp["id"])
}

func TestDeviceHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := NewDeviceHandler(&mockDeviceService{})
	r.DELETE("/devices/:id", h.Delete)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/devices/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Device deleted successfully", resp["message"])
}
