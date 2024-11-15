package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/danielmarioreynaldi/api-gateway/config"
	"github.com/danielmarioreynaldi/api-gateway/http"
	"github.com/danielmarioreynaldi/api-gateway/internal"
)

func main() {
	cfgFile := os.Getenv(config.CFG_FILE)
	cfg := config.LoadConfigFile(cfgFile)

	fmt.Println("Hello, world!")
	httpServer := http.NewHttpServer(cfg.HttpConfigs)

	httpServer.Router.HandleFunc("/{any}", internal.ForwardRequest)

	go httpServer.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	defer httpServer.Stop()

	//connect to redis
	//read rules and store in cache
	//implement rate limiting
}
