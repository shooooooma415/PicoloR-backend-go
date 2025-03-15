package color

import (
	"picolor-backend/app/domain/room"
)

type ColorCode string

type ColorId int

type Color struct {
	ColorId ColorId
	Color   ColorCode
	RoomId  room.RoomId
}
