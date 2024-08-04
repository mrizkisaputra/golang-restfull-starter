package dto

type ProductRequestBody struct {
	Item     string `json:"item" validate:"required,min=1,max=200"`
	Price    int    `json:"price" validate:"required,gte=0"`
	Quantity int    `json:"quantity" validate:"required,gte=1,lte=10"`
}
