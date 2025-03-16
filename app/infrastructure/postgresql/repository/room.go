package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/room"
	"time"
)

type RoomRepositoryImpl struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{db: db}
}

func (q *RoomRepositoryImpl) CreateRoom(createRoom room.Room) (*room.Room, error) {
	query := `
		INSERT INTO rooms (is_start, is_finish, start_at)
		VALUES ($1, $2, $3)
		RETURNING id, is_start, is_finish, start_at
	`

	createRoom.StartAt = time.Now()

	var createdRoom room.Room

	err := q.db.QueryRow(
		query,
		createRoom.IsStart,
		createRoom.IsFinish,
		createRoom.StartAt,
	).Scan(
		&createdRoom.RoomID,
		&createdRoom.IsStart,
		&createdRoom.IsFinish,
		&createdRoom.StartAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create room: %w", err)
	}

	return &createdRoom, nil
}

func (q *RoomRepositoryImpl) CreateRoomMember(user room.RoomMember) (*room.RoomMember, error) {
	query := `
		INSERT INTO room_members (user_id, room_id)
		VALUES ($1, $2)
		RETURNING user_id, room_id
		`

	var createdRoomMember room.RoomMember

	err := q.db.QueryRow(
		query,
		user.UserID,
		user.RoomID,
	).Scan(
		&createdRoomMember.UserID,
		&createdRoomMember.RoomID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create room member:%w", err)
	}
	return &createdRoomMember, nil
}
