package user

import (
	"github.com/mdhishaamakhtar/learnFiber/pkg"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"gorm.io/gorm"
)

type Repository interface {
	AddUser(user *models.User) error

	GetUserDetailsById(id string) (*models.User, error)

	GetUserDetailsByEmail(email string) (*models.User, error)
}

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r repo) AddUser(user *models.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}

func (r repo) GetUserDetailsById(id string) (*models.User, error) {
	var userDetails models.User
	err := r.DB.Where("id=?", id).Find(&userDetails).Error
	if err != nil {
		return nil, err
	}
	userDetails.Password = ""
	return &userDetails, nil
}

func (r repo) GetUserDetailsByEmail(email string) (*models.User, error) {
	var userDetails models.User
	err := r.DB.Where("email=?", email).Find(&userDetails).Error
	if err != nil {
		return nil, err
	}
	return &userDetails, nil
}
