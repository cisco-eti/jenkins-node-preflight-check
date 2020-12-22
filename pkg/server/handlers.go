package server

import (
	"net/http"
	"os"
	"path/filepath"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Error 401
// @Router / [get]
func RootHandler(w http.ResponseWriter, _ *http.Request) {
	logger := utils.LogInit()
	logger.Info("Home / request received")

	_ = utils.OKResponse(w, "root")
	return
}

// Get godoc
// @Summary Get Ping
// @Description get helloworld status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/ping request received")

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
func DocsHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/docs request received")

	filePath, err := filepath.Abs("./docs/openapi.yaml")
	if err != nil {
		utils.ServerErrorResponse(w, err.Error())
		return
	}

	logger.Info("filePath: %s", filePath)
	http.ServeFile(w, r, filePath)
}
