package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) CreateUser(userName auth.UserName) (*auth.UserID, error) {
	createdUser, err := c.authRepo.CreateUser(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &createdUser.ID, nil
}
