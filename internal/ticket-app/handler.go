package ticketapp

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/log"
)

func createFiestaEventHandler(c *fiber.Ctx) error {

	return nil
}

func getFiestaEventHandler(ctx *fiber.Ctx) error {
	ctx.Context().SetStatusCode(200)
	ctx.Context().SetBodyString("getFiestaEventHandler: ok")
	return nil
}

func makePurchaseHandler(ctx *fiber.Ctx) error {
	return nil
}
