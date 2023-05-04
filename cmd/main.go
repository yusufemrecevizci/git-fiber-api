package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yusufemrecevizci/go-fiber-api/pkg/common/config"
	"github.com/yusufemrecevizci/go-fiber-api/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db.Init(c.DBUrl)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(c.Port)
	})

	// books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
