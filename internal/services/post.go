package services

import (
	"github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/hanzohasashi17/blog-api/internal/repository"
)

type IPostService interface {
	CreatePost(title string, content string, author string) (int64, error)
	GetAllPost() ([]models.Post, error)
	GetPostById(id int) (*models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id int) error
}

type PostService struct {
	repo repository.IPostRepository
}

func NewPostService(repo repository.IPostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(title string, content string, author string) (int64, error) {
	return s.repo.Create(title, content, author)
}

func (s *PostService) GetAllPost() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetPostById(id int) (*models.Post, error) {
	return s.repo.GetById(id)
}

func (s *PostService) UpdatePost(post models.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}
