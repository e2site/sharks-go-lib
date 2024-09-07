package server

import (
	"context"
	"errors"
	"github.com/e2site/sharks-go-lib/middleware"
	"github.com/e2site/sharks-go-lib/otl"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"os"
	"os/signal"
)

func CreateServer(serviceName string, jaegerAddress string, handler func(r *gin.Engine)) {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := otl.SetupOTelSDK(ctx, serviceName)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))
	r.Use(middleware.TraceIDMiddleware())
	r.Use(middleware.TracingMiddleware())
	handler(r)
	r.Run()
}
