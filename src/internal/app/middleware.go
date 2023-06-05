package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (app *App) initMiddleware() {
	app.server.Use(recover.New())
	app.server.Use(logger.New())
	app.server.Use(cors.New())
	app.server.Use(compress.New())

	app.server.Use(app.config.PrefixV1, filesystem.New(filesystem.Config{
		Root: http.FS(app.embedFs),
	}))
}
