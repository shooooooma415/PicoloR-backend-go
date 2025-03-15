package room

type RoomId int

type Room struct {
	RoomId RoomId
	IsStart bool
	IsFinish bool
}

type RoomMembers struct {
	RoomId RoomId
	UserId int
}