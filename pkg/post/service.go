package post

import "github.com/mdhishaamakhtar/learnFiber/pkg/models"

type Service interface {
	AddPost(post *models.Post) error

	GetAllPosts(userId string) (*[]models.PostDetails, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s service) AddPost(post *models.Post) error {
	return s.repo.AddPost(post)
}

func (s service) GetAllPosts(userId string) (*[]models.PostDetails, error) {
	return s.repo.GetAllPosts(userId)
}
