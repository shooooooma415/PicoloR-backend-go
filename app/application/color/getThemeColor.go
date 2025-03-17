package color

import (
	"fmt"
	"picolor-backend/app/domain/color"
	"picolor-backend/app/domain/auth"
)

func (c *ColorServiceImpl) GetThemeColor(roomID auth.RoomID) ([]color.Color, error) {
	ColorIDs, _ := c.colorRepo.FindColorIDsByRoomID(roomID)
	var ThemeColors []color.Color
	for _, colorID := range ColorIDs {
		color, _ := c.colorRepo.FindThemeColorByColorID(colorID)
		ThemeColors = append(ThemeColors, *color)
	}
	return ThemeColors, nil
}
