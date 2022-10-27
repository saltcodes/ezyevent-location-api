package handler

import (
	"context"
	"ezyevent-location-api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db = database.Config()
var ctx = context.TODO()

func InitRoutes(app *fiber.App) {

	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello ezyevents location API v1.0")
	})

	app.Get("/events", ListEvents)
	app.Post("/events", CreateEvent)
}
