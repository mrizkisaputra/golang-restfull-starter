package dto

type ProductRequestBody struct {
	Item     string `json:"item" validate:"required,min=4,max=200"`
	Price    int    `json:"price" validate:"gte=1,required"`
	Quantity int    `json:"quantity" validate:"required,gte=1,lte=10"`
}
