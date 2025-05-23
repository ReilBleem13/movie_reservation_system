package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFilmsHandler(c *gin.Context) {
	films, err := h.services.Clientage.GetFilms()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"films": films,
	})
}

func (h *Handler) GetAvailableSeatsHandler(c *gin.Context) {
	filmSessionID := c.Param("id")
	filmSessionIDInt, err := strconv.Atoi(filmSessionID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid film session id" + err.Error()})
		return
	}

	seats, err := h.services.Clientage.GetAvailableSeats(filmSessionIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"seats": seats})
}

func (h *Handler) GetReservationHandler(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	seatID := c.Query("seat")
	seatIDInt, err := strconv.Atoi(seatID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	filmSession := c.Query("film_session")
	filmSessionIDInt, err := strconv.Atoi(filmSession)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	reservations, err := h.services.Clientage.ReserveSeat(userID, filmSessionIDInt, seatIDInt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"reservations": reservations})
}

func (h *Handler) CancelReservation(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	seatID := c.Query("seat")
	seatIDInt, err := strconv.Atoi(seatID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	filmSession := c.Query("film_session")
	filmSessionIDInt, err := strconv.Atoi(filmSession)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Clientage.CancelReservation(userID, filmSessionIDInt, seatIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "no reservation with userID") {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "deleted"})
}
