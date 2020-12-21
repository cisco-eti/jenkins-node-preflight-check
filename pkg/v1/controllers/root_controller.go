package v1controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// V1RootController struct
type V1RootController struct {
}

// AddRoutes add home routes to the Mux router
func (controller *V1RootController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controller.Get).Methods("GET")
	return router
}

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Error 401
// @Router / [get]
func (controller *V1RootController) Get(w http.ResponseWriter, _ *http.Request) {
	logger := utils.LogInit()
	logger.Info("Home / request received")

	_ = utils.UnauthorizedResponse(w)
}
