package room

import (
	"picolor-backend/app/domain/auth"
)

func (c *RoomServiceImpl) DeleteRoom(roomID auth.RoomID) (*auth.RoomID, error) {
	return c.roomRepo.DeleteRoom(roomID)
}
