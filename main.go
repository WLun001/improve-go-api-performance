package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const version = "1.0"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("version: %s", version))
	})

	app.Get("/cars", func(ctx *fiber.Ctx) error {
		return ctx.JSON(carList)
	})

	improved := app.Group("/improved")

	improved.Get("/cars", func(ctx *fiber.Ctx) error {
		var cars []Car
		for _, car := range carList {
			cars = append(cars, Car{ID: car.ID, Name: car.Name})
		}
		// only query required data
		// and return
		return ctx.JSON(cars)
	})

	log.Println(app.Listen(":3000"))
}
