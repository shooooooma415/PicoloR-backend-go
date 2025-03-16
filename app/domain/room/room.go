package room

import (
	"picolor-backend/app/domain/auth"
	"time"
)


type Room struct {
	RoomID    auth.RoomID
	IsStart   bool
	IsFinish  bool
	StartedAT time.Time
}

type RoomMember struct {
	RoomID auth.RoomID
	UserID auth.UserID
}
