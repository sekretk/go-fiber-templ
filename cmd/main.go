package main

import (
	"boy/go-fiber-templ/internal/home"
	"boy/go-fiber-templ/internal/home/config"
	"boy/go-fiber-templ/pkg/logger"
	"fmt"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init()
	dbConf := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	fmt.Println(dbConf)

	fmt.Println("Start")

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())

	home.NewHandler(app, customLogger)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
