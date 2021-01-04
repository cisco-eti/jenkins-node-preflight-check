package pkg

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	slokmiddleware "github.com/slok/go-http-metrics/middleware"
	slokstd "github.com/slok/go-http-metrics/middleware/std"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/controllers"
	etimiddleware "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/middleware"
	v1 "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1"
)

func Router() *chi.Mux {
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

	mdlw := slokmiddleware.New(slokmiddleware.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
	})
	r.Use(slokstd.HandlerProvider("", mdlw))

	r.Method("GET", "/", http.HandlerFunc(controllers.RootHandler))
	r.Method("GET", "/metrics", http.HandlerFunc(controllers.MetricsHandler))
	r.Method("GET", "/ping", http.HandlerFunc(controllers.PingHandler))
	r.Method("GET", "/docs", http.HandlerFunc(controllers.DocsHandler))
	r.Mount("/v1", v1.Router())

	return r
}
