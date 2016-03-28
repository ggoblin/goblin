package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetAllMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "World")
	}
}
