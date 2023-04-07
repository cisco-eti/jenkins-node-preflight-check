package server

import (
	"gorm.io/gorm"

	"wwwin-github.cisco.com/eti/idpadapter"
	etilogger "wwwin-github.cisco.com/eti/sre-go-logger"
	v1auth "wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/server/v1/auth"
	v1device "wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/server/v1/device"
	v1pet "wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/server/v1/pet"
)

type Server struct {
	log      *etilogger.Logger
	v1auth   *v1auth.Auth
	v1device *v1device.Device
	v1pet    *v1pet.Pet
}

func New(l *etilogger.Logger, db *gorm.DB,
	ipa *idpadapter.IdentityProviderAdapter) *Server {
	return &Server{
		log:      l,
		v1auth:   v1auth.New(l, db, ipa),
		v1device: v1device.New(l, db),
		v1pet:    v1pet.New(l, db),
	}
}
