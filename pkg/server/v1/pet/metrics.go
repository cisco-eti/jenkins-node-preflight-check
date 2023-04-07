package pet

import "wwwin-github.cisco.com/eti/sre-go-sre-go-helloworld.git.git/pkg/metrics"

func PetTypeCounter(PetType string) {
	metrics.PetTypeCounter.WithLabelValues(PetType).Add(1)
}

func PetFamilyCounter(PetFamily string) {
	metrics.PetFamilyCounter.WithLabelValues(PetFamily).Add(1)
}
