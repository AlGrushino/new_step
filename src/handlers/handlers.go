package handlers

import (
	"step/handlers/users"
	"step/service"

	"github.com/gin-gonic/gin"
)

type UsersHandler interface {
	GetUsername(c *gin.Context)
}

type Handler struct {
	UsersHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UsersHandler: users.NewUsersHandler(service),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		users := api.Group("users")
		{
			users.GET(":user_id", h.GetUsername)
		}
	}
	return router
}
