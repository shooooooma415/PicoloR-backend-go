package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/auth"
)

type ColorRepositoryImpl struct {
	db *sql.DB
}

func NewColorRepository(db *sql.DB) *ColorRepositoryImpl {
	return &ColorRepositoryImpl{db: db}
}

func (q *ColorRepositoryImpl) DeleteThemeColors(roomID auth.RoomID) (*auth.RoomID, error) {
	query := `
		DELETE FROM room_colors
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
		return nil, fmt.Errorf("failed to delete theme colors:%w", err)
	}
	return &deletedRoomID, nil
}
