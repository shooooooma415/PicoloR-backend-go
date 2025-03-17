package color

import "picolor-backend/app/domain/auth"

type Color struct {
	ColorId   auth.ColorID
	ColorCode auth.ColorCode
	RoomId    auth.RoomID
}

type CreateColor struct {
	ColorCodes []auth.ColorCode
	RoomId     auth.RoomID
}
