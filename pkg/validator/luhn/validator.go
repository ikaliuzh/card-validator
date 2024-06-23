package luhn

import (
	"context"
	"fmt"

	"github.com/ikaliuzh/card-validator/pkg/card"
)

var ()

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(_ context.Context, card card.Card) error {
	checkDigit := luhnAlgorithm(card.Number[:len(card.Number)-1])

	if checkDigit != card.Number[len(card.Number)-1] {
		return fmt.Errorf("%w: card number failed Luhn algorithm validation", ErrInvalidCardNumber)
	}
	return nil
}

func luhnAlgorithm(cardNumber card.CardNumber) int {
	sum := 0
	isSecond := true

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := cardNumber[i]

		if isSecond == true {
			digit = digit * 2
		}

		sum += digit / 10
		sum += digit % 10

		isSecond = !isSecond
	}
	return (10 - (sum % 10)) % 10
}
