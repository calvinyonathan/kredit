package model

import "time"

type Skala_Rental_Tabs struct {
	Custcode       string    `json:"custcode" gorm:"not null; type: varchar(25); unique"`
	Counter        int8      `json:"counter" gorm:"not null; type: smallint"`
	Osbalance      float64   `json:"osbalance" gorm:"type: money"`
	End_Balance    float64   `json:"end_balance" gorm:"type: money"`
	Due_Date       time.Time `json:"due_date"`
	Eff_Rate       float32   `json:"eff_rate" gorm:"type: float"`
	Rental         float64   `json:"rental" gorm:"type: money"`
	Principle      float64   `json:"principle" gorm:"type: money"`
	Interest       float64   `json:"interest" gorm:"type: money"`
	InputDate      time.Time `json:"inputdate"`
	InputBy        string    `json:"inputby" gorm:"type: varchar(50)"`
	LastModified   time.Time `json:"lastmodified"`
	ModifiedBy     string    `json:"modifiedby" gorm:"type: varchar(50)"`
	Payment_Date   time.Time `json:"payment_date"`
	Penalty        float64   `json:"penalty" gorm:"type: money"`
	Payment_Amount float64   `json:"payment_amount" gorm:"type: money"`
	Payment_Type   int8      `json:"payment_type" gorm:"type:smallint"`
}

func (m *Skala_Rental_Tabs) TableName() string {
	return "skala_rental_tab"
}
