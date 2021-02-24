package auth

import (
	"context"
	"net/http"

	"gorm.io/gorm"

	"wwwin-github.cisco.com/eti/idpadapter"
	etilogger "wwwin-github.cisco.com/eti/sre-go-logger"
)

const (
	label            = "okta"
	clientID         = "..."
	clientSecret     = "..."
	issuer           = "https://dev-....okta.com/oauth2/default"
	audience         = "api://default"
	loginCallback    = "http://localhost:5000/auth/login/token"
	issuerLogoutPath = "/v1/logout"
)

type Auth struct {
	log        *etilogger.Logger
	db         *gorm.DB
	idpAdapter *idpadapter.IdentityProviderAdapter
}

func New(l *etilogger.Logger, db *gorm.DB, httpClient *http.Client) (*Auth,
	error) {
	var ipa *idpadapter.IdentityProviderAdapter
	var err error

	if httpClient != nil {
		// a hacky way to skip over initializing in unit tests. the real way to
		// do this is to implement a custom http.RoundTripper for a
		// http.Client.Transport
		ipa, err = idpadapter.New(
			context.Background(),
			etilogger.NewNop(),
			httpClient,
			label,
			clientID,
			clientSecret,
			issuer,
			audience,
			loginCallback,
			"",
			issuerLogoutPath,
		)
		if err != nil {
			return nil, err
		}
	}

	return &Auth{
		log:        l,
		db:         db,
		idpAdapter: ipa,
	}, nil
}
