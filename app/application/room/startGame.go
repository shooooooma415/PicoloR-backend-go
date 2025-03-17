package room

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
	"time"
)

type StartGame struct {
	RoomID  auth.RoomID
	StartAt time.Time
}

func (c *RoomServiceImpl) StartGame(param StartGame) (*room.Room, error) {
	c.roomRepo.UpdateIsStart(param.RoomID)

	startedRoom, _ := c.roomRepo.CreateStartAt(room.Room{
		RoomID:  param.RoomID,
		StartAt: param.StartAt,
	})
	return startedRoom, nil
}
