package main

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/opentracing/api"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/opentracing/middleware"
)

func main() {
	closer := middleware.ProviderTracer()
	defer func(closer io.Closer) {
		_ = closer.Close()
	}(closer)
	opentracing.GlobalTracer()

	_ = api.NewRouter().Serve(":9000")
}
