package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// HomeController struct
type HomeController struct {
}

// AddRoutes add home routes to the Mux router
func (controller *HomeController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controller.Get).Methods("GET")
	return router
}

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Router / [get]
func (controller *HomeController) Get(w http.ResponseWriter, _ *http.Request) {
	logger = utils.LogInit()
	logger.Info("home / request received")
	res := utils.HTTPResponse{ResponseWriter: w}
	res.OKResponse("FrontDesk")
}
