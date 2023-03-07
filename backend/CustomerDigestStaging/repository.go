package CustomerDigestStaging

import (
	"calvin/kredit/model"
	"fmt"
	"log"
	"regexp"
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
func (r *repository) insertStagingError(SeReff string, SeCreateDate time.Time, BranchCode string, Company string, Ppk string, Name string, Reason string) {
	Staging_Error := model.Staging_Errors{
		SeReff:       SeReff,
		SeCreateDate: SeCreateDate,
		BranchCode:   BranchCode,
		Company:      Company,
		Ppk:          Ppk,
		Name:         Name,
		ErrorDesc:    Reason,
	}
	r.db.Create(&Staging_Error)
}
func (r *repository) insertCustomerDataTab(Custcode string, ppk string, name string, address1 string, address2 string,
	city string, zip string, birthPlace string, birthDate time.Time, ID_type int8, ID_Number string, MobileNo string,
	Drawdown_Date time.Time, TglPkChanneling time.Time, MotherMaidenName string, ChannelingCompany string) {
	Customer := model.Customer_Data_Tabs{
		Custcode:           Custcode,
		PPK:                ppk,
		Name:               name,
		Address1:           address1,
		Address2:           address2,
		City:               city,
		Zip:                zip,
		Birth_Place:        birthPlace,
		Birth_Date:         birthDate,
		ID_Type:            ID_type,
		ID_Number:          ID_Number,
		Mobile_No:          MobileNo,
		Drawdown_date:      Drawdown_Date,
		Tgl_pk_channelling: TglPkChanneling,
		Mother_maiden_name: MotherMaidenName,
		Channeling_Company: ChannelingCompany,
		Approval_Status:    "9",
	}
	r.db.Create(&Customer)
}
func (r *repository) insertLoanDataTab(Custcode string, Branch string, OTR float64, DownPayment float64, LoanAmount float64, LoanPeriod string, InterestFlat float32, InterestEffective float32,
	EffectivePaymentType int8, MonthlyPayment float64, inputDate time.Time, InputDate2 time.Time) {
	Loan := model.Loan_Data_Tabs{
		Custcode:             Custcode,
		Branch:               Branch,
		OTR:                  OTR,
		DownPayment:          DownPayment,
		LoanAmount:           LoanAmount,
		LoanPeriod:           LoanPeriod,
		InterestType:         1,
		InterestFlat:         InterestFlat,
		InterestEffective:    InterestEffective,
		EffectivePaymentType: EffectivePaymentType,
		AdminFee:             30,
		MonthlyPayment:       MonthlyPayment,
		InputDate:            inputDate,
		LastModified:         time.Now(),
		ModifiedBy:           "system",
		InputDate2:           InputDate2,
		InputBy:              "system",
		LastModified2:        time.Now(),
		ModifiedBy2:          "system",
	}
	r.db.Create(&Loan)
}
func (r *repository) insertVehicleDataTab(Custcode string, Brand int64, Type string, Year string, Golongan int8, Jenis string, Status int8,
	Color string, PoliceNo string, EngineNo string, ChasisNo string, Bpkb string, RegisterNo string, Stnk string, StnkAddress1 string,
	StnkAddress2 string, StnkCity string, DealerID int, InputDate time.Time, Inputby string,
	LastModified time.Time, Modifiedby string, TglStnk time.Time, TglBpkb time.Time, TglPolis time.Time, PolisNo string, CollateralID int64,
	Ketagunan string, AgunanLbu string, Dealer string, AddressDealer1 string, AddressDealer2 string, CityDealer string) {
	Vehicle := model.Vehicle_Data_Tabs{
		Custcode:       Custcode,
		Brand:          Brand,
		Type:           Type,
		Year:           Year,
		Golongan:       Golongan,
		Jenis:          Jenis,
		Status:         Status,
		Color:          Color,
		PoliceNo:       PoliceNo,
		EngineNo:       EngineNo,
		ChasisNo:       ChasisNo,
		Bpkb:           Bpkb,
		RegisterNo:     RegisterNo,
		Stnk:           Stnk,
		StnkAddress1:   StnkAddress1,
		StnkAddress2:   StnkAddress2,
		StnkCity:       StnkCity,
		DealerID:       DealerID,
		Inputdate:      InputDate,
		Inputby:        Inputby,
		Lastmodified:   LastModified,
		Modifiedby:     Modifiedby,
		TglStnk:        TglStnk,
		TglBpkb:        TglBpkb,
		TglPolis:       TglPolis,
		PolisNo:        PolisNo,
		CollateralID:   CollateralID,
		Ketagunan:      Ketagunan,
		AgunanLbu:      AgunanLbu,
		Dealer:         Dealer,
		AddressDealer1: AddressDealer1,
		AddressDealer2: AddressDealer2,
		CityDealer:     CityDealer,
	}
	r.db.Create(&Vehicle)
}
func (r *repository) GetCustomer() ([]model.Staging_Customers, error) {
	var Staging []model.Staging_Customers
	var Config model.Config_Properties
	var ID_Tab model.ID_Tabs
	var Company model.Mst_Company_Tabs
	var Reason string
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
			Reason = "Validasi Salah"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}
		r.validateCompany(item.ScCompany)
		if !r.validateCompany(item.ScCompany) {
			r.validateError(item.ID)
			Reason = "Validasi Company"
			continue
		}
		r.validateBranch(item.ScBranchCode)
		if !r.validateBranch(item.ScBranchCode) {
			r.validateError(item.ID)
			Reason = "Validasi Branch"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}

		r.validateTglPK(item.LoanTglPk, int(currentTime.Month()))
		if !r.validateTglPK(item.LoanTglPk, int(currentTime.Month())) {
			r.validateError(item.ID)
			Reason = "Validasi Tanggal PK"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}

		if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
			r.validateError(item.ID)
			Reason = "Validasi Customer Type dan Number"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}
		regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
		if matched := regex.MatchString(item.CustomerName); matched {
			r.validateError(item.ID)
			Reason = "Validasi Spesial Karakter"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}

		if item.VehicleBpkb == "" || item.VehicleStnk == "" || item.VehicleEngineNo == "" || item.VehicleChasisNo == "" {
			r.validateError(item.ID)
			Reason = "Validasi Kosong"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}
		r.validateEngineNo(item.VehicleEngineNo)
		if !r.validateEngineNo(item.VehicleEngineNo) {
			r.validateError(item.ID)
			Reason = "Validasi Engine No"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}
		r.validateChasisNo(item.VehicleChasisNo)
		if !r.validateChasisNo(item.VehicleChasisNo) {
			r.validateError(item.ID)
			Reason = "Validasi Chasis No"
			r.insertStagingError(item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName, Reason)
			continue
		}

		//Update SC_FLAG ke 1
		r.db.Model(&Staging).Where("id=?", item.ID).Update("sc_flag", "1")

		r.db.Where("parameter = ?", "appCustCode").First(&Config)
		appCode := Config.Value

		r.db.Where("company_short_name = ? ", item.ScCompany).First(&Company)
		companyCode := Company.Company_Code

		r.db.Where("code = ?", appCode).First(&ID_Tab)
		appCustCodeLen := ID_Tab.Digit
		appCustCodeSeq := ID_Tab.Value

		appCustCodeSeqNew := "0000000000" + fmt.Sprintf("%d", appCustCodeSeq)
		appCustCodeSeqNewLen := len(appCustCodeSeqNew)
		appCustCodeSeqNew = appCustCodeSeqNew[appCustCodeSeqNewLen-appCustCodeLen:]

		Now := time.Now()
		month := int(Now.Month())
		monthString := ""
		year := int(Now.Year())
		yearString := fmt.Sprintf("%d", year)
		if month < 10 {
			monthString = "0" + fmt.Sprintf("%d", month)
		} else {
			monthString = fmt.Sprintf("%d", month)
		}
		newCustcode := appCode + companyCode + yearString + monthString + appCustCodeSeqNew
		appCustCodeSeq += 1
		r.db.Model(&ID_Tab).Where("code=?", appCode).Update("value", appCustCodeSeq)

		//Convert Data - Data
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
		InterestFlat, err := strconv.ParseFloat(item.LoanInterestFlatChanneling, 32)
		if err != nil {
			fmt.Println(err)
		}
		InterestEffective, err := strconv.ParseFloat(item.LoanInterestEffectiveChanneling, 32)
		if err != nil {
			fmt.Println(err)
		}
		EffectivePaymentType, err := strconv.ParseInt(item.LoanEffectivePaymentType, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		VehicleType, err := strconv.ParseInt(item.VehicleType, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		Status, err := strconv.ParseInt(item.VehicleStatus, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		VehicleDealerID, err := strconv.ParseInt(item.VehicleDealerID, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		VehicleTglStnk, err := time.Parse("2006-01-02 15:04:05", item.VehicleTglStnk)
		if err != nil {
			fmt.Println(err)
		}
		VehicleTglBpkb, err := time.Parse("2006-01-02 15:04:05", item.VehicleTglStnk)
		if err != nil {
			fmt.Println(err)
		}
		CollateralTypeID, err := strconv.ParseInt(item.CollateralTypeID, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		OTR, err := strconv.ParseFloat(item.LoanOtr, 32)
		if err != nil {
			fmt.Println(err)
		}

		LoanDownPayment, err := strconv.ParseFloat(item.LoanDownPayment, 64)
		if err != nil {
			fmt.Println(err)
		}
		LoanLoanAmountChanneling, err := strconv.ParseFloat(item.LoanLoanAmountChanneling, 64)
		if err != nil {
			fmt.Println(err)
		}
		LoanMonthlyPaymentChanneling, err := strconv.ParseFloat(item.LoanMonthlyPaymentChanneling, 64)
		if err != nil {
			fmt.Println(err)
		}

		//Insert Data to Customer Data Tab , Loan Data Tab , Skala Rental Tab
		r.insertCustomerDataTab(newCustcode, item.CustomerPpk, item.CustomerName, item.CustomerAddress1, item.CustomerAddress2, item.CustomerCity, item.CustomerZip,
			item.CustomerBirthPlace, date, int8(idType), item.CustomerIDNumber, item.CustomerMobileNo, drawdownDate, tglPkChanneling, item.CustomerMotherMaidenName, item.ScCompany)

		r.insertLoanDataTab(newCustcode, item.ScBranchCode, OTR, LoanDownPayment, LoanLoanAmountChanneling, item.LoanLoanPeriodChanneling, float32(InterestFlat),
			float32(InterestEffective), int8(EffectivePaymentType), LoanMonthlyPaymentChanneling, item.ScCreateDate, item.ScCreateDate)

		r.insertVehicleDataTab(newCustcode, int64(VehicleType), item.VehicleBrand, item.VehicleYear, 1, item.VehicleJenis, int8(Status), item.VehicleColor,
			item.VehiclePoliceNo, item.VehicleEngineNo, item.VehicleChasisNo, item.VehicleBpkb, "1", item.VehicleStnk, "", "", "", int(VehicleDealerID), time.Now(), "system",
			time.Now(), "system", VehicleTglStnk, VehicleTglBpkb, time.Now(), item.VehiclePoliceNo, CollateralTypeID, "", "", item.VehicleDealer,
			item.VehicleAddressDealer1, item.VehicleAddressDealer2, item.VehicleCityDealer)
		//r.db.Exec(" INSERT INTO loan_data_tab(custcode,branch,otr,down_payment) VALUES($1 , $2 , $3 ,$4 )", "12223545", item.ScBranchCode, loanOTR, item.LoanDownPayment)

	}

	return Staging, nil
}
