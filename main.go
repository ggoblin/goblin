package main

import (
	"github.com/ggoblin/goblin/goblin"

	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

func StartServer(config *goblin.SiteConfig) {
	gb := goblin.NetGoblin(config)
	gb.StartServer()
}

func main() {
	configPtr := flag.String("c", "config.toml", "Config file path.")
	flag.Parse()
	if len(*configPtr) < 1 {
		log.Error("Config file path must set.Use -h to get some help.")
	}

	var config goblin.SiteConfig
	if _, err := toml.DecodeFile(*configPtr, &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", config)
	StartServer(&config)
}
