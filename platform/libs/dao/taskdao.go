package dao

import (
	_ "database/sql"
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/ggoblin/goblin/platform/libs/utils"
	_ "github.com/lib/pq"
)

func AddTask(task model.Task) (bool, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return false, err
	}
	defer db.Close()
	task.SetTime()
	log.Infof("Add new Task %#v", task)
	db.Create(&task)
	result := db.NewRecord(task)
	return !result, nil
}

func GetTasksByIterations(iterationId string) ([]model.Task, error) {
	db, err := utils.GetDefaultDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	log.Debugf("Get Tasks by iterationId: %s", iterationId)
	var tasks []model.Task
	db.Where("iteration_id = ?", iterationId).Find(&tasks)
	return tasks, nil
}
