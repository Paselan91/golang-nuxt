package config

import (
	// "app/src/domain"
	"github.com/jinzhu/gorm"
	"log"
	// "time"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	log.Println("Migration start")
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	log.Println("Migration has been processed")
	return conn, nil
}

func Seeds() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return nil, err
}
