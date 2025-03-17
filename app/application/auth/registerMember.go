package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)

type RegisterMember struct {
	RoomID auth.RoomID
	UserID auth.UserID
}

func (c *AuthServiceImpl) RegisterMember(registerInfo RegisterMember) (*auth.UserID, error) {


	joinRoom := room.RoomMember{
		RoomID: registerInfo.RoomID,
		UserID: registerInfo.UserID,
	}

	roomMember, err := c.roomRepo.CreateRoomMember(joinRoom)

	if err != nil {
		return nil, fmt.Errorf("failed to create room member: %w", err)
	}

	return &roomMember.UserID, nil
}
