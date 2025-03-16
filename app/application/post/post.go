package post

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
)

type PostServiceImpl struct {
	postRepo post.PostRepository
	authRepo auth.AuthRepository
}

func NewAuthService(postRepo post.PostRepository, authRepo auth.AuthRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepo: postRepo,
		authRepo: authRepo,
	}
}
