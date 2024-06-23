package validator

import (
	"context"

	"github.com/ikaliuzh/card-validator/pkg/card"
)

//go:generate mockgen -source=validator.go -destination=mock/mock_validator.go
type Validator interface {
	Validate(context.Context, card.Card) error
}
