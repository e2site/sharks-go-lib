package client

import "github.com/go-resty/resty/v2"

func CreateWithTelegramId(tracerName string, headerName string, headerValue string) *resty.Client {
	client := CreateClient(tracerName)
	client.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		headerName:     headerValue,
	})
	return client
}
