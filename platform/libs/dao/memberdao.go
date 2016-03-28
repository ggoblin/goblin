package dao

import (
	_ "database/sql"
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/ggoblin/goblin/platform/libs/utils"
	_ "github.com/lib/pq"
)

func GetAllMembers() ([]model.Member, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var members []model.Member
	db.Find(&members)
	return members, nil
}

func AddNewMember(member model.Member) (bool, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return false, err
	}
	defer db.Close()
	log.Infof("%#v", member)
	result := db.NewRecord(member)
	db.Create(&member)
	return result, nil
}
