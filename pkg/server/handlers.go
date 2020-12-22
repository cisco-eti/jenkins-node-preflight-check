package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	pong := os.Getenv("HOSTNAME")
	if pong == "" {
		pong = "HOSTNAME env missing... pong!"
	}

	_ = utils.OKResponse(w, pong)
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
