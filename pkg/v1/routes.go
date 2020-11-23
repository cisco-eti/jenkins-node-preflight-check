package v1

import (
	"github.com/gorilla/mux"
	//"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/middleware"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers"
)

// AddRoutes add version 1 routes to the Mux router
func AddRoutes(router *mux.Router) *mux.Router {
	homeCtrl := controllers.HomeController{}
	metricCtrl := controllers.MetricsController{}
	devZoneCtrl := controllers.DeviceController{}
	pingCtrl := controllers.PingController{}
	v1Router := router.PathPrefix("/").Subrouter()
	//v1Router.Use(middleware.OAuthMiddleware)
	homeCtrl.AddRoutes(v1Router)
	metricCtrl.AddRoutes(v1Router)
	devZoneCtrl.AddRoutes(v1Router)
	pingCtrl.AddRoutes(v1Router)
	return router
}
