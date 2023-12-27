package main

import (
	"distributed_database_server/config"
	"os"

	"bitbucket.org/hasaki-tech/zeus/package/graceful"
	"go.uber.org/zap"
)

func main() {
	configPath := cconfig.GetConfigPath(os.Getenv("ENVIRONMENT"))
	config, err := config.GetConfig(configPath)
	if err != nil {
		panic(err)
	}

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	s := application.NewServer(config)

	go func() {
		s.Start()
	}()
	defer s.Stop()

	graceful.WaitShutdown()
}
