package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/utils"
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
	res := utils.HTTPResponse{ResponseWriter: w}
	res.OKResponse("Demo Home")
}
