package config

import (
	"fmt"
	"github.com/niT-Tin/prac-hackthon-group4/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDB *gorm.DB

type MySQLConfig struct {
	UserName  string
	Password  string
	Host      string
	Port      int
	DB        string
	Arguments string
}

func G4AutoMigrate() (*gorm.DB, error) {
	init, err := Init()
	if err != nil {
		return nil, err
	}
	return init, init.AutoMigrate(&models.User{}, &models.Text{})
}

func Init() (*gorm.DB, error) {
	open, err := gorm.Open(mysql.Open(InitMySqlDSN()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	return open, nil
}

func InitMySqlDSN() string {
	sqlConfig := MySQLConfig{
		UserName:  "hackliu",
		Password:  "Liu123321",
		Host:      "rm-uf6i3ox1ocm68r52hco.sqlConfig.rds.aliyuncs.com",
		Port:      3306,
		DB:        "hackthon",
		Arguments: "charset=utf8mb4&parseTime=True",
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		sqlConfig.UserName,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.DB,
		sqlConfig.Arguments)
}
