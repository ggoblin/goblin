package goblin

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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

	gb.server.Static("/static", "goblin/data/static")
	gb.SetupHandler()
	gb.server.Run(":" + strconv.Itoa(gb.config.Port))
}

func (gb *Goblin) SetupHandler() {
	t := &Template{
		templates: template.Must(template.ParseGlob("goblin/data/templates/*.html")),
	}
	gb.server.SetRenderer(t)

	gb.server.Get("/", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "index", "")
	})
}
