package DrawdownReport

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetDrawdownReport(c *gin.Context) {
	branch := c.Query("branch")
	company := c.Query("company")
	startdate := c.Query("startdate")
	enddate := c.Query("enddate")
	req := GetSearchRequest{Branch: branch, Company: company, StartDate: startdate, EndDate: enddate}
	user, status, err := h.Service.GetDrawdownReport(req)
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
