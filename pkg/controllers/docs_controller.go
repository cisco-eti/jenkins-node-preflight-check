package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// DocsController struct
type DocsController struct {
}

// AddRoutes add home routes to the Mux router
func (controller *DocsController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/docs", controller.Get).Methods("GET")
	return router
}

// Get godoc
// @Summary Get API Docs
// @Description get Swagger API documentation
// @Produce Yaml
// @Success 200
// @Router /docs [get]
func (controller *DocsController) Get(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/docs request received")
	filePath, _ := filepath.Abs("./docs/openapi.yaml")
	logger.Info("filePath: %s", filePath)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filePath)
}
