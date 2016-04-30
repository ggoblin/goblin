package dao

import (
	_ "database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/ggoblin/goblin/platform/libs/utils"
	_ "github.com/lib/pq"
)

func GetAllIterations() ([]*model.Iteration, error) {
	db, err := utils.GetDefaultDb()
	log.Debug("Start Get all iterations.")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var iterations []model.Iteration
	err = db.Find(&iterations).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	ret := []*model.Iteration{}
	for _, i := range iterations {
		i.SetDate()
		ret = append(ret, &i)
	}
	log.Debug("Start Get all iterations.")
	return ret, nil
}

func AddIteration(iteration model.Iteration) (bool, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return false, err
	}
	defer db.Close()
	iteration.SetTime()
	log.Infof("Add new iteration %#v", iteration)
	err = db.Create(&iteration).Error
	if err != nil {
		return false, err
	}
	result := db.NewRecord(iteration)
	return !result, nil
}

func GetIterationById(id string) (*model.Iteration, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var iterations []model.Iteration
	db.Where("iteration_id = ?", id).Find(&iterations)
	if len(iterations) != 1 {
		return nil, fmt.Errorf("Iteration Id %s not found. %#v", id, iterations)
	}
	ret := &iterations[0]
	ret.SetDate()
	return ret, nil
}
