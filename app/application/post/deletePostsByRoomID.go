package post

import "picolor-backend/app/domain/auth"

func (c *PostServiceImpl) DeletePostsByRoomID(roomID auth.RoomID) (*auth.RoomID, error) {
	return c.postRepo.DeletePostByRoomID(roomID)
}
