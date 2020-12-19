package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

const (
	SharedAccessKeyHeader = "X-API-ACCESS-KEY"
	sharedAccessKeySecret = "c94bcd16-5e7c-4d41-95b0-70c9610e5663"
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
func SharedKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(SharedAccessKeyHeader) != sharedAccessKeySecret {
			log.Info("Shared key not authorized")

			err := utils.UnauthorizedResponse(w)
			if err != nil {
				log.Warn(err)
			}
			return
		}
		log.Info("Authorized request for " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
