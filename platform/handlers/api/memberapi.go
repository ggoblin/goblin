package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/dao"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Info("Start to query all members")
		members, err := dao.GetAllMembers()
		if err != nil {
			log.Error("Get all member fail.%s", err)
			return c.JSON(http.StatusInternalServerError, "Fail")
		}
		return c.JSON(http.StatusOK, members)
	}
}

func CreateMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		member := &model.Member{}
		if err := c.Bind(member); err != nil {
			log.Error(err)
			return err
		}
		log.Infof("Create member %#v", member)
		result, err := dao.AddNewMember(*member)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func GetMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		memberId := c.Param("id")
		log.Debugf("Get Member %s", memberId)
		result, err := dao.GetMember(memberId)
		if err != nil {
			log.Error(err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}
