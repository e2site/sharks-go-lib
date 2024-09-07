package client

import (
	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
)

func CreateClient(tracerName string) *resty.Client {
	cli := resty.New()
	opts := []otelresty.Option{otelresty.WithTracerName("tracerName")}

	otelresty.TraceClient(cli, opts...)
	return cli
}
