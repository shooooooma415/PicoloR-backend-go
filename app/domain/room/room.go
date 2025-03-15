package room

import "time"

type RoomID int

type Room struct {
	RoomId RoomID
	IsStart bool
	IsFinish bool
	StartedAT time.Time
}

type RoomMembers struct {
	RoomId RoomID
	UserId int
}