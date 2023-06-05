package app

import (
	"embed"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	config "github.com/itolog/yodi-api/src/configs"
	"github.com/itolog/yodi-api/src/pkg/logging"
)

type App struct {
	server  *fiber.App
	config  *config.Config
	Logging *logging.Logger
	embedFs *embed.FS
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
		server: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
			Prefork:     true,
		}),
		Logging: logging.GetLogger(),
	}
}

func (app *App) Start() error {
	app.initMiddleware()
	app.server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	app.server.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	return app.server.Listen(":" + app.config.Port)
}
