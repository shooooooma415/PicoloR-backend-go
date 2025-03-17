package color

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *ColorServiceImpl) GetThemeColor(roomID auth.RoomID) ([]auth.ColorCode, error) {
	colors, err := c.colorRepo.GetThemeColorsByRoomID(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get theme colors: %w", err)
	}
	themeColors := make([]auth.ColorCode, len(colors))

	for i, col := range colors {
		themeColors[i] = col.ColorCode
	}
	return themeColors, nil
}
