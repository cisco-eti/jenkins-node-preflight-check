package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	etilogger "wwwin-github.cisco.com/eti/sre-go-logger"
)

func TestRoutes_GetDeviceZones(t *testing.T) {
	t.Run("returns Device_A zone", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/v1/device/A", nil)
		require.NoError(t, err)

		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")

		s := New(etilogger.NewNop(), nil)
		router := s.Router()
		router.ServeHTTP(response, request)

		apiRes := models.APIResponse{}
		err = json.Unmarshal(response.Body.Bytes(), &apiRes)
		require.NoError(t, err, "response body: %s", string(response.Body.Bytes()))

		want := "Plumbing"
		assert.Equal(t, want, apiRes.Data)
	})

	t.Run("returns Device_B zone", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/v1/device/B", nil)
		require.NoError(t, err)

		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")

		s := New(etilogger.NewNop(), nil)
		router := s.Router()
		router.ServeHTTP(response, request)

		apiRes := models.APIResponse{}
		err = json.Unmarshal(response.Body.Bytes(), &apiRes)
		require.NoError(t, err, "response body: %s", string(response.Body.Bytes()))

		want := "Gardening"
		assert.Equal(t, want, apiRes.Data)
	})
}
