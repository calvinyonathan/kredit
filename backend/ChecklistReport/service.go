package ChecklistReport

import (
	"calvin/kredit/model"
	"log"
	"net/http"
)

type Service interface {
	SearchChecklistReport(data GetSearchRequest) ([]response, int, error)
	GetBranch() ([]model.Branch_Tabs, int, error)
	GetCompany() ([]model.Mst_Company_Tabs, int, error)
	UpdateCustomer(req []PpkRequest) (int, error)
}
type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s *service) GetBranch() ([]model.Branch_Tabs, int, error) {
	Branch, err := s.repo.GetBranch()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return Branch, http.StatusOK, nil
}
func (s *service) GetCompany() ([]model.Mst_Company_Tabs, int, error) {
	Company, err := s.repo.GetCompany()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return Company, http.StatusOK, nil
}
func (s *service) GetPpk(data PpkRequest) (model.Customer_Data_Tabs, int, error) {

	link, err := s.repo.GetPpk(data.Ppk)
	if err != nil {
		return model.Customer_Data_Tabs{}, http.StatusInternalServerError, err
	}

	return link, http.StatusOK, nil
}
func (s *service) UpdateCustomer(req []PpkRequest) (int, error) {
	err := s.repo.UpdateCustomer(req)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
func (s *service) SearchChecklistReport(data GetSearchRequest) ([]response, int, error) {
	user, err := s.repo.SearchChecklistReport(data.Branch, data.Company, data.StartDate, data.EndDate)
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
