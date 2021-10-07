package telemetry

import (
	"context"
	"net/http"

	"go.opencensus.io/stats"
)

func RequestCountMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stats.Record(context.Background(), TotalConnections.M(1))
		next.ServeHTTP(w, r)
	})
}
