package ChecklistReport

import (
	"calvin/kredit/model"
	"log"
	"net/http"
)

type Service interface {
	GetChecklistReport() ([]response, int, error)
	GetBranch() ([]model.Branch_Tabs, int, error)
	GetCompany() ([]model.Mst_Company_Tabs, int, error)
	UpdateCustomer(customer PpkRequest) (model.Customer_Data_Tabs, int, error)
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
func (s *service) UpdateCustomer(customer PpkRequest) (model.Customer_Data_Tabs, int, error) {
	Customers := model.Customer_Data_Tabs{
		Approval_Status: "0",
		PPK:             customer.Ppk,
	}
	res, err := s.repo.UpdateCustomer(Customers)
	if err != nil {
		return model.Customer_Data_Tabs{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}

func (s *service) GetChecklistReport() ([]response, int, error) {
	user, err := s.repo.GetChecklistReport()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}
