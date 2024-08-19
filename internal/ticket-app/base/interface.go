package base

import (
	"gowit-task/internal/ticket-app/model"

	"github.com/gofiber/fiber/v2"
)

type FrontReq interface {
	GetName() string
	GetDesc() string
	GetAllocation() uint32
}

type PurchaseReq interface {
	GetId() int
	GetUserId() string
	GetQuantity() int
}

type PurchaseExtractor interface {
	Extract(ctx *fiber.Ctx) (PurchaseReq, error)
}

type PurchaseResponseWriter interface {
	Write(ctx *fiber.Ctx, statusCode int)
}

type FiestaRetrieveReq interface {
	GetId() int
}

type FiestaCreator interface {
	CreateFiesta(req FrontReq) (model.FiestaInfo, error)
}

type FiestaGetter interface {
	GetFiesta()
}

type RepositoryReader interface {
	Read(id int) (model.FiestaInfo, error)
}
type RepositoryWriter interface {
	Write(fiestaInfo model.FiestaInfo) (int, error)
}

type RepositoryPurchaseMaker interface {
	MakePurchase(id, quantity int, userId string) error
}

type ResponseWriter interface {
	Write(ctx *fiber.Ctx, statusCode int, info *model.FiestaInfo)
}

type RequestExtractor interface {
	Extract(ctx *fiber.Ctx) (FrontReq, error)
}

type RequestRetrieveExtractor interface {
	Extract(ctx *fiber.Ctx) (FiestaRetrieveReq, error)
}
