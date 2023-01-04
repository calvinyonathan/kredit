package model

import "time"

type Customer_Data_Tabs struct {
	Custcode           string    `json:"custcode" gorm:"type:varchar(35); not null; unique"`
	PPK                string    `json:"ppk" gorm:"type:varchar(20)"`
	Name               string    `json:"name" gorm:"type:varchar(100)"`
	Address1           string    `json:"address1" gorm:"type:varchar(100)"`
	Address2           string    `json:"address2" gorm:"type:varchar(100)"`
	City               string    `json:"city" gorm:"type:varchar(30)"`
	Zip                string    `json:"zip" gorm:"type:varchar(6)"`
	Birth_Place        string    `json:"birth_place" gorm:"type:varchar(20)"`
	Birth_Date         time.Time `json:"birth_date"`
	ID_Type            int8      `json:"id_type"`
	ID_Number          string    `json:"id_number" gorm:"type:varchar(30)"`
	Mobile_No          string    `json:"mobile_no" gorm:"type:varchar(20)"`
	Drawdown_date      time.Time `json:"drawdown_date"`
	Tgl_pk_channelling time.Time `json:"tgl_pk_channeling"`
	Mother_maiden_name string    `json:"mother_maiden_name" gorm:"type: varchar(100)"`
	Channeling_Company string    `json:"channeling_company" gorm:"type: varchar(100)"`
	Approval_Status    string    `json:"approval_status" gorm:"type: varchar(2)"`
}

func (m *Customer_Data_Tabs) TableName() string {
	return "customer_data_tab"
}
