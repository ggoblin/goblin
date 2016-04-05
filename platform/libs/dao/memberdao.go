package dao

import (
	_ "database/sql"
	"fmt"
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

func GetMember(id string) (*model.Member, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var members []model.Member
	db.Where(&model.Member{MemberId: id}).Find(&members)
	if len(members) != 1 {
		return nil, fmt.Errorf("Get Member by id %s error.Result %#v", id, members)
	}
	return &members[0], nil
}

func AddNewMember(member model.Member) (bool, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return false, err
	}
	defer db.Close()
	log.Infof("%#v", member)
	db.Create(&member)
	result := db.NewRecord(member)
	return result, nil
}
