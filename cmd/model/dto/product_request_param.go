package dto

type ProductRequestParam struct {
	Id string `validate:"min=10"`
}
