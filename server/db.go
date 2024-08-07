package server

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DbConfig struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DbName   string `json:"db_name"`
	}
)

func GetDefaultDbConfig() (cfg *DbConfig) {
	cfg = &DbConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "",
		DbName:   "boilerplate_dev",
	}
	return
}

func NewDb(cfg *DbConfig) (db *gorm.DB, err error) {
	if cfg == nil {
		cfg = GetDefaultDbConfig()
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return
	}
	return
}
