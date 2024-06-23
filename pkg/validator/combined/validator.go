package combined

import (
	"context"

	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/validator"
)

// Validator is a combined validator that runs multiple validators in sequence. It stops validation at the first error.
type Validator struct {
	validators []validator.Validator
}

func New(validators ...validator.Validator) *Validator {
	return &Validator{validators: validators}
}

func (v *Validator) Validate(ctx context.Context, card card.Card) error {
	for _, val := range v.validators {
		if err := val.Validate(ctx, card); err != nil {
			return err
		}
	}
	return nil
}
