package color

import (
	"picolor-backend/app/domain/auth"
)

type ColorRepository interface {
	FindThemeColorsByRoomID(roomId auth.RoomID) ([]Color, error)
	FindColorIDsByRoomID(roomId auth.RoomID) ([]auth.ColorID, error)
	FindThemeColorByColorID(colorId auth.ColorID) (*Color, error)
	DeleteThemeColors(roomId auth.RoomID) (*auth.RoomID, error)
}
