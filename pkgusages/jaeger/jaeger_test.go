package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func TestJaeger(t *testing.T) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "{host}:6831", // replace host
		},
	}

	closer, err := cfg.InitGlobalTracer(
		"serviceName",
	)

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}

	var ctx = context.TODO()
	span1, ctx := opentracing.StartSpanFromContext(ctx, "span_1")
	time.Sleep(time.Second / 2)
	span11, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second / 2)
	span11.Finish()
	span1.Finish()
	defer closer.Close()
}
