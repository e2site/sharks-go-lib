package middleware

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	log2 "gitlab.com/shark-game/backend/shark-go-lib/log"
	"io"
)

func InitJaeger(service string, jaegerAddress string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1, // Всегда выбираем 100% запросов для трассировки
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerAddress, // Адрес Jaeger агента
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	log2.CheckEndLogFatal(err)

	opentracing.SetGlobalTracer(tracer) // Установка глобального трейсера

	return tracer, closer
}
