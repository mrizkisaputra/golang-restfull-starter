package dto

type ProductResponseBody struct {
	Id       string `json:"id"`
	Item     string `json:"item"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
