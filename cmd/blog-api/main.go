package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hanzohasashi17/blog-api/internal/config"
	"github.com/hanzohasashi17/blog-api/internal/handlers"
	"github.com/hanzohasashi17/blog-api/internal/repositories"
	"github.com/hanzohasashi17/blog-api/internal/services"
	"github.com/hanzohasashi17/blog-api/internal/database/sqlite"
	"github.com/hanzohasashi17/blog-api/lib/logger/sl"
)

func main() {
	// run config
	cfg := config.MustLoad()

	// setup logger
	log := sl.SetupLogger()

	// run database && migration
	db, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	defer db.Db.Close()

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", handler.CreatePostHandler(postService))
		r.Get("/", handler.GetAllPostHandler(postService))
		r.Get("/{id}", handler.GetPostByIdHandler(postService))
		r.Get("/{author}", handler.GetPostByAuthorHandler(postService))
		r.Put("/", handler.UpdatePostHandler(postService))
		r.Delete("/{id}", handler.DeletePostHandler(postService))
	})

	log.Info("Server run on 8080 port...")
	http.ListenAndServe(":8080", r)
}
