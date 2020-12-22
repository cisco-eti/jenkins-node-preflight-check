package pet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/datastore"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get All Pets
// @Description Get all pets in the pet store
// @Tags getallpets
// @Produce json
// @Success 200 {object} models.Pets
// @Failure 404 {object} models.Error
// @Router /pet [get]
func GetAllPets(w http.ResponseWriter, r *http.Request) {
	db, _ := datastore.DbConn()
	var pets []models.Pet
	db.Find(&pets)
	fmt.Println("{}", pets)
	json.NewEncoder(w).Encode(pets)
}

func PostAllPets(w http.ResponseWriter, r *http.Request) {
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

// Get godoc
// @Summary Get Pet by ID
// @Description Get one Pet by ID
// @Tags getpetid
// @Produce json
// @Success 200 {object} models.Pet
// @Failure 404 {object} models.Error
// @Router /pet [get]
func GetPetByID(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	petID := chi.URLParam(r, "petID")
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
func PostPetByID(w http.ResponseWriter, r *http.Request) {
	petID := chi.URLParam(r, "petID")
	logger := utils.LogInit()
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

// Get godoc
// @Summary Delete Pet by ID
// @Description Delete Pet by ID
// @Tags deletepetid
// @Success 200 {object}
// @Failure 404 {object} models.Error
// @Router /pet [delete]
func DeletePetByID(w http.ResponseWriter, r *http.Request) {
	petID := chi.URLParam(r, "petID")
	logger := utils.LogInit()
	logger.Info("DeletePetByID PetID:" + petID)

	// Do something with the Person struct...
	db, _ := datastore.DbConn()
	fmt.Printf("Delete petID %s\n", petID)

	var pet models.Pet
	db.Delete(&pet, petID)

	fmt.Fprintf(w, "Pet %s successfully updated", petID)
}
