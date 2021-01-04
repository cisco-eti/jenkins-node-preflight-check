package controllers

import (
	"net/http"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Error 401
// @Router / [get]
func RootHandler(w http.ResponseWriter, _ *http.Request) {
	logger := utils.LogInit()
	logger.Info("Home / request received")

	_ = utils.OKResponse(w, nil)
	return
}
