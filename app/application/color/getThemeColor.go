package color

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
)

func (c *ColorServiceImpl) GetThemeColors(roomID auth.RoomID) ([]color.Color, error) {
	ColorIDs, err := c.colorRepo.FindColorIDsByRoomID(roomID)

	if err != nil {
		return nil, fmt.Errorf("failed to get color ids: %w", err)
	}
	var ThemeColors []color.Color
	for _, colorID := range ColorIDs {
		color, _ := c.colorRepo.FindThemeColorByColorID(colorID)
		ThemeColors = append(ThemeColors, *color)
	}
	return ThemeColors, nil
}
