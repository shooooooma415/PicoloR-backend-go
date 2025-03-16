package color

import "picolor-backend/app/domain/auth"

type ColorRepository interface {
	CreateColors(CreateColor CreateColor) (*CreateColor, error)
	GetColors(roomId auth.RoomID) ([]Color, error)
	DeleteColors(roomId auth.RoomID) (*auth.RoomID, error)
}
