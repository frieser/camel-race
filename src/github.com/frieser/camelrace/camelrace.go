package main

import (
	"flag"
	"github.com/frieser/camelrace/race"
)

//Main function. Parse command flags an boot the app
func main() {
	configFilePath := flag.String("config", "config.yml", "Path to the config file in Yml format")

	config := race.LoadFromFile(*configFilePath)

	race.Build(config).Start()

}
