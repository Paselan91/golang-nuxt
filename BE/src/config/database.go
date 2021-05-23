/**
 * database と接続を行う
 */
package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// ConfigDB db seting
type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

// ConnectDB returns initialized gorm.DB
func ConnectDB() (*gorm.DB, error) {
	config := ConfigDB{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_CONTAINER_NAME"),
		Port:     os.Getenv("MYSQL_PORT"),
		Dbname:   os.Getenv("MYSQL_DATABASE"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		panic("failed to connect DB")
		return nil, err
	}
	return db, nil
}
