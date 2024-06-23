package expirydate

import (
	"time"
)

//go:generate mockgen -destination=mocks/mock_dateprovider.go -package=mocks github.com/ikaliuzh/card-validator/pkg/validator/expirydate DateProvider
type DateProvider interface {
	CurrentYear() int
	CurrentMonth() int
}

type dateProvider struct{}

func (p dateProvider) CurrentYear() int {
	return time.Now().Year()
}

func (p dateProvider) CurrentMonth() int {
	return int(time.Now().Month())
}
