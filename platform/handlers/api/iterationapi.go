package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/dao"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllIterations() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Start to query all iterations.")
		iterations, err := dao.GetAllIterations()
		if err != nil {
			log.Error("Get all iterations fail.%s", err)
			return c.JSON(http.StatusInternalServerError, "Fail")
		}
		log.Debugf("Get iterations %#v", iterations)
		return c.JSON(http.StatusOK, iterations)
	}
}

func CreateIteration() echo.HandlerFunc {
	return func(c echo.Context) error {
		iteration := &model.Iteration{}
		if err := c.Bind(iteration); err != nil {
			log.Error(err)
			return err
		}
		log.Infof("Create iteration %#v", iteration)
		result, err := dao.AddIteration(*iteration)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func GetIterationTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		itId := c.Param("id")
		log.Debugf("Get Tasks by iteration %s", itId)
		result, err := dao.GetTasksByIterations(itId)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

// Maybe dont need this.
func AddIterationTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		itId := c.Param("id")
		log.Debugf("Create task to iteration %s", itId)
		// TODO Not done
		result, err := dao.GetIterationById(itId)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}
