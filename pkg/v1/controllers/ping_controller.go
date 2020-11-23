package controllers

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
// @Description get hello-world status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func (controller *PingController) Get(w http.ResponseWriter, r *http.Request) {
	logger = utils.LogInit()

	logger.Info("/ping request received")

	res := utils.HTTPResponse{ResponseWriter: w}

	res.OKResponse("Hello-world")

}
