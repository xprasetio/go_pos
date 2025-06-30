package user_domain

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (*User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE email = ?", email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (name, email, password, role_id, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", user.Name, user.Email, user.Password, user.RoleID)
	return err
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password, role_id, created_at, updated_at FROM users WHERE id = ?", id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
