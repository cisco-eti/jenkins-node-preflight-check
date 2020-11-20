package v1

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	log "sqbu-github.cisco.com/Nyota/frontline-go-logger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/models"
	"testing"
)

var testCtx log.TraceContext

func setupContext() log.TraceContext {
	ctx := context.WithValue(context.Background(), log.UserIDKey, "User-1234")
	ctx = context.WithValue(ctx, log.TenantIDKey, "Tenant-1235")
	trace := log.NewTraceContextWithParent(ctx, "1234567890")
	return trace
}

func TestMain(m *testing.M) {
	log.LogInitTest("sre-go-helloworld")
	testCtx = setupContext()
	code := m.Run()
	os.Exit(code)
}

func TestRoutes_GetDeviceZones(t *testing.T) {
	t.Run("returns Device_A zone", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/device/A", nil)
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
		request, _ := http.NewRequest(http.MethodGet, "/device/B", nil)
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
