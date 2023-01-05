package model

type ID_Tabs struct {
	Code        string `json:"Code" gorm:"type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(250)"`
	Digit       int    `json:"digit" gorm:"type:int"`
	Value       int    `json:"value" gorm:"type:int"`
}

func (m *ID_Tabs) TableName() string {
	return "id_tab"
}
