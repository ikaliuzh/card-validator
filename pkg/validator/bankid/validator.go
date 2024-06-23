package bankid

import (
	"context"
	"errors"
	"fmt"

	"github.com/ikaliuzh/card-validator/pkg/card"
)

var (
	ErrInvalidCardNumberIIN = errors.New("invalid card BIN")
)

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
		return fmt.Errorf("%w: %q is not a supported IIN", ErrInvalidCardNumberIIN, iin)
	}
	return nil
}
