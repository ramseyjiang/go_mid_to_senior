package middleware

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type MiddleRuntimeFunc struct {
	err error
}

func (m *MiddleRuntimeFunc) MiddleClaimWrapTrace(f runtime.HandlerFunc) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		ctx := r.Context()
		funcName := GetFuncName()
		span := StartSpan(ctx, funcName)
		if span != nil {
			defer span.Finish()
		}
		ctx = context.WithValue(ctx, SpanStr, span)
		r = r.WithContext(ctx)
		f(w, r, pathParams)
	}
}
