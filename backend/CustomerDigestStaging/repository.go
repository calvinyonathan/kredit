package CustomerDigestStaging

import (
	"calvin/kredit/model"
	"fmt"
	"log"
	"strconv"
	"time"

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
func (r *repository) validateCustomer(ppk string) bool {
	var Customer []model.Customer_Data_Tabs
	res2 := r.db.Where("ppk= ?", ppk).First(&Customer)
	if res2.Error != nil {
		return true
	} else {
		return false
	}
}
func (r *repository) validateCompany(ScCompany string) bool {
	var Company []model.Mst_Company_Tabs
	res3 := r.db.Where("company_short_name = ?", ScCompany).First(&Company)
	if res3.Error != nil {
		return false
	} else {
		return true
	}
}
func (r *repository) validateBranch(BranchCode string) bool {
	var Branch []model.Branch_Tabs
	res4 := r.db.Where("code=?", BranchCode).First(&Branch)
	if res4.Error != nil {
		return false
	} else {
		return true
	}
}
func (r *repository) validateEngineNo(EngineNo string) bool {
	var Vehicle []model.Vehicle_Data_Tabs
	res5 := r.db.Where("engine_no = ?", EngineNo).First(&Vehicle)
	if res5.Error != nil {
		return true
	} else {
		return false
	}
}
func (r *repository) validateChasisNo(ChasisNo string) bool {
	var Vehicle []model.Vehicle_Data_Tabs
	res5 := r.db.Where("chasis_no = ?", ChasisNo).First(&Vehicle)
	if res5.Error != nil {
		return true
	} else {
		return false
	}
}
func (r *repository) validateTglPK(TglPK string, currentMonth int) bool {
	// Parse the date string
	date, err := time.Parse("2006-01-02", TglPK)
	if err != nil {
		fmt.Println(err)
	}
	// Get the month
	month := date.Month()
	if currentMonth == int(month) {
		return true
	} else {
		return false
	}

}
func (r *repository) validateError(ID int64) {
	var Staging []model.Staging_Customers
	r.db.Model(&Staging).Where("id=?", ID).Update("sc_flag", "8")

}
func (r *repository) GetCustomer() ([]model.Staging_Customers, error) {
	var Staging []model.Staging_Customers
	currentTime := time.Now()

	res := r.db.Where("sc_flag= ? AND sc_create_date = ?", "0", currentTime.Format("2006-01-02")).Find(&Staging)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}

	for _, item := range Staging {
		r.validateCustomer(item.CustomerPpk)
		if !r.validateCustomer(item.CustomerPpk) {
			r.validateError(item.ID)
			continue
		}
		// r.validateCompany(item.ScCompany)
		// if !r.validateCompany(item.ScCompany) {
		// 	r.validateError(item.ID)
		// 	continue
		// }
		// r.validateBranch(item.ScBranchCode)
		// if !r.validateBranch(item.ScBranchCode) {
		// 	r.validateError(item.ID)
		// 	continue
		// }

		// r.validateTglPK(item.LoanTglPk, int(currentTime.Month()))
		// if !r.validateTglPK(item.LoanTglPk, int(currentTime.Month())) {
		// 	r.validateError(item.ID)
		// 	continue
		// }

		// if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
		// 	r.validateError(item.ID)
		// 	continue
		// }
		// regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
		// if matched := regex.MatchString(item.CustomerName); !matched {
		// 	r.validateError(item.ID)
		// 	continue
		// }

		// if item.VehicleBpkb == "" || item.VehicleStnk == "" || item.VehicleEngineNo == "" || item.VehicleChasisNo == "" {
		// 	r.validateError(item.ID)
		// 	continue
		// }
		// r.validateEngineNo(item.VehicleEngineNo)
		// if !r.validateEngineNo(item.VehicleEngineNo) {
		// 	r.validateError(item.ID)
		// 	continue
		// }
		// r.validateChasisNo(item.VehicleChasisNo)
		// if !r.validateChasisNo(item.VehicleChasisNo) {
		// 	r.validateError(item.ID)
		// 	continue
		// }

		r.db.Model(&Staging).Where("id=?", item.ID).Update("sc_flag", "1")
		date, err := time.Parse("2006-01-02 15:04:05", item.CustomerBirthDate)
		if err != nil {
			fmt.Println(err)
		}
		idType, err := strconv.ParseInt(item.CustomerIDType, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		drawdownDate, err := time.Parse("2006-01-02", item.LoanTglPk)
		if err != nil {
			fmt.Println(err)
		}
		tglPkChanneling, err := time.Parse("2006-01-02", item.LoanTglPkChanneling)
		if err != nil {
			fmt.Println(err)
		}
		Customer := model.Customer_Data_Tabs{
			Custcode:           "1234567890",
			PPK:                item.CustomerPpk,
			Name:               item.CustomerName,
			Address1:           item.CustomerAddress1,
			Address2:           item.CustomerAddress2,
			City:               item.CustomerCity,
			Zip:                item.CustomerZip,
			Birth_Place:        item.CustomerBirthPlace,
			Birth_Date:         date,
			ID_Type:            int8(idType),
			ID_Number:          item.CustomerIDNumber,
			Mobile_No:          item.CustomerMobileNo,
			Drawdown_date:      drawdownDate,
			Tgl_pk_channelling: tglPkChanneling,
			Mother_maiden_name: item.CustomerMotherMaidenName,
			Channeling_Company: item.ScCompany,
			Approval_Status:    "9",
		}
		r.db.Create(&Customer)

	}

	return Staging, nil
}
