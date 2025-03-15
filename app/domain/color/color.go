package color

import (
	"picolor-backend/app/domain/room"
)

type ColorCode string

type ColorID int

type Color struct {
	ColorId ColorID
	Color   ColorCode
	RoomId  room.RoomID
}

type CreateColor struct {
	ColorCodes []ColorCode
	RoomId    room.RoomID
}