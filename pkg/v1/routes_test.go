package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	//log "wwwin-github.cisco.com/eti/sre-go-logger"
	"testing"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
)

func TestRoutes_GetDeviceZones(t *testing.T) {
	t.Run("returns Device_A zone", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/v1/device/A", nil)
		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")
		router := mux.NewRouter()
		AddRoutes(router)
		router.ServeHTTP(response, request)
		apiRes := models.APIResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &apiRes)
		want := "Plumbing"
		assert.Equal(t, want, apiRes.Data)
	})

	t.Run("returns Device_B zone", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/v1/device/B", nil)
		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")
		router := mux.NewRouter()
		AddRoutes(router)
		router.ServeHTTP(response, request)
		apiRes := models.APIResponse{}
		_ = json.Unmarshal(response.Body.Bytes(), &apiRes)
		want := "Gardening"
		assert.Equal(t, want, apiRes.Data)
	})
}
