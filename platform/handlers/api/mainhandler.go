package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func SetRouting(e *echo.Echo) {
	log.Info("Start set api's routing")
	g := e.Group("/api")

	// Members
	g.Get("/members", GetAllMembers())
	g.Get("/members/:id", GetMember())
	g.Post("/members", CreateMember())

	// Iterations
	g.Get("/iterations", GetAllIterations())
	g.Get("/iterations/:id/tasks", GetIterationTasks())
	g.Post("/iterations", CreateIteration())

	// Task
	g.Post("/tasks", CreateTask())
}
