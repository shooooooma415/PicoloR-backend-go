package color

import (
	"fmt"
	"picolor-backend/app/domain/color"
	"picolor-backend/app/domain/auth"
)

func (c *ColorServiceImpl) GetThemeColor(roomID auth.RoomID) ([]color.Color, error) {
	ColorIDs, err := c.colorRepo.GetColorIDsByRoomID(roomID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return themeColors, nil
}
