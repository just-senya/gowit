package ticketapp

import (
	"context"
	"errors"

	"gowit-task/internal/ticket-app/purchase"
	"gowit-task/internal/ticket-app/repository/postgres"
	"gowit-task/internal/ticket-app/ticket_create"
	"gowit-task/internal/ticket-app/ticket_retrieve"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	jsonDbPath = "/Users/arsen/Documents/interview/gowit/repository/data.json"
)

type App struct {
	app *fiber.App
	// config
}

func NewApp() App {
	app := fiber.New()
	return App{app: app}
}

func (app App) Register() error {
	app.app.Use(recover.New())

	postgresDb, err := postgres.NewPostgresStore()
	if err != nil {
		return err
	}
	postgresDb.Init()

	extr := ticket_create.NewCreateExtractor()
	fiestaEvent := ticket_create.NewFiestaEventCreator()
	respWriter := ticket_create.NewRespWriter()

	fiestaCreator := NewFiestaCreateProcessor(extr, fiestaEvent, respWriter, postgresDb)
	if fiestaCreator == nil {
		return errors.New("fail to register app: fiestaCreator")
	}

	getExtractor := ticket_retrieve.NewTicketRetrieveExtractor()
	fiestaRetriever := NewFiestaGetProcessor(getExtractor, postgresDb, respWriter)
	if fiestaRetriever == nil {
		return errors.New("fail to register app: fiestaRetriever")
	}

	purchaseExtractor := purchase.NewPurchaseMakeExtractor()
	purchaseRespWriter := purchase.NewRespWriter()
	purchaseMaker := NewPurchaseMakeProcessor(purchaseExtractor, postgresDb, purchaseRespWriter)

	app.app.Route("/tickets", func(router fiber.Router) {
		router.Post("", fiestaCreator.Handle)
		router.Get("/:id<int>", fiestaRetriever.Handle)
		router.Post("/:id<int>/purchases", purchaseMaker.Handle)
	}, "tickets.")

	return nil
}

func (app App) Resolve(ctx context.Context) error {
	return app.app.Listen(":3000")
}

func (app App) Release() error {
	return nil
}
