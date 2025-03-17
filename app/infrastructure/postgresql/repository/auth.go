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

func (q *AuthRepositoryImpl) CreateUser(userName auth.UserName) (*auth.User, error) {
	query := `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id, name
		`

	var createdUser auth.User
		
	err := q.db.QueryRow(
		query,
		userName,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user:%w", err)
	}
	return &createdUser, nil
}

func (q *AuthRepositoryImpl) DeleteUserByUserID(userID auth.UserID) (*auth.User, error) {
	query := `
		DELETE FROM users 
		WHERE id = $1
		Returning id, name
		`

	var deletedUser auth.User
		
	err := q.db.QueryRow(
		query,
		userID,
	).Scan(
		&deletedUser.ID,
		&deletedUser.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete user:%w", err)
	}
	return &deletedUser, nil
}

func (q *AuthRepositoryImpl) FindUserByUserID(userID auth.UserID) (*auth.User, error) {
	query := `
		SELECT id, name
		FROM users
		WHERE id = $1
		`

	var foundUser auth.User
		
	err := q.db.QueryRow(
		query,
		userID,
	).Scan(
		&foundUser.ID,
		&foundUser.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to find user:%w", err)
	}
	return &foundUser, nil
}
