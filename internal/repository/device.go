package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"newaccess/internal/dto"

	_ "github.com/mattn/go-sqlite3"
)

type DeviceRepository interface {
	Create(ctx context.Context, user *dto.DeviceRequest) (int, error)
	List(ctx context.Context) ([]dto.DeviceResponse, error)
	FindByID(ctx context.Context, id int) (*dto.DeviceResponse, error)
	Update(ctx context.Context, user *dto.DeviceUpdateRequest) error
	Delete(ctx context.Context, id int) error
}

type deviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (r *deviceRepository) Create(ctx context.Context, device *dto.DeviceRequest) (int, error) {
	query := `INSERT INTO device (
		name,
		server_ip,
		ip,
		port,
		uniorg,
		timezone
	) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(
		ctx,
		query,
		device.Name,
		device.ServerIP,
		device.IP,
		device.Port,
		device.Uniorg,
		device.Timezone,
	)

	if err != nil {
		return 0, fmt.Errorf("error creating device: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}
	return int(id), nil
}

func (r *deviceRepository) List(ctx context.Context) ([]dto.DeviceResponse, error) {
	query := `SELECT id, name, server_ip, ip, port, uniorg, timezone FROM device`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing devices: %w", err)
	}
	defer rows.Close()

	var devices []dto.DeviceResponse
	for rows.Next() {
		var device dto.DeviceResponse
		if err := rows.Scan(
			&device.ID,
			&device.Name,
			&device.ServerIP,
			&device.IP,
			&device.Port,
			&device.Uniorg,
			&device.Timezone,
		); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return devices, nil
}

func (r *deviceRepository) FindByID(ctx context.Context, id int) (*dto.DeviceResponse, error) {
	query := `SELECT id, name, server_ip, ip, port, uniorg, timezone FROM device WHERE id = ?`
	var device dto.DeviceResponse
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&device.ID,
		&device.Name,
		&device.ServerIP,
		&device.IP,
		&device.Port,
		&device.Uniorg,
		&device.Timezone,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error finding device by id: %w", err)
	}
	return &device, nil
}

func (r *deviceRepository) Update(ctx context.Context, device *dto.DeviceUpdateRequest) error {
	query := `UPDATE devices SET 
		name = ?,
		server_ip = ?,
		ip = ?,
		pin = ?,
		port = ?,
		uniorg = ?,
		timezone = ?,
	WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query,
		device.Name,
		device.ServerIP,
		device.IP,
		device.Port,
		device.Uniorg,
		device.Timezone,
		device.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating device: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *deviceRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM devices WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting device: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
