package device

import (
	"net/http"

	"github.com/gorilla/mux"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/metrics"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
	etilog "wwwin-github.cisco.com/eti/sre-go-logger"
)

var logger *etilog.Logger

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
	logger = utils.LogInit()

	vars := mux.Vars(r)
	deviceID := vars["deviceID"]

	logger.Info("DeviceZoneHandler deviceId:" + deviceID)
	metrics.DeviceCounter.WithLabelValues(deviceID).Add(1)

	var err error
	switch deviceID {
	case "A":
		err = utils.OKResponse(w, "Plumbing")
	case "B":
		err = utils.OKResponse(w, "Gardening")
	case "C":
		err = utils.OKResponse(w, "Lighting")
	case "D":
		err = utils.OKResponse(w, "FrontDesk")
	default:
		err = utils.NotFoundResponse(w, "Unknown Device Specified")
	}

	if err != nil {
		logger.Warn("error: %s", err)
	}
}
