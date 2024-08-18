package extracttypes

import dto "gowit-task/internal/ticket-app/dto"

type PurchaseBody struct {
	UserId   string
	Quantity int
	Id       int
}

func (p PurchaseBody) GetUserId() string {
	return p.UserId
}

func (p PurchaseBody) GetQuantity() int {
	return p.Quantity
}

func (p PurchaseBody) GetId() int {
	return p.Id
}

func ExtractPurchaseBody(source *dto.MakePurchaseReqDTO, dest *PurchaseBody) {
	dest.Id = source.Id
	dest.UserId = source.UserId
	dest.Quantity = source.Quantity
}
