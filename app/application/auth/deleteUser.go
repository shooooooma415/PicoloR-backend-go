package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) DeleteUserByUserID(userID auth.UserID) (*auth.UserID, error) {
	deletedUser, err := c.authRepo.DeleteUserByUserID(userID)

	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}
	return &deletedUser.ID, nil
}
