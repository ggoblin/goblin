package main

import (
	"github.com/ggoblin/goblin/platform/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"strconv"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Goblin struct {
	server *echo.Echo
	config *SiteConfig
}

func NewGoblin(config *SiteConfig) *Goblin {
	gb := new(Goblin)
	gb.config = config
	gb.server = echo.New()
	return gb
}

func (gb *Goblin) StartServer() {
	// Middleware
	gb.server.Use(mw.Logger())
	gb.server.Use(mw.Recover())

	gb.server.Static("/static", "data/static")
	gb.SetupHandler()
	gb.server.Run(standard.New(":" + strconv.Itoa(gb.config.Port)))
}

func (gb *Goblin) SetupHandler() {
	t := &Template{
		templates: template.Must(template.ParseGlob("data/templates/*.html")),
	}
	gb.server.SetRenderer(t)

	gb.server.Get("/", handlers.Index())
}
