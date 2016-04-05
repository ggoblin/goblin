package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"os"
)

func StartServer(config *SiteConfig) {
	gb := NewGoblin(config)
	err := gb.Init()
	if err != nil {
		log.Errorf("Goblin init fail. %s", err)
		return
	}
	gb.StartServer()
}

func main() {
	configPtr := flag.String("c", "config.toml", "Config file path.")
	debugPtr := flag.Bool("d", false, "Enable debug level.")

	flag.Parse()
	if len(*configPtr) < 1 {
		log.Error("Config file path must set.Use -h to get some help.")
	}

	if *debugPtr {
		log.SetLevel(log.DebugLevel)
	}

	var config SiteConfig
	if _, err := toml.DecodeFile(*configPtr, &config); err != nil {
		fmt.Println(err)
		return
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl != "" {
		config.DBConnection = dbUrl
	}

	fmt.Printf("%#v \n", config)
	StartServer(&config)
}
