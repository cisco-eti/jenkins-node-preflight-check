package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
)

func TestRoutes_GetDeviceZones(t *testing.T) {
	t.Run("returns Device_A zone", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/device/A", nil)
		require.NoError(t, err)

		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")

		router := Router()
		router.ServeHTTP(response, request)

		apiRes := models.APIResponse{}
		err = json.Unmarshal(response.Body.Bytes(), &apiRes)
		require.NoError(t, err, "response body: %s", string(response.Body.Bytes()))

		want := "Plumbing"
		assert.Equal(t, want, apiRes.Data)
	})

	t.Run("returns Device_B zone", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/device/B", nil)
		require.NoError(t, err)

		response := httptest.NewRecorder()
		request.Header.Set("Authorization", "Bearer 123456")

		router := Router()
		router.ServeHTTP(response, request)

		apiRes := models.APIResponse{}
		err = json.Unmarshal(response.Body.Bytes(), &apiRes)
		require.NoError(t, err, "response body: %s", string(response.Body.Bytes()))

		want := "Gardening"
		assert.Equal(t, want, apiRes.Data)
	})
}
