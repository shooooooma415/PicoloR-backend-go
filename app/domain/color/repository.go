package color

import "picolor-backend/app/domain/auth"

type ColorRepository interface {
	GetThemeColors(roomId auth.RoomID) ([]Color, error)
	DeleteThemeColors(roomId auth.RoomID) (*auth.RoomID, error)
}
