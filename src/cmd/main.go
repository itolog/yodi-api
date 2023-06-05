package main

import (
	config "github.com/itolog/yodi-api/src/configs"
	"github.com/itolog/yodi-api/src/internal/app"
)

func main() {
	cfg := config.NewConfig()
	server := app.NewApp(cfg)
	server.Logging.Info("Start server on port: ", cfg.Port)

	if err := server.Start(); err != nil {
		server.Logging.Fatal(err)
	}
}
