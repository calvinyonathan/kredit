package DrawdownReport

import (
	"log"
	"net/http"
)

type Service interface {
	GetDrawdownReport() ([]response, int, error)
}
type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s *service) GetDrawdownReport() ([]response, int, error) {
	user, err := s.repo.GetDrawdownReport()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
