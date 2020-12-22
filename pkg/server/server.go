package server

import (
	"gorm.io/gorm"

	v1device "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/device"
	v1pet "wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/server/v1/pet"
	etilogger "wwwin-github.cisco.com/eti/sre-go-logger"
)

type Server struct {
	log      *etilogger.Logger
	v1device *v1device.Device
	v1pet    *v1pet.Pet
}

func New(l *etilogger.Logger, db *gorm.DB) *Server {
	return &Server{
		log:      l,
		v1device: v1device.New(l, db),
		v1pet:    v1pet.New(l, db),
	}
}
