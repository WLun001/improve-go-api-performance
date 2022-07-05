package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func prettyJSON(ctx *fiber.Ctx, data interface{}) error {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return ctx.SendString(string(res))
}

func longTask() {
	fmt.Println("start long task")
	time.Sleep(time.Second * 5)
	fmt.Println("done long task")
}
