package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Error 401
// @Router / [get]
func (s *Server) RootHandler(w http.ResponseWriter, _ *http.Request) {
	s.log.Info("/ request received")

	_ = utils.OKResponse(w, "root")
	return
}

// Get godoc
// @Summary Get Ping
// @Description get helloworld status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/ping request received")

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "HOSTNAME env missing"
	}

	// Pure sample data, value need to be changed
	myService := models.Service{
		ServiceName:  "HelloWorld",
		ServiceType:  "OPTIONAL",
		ServiceState: "online",
		Message:      "Healthy",
		ServiceInstance: models.ServiceInstance{
			InstanceID: hostname,
			Host:       "172.18.231.5",
			Port:       21455,
		},
		LastUpdated:    "2020-10-20T08:42:07.290Z",
		BaseURL:        "http://helloworld.int.scratch-aws-1.prod.eticloud.io/",
		DurationPretty: "91ms",
		Duration:       91350005,
		UpstreamServices: []models.Service{
			{
				ServiceName:  "Postgres",
				ServiceType:  "REQUIRED",
				ServiceState: "online",
				ServiceInstance: models.ServiceInstance{
					InstanceID: "e3c16830-2c65-c6c8-68ab-30d728d6179e[9]",
					Host:       "172.18.244.25",
					Port:       5432,
				},
				Message:          "PostgresDataSource is online",
				LastUpdated:      "2021-01-06T13:43:42.984Z",
				DurationPretty:   "3ms",
				Duration:         3631684,
				UpstreamServices: []models.Service{},
				DefaultCharset:   "UTF-8",
			},
		},
		DefaultCharset: "UTF-8",
	}
	_ = utils.OKResponse(w, myService)
	return
}

// Get godoc
// @Summary Get API Docs
// @Description get Swagger API documentation
// @Produce Yaml
// @Success 200
// @Router /docs [get]
func (s *Server) DocsHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/docs request received")

	filePath, err := filepath.Abs("./docs/openapi.yaml")
	if err != nil {
		utils.ServerErrorResponse(w, err.Error())
		return
	}

	s.log.Info("filePath: %s", filePath)
	http.ServeFile(w, r, filePath)
}

// Get godoc
// @Summary Get Prometheus Metrics
// @Description get helloworld status
// @Produce Yaml
// @Success 200
// @Router /metrics [get]
func (s *Server) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/metrics request received")

	promhttp.Handler().ServeHTTP(w, r)
}
