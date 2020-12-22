package v1

import (
	"net/http"

	"github.com/go-chi/chi"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/middleware"
	v1controllers "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers/device"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/v1/controllers/pets"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.OAuthMiddleware)

	r.Method("GET", "/", http.HandlerFunc(v1controllers.RootHandler))
	r.Method("GET", "/device/{deviceID}",
		http.HandlerFunc(device.GetDeviceHandler))
	r.Mount("/pet", pets.Router())

	return r
}
