package room

import "picolor-backend/app/domain/auth"

type RoomRepository interface {
	CreateRoom(room Room) (*Room, error)
	CreateRoomMember(user RoomMember) (*RoomMember, error)
	DeleteRoomMemberByUserID(userID auth.UserID) (*RoomMember, error)
	DeleteRoomMembersByRoomID(roomID auth.RoomID) (*auth.RoomID, error)
	// StartRoom(roomId auth.RoomID) (*Room, error)
	// FinishRoom(roomId auth.RoomID) (*Room, error)
	// GetRoom(roomId auth.RoomID) (*Room, error)
	// GetIsStart(roomId auth.RoomID) (*Room, error)
	// GetIsFinish(roomId auth.RoomID) (*Room, error)
	DeleteRoom(roomID auth.RoomID) (*auth.RoomID, error)
}
