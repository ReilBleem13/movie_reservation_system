package repository

import (
	"mrs_project/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email string) (models.UserWithRole, error)
}

type Clientage interface {
	GetFilms() ([]models.AboutFilm, error)
	GetAvailableSeats(filmSessionID int) ([]models.FreeSeat, error)
	ReserveSeat(userID, filmSessionID, seatID int) (models.Reservation, error)
	DeleteReservation(userID, filmSessionID, seatID int) error
}

type Admin interface {
	AddFilm(film models.Film) error
	DeleteFilm(filmID int) error
	UpdateFilm(filmID int, update models.Film) error
}

type Repository struct {
	Authorization
	Clientage
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Clientage:     NewClientagePostgres(db),
		Admin:         NewAdminPostgres(db),
	}
}
