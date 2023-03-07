package api

import (
	"fmt"
	"os"

	// "github.com/glebarez/sqlite"
	"calvin/kredit/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		host := os.Getenv("host")
		port := os.Getenv("port_db")
		user := os.Getenv("user")
		password := os.Getenv("password")
		dbname := os.Getenv("dbname")
		config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	if os.Getenv("AUTO_MIGRATE") == "Y" {
		if err := db.AutoMigrate(model.Config_Properties{}, model.ID_Tabs{}, model.Customer_Data_Tabs{}, model.Mst_Company_Tabs{}, model.Branch_Tabs{}, model.Loan_Data_Tabs{}, model.Skala_Rental_Tabs{}, model.Vehicle_Data_Tabs{}, model.Staging_Customers{}, model.Staging_Errors{}, model.Officer_Tab{}); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}
	}
	return db, err
}
