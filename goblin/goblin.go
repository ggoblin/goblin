package goblin

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

type Goblin struct {
	server *echo.Echo
	config *SiteConfig
}

func NetGoblin(config *SiteConfig) *Goblin {
	gb := new(Goblin)
	gb.config = config
	gb.server = echo.New()
	return gb
}

func (gb *Goblin) StartServer() {
	// Middleware
	gb.server.Use(mw.Logger())
	gb.server.Use(mw.Recover())
	gb.server.Get("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	gb.server.Run(":" + strconv.Itoa(gb.config.Port))
}
