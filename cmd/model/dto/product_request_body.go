package dto

type ProductRequestBody struct {
	Item     string `json:"item"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
