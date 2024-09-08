package client

import (
	"context"
	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
)

func CreateClient(ctx context.Context, tracerName string) *resty.Client {
	cli := resty.New()

	opts := []otelresty.Option{otelresty.WithTracerName(tracerName), otelresty.WithSkipper(func(r *resty.Request) bool {
		r.SetContext(ctx)
		return false
	})}

	otelresty.TraceClient(cli, opts...)

	return cli
}
