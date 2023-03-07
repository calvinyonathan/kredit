package Login

import (
	"calvin/kredit/model"
	"log"

	"gorm.io/gorm"
)

type LoginRepository interface {
	GetLogin(nik string, password string) ([]model.Officer_Tab, error)
	UpdatePassword(nik string, oldpassword string, password string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLogin(nik string, password string) ([]model.Officer_Tab, error) {
	var Login []model.Officer_Tab
	res := r.db.Where("nik = ? and password = ?", nik, password).Find(&Login)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return Login, nil

}

func (r *repository) UpdatePassword(nik string, oldpassword string, password string) error {
	var Login []model.Officer_Tab
	res1 := r.db.Where("nik = ? and password = ?", nik, oldpassword).First(&Login)
	if res1.Error != nil {
		log.Println("Get Data error : ", res1.Error)
		return res1.Error
	} else {
		res := r.db.Where("nik=?", nik).Updates(model.Officer_Tab{
			Password: password,
		})
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
