package model

type Officer_Tab struct {
	Nik      string `json:"nik" gorm:"type:varchar(50);not null"`
	Password string `json:"password" gorm:"type:varchar(150)"`
}

func (m *Officer_Tab) TableName() string {
	return "officer_tab"
}
