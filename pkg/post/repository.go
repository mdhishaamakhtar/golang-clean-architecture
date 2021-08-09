package post

import (
	"github.com/mdhishaamakhtar/learnFiber/pkg"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	AddPost(post *models.Post) error

	GetAllPosts(userId string) (*[]models.PostDetails, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r repo) AddPost(post *models.Post) error {
	err := r.DB.Create(post).Error
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}

func (r repo) GetAllPosts(userId string) (*[]models.PostDetails, error) {
	var posts []models.PostDetails
	err := r.DB.Model(&models.Post{}).Select("id, created_at, updated_at, title, description, user_id").Where("user_id=?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
