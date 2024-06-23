package errorcodes

import (
	"errors"
)

type ValidationError struct {
	statusCode string
	error      error
}

func (e ValidationError) StatusCode() string {
	return e.statusCode
}

func (e ValidationError) Error() string {
	return e.error.Error()
}

const (
	// ExpirationDateOther is returned for all expiration date validation errors that don't have more specific
	// error codes.
	ExpirationDateOther = "100"
	// ExpirationDateInvalidFormat is returned for all errors related to invalid date format of the expiration
	// month or expiration year.
	ExpirationDateInvalidFormat = "101"
	// Expired is returned when the card is expired.
	Expired = "102"

	// CardNumberOther is returned for all card number validation errors that don't have more specific error codes.
	CardNumberOther = "200"
	// CardNumberInvalidFormat is returned when the card number has an invalid format, i.e. it contains non-digit
	// symbols.
	CardNumberInvalidFormat = "201"
	// CardNumberLuhnFailed is returned when the card number fails the Luhn algorithm validation.
	CardNumberLuhnFailed = "202"
	// CardNumberInvalidIIN is returned when the card number has an invalid Issuer Identification Number.
	CardNumberInvalidIIN = "203"

	// Other is returned for all validation errors that don't have more specific error codes.
	Other = "999"
)

var (
	ErrExpirationDate              = errors.New("expiration date validation failed")
	ErrExpirationDateInvalidFormat = errors.New("invalid expiration date format")
	ErrExpired                     = errors.New("card is expired")

	ErrInvalidCardNumber       = errors.New("card number validation failed")
	ErrInvalidCardNumberFormat = errors.New("invalid card number format")
	ErrCardNumberLuhnFailed    = errors.New("card number failed Luhn algorithm validation")
	ErrInvalidCardNumberIIN    = errors.New("invalid card IIN")

	ErrOther = errors.New("validation error")
)

// Extract extracts the error code from the error. If the error is not a validation error, it returns false.
// A provided error should wrap one of the validation errors defined in this package to be treated as an error
// with a valid error code.
func Extract(err error) (string, bool) {
	if errors.Is(err, ErrExpirationDate) {
		return ExpirationDateOther, true
	}
	if errors.Is(err, ErrExpirationDateInvalidFormat) {
		return ExpirationDateInvalidFormat, true
	}
	if errors.Is(err, ErrExpired) {
		return Expired, true
	}
	if errors.Is(err, ErrInvalidCardNumber) {
		return CardNumberOther, true
	}
	if errors.Is(err, ErrInvalidCardNumberFormat) {
		return CardNumberInvalidFormat, true
	}
	if errors.Is(err, ErrCardNumberLuhnFailed) {
		return CardNumberLuhnFailed, true
	}
	if errors.Is(err, ErrInvalidCardNumberIIN) {
		return CardNumberInvalidIIN, true
	}
	if errors.Is(err, ErrOther) {
		return Other, true
	}
	return "", false
}
