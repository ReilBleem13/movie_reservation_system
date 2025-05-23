package service

import (
	"mrs_project/pkg/models"
	"mrs_project/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CheckUser(email, password string) (int, string, error)
}

type Clientage interface {
	GetFilms() ([]models.AboutFilm, error)
	GetAvailableSeats(filmSessionID int) ([]models.FreeSeat, error)
	ReserveSeat(userID, filmSessionID, seatID int) (models.Reservation, error)
	CancelReservation(userID, filmSessionID, seatID int) error
}

type Admin interface {
	AddFilm(film models.Film) error
	DeleteFilm(filmID int) error
	UpdateFilm(filmID int, update models.Film) error
}

type Service struct {
	Authorization
	Clientage
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Clientage:     NewClientageService(repos.Clientage),
		Admin:         NewAdminService(repos.Admin),
	}
}
