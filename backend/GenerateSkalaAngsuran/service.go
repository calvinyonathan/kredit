package GenerateSkalaAngsuran

import (
	"calvin/kredit/model"
	"log"
	"net/http"
)

type Service interface {
	GetSkalaAngsuran() ([]model.Customer_Data_Tabs, int, error)
}
type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}
func (s *service) GetSkalaAngsuran() ([]model.Customer_Data_Tabs, int, error) {
	user, err := s.repo.GetSkalaAngsuran()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
