package CustomerDigestStaging

import (
	"calvin/kredit/model"
	"log"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomer() ([]model.Staging_Customers, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCustomer() ([]model.Staging_Customers, error) {
	var Staging []model.Staging_Customers
	var Customer []model.Customer_Data_Tabs
	res := r.db.Where("sc_flag= ?", "0").Find(&Staging)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}

	for _, item := range Staging {
		var sc_flag = "0"
		res2 := r.db.Find(&Customer)

		if res2.Error != nil {
			log.Println("Get Data error : ", res2.Error)
			return nil, res2.Error
		}

		for _, item2 := range Customer {
			if item2.PPK == item.CustomerPpk {
				sc_flag = "8"
			} else {
				sc_flag = "0"
			}
		}

		if sc_flag == "0" {
			sc_flag = "1"
			r.db.Model(&Staging).Update("sc_flag", sc_flag)
		} else if sc_flag == "8" {
			r.db.Model(&Staging).Update("sc_flag", sc_flag)
		}

	}

	return Staging, nil
}
