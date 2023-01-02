package middleware

import (
	"context"
	"io"
	"runtime"
	"strings"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/opentracing/opentracing-go"
)

const SpanStr = "ParentSpan"

func ProviderTracer() io.Closer {
	var cfg = jaegercfg.Configuration{
		ServiceName: "trace-test",
		Sampler: &jaegercfg.SamplerConfig{
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		Tags: []opentracing.Tag{{Key: "env", Value: "ramsey-test"}, {Key: "projectEnv", Value: "localhost-docker"}},
	}

	jMetricsFactory := metrics.NullFactory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		panic(err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
	return closer
}

func StartTrace() {
	// initialize the global singleton for tracing...
	tracer, reporter, closer, err := ginopentracing.InitTracing(
		"trace-test",
		"localhost:6831",
		ginopentracing.WithEnableInfoLog(false),
	)
	if err != nil {
		panic("unable to init tracing")
	}
	defer func(closer io.Closer) {
		_ = closer.Close()
	}(closer)
	defer reporter.Close()

	opentracing.SetGlobalTracer(tracer)
	opentracing.GlobalTracer()
}

func StartSpan(ctx context.Context, funcName string) (span opentracing.Span) {
	cSpan := ctx.Value(SpanStr)
	if cSpan != nil {
		parentSpan := cSpan.(opentracing.Span)
		span = opentracing.GlobalTracer().StartSpan(
			funcName,
			opentracing.ChildOf(parentSpan.Context()),
		)
	}
	return span
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	namePath := runtime.FuncForPC(pc).Name()
	funcNameSlice := strings.Split(namePath, "/")
	return funcNameSlice[len(funcNameSlice)-1]
}
