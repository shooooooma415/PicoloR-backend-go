package room

import "time"

type RoomID int

type Room struct {
	RoomID RoomID
	IsStart bool
	IsFinish bool
	StartedAT time.Time
}

type RoomMember struct {
	RoomID RoomID
	UserID int
}