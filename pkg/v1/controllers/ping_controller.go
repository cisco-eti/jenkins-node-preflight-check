package controllers

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	//log "sqbu-github.cisco.com/Nyota/frontline-go-logger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/utils"
)

type ServiceInstance struct {
	instanceId string
	host       string
	port       int
}

type PingController struct {
	serviceName string
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
	res := utils.HTTPResponse{ResponseWriter: w}
	res.OKResponse("Hello-world")
}
