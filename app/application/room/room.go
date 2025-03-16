package room

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)

type RoomServiceImpl struct {
	authRepo auth.AuthRepository
	roomRepo room.RoomRepository
}

func NewRoomService(authRepo auth.AuthRepository, roomRepo room.RoomRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		authRepo: authRepo,
		roomRepo: roomRepo,
	}
}
