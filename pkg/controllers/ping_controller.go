package controllers

import (
	//"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

type PingController struct {
}

// AddRoutes add home routes to the Mux router
func (controller *PingController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ping", controller.Get).Methods("GET")
	return router
}

// Get godoc
// @Summary Get Ping
// @Description get helloworld status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func (controller *PingController) Get(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/ping request received")
	res := utils.HTTPResponse{ResponseWriter: w}
	res.OKResponse(os.Getenv("HOSTNAME"))
}
