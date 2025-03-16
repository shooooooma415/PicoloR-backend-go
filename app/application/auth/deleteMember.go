package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) DeleteMember(deleteInfo auth.DeleteUser) (*auth.UserID, error) {

	deletedUser, err := c.authRepo.DeleteUser(deleteInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to delete user:%w", err)
	}

	deleteMember, err := c.roomRepo.DeleteRoomMember(deletedUser.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to delete room member:%w", err)
	}
	return &deleteMember.UserID, nil
}
