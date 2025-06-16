package repository

import (
	"context"
	"database/sql"
	"testing"

	"newaccess/internal/dto"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	// Cria tabela igual ao schema real
	_, err = db.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		profile TEXT,
		document TEXT,
		pin INTEGER,
		coercion INTEGER,
		card_number INTEGER,
		status BOOLEAN,
		work_start TEXT,
		work_end TEXT
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	return db
}

func TestUserRepository_CreateAndFindByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	user := &dto.UserRequest{
		Name:       "João Silva",
		Profile:    "admin",
		Document:   "123456789",
		Pin:        123456,
		Coercion:   1234567,
		CardNumber: 456789,
		Status:     true,
		WorkStart:  "08:00",
		WorkEnd:    "17:00",
	}

	id, err := repo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if id <= 0 {
		t.Fatalf("invalid id returned: %d", id)
	}

	saved, err := repo.FindByID(context.Background(), id)
	if err != nil {
		t.Fatalf("FindByID failed: %v", err)
	}
	if saved.Name != user.Name {
		t.Errorf("Name mismatch: got %s, want %s", saved.Name, user.Name)
	}
	if saved.Profile != user.Profile {
		t.Errorf("Profile mismatch: got %s, want %s", saved.Profile, user.Profile)
	}
	if saved.Document != user.Document {
		t.Errorf("Document mismatch: got %s, want %s", saved.Document, user.Document)
	}
	if saved.CardNumber != user.CardNumber {
		t.Errorf("CardNumber mismatch: got %d, want %d", saved.CardNumber, user.CardNumber)
	}
	if saved.Status != user.Status {
		t.Errorf("Status mismatch: got %v, want %v", saved.Status, user.Status)
	}
	if saved.WorkStart != user.WorkStart {
		t.Errorf("WorkStart mismatch: got %s, want %s", saved.WorkStart, user.WorkStart)
	}
	if saved.WorkEnd != user.WorkEnd {
		t.Errorf("WorkEnd mismatch: got %s, want %s", saved.WorkEnd, user.WorkEnd)
	}
	// Pin e Coercion não estão em UserResponse; se quiser testar, adicione no DTO e na query.
}

func TestUserRepository_PinExists(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	// Cria usuário
	user := &dto.UserRequest{
		Name:       "Maria",
		Profile:    "user",
		Document:   "987654321",
		Pin:        111222,
		Coercion:   7654321,
		CardNumber: 123456,
		Status:     true,
		WorkStart:  "09:00",
		WorkEnd:    "18:00",
	}
	_, err := repo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Pin existe
	resp, err := repo.PinExists(context.Background(), "111222")
	if err != nil {
		t.Fatalf("PinExists failed: %v", err)
	}
	if resp.Name != "Maria" {
		t.Errorf("Name mismatch: got %s, want %s", resp.Name, "Maria")
	}

	// Pin não existe
	_, err = repo.PinExists(context.Background(), "999999")
	if err == nil {
		t.Errorf("Esperava erro para pin inexistente, mas não houve erro")
	}
}
