package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hanzohasashi17/blog-api/internal/models"
)

type IPostService interface {
	CreatePost(title string, content string, author string) (int64, error)
	// GetAll() ([]models.Post, error)
	// GetById(id int) (*models.Post, error)
	// Update(id int) error
	// Delete(id int) error
}

func CreatePostHandler(s IPostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPost models.Post
		if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		postId, err := s.CreatePost(newPost.Title, newPost.Content, newPost.Author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(postId)
	}
}