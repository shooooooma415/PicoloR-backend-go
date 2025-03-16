package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/auth"
)

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (q *PostRepositoryImpl) DeletePostByRoomID(roomID auth.RoomID) (*auth.RoomID, error) {
	query := `
		DELETE FROM posts
		WHERE room_id = $1
		RETURNING room_id
		`

	var deletedRoomID auth.RoomID

	err := q.db.QueryRow(
		query,
		roomID,
	).Scan(
		&deletedRoomID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete posts:%w", err)
	}
	return &deletedRoomID, nil
}
