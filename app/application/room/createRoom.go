package room

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)

func (c *RoomServiceImpl) CreateRoom() (*auth.RoomID, error) {
	CreateRoom := room.Room{
		IsStart:  false,
		IsFinish: false,
	}

	CreatedRoom, err := c.roomRepo.CreateRoom(CreateRoom)

	if err != nil {
		return nil, fmt.Errorf("failed to create room:%w", err)
	}
	return &CreatedRoom.RoomID, nil
}
