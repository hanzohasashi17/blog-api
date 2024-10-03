package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hanzohasashi17/blog-api/internal/models"
)

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