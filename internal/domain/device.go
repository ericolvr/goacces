package domain

type Device struct {
	ID       int    `json:"id"`
	ServerIP string `json:"server_i"`
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Port     int
	Uniorg   string `json:"uniorg"`
	Timezone string `json:"timezone"`
}
