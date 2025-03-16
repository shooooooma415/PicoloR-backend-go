package room

import (
	"picolor-backend/app/domain/auth"
	"sync"
)

func (c *RoomServiceImpl) ResetRoom(roomID auth.RoomID) (*auth.RoomID, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var deletedRoomID *auth.RoomID

	go func() {
		defer wg.Done()
		deletedRoomID, _ = c.roomRepo.DeleteRoomMembersByRoomID(roomID)
	}()

	go func() {
		defer wg.Done()
		deletedRoomID, _ = c.postRepo.DeletePostByRoomID(roomID)
	}()

	wg.Wait()
	return deletedRoomID, nil
}
