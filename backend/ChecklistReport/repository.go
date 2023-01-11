package ChecklistReport

import (
	"calvin/kredit/model"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	SearchChecklistReport(branch string, company string, startDate string, endDate string) ([]response, error)
	GetPpk(ppk string) (model.Customer_Data_Tabs, error)
	UpdateCustomer(customer model.Customer_Data_Tabs) (model.Customer_Data_Tabs, error)
	GetBranch() ([]model.Branch_Tabs, error)
	GetCompany() ([]model.Mst_Company_Tabs, error)
}
type repository struct {
	db *gorm.DB
}
type response struct {
	RowNumber          int
	Name               string
	Ppk                string
	Channeling_Company string
	Otr                string
	Loan_Amount        string
	DrawdownDate       time.Time
	LoanPeriod         string
	InterestEffective  float32
	MonthlyPayment     string
	CollateralID       int64
	Branch             string
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetBranch() ([]model.Branch_Tabs, error) {
	var Branch []model.Branch_Tabs
	res := r.db.Order("code asc").Find(&Branch)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return Branch, nil
}
func (r *repository) GetCompany() ([]model.Mst_Company_Tabs, error) {
	var Company []model.Mst_Company_Tabs
	res := r.db.Order("company_code asc").Find(&Company)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return Company, nil
}

func (r *repository) GetPpk(ppk string) (model.Customer_Data_Tabs, error) {
	var Link model.Customer_Data_Tabs
	if err := r.db.Where("ppk = ?", ppk).First(&Link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Customer_Data_Tabs{}, errors.New("link not found")
		}
		return model.Customer_Data_Tabs{}, err
	}
	return Link, nil
}
func (r *repository) UpdateCustomer(customer model.Customer_Data_Tabs) (model.Customer_Data_Tabs, error) {
	_, err := r.GetPpk(customer.PPK)
	if err != nil {
		return model.Customer_Data_Tabs{}, errors.New("wrong data insurance")
	}
	res := r.db.Where("ppk=?", customer.PPK).Updates(model.Customer_Data_Tabs{
		Approval_Status: "0",
	})
	if res.Error != nil {
		return model.Customer_Data_Tabs{}, res.Error
	}
	return customer, nil
}
func (r *repository) SearchChecklistReport(branch string, company string, startDate string, endDate string) ([]response, error) {

	//res := r.db.Find(&Customer)
	query1 := ""
	query2 := ""
	if branch == "000" {
		query1 = "and ldt.branch like $1 "
		branch = "%%"
	} else {
		query1 = "and ldt.branch = $1 "
	}

	if company == "All Company" {
		query2 = "and cdt.channeling_company like $2 "
		company = "%%"
	} else {
		query2 = "and cdt.channeling_company = $2 "
	}

	res, err := r.db.Raw(`Select ROW_NUMBER() OVER (Order by cdt.name) AS RowNumber,cdt.ppk, cdt.name,cdt.channeling_company , 
	ldt.otr , ldt.loan_amount,cdt.drawdown_date , 
	ldt.loan_period , ldt.interest_effective, monthly_payment,vdt.collateral_id,ldt.branch from 
	customer_data_tab cdt left join Loan_Data_Tab ldt on cdt.custcode = ldt.custcode
	left join vehicle_data_tab vdt on cdt.custcode = vdt.custcode where cdt.approval_status ='0' and drawdown_date between $3 and $4 `+query1+query2, branch, company, startDate, endDate).Rows()
	data := []response{}
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var rowNumber int
		var ppk string
		var name string
		var channeling_company string
		var otr string
		var loan_amount string
		var drawdownDate time.Time
		var LoanPeriod string
		var InterestEffective float32
		var MonthlyPayment string
		var CollateralID int64
		var Branch string
		if err := res.Scan(&rowNumber, &ppk, &name, &channeling_company, &otr, &loan_amount, &drawdownDate, &LoanPeriod, &InterestEffective, &MonthlyPayment, &CollateralID, &Branch); err != nil {
			panic(err)
		}
		customer := response{
			RowNumber:          rowNumber,
			Ppk:                ppk,
			Name:               name,
			Channeling_Company: channeling_company,
			Otr:                otr,
			Loan_Amount:        loan_amount,
			DrawdownDate:       drawdownDate,
			LoanPeriod:         LoanPeriod,
			InterestEffective:  InterestEffective,
			MonthlyPayment:     MonthlyPayment,
			CollateralID:       CollateralID,
			Branch:             Branch,
		}
		data = append(data, customer)

	}

	return data, nil
}
