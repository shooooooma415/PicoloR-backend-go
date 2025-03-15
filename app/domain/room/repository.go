package room

type RoomRepository interface {
	CreateRoom() (*Room, error)
	GetRoom(roomId RoomID) (*Room, error)
	GetIsStart(roomId RoomID) (*Room, error)
	GetIsFinish(roomId RoomID) (*Room, error)
	DeleteRoom(roomId RoomID) (*Room, error)
}