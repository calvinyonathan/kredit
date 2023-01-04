package api

import (
	"calvin/kredit/CustomerDigestStaging"

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

}
