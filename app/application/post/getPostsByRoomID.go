package post

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
)

func (s *PostServiceImpl) GetPostsByRoomID(roomID auth.RoomID) ([]post.GetPost, error) {
	posts, err := s.postRepo.FindPostsByRoomID(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	var getPosts []post.GetPost

	for _, p := range posts {
		user, _ := s.authRepo.FindUserByUserID(p.UserID)
		color, _ := s.colorRepo.FindThemeColorByColorID(p.ColorID)
		getPosts = append(getPosts, post.GetPost{
			UserName:   user.Name,
			Rank:       p.Rank,
			RoomID:     p.RoomID,
			Color:      color.ColorCode,
			Image:      p.Image,
			PostedTime: p.PostedTime,
		})
	}
	return getPosts, nil
}
