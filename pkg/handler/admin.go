package handler

import (
	"mrs_project/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddFilmHandler(c *gin.Context) {
	var film models.Film
	if err := c.BindJSON(&film); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Admin.AddFilm(film); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "added"})
}

func (h *Handler) DeleteFilmHandler(c *gin.Context) {
	filmID_param := c.Param("id")

	filmID, err := strconv.Atoi(filmID_param)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Admin.DeleteFilm(filmID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "deleted"})
}

func (h *Handler) UpdateFilmHandler(c *gin.Context) {
	filmID_param := c.Param("id")

	filmID, err := strconv.Atoi(filmID_param)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var film models.Film
	if err := c.BindJSON(&film); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Admin.UpdateFilm(filmID, film); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "updated"})
}
