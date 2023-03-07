package api

import (
	"calvin/kredit/ChecklistReport"
	"calvin/kredit/CustomerDigestStaging"
	"calvin/kredit/DrawdownReport"
	"calvin/kredit/GenerateSkalaAngsuran"
	"calvin/kredit/Login"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	loginRepo := Login.NewRepository(s.DB)
	loginService := Login.NewService(loginRepo)
	loginHandler := Login.NewHandler(loginService)
	s.Router.GET("/login", loginHandler.GetLogin)
	s.Router.PUT("/updatePassword", loginHandler.UpdatePassword)

	customerRepo := CustomerDigestStaging.NewRepository(s.DB)
	customerService := CustomerDigestStaging.NewService(customerRepo)
	customerHandler := CustomerDigestStaging.NewHandler(customerService)
	s.Router.GET("/", customerHandler.GetCustomer)

	checklistRepo := ChecklistReport.NewRepository(s.DB)
	checklistService := ChecklistReport.NewService(checklistRepo)
	checkListHandler := ChecklistReport.NewHandler(checklistService)
	s.Router.GET("/checklistBranch", checkListHandler.SearchChecklistReport)
	s.Router.PUT("/approve", checkListHandler.UpdateCustomer)
	s.Router.GET("/getbranch", checkListHandler.GetBranch)
	s.Router.GET("/getcompany", checkListHandler.GetCompany)

	GetSkalaAngsuranRepo := GenerateSkalaAngsuran.NewRepository(s.DB)
	GetSkalaAngsuranService := GenerateSkalaAngsuran.NewService(GetSkalaAngsuranRepo)
	GetSkalaAngsuranHandler := GenerateSkalaAngsuran.NewHandler(GetSkalaAngsuranService)
	s.Router.GET("/skala", GetSkalaAngsuranHandler.GetSkalaAngsuran)

	drawdownReportRepo := DrawdownReport.NewRepository(s.DB)
	drawdownReportService := DrawdownReport.NewService(drawdownReportRepo)
	drawdownReportHandler := DrawdownReport.NewHandler(drawdownReportService)
	s.Router.GET("/drawdown", drawdownReportHandler.GetDrawdownReport)

}
