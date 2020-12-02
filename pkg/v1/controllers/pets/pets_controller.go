package pets

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/datastore"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	etilog "wwwin-github.cisco.com/eti/sre-go-logger"
)

var logger *etilog.Logger

// DeviceController struct
type PetController struct {
}

// AddRoutes
func (controller *PetController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/pet", controller.GetAllPets).Methods("GET")
	router.HandleFunc("/pet", controller.PostAllPets).Methods("POST")
	return router
}

// Get godoc
// @Summary Get All Pets
// @Description Get all pets in the pet store
// @Tags getallpets
// @Produce json
// @Success 200 {object} models.Pets
// @Failure 404 {object} models.Error
// @Router /pet [get]
func (controller *PetController) GetAllPets(w http.ResponseWriter, r *http.Request) {
	db, _ := datastore.DbConn()
	var pets []models.Pet
	db.Find(&pets)
	fmt.Println("{}", pets)
	json.NewEncoder(w).Encode(pets)
}

func (controller *PetController) PostAllPets(w http.ResponseWriter, r *http.Request) {
	// Declare a new Pet struct.
	var p models.Pet

	//// Try to decode the request body into the struct. If there is an error,
	//// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do something with the Person struct...
	db, _ := datastore.DbConn()
	fmt.Println("Adding new pet to db")

	db.Create(&models.Pet{Name: p.Name, Family: p.Family, Type: p.Type})

	PetFamilyCounter(p.Family)
	PetTypeCounter(p.Type)
	fmt.Fprintf(w, "New Pet %s added Successfully", p.Name)
}