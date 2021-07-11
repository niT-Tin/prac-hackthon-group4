package config

import (
	"fmt"
)

type MySQLConfig struct {
	UserName  string
	Password  string
	Host      string
	Port      int
	DB        string
	Arguments string
}

func InitMySqlDSN() string {
	sqlConfig := MySQLConfig{
		UserName:  "hackliu",
		Password:  "Liu123321",
		Host:      "rm-uf6i3ox1ocm68r52hco.mysql.rds.aliyuncs.com",
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
