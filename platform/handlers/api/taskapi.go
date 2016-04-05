package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/dao"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/labstack/echo"
	"net/http"
)

func CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		task := &model.Task{}
		if err := c.Bind(task); err != nil {
			log.Error(err)
			return err
		}
		log.Infof("Create task %#v", task)
		result, err := dao.AddTask(*task)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}
