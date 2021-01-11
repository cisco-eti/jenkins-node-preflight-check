package device

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Method("GET", "/{deviceID}", http.HandlerFunc(GetDeviceHandler))

	return r
}
