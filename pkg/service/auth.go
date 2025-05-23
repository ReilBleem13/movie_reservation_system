package service

import (
	"mrs_project/pkg/models"
	"mrs_project/pkg/repository"
	"mrs_project/pkg/utils"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) CheckUser(email, password string) (int, string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		return 0, "", err
	}

	if utils.CheckPasswordHash(password, user.Password) {
		return 0, "", err
	}
	return user.ID, user.Role, nil
}
