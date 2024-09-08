package client

import (
	"context"
	"github.com/go-resty/resty/v2"
)

func CreateWithTelegramId(ctx context.Context, tracerName string, headerName string, headerValue string) *resty.Client {
	client := CreateClient(ctx, tracerName)
	client.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		headerName:     headerValue,
	})
	return client
}
