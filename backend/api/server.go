package api

import (
	"calvin/kredit/CustomerDigestStaging"
	"calvin/kredit/GenerateSkalaAngsuran"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}
	CustomerDigestStaging := CustomerDigestStaging.NewRepository(s.DB)
	GenerateSkalaAngsuran := GenerateSkalaAngsuran.NewRepository(s.DB)
	c := cron.New()
	c.AddFunc("@every 1m", func() { CustomerDigestStaging.GetCustomer() })
	c.AddFunc("@every 1m", func() { GenerateSkalaAngsuran.GetSkalaAngsuran() })
	c.Start()
	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}

}
