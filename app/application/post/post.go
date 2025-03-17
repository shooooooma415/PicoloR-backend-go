package post

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
	"picolor-backend/app/domain/post"
)

type PostServiceImpl struct {
	postRepo  post.PostRepository
	authRepo  auth.AuthRepository
	colorRepo color.ColorRepository
}

func NewPostService(postRepo post.PostRepository, authRepo auth.AuthRepository, colorRepo color.ColorRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepo:  postRepo,
		authRepo:  authRepo,
		colorRepo: colorRepo,
	}
}
