package controllers

import (
	"net/http"
	"path/filepath"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods",
		"POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}
