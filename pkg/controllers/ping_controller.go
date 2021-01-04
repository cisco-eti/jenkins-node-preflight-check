package controllers

import (
	"net/http"
	"os"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get Ping
// @Description get helloworld status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.LogInit()
	logger.Info("/ping request received")

	_ = utils.OKResponse(w, os.Getenv("HOSTNAME"))
	return
}
