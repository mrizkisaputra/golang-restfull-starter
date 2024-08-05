package dto

type WebApiResponseError struct {
	Status           string `json:"status"`
	Error            any	`json:"error"`
	TraceId          string	`json:"trace_id"`
	DocumentationURL string	`json:"documentation_url"`
}
