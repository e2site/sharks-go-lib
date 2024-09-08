package server

import (
	"context"
	"errors"
	"github.com/e2site/sharks-go-lib/otl"
	"os"
	"os/signal"
)

func CreateAMQPServer(serviceName string, mainFunc func()) {
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

	mainFunc()
}
