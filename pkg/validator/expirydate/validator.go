package expirydate

import (
	"context"
	"errors"
	"fmt"

	"github.com/ikaliuzh/card-validator/pkg/card"
)

var (
	ErrExpired = errors.New("card is expired")
)

type Validator struct {
	DateProvider DateProvider
}

func New() *Validator {
	return &Validator{
		DateProvider: dateProvider{},
	}
}

func (v *Validator) Validate(_ context.Context, card card.Card) error {
	currentMonth := v.DateProvider.CurrentMonth()
	currentYear := v.DateProvider.CurrentYear()

	if card.ExpirationYear < currentYear || (card.ExpirationYear == currentYear && card.ExpirationMonth < currentMonth) {
		return fmt.Errorf("%w: got %d/%d", ErrExpired, card.ExpirationMonth, card.ExpirationYear)
	}

	return nil
}
