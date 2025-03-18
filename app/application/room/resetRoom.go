package room

import (
	"fmt"
	"log"
	"picolor-backend/app/domain/auth"
	"sync"
)

func (c *RoomServiceImpl) ResetRoom(roomID auth.RoomID) (*auth.RoomID, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	var deletedRoomID *auth.RoomID
	var mu sync.Mutex

	go func() {
		defer wg.Done()
		roomID, err := c.roomRepo.DeleteRoomMembersByRoomID(roomID)
		if err != nil {
			log.Printf("Error deleting room members by room ID: %v", err)
			return
		}
		mu.Lock()
		deletedRoomID = roomID
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		roomID, err := c.postRepo.DeletePostByRoomID(roomID)
		if err != nil {
			log.Printf("Error deleting posts by room ID: %v", err)
			return
		}
		mu.Lock()
		deletedRoomID = roomID
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		roomID, err := c.colorRepo.DeleteThemeColors(roomID)
		if err != nil {
			log.Printf("Error deleting theme colors by room ID: %v", err)
			return
		}
		mu.Lock()
		deletedRoomID = roomID
		mu.Unlock()
	}()

	wg.Wait()
	if deletedRoomID == nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", roomID)
	}
	return deletedRoomID, nil
}
