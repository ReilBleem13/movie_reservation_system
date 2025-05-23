package handler

import (
	"mrs_project/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	adminGroup := router.Group("/admin")
	adminGroup.Use(h.AuthMiddleWare("admin"))
	{
		adminGroup.GET("/add", h.AddFilmHandler)
		adminGroup.DELETE("/delete/:id", h.DeleteFilmHandler)
		adminGroup.PUT("/update/:id", h.UpdateFilmHandler)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	film := router.Group("/film")
	film.Use(h.AuthMiddleWare(""))
	{
		film.GET("/list", h.GetFilmsHandler)
		film.GET("/seats/:id", h.GetAvailableSeatsHandler)
		film.GET("/reservations", h.GetReservationHandler)
		film.DELETE("/reservations/cancel", h.CancelReservation)

	}

	return router
}
