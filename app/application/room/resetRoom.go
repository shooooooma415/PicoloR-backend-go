package room

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *RoomServiceImpl) ResetRoom(roomID auth.RoomID) (*auth.RoomID, error) {
	var deletedRoomID *auth.RoomID

	tempRoomID, err := c.roomRepo.DeleteRoomMembersByRoomID(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", err)
	}
	roomID = *tempRoomID
	deletedRoomID = tempRoomID

	tempPostRoomID, err := c.postRepo.DeletePostByRoomID(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", err)
	}
	roomID = *tempPostRoomID
	deletedRoomID = tempPostRoomID

	_, err = c.colorRepo.DeleteThemeColors(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", err)
	}

	if deletedRoomID == nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", roomID)
	}

	_, err = c.roomRepo.UpdateIsStartAndIsFinishToFalse(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to reset room with ID: %v", err)
	}
	return deletedRoomID, nil
}
