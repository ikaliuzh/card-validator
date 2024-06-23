package bankid

import (
	"context"
	"fmt"

	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/errorcodes"
)

// Validator validates the issuer ID of a card. This is a simple validator that checks if the issuer ID is known to it.
// It treats the first 6 digits of the card number as the IIN.
type Validator struct {
	knownIssuerIDs map[string]struct{}
}

func NewWithDefaultKnownIINs() *Validator {
	return &Validator{knownIssuerIDs: map[string]struct{}{
		"440066": {},
		"440393": {},
		"423223": {},
		"414720": {},
		"434769": {},
		"400022": {},
		"414709": {},
		"542418": {},
		"410039": {},
		"481582": {},
	}}
}

func (v *Validator) Validate(_ context.Context, c card.Card) error {
	iin := c.Number[:6].ToString()
	if _, ok := v.knownIssuerIDs[iin]; !ok {
		return fmt.Errorf("%w: %q is not a supported IIN", errorcodes.ErrInvalidCardNumberIIN, iin)
	}
	return nil
}
