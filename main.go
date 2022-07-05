package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

const (
	version = "1.0"
	user    = "john"
)

func main() {
	app := fiber.New(fiber.Config{BodyLimit: utils.ConvertToBytes("2mb")})
	store := session.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("version: %s", version))
	})

	app.Get("/heroes", func(ctx *fiber.Ctx) error {
		hero, err := toHeroStruct()
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		return prettyJSON(ctx, hero)
	})

	app.Get("/rockets", func(ctx *fiber.Ctx) error {
		return ctx.JSON(findPayloadWeightMap("leo"))
	})

	app.Get("/longTask", func(ctx *fiber.Ctx) error {
		longTask()
		return ctx.SendStatus(fiber.StatusAccepted)
	})

	improved := app.Group("/improved")

	improved.Get("/heroes", func(ctx *fiber.Ctx) error {
		hero, err := toHeroStruct()
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		// remove members if not using it
		hero.Members = nil
		return prettyJSON(ctx, hero)
	})

	improved.Get("/longTask", func(ctx *fiber.Ctx) error {
		go longTask()
		return ctx.SendStatus(fiber.StatusAccepted)
	})

	authSession := app.Group("/authSession")
	authSession.Post("/login", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		sess.Set("name", user)
		sess.Set("secret", utils.UUIDv4())
		// Save session
		if err := sess.Save(); err != nil {
			return fiber.NewError(500, err.Error())
		}
		return ctx.SendString(fmt.Sprintf("Welcome %v, secret has been generated", user))
	})
	authSession.Get("/content", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		name := sess.Get("name")
		data := sess.Get("secret")
		if name == nil || data == nil {
			return ctx.SendStatus(401)
		}

		return ctx.SendString(fmt.Sprintf("Secret of %v is %v", name, data))
	})
	authSession.Post("/logout", func(ctx *fiber.Ctx) error {
		sess, err := store.Get(ctx)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		sess.Delete("name")

		// Destroy session
		if err := sess.Destroy(); err != nil {
			return fiber.NewError(500, err.Error())
		}
		return ctx.SendString("logged out")
	})

	authJWT := app.Group("/authJWT")
	authJWT.Post("/login", func(ctx *fiber.Ctx) error {
		token, err := createJWTToken(user, utils.UUIDv4())
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		ctx.Cookie(&fiber.Cookie{
			Name:        "jwt",
			Value:       token.Token,
			MaxAge:      token.ExpiresAt.Second(),
			Expires:     token.ExpiresAt,
			SessionOnly: false,
		})
		return ctx.JSON(
			fiber.Map{
				"token":   token,
				"message": fmt.Sprintf("Welcome %v, secret has been generated", user),
			})
	})
	authJWT.Get("/content", func(ctx *fiber.Ctx) error {
		jwtCookies := ctx.Cookies("jwt")
		if jwtCookies == "" {
			return ctx.SendStatus(401)
		}
		claims, err := verifyJWTToken(jwtCookies)
		if err != nil {
			return ctx.Status(401).JSON(err)
		}

		return ctx.SendString(fmt.Sprintf("Secret of %v is %v", claims.Subject, claims.Secret))
	})
	authJWT.Post("/logout", func(ctx *fiber.Ctx) error {
		ctx.ClearCookie("jwt")
		return ctx.SendString("logged out")
	})

	log.Println(app.Listen(":3000"))
}
