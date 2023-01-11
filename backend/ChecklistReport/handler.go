package ChecklistReport

import (
	"fmt"
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
func (h *Handler) GetBranch(c *gin.Context) {
	Branch, status, err := h.Service.GetBranch()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    Branch,
	})
}
func (h *Handler) GetCompany(c *gin.Context) {
	Company, status, err := h.Service.GetCompany()
	if err != nil {
		log.Println("Error handler Get : ", err)
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    Company,
	})
}
func (h *Handler) UpdateCustomer(c *gin.Context) {

	ppk := c.Query("ppk")
	if ppk == "" {
		messageErr := []string{"Param data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	reqFix := PpkRequest{
		Ppk: ppk,
	}
	customer, status, err := h.Service.UpdateCustomer(reqFix)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
			"code":    "99",
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "Update Success",
		"data":    customer,
		"code":    "00",
	})
	fmt.Println(customer)
}

func (h *Handler) SearchChecklistReport(c *gin.Context) {
	branch := c.Query("branch")
	company := c.Query("company")
	startdate := c.Query("startdate")
	enddate := c.Query("enddate")
	req := GetSearchRequest{Branch: branch, Company: company, StartDate: startdate, EndDate: enddate}
	user, status, err := h.Service.SearchChecklistReport(req)
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
