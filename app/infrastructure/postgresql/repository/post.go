package repository

import (
	"database/sql"
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
)

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (q *PostRepositoryImpl) DeletePostByRoomID(roomID auth.RoomID) (*auth.RoomID, error) {
	query := `
		DELETE FROM posts
		WHERE room_id = $1
		RETURNING room_id
		`

	var deletedRoomID auth.RoomID

	err := q.db.QueryRow(
		query,
		roomID,
	).Scan(
		&deletedRoomID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to delete posts:%w", err)
	}
	return &deletedRoomID, nil
}

func (q *PostRepositoryImpl) GetPostsByRoomID(roomID auth.RoomID) ([]post.Post, error) {
	query := `
		SELECT user_id, room_id, color_id, rank, image, posted_time 
		FROM posts
		WHERE room_id = $1
		`

	rows, err := q.db.Query(
		query,
		roomID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts:%w", err)
	}
	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var post post.Post
		err := rows.Scan(
			&post.UserID,
			&post.RoomID,
			&post.ColorID,
			&post.Rank,
			&post.Image,
			&post.PostedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan posts:%w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}
