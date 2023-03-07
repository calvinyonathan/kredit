package Login

import (
	"calvin/kredit/model"
	"log"
	"net/http"
)

type Service interface {
	GetLogin(data loginRequest) ([]model.Officer_Tab, int, error)
	UpdatePassword(data updatePasswordRequest) (int, error)
}
type service struct {
	repo LoginRepository
}

func NewService(repo LoginRepository) *service {
	return &service{repo}
}

func (s *service) GetLogin(data loginRequest) ([]model.Officer_Tab, int, error) {

	login, err := s.repo.GetLogin(data.Nik, data.Password)
	if err != nil {
		return []model.Officer_Tab{}, http.StatusInternalServerError, err
	}

	return login, http.StatusOK, nil
}
func (s *service) UpdatePassword(data updatePasswordRequest) (int, error) {
	err := s.repo.UpdatePassword(data.Nik, data.OldPassword, data.Password)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
