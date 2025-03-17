package color

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
)

func (c *ColorServiceImpl) GetThemeColor(roomID auth.RoomID) ([]color.ColorCode, error) {
	colors, err := c.colorRepo.GetThemeColorsByRoomID(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get theme colors: %w", err)
	}
	themeColors := make([]color.ColorCode, len(colors))

	for i, col := range colors {
		themeColors[i] = col.ColorCode
	}
	return themeColors, nil
}
