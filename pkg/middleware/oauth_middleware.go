package middleware

import (
	"log"
	"net/http"
)

func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Println(r.RequestURI + token)
		if token != "Bearer 123456" && r.RequestURI != "/metrics" {
			// Validate authorization token here
			unauthorized := "401 Unauthorized request"
			log.Println("Unauthorized request for " + r.RequestURI)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(unauthorized))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("Authorized request for " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
