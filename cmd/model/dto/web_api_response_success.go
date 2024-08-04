package dto

type WebApiResponseSuccess struct {
	Status string                `json:"status"`
	Data   []ProductResponseBody `json:"data"`
}
