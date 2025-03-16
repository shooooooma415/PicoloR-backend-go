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

func (q *RoomRepositoryImpl) CreateRoomMember(user room.RoomMember) (*room.RoomMember, error) {
	query := `
		INSERT INTO room_members (user_id, room_id)
		VALUES $1, $2
		RETURNING *
		`

	var createdRoomMember room.RoomMember

	err := q.db.QueryRow(
		query,
		user.UserID,
		user.RoomID ,
	).Scan(
		&createdRoomMember,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create room member:%w", err)
	}
	return &createdRoomMember, nil
}

