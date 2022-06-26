package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

const version = "1.0"

func main() {
	app := fiber.New(fiber.Config{BodyLimit: utils.ConvertToBytes("2mb")})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("version: %s", version))
	})

	app.Get("/heroes", func(ctx *fiber.Ctx) error {
		hero, err := toHeroStruct()
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		return ctx.JSON(hero)
	})

	improved := app.Group("/improved")

	improved.Get("/heroes", func(ctx *fiber.Ctx) error {
		hero, err := toHeroStruct()
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		// remove members if not using it
		hero.Members = nil
		return ctx.JSON(hero)
	})

	log.Println(app.Listen(":3000"))
}
