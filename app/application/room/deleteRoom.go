package room

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *RoomServiceImpl) DeleteRoom(roomID auth.RoomID) (*auth.RoomID, error) {

	DeletedRoomID, err := c.roomRepo.DeleteRoom(roomID)

	if err != nil {
		return nil, fmt.Errorf("failed to delete room:%w", err)
	}
	return DeletedRoomID, nil
}
