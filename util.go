package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func prettyJSON(ctx *fiber.Ctx, data interface{}) error {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.SendString(string(res))
}
