package GenerateSkalaAngsuran

import (
	"calvin/kredit/model"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetSkalaAngsuran() ([]model.Customer_Data_Tabs, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetSkalaAngsuran() ([]model.Customer_Data_Tabs, error) {
	var customer []model.Customer_Data_Tabs
	var loan model.Loan_Data_Tabs
	var skala model.Skala_Rental_Tabs
	res := r.db.Where("approval_status = ?", "0").Find(&customer)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	for _, item := range customer {

		res := r.db.Where("custcode = ?", item.Custcode).First(&loan)
		if res.Error != nil {
			log.Println("Get Data error : ", res.Error)
			return nil, res.Error
		}
		LoanPeriod, err := strconv.ParseInt(loan.LoanPeriod, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		OsBalanace := loan.LoanAmount

		MonthlyPayment := loan.MonthlyPayment
		dataSkalaRental := make([]model.Skala_Rental_Tabs, LoanPeriod+1)
		TimeNow := time.Now()
		for i := range dataSkalaRental {
			if i == 0 {
				Skala := model.Skala_Rental_Tabs{
					Custcode:    item.Custcode,
					Counter:     int8(i),
					Osbalance:   OsBalanace,
					End_Balance: OsBalanace,
					Due_Date:    TimeNow,
					Eff_Rate:    loan.InterestEffective,
					Rental:      MonthlyPayment,
					Principle:   0,
					Interest:    0,
					InputDate:   TimeNow,
				}
				r.db.Create(&Skala)
			} else {
				Interest := math.Round(OsBalanace * float64(loan.InterestEffective) * 30 / 36000)
				Principle := math.Round(MonthlyPayment - Interest)
				End_Balance := math.Round(math.Round(OsBalanace) - math.Round(Principle))
				// Interest := math.Floor(OsBalanace * float64(loan.InterestEffective) * 30 / 36000)
				// Principle := MonthlyPayment - Interest
				// End_Balance := OsBalanace - Principle
				DueDate := TimeNow.AddDate(0, i, 0)
				Skala := model.Skala_Rental_Tabs{
					Custcode:    item.Custcode,
					Counter:     int8(i),
					Osbalance:   OsBalanace,
					End_Balance: End_Balance,
					Due_Date:    DueDate,
					Eff_Rate:    loan.InterestEffective,
					Rental:      MonthlyPayment,
					Principle:   Principle,
					Interest:    Interest,
					InputDate:   TimeNow,
				}
				r.db.Create(&Skala)
				OsBalanace = End_Balance
			}
		}
		var maxCounter float64
		row := r.db.Model(&model.Skala_Rental_Tabs{}).Where("custcode = ?", item.Custcode).Select("max(counter)").Row()
		row.Scan(&maxCounter)
		fmt.Println(maxCounter)
		res = r.db.Where("custcode = ? and counter = ?", item.Custcode, maxCounter).First(&skala)
		if res.Error != nil {
			log.Println("Get Data error : ", res.Error)
			return nil, res.Error
		}
		r.db.Exec("UPDATE SKALA_RENTAL_TAB SET end_balance = $1 where custcode = $2 and counter = $3", 0, item.Custcode, maxCounter)
		r.db.Model(&customer).Where("custcode=?", item.Custcode).Update("Approval_Status", "1")
	}
	return customer, nil
}
