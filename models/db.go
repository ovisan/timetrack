package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DB struct {
	*gorm.DB
}

func InitDB(dataSource string) {
	os.Remove(dataSource)
	db, err := gorm.Open("sqlite3", dataSource)
	if err != nil {
		log.Panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Panic(err)
	}
	defer db.Close()
}

func NewDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return &db, nil
}
