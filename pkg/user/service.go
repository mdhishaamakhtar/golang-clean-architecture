package user

import (
	"github.com/mdhishaamakhtar/learnFiber/pkg"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	AddUser(user *models.User) error

	Login(email string, password string) (*models.User, error)

	GetUserDetailsById(id string) (*models.User, error)

	GetUserDetailsByEmail(email string) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s service) AddUser(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return pkg.ErrDatabase
	}
	user.Password = string(hash)
	return s.repo.AddUser(user)
}

func (s service) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserDetailsByEmail(email)
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(password, user.Password) {
		return user, nil
	}
	return nil, pkg.ErrNotFound
}

func (s service) GetUserDetailsById(id string) (*models.User, error) {
	return s.repo.GetUserDetailsById(id)
}

func (s service) GetUserDetailsByEmail(email string) (*models.User, error) {
	return s.repo.GetUserDetailsByEmail(email)
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
