package api

import (
	"github.com/labstack/echo"
)

func SetRouting(e *echo.Echo) {
	g := e.Group("/api")
	g.Get("/members", GetAllMembers())
}
