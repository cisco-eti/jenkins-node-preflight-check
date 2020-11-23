package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// OAuthMiddleware adds Mux middleware operations to authenticate requests
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := utils.HTTPResponse{ResponseWriter: w}
		token := r.Header.Get("Authorization")
		log.Info(r.RequestURI + token)
		if token != "Bearer 123456" && r.RequestURI != "/metrics" {
			res.UnauthorizedResponse()
			return
		}
		log.Info("Authorized request for " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
