package dao

import (
	"Group4/config"
	"Group4/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func G4AutoMigrate() (*gorm.DB, error) {
	var err error
	mysqlDB, err = Init()
	if err != nil {
		return nil, err
	}
	return mysqlDB, mysqlDB.AutoMigrate(&models.User{}, &models.Text{})
}

func Init() (*gorm.DB, error) {
	open, err := gorm.Open(mysql.Open(config.InitMySqlDSN()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	return open, nil
}
