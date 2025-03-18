package room

import (
	"fmt"
	"log"
	"picolor-backend/app/domain/auth"
)

func (c *RoomServiceImpl) ResetRoom(roomID auth.RoomID) (*auth.RoomID, error) {
	var deletedRoomID *auth.RoomID

	tempRoomID, err := c.roomRepo.DeleteRoomMembersByRoomID(roomID)
	if err != nil {
		log.Printf("Error deleting room members by room ID: %v", err)
		return nil, err
	}
	roomID = *tempRoomID
	deletedRoomID = tempRoomID

	tempPostRoomID, err := c.colorRepo.DeleteThemeColors(roomID)
	if err != nil {
		log.Printf("Error deleting posts by room ID: %v", err)
		return nil, err
	}
	roomID = *tempPostRoomID
	deletedRoomID = tempPostRoomID

	_, err = c.postRepo.DeletePostByRoomID(roomID)
	if err != nil {
		log.Printf("Error deleting theme colors by room ID: %v", err)
		return nil, err
	}

	if deletedRoomID == nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", roomID)
	}
	return deletedRoomID, nil
}
