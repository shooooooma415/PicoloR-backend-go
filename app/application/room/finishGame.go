package room

import (
	"picolor-backend/app/domain/auth"
)

func (r *RoomServiceImpl) FinishGame(roomId auth.RoomID) (*auth.RoomID, error) {
	FinishedRoom, _ := r.roomRepo.UpdateIsFinish(roomId)
	return &FinishedRoom.RoomID, nil
}
