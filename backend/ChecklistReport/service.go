package ChecklistReport

import (
	"log"
	"net/http"
)

type Service interface {
	GetChecklistReport() ([]response, int, error)
}
type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}
func (s *service) GetChecklistReport() ([]response, int, error) {
	user, err := s.repo.GetChecklistReport()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
