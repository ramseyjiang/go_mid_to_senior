package api

import (
	"log"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/opentracing/middleware"
	"github.com/uber/jaeger-client-go"
)

func SpanFromParent(c *gin.Context) {
	var span opentracing.Span
	if cSpan, ok := c.Get("tracing-context"); ok {
		span = ginopentracing.StartSpanWithParent(
			cSpan.(opentracing.Span).Context(),
			"Span from parent",
			c.Request.Method, c.Request.URL.Path,
		)
	}

	// get traceID
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		log.Println(sc.TraceID())
	}

	defer span.Finish()
	c.String(200, "Span from parent")
}

func SpanFromHeader(c *gin.Context) {
	span := ginopentracing.StartSpanWithHeader(
		&c.Request.Header,
		"Span from header",
		c.Request.Method,
		c.Request.URL.Path,
	)
	defer span.Finish()
	c.String(200, "Span from header")
}

func SpanHasFuncName(c *gin.Context) {
	funcName := middleware.GetFuncName()
	var span opentracing.Span
	if cSpan, ok := c.Get("tracing-context"); ok {
		span = ginopentracing.StartSpanWithParent(
			cSpan.(opentracing.Span).Context(),
			funcName,
			c.Request.Method, c.Request.URL.Path,
		)
	}

	// get traceID
	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		log.Println(sc.TraceID())
	}

	defer span.Finish()
	c.String(200, "Span customise the operationName")
}
