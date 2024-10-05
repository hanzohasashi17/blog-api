package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/hanzohasashi17/blog-api/internal/services"
)

// type PostHandler struct {
// 	service services.IPostService
// }

func CreatePostHandler(s services.IPostService) http.HandlerFunc {
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

func GetAllPostHandler(s services.IPostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		pageSizeStr := r.URL.Query().Get("page_size")
		if pageSizeStr == "" {
			pageSizeStr = "10"
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 0 {
			http.Error(w, "Invalid page size", http.StatusBadRequest)
			return
		}

		posts, err := s.GetAllPost(page, pageSize)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(posts)
	}
}

func GetPostByIdHandler(s services.IPostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
		}

		post, err := s.GetPostById(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
                http.Error(w, "Post not found", http.StatusNotFound)
            } else {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostHandler(s services.IPostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post models.Post

		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
            return
		}

		if err := s.UpdatePost(post); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
                http.Error(w, "Post not found", http.StatusNotFound)
            } else {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
		}

		json.NewEncoder(w).Encode(post)
	}
}

func DeletePostHandler(s services.IPostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
		}

		if err := s.DeletePost(id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
                http.Error(w, "Post not found", http.StatusNotFound)
            } else {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}