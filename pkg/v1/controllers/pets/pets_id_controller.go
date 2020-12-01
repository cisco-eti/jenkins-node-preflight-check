package pets

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/datastore"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// DeviceController struct
type PetIDController struct {
}

// AddRoutes
func (controller *PetIDController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/pet/{petID}", controller.GetPetByID).Methods("GET")
	router.HandleFunc("/pet/{petID}", controller.PostPetByID).Methods("POST")
	return router
}

// Get godoc
// @Summary Get Pet by ID
// @Description Get one Pet by ID
// @Tags getpetid
// @Produce json
// @Success 200 {object} models.Pet
// @Failure 404 {object} models.Error
// @Router /pet [get]
func (controller *PetIDController) GetPetByID(w http.ResponseWriter, r *http.Request) {
	logger = utils.LogInit()
	vars := mux.Vars(r)
	petID := vars["petID"]
	logger.Info("GetPetByID PetID:" + petID)

	db, _ := datastore.DbConn()
	var pet models.Pet
	db.Find(&pet, petID)
	json.NewEncoder(w).Encode(pet)
}

// Get godoc
// @Summary Update Pet by ID
// @Description Update one Pet by ID
// @Tags updatepetid
// @Success 200 {object}
// @Failure 404 {object} models.Error
// @Router /pet [get]
func (controller *PetIDController) PostPetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petID := vars["petID"]
	logger.Info("PostPetByID PetID:" + petID)

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

	db.Save(&p)

	PetFamilyCounter(p.Family)
	PetTypeCounter(p.Type)
	fmt.Fprintf(w, "Pet %s successfully updated", p.Name)
}
