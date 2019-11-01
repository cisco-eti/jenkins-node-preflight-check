package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETdeviceZone(t *testing.T) {
	t.Run("returns Device_A zone", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/deviceZone/A", nil)
		response := httptest.NewRecorder()

		getDeviceZone(response, request)

		got := response.Body.String()
		want := "Plumbing"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Device_B zone", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/deviceZone/B", nil)
		response := httptest.NewRecorder()

		getDeviceZone(response, request)

		got := response.Body.String()
		want := "Gardening"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
