package color

import "picolor-backend/app/domain/auth"

type Color struct {
	ColorID   auth.ColorID
	ColorCode auth.ColorCode
	RoomID    auth.RoomID
}

type CreateColor struct {
	ColorCodes []auth.ColorCode
	RoomID     auth.RoomID
}
