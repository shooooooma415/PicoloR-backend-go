package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/room"
)

type RoomRepositoryImpl struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (q *RoomRepositoryImpl) CreateRooms() (*room.Room, error) {
	query := `
		INSERT INTO rooms
		RETURNING *
		`

	var createdRoom room.Room

	err := q.db.QueryRow(
		query,
	).Scan(
		&createdRoom,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create room:%w", err)
	}
	return &createdRoom, nil
}
