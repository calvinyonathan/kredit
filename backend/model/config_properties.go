package model

type Config_Properties struct {
	ID        int64  `json:"id" gorm:"type:bigint;AUTO_INCREMENT;NOT NULL"`
	Parameter string `json:"parameter" gorm:"type:varchar(250)"`
	Value     string `json:"value" gorm:"type:varchar(250)"`
	Company   string `json:"company" gorm:"type:varchar(50)"`
}

func (m *Config_Properties) TableName() string {
	return "config_property"
}
