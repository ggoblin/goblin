package utils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Dbc string
var Dbtype string

func GetDefaultDb() (*gorm.DB, error) {
	if Dbc == "" {
		return nil, fmt.Errorf("Dbconnection string is empty.")
	}
	dbc, err := gorm.Open(Dbtype, Dbc)
	if err != nil {
		return nil, err
	}
	dbc.SetLogger(&GormLogger{})
	dbc.LogMode(true)
	return dbc, nil
}

func AutoMigrate() error {
	db, err := GetDefaultDb()
	if err != nil {
		return err
	}
	log.Info("Start migrate databse.")
	db.AutoMigrate(&model.Member{}, &model.Task{}, &model.Iteration{})
	log.Info("Migrate database done.")
	return nil
}

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm", "type": "sql"}).Print(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
