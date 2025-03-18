package room

import (
	"fmt"
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

	startedRoom, err := c.roomRepo.CreateStartAt(room.Room{
		RoomID:  param.RoomID,
		StartAt: param.StartAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start game: %w", err)
	}
	return startedRoom, nil
}
