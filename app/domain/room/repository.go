package room

import "picolor-backend/app/domain/auth"

type RoomRepository interface {
	CreateRoom(room Room) (*Room, error)
	CreateRoomMember(user RoomMember) (*RoomMember, error)
	DeleteRoomMembersByRoomID(roomID auth.RoomID) (*auth.RoomID, error)
	UpdateIsStart(roomId auth.RoomID) (*Room, error)
	UpdateIsFinish(roomId auth.RoomID) (*Room, error)
	UpdateIsStartAndIsFinishToFalse(roomId auth.RoomID) (*Room, error)
	CreateStartAt(room Room) (*Room, error)
	DeleteRoom(roomID auth.RoomID) (*auth.RoomID, error)
}
