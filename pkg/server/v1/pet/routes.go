package pet

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Method("GET", "/", http.HandlerFunc(GetAllPets))
	r.Method("POST", "/", http.HandlerFunc(PostAllPets))
	r.Method("GET", "/{petID}", http.HandlerFunc(GetPetByID))
	r.Method("POST", "/{petID}", http.HandlerFunc(PostPetByID))
	r.Method("DELETE", "/{petID}", http.HandlerFunc(DeletePetByID))

	return r
}
