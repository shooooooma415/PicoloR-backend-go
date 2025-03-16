package color

import "picolor-backend/app/domain/auth"

type ColorCode string

type ColorID int

type Color struct {
	ColorId ColorID
	Color   ColorCode
	RoomId  auth.RoomID
}

type CreateColor struct {
	ColorCodes []ColorCode
	RoomId     auth.RoomID
}
