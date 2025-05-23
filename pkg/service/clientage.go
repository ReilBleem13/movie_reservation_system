package service

import (
	"mrs_project/pkg/models"
	"mrs_project/pkg/repository"
)

type ClientageService struct {
	repo repository.Clientage
}

func NewClientageService(repo repository.Clientage) *ClientageService {
	return &ClientageService{
		repo: repo,
	}
}

func (c *ClientageService) GetFilms() ([]models.AboutFilm, error) {
	return c.repo.GetFilms()
}

func (c *ClientageService) GetAvailableSeats(filmSessionID int) ([]models.FreeSeat, error) {
	return c.repo.GetAvailableSeats(filmSessionID)
}

func (c *ClientageService) ReserveSeat(userID, filmSessionID, seatID int) (models.Reservation, error) {
	return c.repo.ReserveSeat(userID, filmSessionID, seatID)
}

func (c *ClientageService) CancelReservation(userID, filmSessionID, seatID int) error {
	return c.repo.DeleteReservation(userID, filmSessionID, seatID)
}
