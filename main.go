package main

import (
	"github.com/go-martini/martini"

	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"strconv"
)

func StartServer(config Config) {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.RunOnAddr(":" + strconv.Itoa(config.Port))

}

func main() {
	configPtr := flag.String("c", "", "Config file path.")
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
