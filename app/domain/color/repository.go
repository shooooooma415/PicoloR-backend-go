package color

import "picolor-backend/app/domain/room"

type ColorRepository interface {
	CreateColors(CreateColor CreateColor) (*CreateColor, error)
	GetColors(roomId room.RoomId) ([]Color, error)
}	