package server

import (
	"github.com/e2site/sharks-go-lib/middleware"
	"github.com/gin-gonic/gin"
	gintrace "github.com/opentracing-contrib/go-gin/ginhttp"
)

func CreateServer(serviceName string, jaegerAddress string, handler func(r *gin.Engine)) {
	tracer, closer := middleware.InitJaeger(serviceName, jaegerAddress)
	defer closer.Close()

	r := gin.Default()
	r.Use(gintrace.Middleware(tracer))
	handler(r)
	r.Run()
}
