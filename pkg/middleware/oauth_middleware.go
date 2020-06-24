package middleware

import (
	"net/http"
	log "sqbu-github.cisco.com/Nyota/frontline-go-logger"
	"sqbu-github.cisco.com/Nyota/go-template/pkg/utils"
)

// OAuthMiddleware adds Mux middleware operations to authenticate requests
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := utils.HTTPResponse{ResponseWriter: w}
		token := r.Header.Get("Authorization")
		log.Tracer.Info(r.RequestURI + token)
		if token != "Bearer 123456" && r.RequestURI != "/metrics" {
			res.UnauthorizedResponse()
			return
		}
		log.Tracer.Info("Authorized request for " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
