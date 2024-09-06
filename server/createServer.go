package server

import (
	"github.com/e2site/sharks-go-lib/middleware"
	"github.com/gin-gonic/gin"
	gintrace "github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
)

func CreateServer(serviceName string, jaegerAddress string, handler func(r *gin.Engine)) {
	tracer, closer := middleware.InitJaeger(serviceName, jaegerAddress)
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	r := gin.Default()
	r.Use(gintrace.Middleware(tracer))
	r.Use(middleware.TracingMiddleware())
	handler(r)
	r.Run()
}
