package Login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
func (h *Handler) GetLogin(c *gin.Context) {
	nik := c.Query("nik")
	password := c.Query("password")
	req := loginRequest{Nik: nik, Password: password}
	user, status, err := h.Service.GetLogin(req)
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    user,
	})
}
func (h *Handler) UpdatePassword(c *gin.Context) {
	var req updatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input data not suitable"})
		return
	}
	status, err := h.Service.UpdatePassword(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": "Bad Request",
			"code":    "99",
		})
	}

	c.JSON(status, gin.H{
		"message": "success",
		"code":    "00",
	})
}
