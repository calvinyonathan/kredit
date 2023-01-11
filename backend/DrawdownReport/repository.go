package DrawdownReport

import (
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetDrawdownReport() ([]response, error)
}
type repository struct {
	db *gorm.DB
}
type response struct {
	RowNumber          int
	Name               string
	Ppk                string
	Channeling_Company string
	Loan_Amount        string
	DrawdownDate       time.Time
	LoanPeriod         string
	InterestEffective  float32
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetDrawdownReport() ([]response, error) {

	//res := r.db.Find(&Customer)
	res, err := r.db.Raw("Select ROW_NUMBER() OVER (Order by cdt.name) AS RowNumber,cdt.ppk, cdt.name,cdt.channeling_company ,  ldt.loan_amount,cdt.drawdown_date , ldt.loan_period , ldt.interest_effective from customer_data_tab cdt left join Loan_Data_Tab ldt on cdt.custcode = ldt.custcode left join vehicle_data_tab vdt on cdt.custcode = vdt.custcode where cdt.approval_status in('0','1') ").Rows()
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
		var loan_amount string
		var drawdownDate time.Time
		var LoanPeriod string
		var InterestEffective float32

		if err := res.Scan(&rowNumber, &ppk, &name, &channeling_company, &loan_amount, &drawdownDate, &LoanPeriod, &InterestEffective); err != nil {
			panic(err)
		}
		customer := response{
			RowNumber:          rowNumber,
			Ppk:                ppk,
			Name:               name,
			Channeling_Company: channeling_company,
			Loan_Amount:        loan_amount,
			DrawdownDate:       drawdownDate,
			LoanPeriod:         LoanPeriod,
			InterestEffective:  InterestEffective,
		}
		data = append(data, customer)

	}

	return data, nil
}
