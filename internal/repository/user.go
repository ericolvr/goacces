package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"newaccess/internal/dto"

	_ "github.com/mattn/go-sqlite3"
)

var ErrNotFound = errors.New("record not found")

type UserRepository interface {
	Create(ctx context.Context, user *dto.UserRequest) (int, error)
	List(ctx context.Context) ([]dto.UserResponse, error)
	FindByID(ctx context.Context, id int) (*dto.UserResponse, error)
	PinExists(ctx context.Context, pin string) (*dto.QueryPinReponse, error)
	Update(ctx context.Context, user *dto.UserUpdateRequest) error
	Delete(ctx context.Context, id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *dto.UserRequest) (int, error) {
	query := `INSERT INTO users (
		name, 
		profile, 
		document, 
		pin, 
		coercion, 
		card_number, 
		status, 
		work_start, 
		work_end
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Profile,
		user.Document,
		user.Pin,
		user.Coercion,
		user.CardNumber,
		user.Status,
		user.WorkStart,
		user.WorkEnd,
	)

	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %w", err)
	}
	return int(id), nil
}

func (r *userRepository) List(ctx context.Context) ([]dto.UserResponse, error) {
	query := `SELECT * FROM users ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	defer rows.Close()

	var users []dto.UserResponse
	for rows.Next() {
		var user dto.UserResponse
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Profile,
			&user.Document,
			&user.Status,
			&user.WorkStart,
			&user.WorkEnd,
		); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*dto.UserResponse, error) {
	query := `SELECT id, name, profile, document, card_number, status, work_start, work_end FROM users WHERE id = ?`
	var user dto.UserResponse
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Profile,
		&user.Document,
		&user.CardNumber,
		&user.Status,
		&user.WorkStart,
		&user.WorkEnd,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error finding user by id: %w", err)
	}
	return &user, nil
}

func (r *userRepository) PinExists(ctx context.Context, pin string) (*dto.QueryPinReponse, error) {
	query := `SELECT name, profile, document FROM users WHERE pin = ? LIMIT 1`
	var resp dto.QueryPinReponse

	err := r.db.QueryRowContext(ctx, query, pin).Scan(&resp.Name, &resp.Profile, &resp.Document)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error checking pin: %w", err)
	}
	return &resp, nil
}

func (r *userRepository) Update(ctx context.Context, user *dto.UserUpdateRequest) error {
	query := `UPDATE users SET 
		name = ?,
		profile = ?,
		document = ?,
		pin = ?,
		coercion = ?,
		card_number = ?,
		status = ?,
		work_start = ?,
		work_end = ?
	WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Profile,
		user.Document,
		user.Pin,
		user.Coercion,
		user.CardNumber,
		user.Status,
		user.WorkStart,
		user.WorkEnd,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
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

func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
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
