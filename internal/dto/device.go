package dto

type DeviceRequest struct {
	Name     string `json:"name" validate:"required,min=5"`
	ServerIP string `json:"server_ip" validate:"required"`
	IP       string `json:"ip" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Uniorg   string `json:"uniorg" validate:"required"`
	Timezone string `json:"timezone" validate:"required"`
}

type DeviceResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=5"`
	ServerIP string `json:"server_ip" validate:"required"`
	IP       string `json:"ip" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Uniorg   string `json:"uniorg" validate:"required"`
	Timezone string `json:"timezone" validate:"required"`
}

type DeviceUpdateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required,min=5"`
	ServerIP string `json:"server_ip" validate:"required"`
	IP       string `json:"ip" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Uniorg   string `json:"uniorg" validate:"required"`
	Timezone string `json:"timezone" validate:"required"`
}
