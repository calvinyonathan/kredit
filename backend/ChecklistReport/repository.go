package ChecklistReport

import (
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetChecklistReport() ([]response, error)
}
type repository struct {
	db *gorm.DB
}
type response struct {
	Custcode          string
	Ppk               string
	Otr               string
	Loan_Amount       string
	DrawdownDate      time.Time
	LoanPeriod        string
	InterestEffective float32
	MonthlyPayment    string
	CollateralID      int64
	Branch            string
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetChecklistReport() ([]response, error) {

	//res := r.db.Find(&Customer)
	res, err := r.db.Raw("Select cdt.ppk, cdt.name , ldt.otr , ldt.loan_amount,cdt.drawdown_date , ldt.loan_period , ldt.interest_effective,monthly_payment,vdt.collateral_id,ldt.branch from customer_data_tab cdt left join Loan_Data_Tab ldt on cdt.custcode = ldt.custcode left join vehicle_data_tab vdt on cdt.custcode = vdt.custcode").Rows()
	data := []response{}
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var ppk string
		var custcode string
		var otr string
		var loan_amount string
		var drawdownDate time.Time
		var LoanPeriod string
		var InterestEffective float32
		var MonthlyPayment string
		var CollateralID int64
		var Branch string
		if err := res.Scan(&ppk, &custcode, &otr, &loan_amount, &drawdownDate, &LoanPeriod, &InterestEffective, &MonthlyPayment, &CollateralID, &Branch); err != nil {
			panic(err)
		}
		customer := response{
			Ppk:               ppk,
			Custcode:          custcode,
			Otr:               otr,
			Loan_Amount:       loan_amount,
			DrawdownDate:      drawdownDate,
			LoanPeriod:        LoanPeriod,
			InterestEffective: InterestEffective,
			MonthlyPayment:    MonthlyPayment,
			CollateralID:      CollateralID,
			Branch:            Branch,
		}
		data = append(data, customer)

	}

	return data, nil
}
