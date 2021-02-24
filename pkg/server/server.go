package server

import (
	"net/http"

	"gorm.io/gorm"

	v1auth "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/auth"
	v1device "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/device"
	v1pet "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/pet"
	etilogger "wwwin-github.cisco.com/eti/sre-go-logger"
)

type Server struct {
	log      *etilogger.Logger
	v1auth   *v1auth.Auth
	v1device *v1device.Device
	v1pet    *v1pet.Pet
}

func New(l *etilogger.Logger, db *gorm.DB, httpClient *http.Client) (*Server,
	error) {
	a, err := v1auth.New(l, db, httpClient)
	if err != nil {
		return nil, err
	}

	return &Server{
		log:      l,
		v1auth:   a,
		v1device: v1device.New(l, db),
		v1pet:    v1pet.New(l, db),
	}, nil
}
