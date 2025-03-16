package auth

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)

type AuthServiceImpl struct {
	authRepo auth.AuthRepository
	roomRepo room.RoomRepository
}

func NewAuthService(authRepo auth.AuthRepository, roomRepo room.RoomRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepo: authRepo,
		roomRepo: roomRepo,
	}
}
