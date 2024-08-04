package dto

type WebApiResponseSuccess struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
