package root

import (
	"github.com/gorilla/mux"
	controllers "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/controllers"
)

// AddRoutes to the root Mux router
func AddRoutes(router *mux.Router) *mux.Router {
	rootCtrl := controllers.RootController{}
	metricCtrl := controllers.MetricsController{}
	pingCtrl := controllers.PingController{}
	Router := router.PathPrefix("/").Subrouter()
	rootCtrl.AddRoutes(Router)
	metricCtrl.AddRoutes(Router)
	pingCtrl.AddRoutes(Router)
	return router
}
