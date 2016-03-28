package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func SetRouting(e *echo.Echo) {
	log.Info("Start set api's routing")
	g := e.Group("/api")
	g.Get("/members", GetAllMembers())
	g.Post("/members", CreateMember())
}
