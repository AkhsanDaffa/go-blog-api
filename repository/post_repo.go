package repository

import (
	"database/sql"
	"go-blog-api/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(title string, content string, userID int) (models.Post, error) {
	var post models.Post

	query := `
	INSERT INTO posts (title, content, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, title, content, created_at, user_id
	`

	err := r.DB.QueryRow(query, title, content, userID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UserID,
	)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	query := `SELECT id, title, content, created_at, user_id FROM posts ORDER BY created_at DESC`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UserID); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetPostByUserID(userID int) ([]models.Post, error) {
	var posts []models.Post

	query := `
	SELECT id, title, content, created_at, user_id
	FROM posts
	WHERE user_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UserID,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
