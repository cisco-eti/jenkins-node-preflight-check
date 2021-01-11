package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	slokmiddleware "github.com/slok/go-http-metrics/middleware"
	slokstd "github.com/slok/go-http-metrics/middleware/std"

	etimiddleware "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/middleware"
	v1device "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/device"
	v1pet "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/pet"
)

func Router(extraMiddleware ...func(http.Handler) http.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type",
			"Content-Length", "Accept-Encoding", "X-CSRF-Token",
			etimiddleware.SharedAccessKeyHeader},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any major browsers
	}))

	for _, mw := range extraMiddleware {
		r.Use(mw)
	}

	r.Method("GET", "/", http.HandlerFunc(RootHandler))
	r.Method("GET", "/ping", http.HandlerFunc(PingHandler))
	r.Method("GET", "/docs", http.HandlerFunc(DocsHandler))

	authedV1Router := chi.NewRouter()
	authedV1Router.Use(etimiddleware.OAuthMiddleware)
	authedV1Router.Mount("/device", v1device.Router())
	authedV1Router.Mount("/pet", v1pet.Router())
	r.Mount("/v1", authedV1Router)

	return r
}

func MetricMiddleware() func(http.Handler) http.Handler {
	mdlw := slokmiddleware.New(slokmiddleware.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
	})
	return slokstd.HandlerProvider("", mdlw)
}
