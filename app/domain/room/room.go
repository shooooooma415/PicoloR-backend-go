package room

type RoomID int

type Room struct {
	RoomId RoomID
	IsStart bool
	IsFinish bool
}

type RoomMembers struct {
	RoomId RoomID
	UserId int
}