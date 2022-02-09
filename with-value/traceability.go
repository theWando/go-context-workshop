package main

import (
	"context"
	"net/http"
)

type traceKey struct{}

func GetTraceabilityHeaders(ctx context.Context) http.Header {
	return ctx.Value(traceKey{}).(http.Header)
}

func SetTraceability(ctx context.Context, r http.Request) context.Context {
	var headers http.Header
	for k, v := range r.Header {
		if k == "x-request-id" {
			for _, s := range v {
				headers.Add(k, s)
			}
		}
	}
	return context.WithValue(ctx, traceKey{}, headers)
}
