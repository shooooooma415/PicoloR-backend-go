package room


type RoomRepository interface {
	CreateRoom() (*Room, error)
	CreateRoomMember(user RoomMember) (*RoomMember, error)
	StartRoom(roomId RoomID) (*Room, error)
	FinishRoom(roomId RoomID) (*Room, error)
	GetRoom(roomId RoomID) (*Room, error)
	GetIsStart(roomId RoomID) (*Room, error)
	GetIsFinish(roomId RoomID) (*Room, error)
	DeleteRoom(roomId RoomID) (*Room, error)
}
