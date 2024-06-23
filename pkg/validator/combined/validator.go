package combined

import (
	"context"

	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/validator"
)

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
