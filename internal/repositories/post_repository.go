package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IPostRepository interface {
	Create(title string, content string, author string) (int64, error)
	GetAll(page, pageSize int) ([]models.Post, error)
	GetById(id int) (*models.Post, error)
	GetByAuthor(author string) ([]models.Post, error)
	Update(post models.Post) error
	Delete(id int) error
}

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *postRepository {
	return &postRepository{db: db}
}

// NEW POST +
func (r *postRepository) Create(title string, content string, author string) (int64, error) {
	op := "repositories.post.Create"

	var newPostId int64

	err := r.db.QueryRow(context.Background(), "INSERT INTO posts(title, content, author) VALUES($1, $2, $3) RETURNING id", title, content, author).Scan(&newPostId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return newPostId, nil
}

// GET ALL POST +
func (r *postRepository) GetAll(page, pageSize int) ([]models.Post, error) {
	op := "repositories.post.GetAll"

	offset := (page - 1) * pageSize

	rows, err := r.db.Query(context.Background(), "SELECT * FROM posts LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

// GET POST BY ID +
func (r *postRepository) GetById(id int) (*models.Post, error) {
	op := "repositories.post.GetById"

	var post models.Post

	row := r.db.QueryRow(context.Background(), "SELECT id, title, content, author, created_at FROM posts WHERE id = $1", id)
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, err
	}

	return &post, nil
}

// GET POST BY AUTHOR +
func (r *postRepository) GetByAuthor(author string) ([]models.Post, error) {
	op := "repositories.post.GetByAuthor"

	rows, err := r.db.Query(context.Background(), "SELECT * FROM posts WHERE author = $1", author)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

// UPDATE POST +
func (r *postRepository) Update(post models.Post) error {
	op := "repositories.post.Update"


	_, err := r.db.Exec(context.Background(), "UPDATE posts SET title=$1, content=$2, author=$3 WHERE id=$4 RETURNING id", post.Title, post.Content, post.Author, post.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("%s, post with id not found: %d", op, post.Id)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// DELETE POST BY ID +
func (r *postRepository) Delete(id int) error {
	op := "repositories.post.Delete"

	_, err := r.db.Exec(context.Background(), "DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("%s, post with id not found: %d", op, id)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	
	return nil
}
