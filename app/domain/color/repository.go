package color

import (
	"picolor-backend/app/domain/auth"
)

type ColorRepository interface {
	GetThemeColorsByRoomID(roomId auth.RoomID) ([]Color, error)
	GetColorIDsByRoomID(roomId auth.RoomID) ([]auth.ColorID, error)
	GetThemeColorByColorID(colorId auth.ColorID) (*Color, error)
	DeleteThemeColors(roomId auth.RoomID) (*auth.RoomID, error)
}
