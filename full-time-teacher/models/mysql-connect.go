package models

import (
	"FullTimeTeacher/config"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type MysqlConfig struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Dbname   string `json:"dbname" mapstructure:"db_name"`
}

const args = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

func MySQLConnection(config *config.Config) error {

	username := config.MySQL.Username
	password := config.MySQL.Password
	host := config.MySQL.Host
	port := config.MySQL.Port
	dbname := config.MySQL.DBName

	dbOpt := fmt.Sprintf(args, username, password, host, port, dbname)

	fmt.Println(dbOpt)
	db, err := gorm.Open(mysql.Open(dbOpt), &gorm.Config{})
	if err != nil {
		return errors.New("db connect failed")
	}

	err = NewUserTable(db)
	if err != nil {
		return errors.New("user_info table create failed")
	}
	return nil
}
