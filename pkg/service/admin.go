package service

import (
	"mrs_project/pkg/models"
	"mrs_project/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (a *AdminService) AddFilm(film models.Film) error {
	return a.repo.AddFilm(film)
}

func (a *AdminService) DeleteFilm(filmID int) error {
	return a.repo.DeleteFilm(filmID)
}

func (a *AdminService) UpdateFilm(filmID int, update models.Film) error {
	return a.repo.UpdateFilm(filmID, update)
}
