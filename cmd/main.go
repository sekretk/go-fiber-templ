package main

import (
	"boy/go-fiber-templ/internal/home"
	"boy/go-fiber-templ/internal/home/config"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init()
	dbConf := config.NewDatabaseConfig()

	log.Println(dbConf)

	fmt.Println("Start")

	app := fiber.New()

	app.Use(recover.New())

	home.NewHandler(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
