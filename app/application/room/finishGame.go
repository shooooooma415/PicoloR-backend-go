package room

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (r *RoomServiceImpl) FinishGame(roomId auth.RoomID) (*auth.RoomID, error) {
	FinishedRoom, err := r.roomRepo.UpdateIsFinish(roomId)
	if err != nil {
		return nil, fmt.Errorf("failed to finish game: %w", err)
	}
	return &FinishedRoom.RoomID, nil
}
