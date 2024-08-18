package repository

type JsonData struct {
	Data []FiestaData `json:"data"`
}

type FiestaData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation uint32 `json:"allocation"`
}
