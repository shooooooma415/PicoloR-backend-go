package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
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

func (q *ColorRepositoryImpl) GetThemeColorsByRoomID(roomID auth.RoomID) ([]color.Color, error) {
	query := `
		SELECT color
		FROM room_colors
		WHERE room_id = $1
		`

	rows, err := q.db.Query(
		query,
		roomID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get theme colors:%w", err)
	}
	defer rows.Close()

	var colors []color.Color
	for rows.Next() {
		var color color.Color
		err := rows.Scan(
			&color.ColorCode,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan theme colors:%w", err)
		}
		colors = append(colors, color)
	}
	return colors, nil
}

func (q *ColorRepositoryImpl) GetThemeColorByColorID(colorID auth.ColorID) (*color.Color, error) {
	query := `
		SELECT color
		FROM room_colors
		WHERE id = $1
		`

	var themeColor color.Color

	err := q.db.QueryRow(
		query,
		colorID,
	).Scan(
		&themeColor.ColorCode,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get theme color:%w", err)
	}
	return &themeColor, nil
}
