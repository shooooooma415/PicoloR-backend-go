package room

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *RoomServiceImpl) CreateRoom() (*auth.RoomID, error) {
	room, err := c.roomRepo.CreateRoom()

	if err != nil {
		return nil, fmt.Errorf("failed to create room:%w", err)
	}
	return &room.RoomID, nil
}
