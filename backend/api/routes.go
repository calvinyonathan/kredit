package api

import (
	"calvin/kredit/ChecklistReport"
	"calvin/kredit/CustomerDigestStaging"
	"calvin/kredit/GenerateSkalaAngsuran"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	customerRepo := CustomerDigestStaging.NewRepository(s.DB)
	customerService := CustomerDigestStaging.NewService(customerRepo)
	customerHandler := CustomerDigestStaging.NewHandler(customerService)
	s.Router.GET("/", customerHandler.GetCustomer)

	checklistRepo := ChecklistReport.NewRepository(s.DB)
	checklistService := ChecklistReport.NewService(checklistRepo)
	checkListHandler := ChecklistReport.NewHandler(checklistService)
	s.Router.GET("/checklist", checkListHandler.GetChecklistReport)

	GetSkalaAngsuranRepo := GenerateSkalaAngsuran.NewRepository(s.DB)
	GetSkalaAngsuranService := GenerateSkalaAngsuran.NewService(GetSkalaAngsuranRepo)
	GetSkalaAngsuranHandler := GenerateSkalaAngsuran.NewHandler(GetSkalaAngsuranService)
	s.Router.GET("/skala", GetSkalaAngsuranHandler.GetSkalaAngsuran)
}
