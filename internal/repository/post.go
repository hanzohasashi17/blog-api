package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/hanzohasashi17/blog-api/internal/storage/sqlite"
)

type IPostRepository interface {
	Create(title string, content string, author string) (int64, error)
	GetAll() ([]models.Post, error)
	GetById(id int) (*models.Post, error)
	Update(post models.Post) error
	Delete(id int) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(storage *sqlite.Storage) *postRepository {
	return &postRepository{db: storage.Db}
}

func (r *postRepository) Create(title string, content string, author string) (int64, error) {
	op := "repository.post.Create"

	stmt, err := r.db.Prepare("INSERT INTO posts(title, content, author) VALUES(?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(title, content, author)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *postRepository) GetAll() ([]models.Post, error) {
	op := "repository.post.GetAll"

	rows, err := r.db.Query("SELECT * FROM posts")
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

func (r *postRepository) GetById(id int) (*models.Post, error) {
	op := "repository.post.GetById"

	var post models.Post

	row := r.db.QueryRow("SELECT id, title, content, author, created_at FROM posts WHERE id = ?", id)
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) Update(post models.Post) error {
	op := "repository.post.Update"

	res, err := r.db.Exec("UPDATE posts SET title=?, content=?, author=? WHERE id=?", post.Title, post.Content, post.Author, post.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

func (r *postRepository) Delete(id int) error {
	op := "repository.post.Delete"

	res, err := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return errors.New("post not found")
	}
	
	return nil
}
