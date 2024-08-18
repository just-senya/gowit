package dto

type FiestaCreateReqDTO struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation uint32 `json:"allocation"`
}

type FrontRespDTO struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation uint32 `json:"allocation"`
}

type MakePurchaseReqDTO struct {
	Id       int
	UserId   string `json:"user_id"`
	Quantity int    `json:"quantity"`
}
