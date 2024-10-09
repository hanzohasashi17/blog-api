package services

import (
	"github.com/hanzohasashi17/blog-api/internal/models"
	"github.com/hanzohasashi17/blog-api/internal/repositories"
)

type IPostService interface {
	CreatePost(title string, content string, author string) (int64, error)
	GetAllPost(page, pageSize int) ([]models.Post, error)
	GetPostById(id int) (*models.Post, error)
	GetPostByAuthor(author string) ([]models.Post, error)
	UpdatePost(post models.Post) error
	DeletePost(id int) error
}

type PostService struct {
	repo repositories.IPostRepository
}

func NewPostService(repo repositories.IPostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(title string, content string, author string) (int64, error) {
	return s.repo.Create(title, content, author)
}

func (s *PostService) GetAllPost(page, pageSize int) ([]models.Post, error) {
	return s.repo.GetAll(page, pageSize)
}

func (s *PostService) GetPostById(id int) (*models.Post, error) {
	return s.repo.GetById(id)
}

func (s *PostService) GetPostByAuthor(author string) ([]models.Post, error) {
	return s.repo.GetByAuthor(author)
}

func (s *PostService) UpdatePost(post models.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.Delete(id)
}
