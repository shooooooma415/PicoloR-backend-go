package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)

func (c *AuthServiceImpl) RegisterMember(joinRoom room.RoomMember) (*auth.UserID, error) {

	roomMember, err := c.roomRepo.CreateRoomMember(joinRoom)

	if err != nil {
		return nil, fmt.Errorf("failed to create room member: %w", err)
	}

	return &roomMember.UserID, nil
}
