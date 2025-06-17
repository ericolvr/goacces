package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"newaccess/internal/dto"

	_ "github.com/mattn/go-sqlite3"
)

// setupTestDBDevice cria um banco em memória e a tabela device igual ao schema usado no DeviceRepository
func setupTestDBDevice(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	_, err = db.Exec(`CREATE TABLE device (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		server_ip TEXT,
		name TEXT,
		ip TEXT,
		port INTEGER,
		uniorg TEXT,
		timezone TEXT
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	// Inserindo 5 registros de exemplo
	for i := 1; i <= 5; i++ {
		_, err = db.Exec(`INSERT INTO device (server_ip, name, ip, port, uniorg, timezone) VALUES (?, ?, ?, ?, ?, ?)`,
			fmt.Sprintf("192.168.0.%d", i),
			fmt.Sprintf("Device %d", i),
			fmt.Sprintf("10.0.0.%d", i),
			8000+i,
			fmt.Sprintf("Org%d", i),
			"America/Sao_Paulo",
		)
		if err != nil {
			t.Fatalf("failed to insert device %d: %v", i, err)
		}
	}
	return db
}

func TestDeviceRepository_CreateAndFindByID(t *testing.T) {
	db := setupTestDBDevice(t)
	repo := NewDeviceRepository(db)

	deviceReq := &dto.DeviceRequest{
		Name:     "Device Teste",
		ServerIP: "192.168.0.1",
		IP:       "10.0.0.1",
		Port:     8080,
		Uniorg:   "Org1",
		Timezone: "America/Sao_Paulo",
	}

	ctx := context.Background()
	id, err := repo.Create(ctx, deviceReq)
	if err != nil {
		t.Fatalf("erro ao criar device: %v", err)
	}
	if id == 0 {
		t.Errorf("id retornado deve ser diferente de zero")
	}

	device, err := repo.FindByID(ctx, id)
	if err != nil {
		t.Fatalf("erro ao buscar device: %v", err)
	}
	if device == nil || device.Name != deviceReq.Name {
		t.Errorf("device retornado inválido: %+v", device)
	}
}

func TestDeviceRepository_List(t *testing.T) {
	db := setupTestDBDevice(t)
	repo := NewDeviceRepository(db)
	ctx := context.Background()
	// Cria dois devices
	_, err := repo.Create(ctx, &dto.DeviceRequest{
		Name:     "Device 1",
		ServerIP: "192.168.0.1",
		IP:       "10.0.0.1",
		Port:     8080,
		Uniorg:   "Org1",
		Timezone: "America/Sao_Paulo",
	})
	if err != nil {
		t.Fatalf("erro ao criar device 1: %v", err)
	}
	_, err = repo.Create(ctx, &dto.DeviceRequest{
		Name:     "Device 2",
		ServerIP: "192.168.0.2",
		IP:       "10.0.0.2",
		Port:     8081,
		Uniorg:   "Org2",
		Timezone: "America/Sao_Paulo",
	})
	if err != nil {
		t.Fatalf("erro ao criar device 2: %v", err)
	}
	devices, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("erro ao listar devices: %v", err)
	}
	if len(devices) != 7 {
		t.Errorf("esperado 7 devices, obtido %d", len(devices))
	}
}
