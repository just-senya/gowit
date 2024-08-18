package purchase

import (
	"gowit-task/internal/ticket-app/base"
	"gowit-task/internal/ticket-app/dto"
	"gowit-task/internal/ticket-app/extracttypes"
)

func convertToPurchaseReq(req *dto.MakePurchaseReqDTO) base.PurchaseReq {
	var res PurchaseMakeExtractor

	extracttypes.ExtractPurchaseBody(req, &res.PurchaseBody)

	return res
}
