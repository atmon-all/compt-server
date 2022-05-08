package main

import (
	"flag"

	"github.com/atmom/compt/common"
	"github.com/atmom/compt/config"
	"github.com/atmom/compt/server"
)

var (
	configPath = flag.String("config", "", "The component server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start component server.
	server.Start(c)
}
