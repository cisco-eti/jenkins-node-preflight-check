package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	log "sqbu-github.cisco.com/Nyota/frontline-go-logger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/services"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/utils"
)

// DeviceController struct
type DeviceController struct {
}

// AddRoutes add device zone routes to the Mux router
func (controller *DeviceController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/device/{deviceID}", controller.Get).Methods("GET")
	return router
}

// Get godoc
// @Summary Get Device Zone
// @Description get device zone by device Id
// @ID get-device-zone
// @Tags DeviceZone
// @Accept json
// @Produce json
// @Param deviceId path string true "Device ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /deviceZone/{deviceId} [get]
func (controller *DeviceController) Get(w http.ResponseWriter, r *http.Request) {
	res := utils.HTTPResponse{ResponseWriter: w}
	vars := mux.Vars(r)
	deviceID := vars["deviceID"]
	log.Tracer.Info("DeviceZoneHandler deviceId:" + deviceID)
	services.Counter.WithLabelValues(deviceID).Add(1)
	switch deviceID {
	case "A":
		res.OKResponse("Plumbing")
	case "B":
		res.OKResponse("Gardening")
	case "C":
		res.OKResponse("Lighting")
	case "D":
		res.OKResponse("FrontDesk")
	default:
		res.NotFoundResponse("Unknown Device Specified")
	}
}