package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "drose"
const DB_PASSWORD = "Daroca288"
const DB_NAME = "contable"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = MysqlConnection()
	return Db
}

func MysqlConnection() *gorm.DB {
	dsnName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	log.Println("Dsn =", dsnName)
	db, err := gorm.Open(mysql.Open(dsnName), &gorm.Config{})
	if err != nil {
		log.Fatal("Error db connection, ", err)
		return nil
	}
	return db

}
