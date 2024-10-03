package repository

import (
	"database/sql"
	"fmt"

	// "github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/hanzohasashi17/blog-api/internal/storage/sqlite"
)

type IPostRepository interface {
	Create(title string, content string, author string) (int64, error)
	// GetAll() ([]models.Post, error)
	// GetById(id int) (*models.Post, error)
	// Update(id int) error
	// Delete(id int) error
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(storage *sqlite.Storage) *PostRepository {
	return &PostRepository{db: storage.Db}
}

func (r *PostRepository) Create(title string, content string, author string) (int64, error) {
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

// func (r *PostRepository) GetAll() ([]models.Post, error) {
// 	op := "repository.post.GetAll"

// 	posts, err := r.db.Query("SELECT * FROM posts")
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", op, err)
// 	}
// 	defer posts.Close()

// 	posts.
// 	if err != nil {
// 		return nil, fmt.Errorf("%s: %w", op, err)
// 	}

// }
