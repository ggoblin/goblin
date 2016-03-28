package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Dbc string
var Dbtype string

func GetDefaultDb() (*gorm.DB, error) {
	return gorm.Open(Dbtype, Dbc)
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
