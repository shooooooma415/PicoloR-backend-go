package color

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
)

func (c *ColorServiceImpl) GetThemeColors(roomID auth.RoomID) ([]color.Color, error) {
	ColorIDs, _ := c.colorRepo.FindColorIDsByRoomID(roomID)
	var ThemeColors []color.Color
	for _, colorID := range ColorIDs {
		color, _ := c.colorRepo.FindThemeColorByColorID(colorID)
		ThemeColors = append(ThemeColors, *color)
	}
	return ThemeColors, nil
}
