package v1

import (
	"github.com/gorilla/mux"
	v1controllers "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers/device"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers/pets"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/middleware"
)

// AddRoutes add version 1 routes to the Mux router
func AddRoutes(router *mux.Router) *mux.Router {
	rootCtrl := v1controllers.V1RootController{}
	devZoneCtrl := device.DeviceController{}
	petCtrl := pets.PetController{}
	petIdCtrl := pets.PetIDController{}
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middleware.OAuthMiddleware)
	rootCtrl.AddRoutes(v1Router)
	devZoneCtrl.AddRoutes(v1Router)
	petCtrl.AddRoutes(v1Router)
	petIdCtrl.AddRoutes(v1Router)
	return router
}
