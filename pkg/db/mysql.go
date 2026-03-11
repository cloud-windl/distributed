package db

import (
	"fmt"

	"distributed/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL() (*gorm.DB, error) {
	host := config.GetEnv("MYSQL_HOST", "127.0.0.1")
	port := config.GetEnv("MYSQL_PORT", "3306")
	user := config.GetEnv("MYSQL_USER", "root")
	password := config.GetEnv("MYSQL_PASSWORD", "123456")
	database := config.GetEnv("MYSQL_DB", "distributed")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
