package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

func StartServer(config Config) {
	e := echo.New()
	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Get("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Run(":" + strconv.Itoa(config.Port))
}

func main() {
	configPtr := flag.String("c", "config.toml", "Config file path.")
	flag.Parse()
	if len(*configPtr) < 1 {
		log.Error("Config file path must set.Use -h to get some help.")
	}

	var config Config
	if _, err := toml.DecodeFile(*configPtr, &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", config)
	StartServer(config)
}

type Config struct {
	Port         int
	DBConnection string
}
