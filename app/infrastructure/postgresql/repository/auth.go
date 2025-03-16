package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/auth"
)

type AuthRepositoryImpl struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (q *AuthRepositoryImpl) CreateUser(joinUser auth.JoinUser) (*auth.User, error) {
	query := `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id, name
		`

	var createdUser auth.User
		
	err := q.db.QueryRow(
		query,
		joinUser.UserName,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user:%w", err)
	}
	return &createdUser, nil
}

func (q *AuthRepositoryImpl) DeleteUserByUserID(deleteUser auth.DeleteUser) (*auth.User, error) {
	query := `
		DELETE FROM users 
		WHERE id = $1
		Returning id, name
		`

	var deletedUser auth.User
		
	err := q.db.QueryRow(
		query,
		deleteUser.UserID,
	).Scan(
		&deletedUser.ID,
		&deletedUser.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete user:%w", err)
	}
	return &deletedUser, nil
}
