package model

type Mst_Company_Tabs struct {
	ID                 int64   `json:"id" gorm:"type:bigint;AUTO_INCREMENT;NOT NULL"`
	Company_Code       string  `json:"company_code" gorm:"type:varchar(5)"`
	Company_Short_Name string  `json:"company_short_name" gorm:"type:varchar(50)"`
	Company_Name       string  `json:"company_name" gorm:"type:varchar(200)"`
	Company_Address1   string  `json:"company_address1" gorm:"type:varchar(200)"`
	Company_Address2   string  `json:"company_address2" gorm:"type:varchar(200)"`
	Company_City       string  `json:"company_city" gorm:"type:varchar(100)"`
	Company_Phone      string  `json:"company_phone" gorm:"type:varchar(50)"`
	Company_Fax        string  `json:"company_fax" gorm:"type:varchar(50)"`
	Bunga_Eff_Min      float32 `json:"bunga_eff_min" gorm:"type:real"`
	Bunga_Eff_Max      float32 `json:"bunga_eff_max" gorm:"type:real"`
	Bunga_Flat_Min     float32 `json:"bunga_flat_min" gorm:"type:real"`
	Bunga_Flat_Max     float32 `json:"bunga_flat_max" gorm:"type:real"`
	LA_min             string  `json:"la_min" gorm:"type:money"`
	LA_max             string  `json:"la_max" gorm:"type:money"`
	Periode_min        string  `json:"periode_min" gorm:"type:varchar(10)"`
	Periode_max        string  `json:"periode_max" gorm:"type:varchar(10)"`
}

func (m *Mst_Company_Tabs) TableName() string {
	return "mst_company_tab"
}
