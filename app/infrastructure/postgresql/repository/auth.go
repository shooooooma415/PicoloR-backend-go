package repository

import (
	"database/sql"
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
		INSERT INTO users name 
		VALUES $1
		Returning id, name
		`

	var createdUser auth.User
		
	err := q.db.QueryRow(
		query,
		joinUser.RoomID,
		joinUser.UserName,
	).Scan(
		&createdUser.Id,
		&createdUser.Name,
	)

	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}
