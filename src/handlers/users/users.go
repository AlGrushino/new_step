package users

import (
	"net/http"
	"step/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *service.Service
}

func NewUsersHandler(service *service.Service) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) GetUsername(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request, it must be like: get_username/user_id=123"})
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	username, err := h.service.GetUsername(c, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": username})
}
