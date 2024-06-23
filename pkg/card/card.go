package card

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/ikaliuzh/card-validator/gen/proto"

	"github.com/ikaliuzh/card-validator/pkg/errorcodes"
)

// Card represents a card.
type Card struct {
	// Number is a card number, i.e. the primary account number.
	Number CardNumber
	// ExpirationMonth is a month when the card expires.
	ExpirationMonth int
	// ExpirationYear is a year when the card expires.
	ExpirationYear int
}

type CardNumber []int

// NewCardFromProto creates a new Card from the protobuf Card message.
func NewCardFromProto(card *proto.Card) (Card, error) {
	cardNumber, err := NumberFromString(strings.TrimSpace(card.CardNumber))
	if err != nil {
		return Card{}, fmt.Errorf("%w: card number is invalid: %w", errorcodes.ErrInvalidCardNumberFormat, err)
	}
	if len(cardNumber) < 8 {
		return Card{}, fmt.Errorf("%w: card number length should be at least 8 digits: got %d",
			errorcodes.ErrInvalidCardNumberFormat, len(cardNumber))
	}
	if len(cardNumber) > 19 {
		return Card{}, fmt.Errorf("%w: card number length should be at most 19 digits: got %d",
			errorcodes.ErrInvalidCardNumberFormat, len(cardNumber))
	}

	month, err := strconv.Atoi(strings.TrimSpace(card.ExpirationMonth))
	if err != nil {
		return Card{}, fmt.Errorf("%w: expiration month is invalid: %w",
			errorcodes.ErrExpirationDateInvalidFormat, err)
	}
	if month < 1 || month > 12 {
		return Card{}, fmt.Errorf("%w: expiration month is invalid: got %d",
			errorcodes.ErrExpirationDateInvalidFormat, month)

	}

	year, err := strconv.Atoi(strings.TrimSpace(card.ExpirationYear))
	if err != nil {
		return Card{}, fmt.Errorf("%w: expiration year is invalid: %w",
			errorcodes.ErrExpirationDateInvalidFormat, err)
	}

	return Card{
		Number:          cardNumber,
		ExpirationMonth: month,
		ExpirationYear:  year,
	}, nil
}

// ToProto converts the Card to the protobuf Card message.
func (c *Card) ToProto() *proto.Card {
	return &proto.Card{
		CardNumber:      c.Number.ToString(),
		ExpirationMonth: strconv.Itoa(c.ExpirationMonth),
		ExpirationYear:  strconv.Itoa(c.ExpirationYear),
	}
}

func NumberFromString(s string) (CardNumber, error) {
	digits := make([]int, len(s))
	for i, r := range s {
		if !unicode.IsDigit(r) {
			return CardNumber{},
				fmt.Errorf("card number must contain only digits, got %q", r)
		}
		digits[i] = int(r - '0')
	}
	return digits, nil
}

func MustGetNumberFromString(s string) CardNumber {
	n, err := NumberFromString(s)
	if err != nil {
		panic(err)
	}
	return n
}

func (cn CardNumber) ToString() string {
	b := make([]byte, len(cn))
	for i, d := range cn {
		b[i] = byte(d) + '0'
	}
	return string(b)
}
