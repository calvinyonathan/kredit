package CustomerDigestStaging

import (
	"log"
	"net/http"

	"calvin/kredit/model"
)

type Service interface {
	GetCustomer() ([]model.Staging_Customers, int, error)
}
type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}
func (s *service) GetCustomer() ([]model.Staging_Customers, int, error) {
	user, err := s.repo.GetCustomer()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
