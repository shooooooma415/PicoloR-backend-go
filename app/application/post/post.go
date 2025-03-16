package post

import (
	"picolor-backend/app/domain/post"
)

type PostServiceImpl struct {
	postRepo post.PostRepository
}

func NewAuthService(postRepo post.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepo: postRepo,
	}
}
