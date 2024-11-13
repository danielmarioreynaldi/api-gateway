package main

import (
	"fmt"
	"os"

	"github.com/danielmarioreynaldi/api-gateway/config"
	"github.com/danielmarioreynaldi/api-gateway/http"
)

func main() {
	cfgFile := os.Getenv(config.CFG_FILE)
	cfg := config.LoadConfigFile(cfgFile)

	fmt.Println("Hello, world!")
	httpServer := http.NewHttpServer(cfg.HttpConfigs)

	go httpServer.Start()
	defer httpServer.Stop()

	//connect to redis
	//read rules and store in cache
	//implement rate limiting
}